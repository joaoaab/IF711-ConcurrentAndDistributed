package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func sendMessage(ch *amqp.Channel) {
	err := ch.Publish(
		"chatglobal", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Oi"), // Mensagem a ser enviada
		})
	failOnError(err, "Failed to publish a message")
}

func receiveMessage(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		"chatglobal", // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	sendMessage(ch) // Sends first message

	innerlimit := 0
	for range msgs {
		// log.Printf(" [x] %s", d.Body)
		innerlimit++
		if innerlimit < 10000 { // Iterações de envios de mensagem
			sendMessage(ch) // Envio sincrono com o recebimento
		} else {
			break
		}
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"chatglobal", // name
		"fanout",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	outlimit := 0
	for outlimit < 50 { // Iterações de teste
		start := time.Now()
		receiveMessage(ch)
		elapsed := time.Since(start)
		fmt.Println(float64(elapsed) / float64(time.Millisecond))

		// Count
		outlimit++
	}

}
