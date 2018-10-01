package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/streadway/amqp"
)

// Operation docstring.
type Operation struct {
	Name   string
	Params []int
}

// AddParam docstring.
func (op *Operation) AddParam(x int) {
	op.Params = append(op.Params, x)
}

// GetParam docstring.
func (op *Operation) GetParam() int {
	r := op.Params[0]
	op.Params = op.Params[1:]
	return r
}

// Response docstring.
type Response struct {
	Name   string
	Result int
}

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

func fibonacciRPC(n int, msgs <-chan amqp.Delivery, chout *amqp.Channel, q amqp.Queue) Response {
	corrID := randomString(32)
	op := new(Operation)
	op.Name = "fib"
	op.AddParam(15)
	outcoming, err := json.Marshal(op)

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
			Body:          outcoming,
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

	var incoming Response
	err = json.Unmarshal(aux, &incoming)

	return incoming
}

func main() {
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

	for i := 0; i < 1000; i++ {
		fibonacciRPC(15, msgs, chout, q)
		time.Sleep(10 * time.Second)
	}
}
