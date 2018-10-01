package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (int, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	chin, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chin.Close()

	chout, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chout.Close()

	q, err := chout.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := chin.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	corrID := randomString(32)

	start := time.Now()
	err = chout.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       q.Name,
			Body:          []byte(strconv.Itoa(n)),
		})
	failOnError(err, "Failed to publish a message")

	var aux []byte
	for d := range msgs {
		if corrID == d.CorrelationId {
			aux = d.Body
			failOnError(err, "Failed to convert body to integer")
			break
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("%.3f\n", float64(elapsed)/float64(time.Millisecond))

	return strconv.Atoi(string(aux))
}

func main() {
	for i := 0; i < 1000; i++ {
		_, err := fibonacciRPC(15)
		failOnError(err, "Failed to handle RPC request")
		time.Sleep(10 * time.Second)
	}
}

func bodyFrom(args []string) int {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "30"
	} else {
		s = strings.Join(args[1:], " ")
	}
	n, err := strconv.Atoi(s)
	failOnError(err, "Failed to convert arg to integer")
	return n
}
