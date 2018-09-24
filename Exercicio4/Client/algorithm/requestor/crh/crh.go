package crh

import (
	"log"
	"math/rand"

	"github.com/streadway/amqp"
)

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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// ClientRequestHandler docstring.
type ClientRequestHandler struct {
	Host          string
	Port          int
	connection    *amqp.Connection
	channel       *amqp.Channel
	queue         amqp.Queue
	correlationID string
	// sentMsgSize    int
	// receiveMsgSize int
}

// Send docstring.
func (crh *ClientRequestHandler) Send(outcoming []byte) {
	// crh.sentMsgSize = len(outcoming)
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	corrID := randomString(32)

	err = ch.Publish(
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

	crh.connection = conn
	crh.channel = ch
	crh.queue = q
	crh.correlationID = corrID
}

// Receive docstring.
func (crh *ClientRequestHandler) Receive() []byte {
	defer crh.connection.Close()
	defer crh.channel.Close()

	msg := []byte("error")
	msgs, err := crh.channel.Consume(
		crh.queue.Name, // queue
		"",             // consumer
		true,           // auto-ack
		false,          // exclusive
		false,          // no-local
		false,          // no-wait
		nil,            // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		if crh.correlationID == d.CorrelationId {
			msg = d.Body
			break
		}
	}

	return msg
}
