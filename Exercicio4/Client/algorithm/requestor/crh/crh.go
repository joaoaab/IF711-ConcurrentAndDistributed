package crh

import (
	"fmt"
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
type ClientRequestHandler interface {
	Setup(host string, port int) error
	Send(outcoming []byte)
	Receive() []byte
	Close()
}

// MiddlewareRequestHandler docstrig.
type MiddlewareRequestHandler struct {
	host          string
	port          int
	connection    *amqp.Connection
	channel       *amqp.Channel
	queue         amqp.Queue
	incoming      <-chan amqp.Delivery
	correlationID string
}

// Setup for MiddlewareRequestHandler.
func (handler *MiddlewareRequestHandler) Setup(host string, port int) error {
	var err error
	handler.host = host
	handler.port = port
	handler.correlationID = randomString(32)

	handler.connection, err = amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%d/", host, port))
	if err != nil {
		return err
	}

	handler.channel, err = handler.connection.Channel()
	if err != nil {
		return err
	}

	handler.queue, err = handler.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // noWait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	handler.incoming, err = handler.channel.Consume(
		handler.queue.Name, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		return err
	}

	return nil
}

// Send for MiddlewareRequestHandler.
func (handler *MiddlewareRequestHandler) Send(outcoming []byte) {
	err := handler.channel.Publish(
		"",          // exchange
		"rpc_queue", // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: handler.correlationID,
			ReplyTo:       handler.queue.Name,
			Body:          outcoming,
		})
	failOnError(err, "Failed to publish a message")
}

// Receive for MiddlewareRequestHandler.
func (handler *MiddlewareRequestHandler) Receive() []byte {
	msg := []byte("error")

	for d := range handler.incoming {
		if handler.correlationID == d.CorrelationId {
			msg = d.Body
			break
		}
	}

	return msg
}

// Close for MiddlewareRequestHandler.
func (handler *MiddlewareRequestHandler) Close() {
	handler.channel.Close()
	handler.connection.Close()
}

// // TcpRequestHandler docstrig.
// type TCPRequestHandler struct {
// 	Host          string
// 	Port          int
// 	connection    *net.Conn
// }

// // UdpRequestHandler docstrig.
// type UDPRequestHandler struct {
// 	Host          string
// 	Port          int
// 	connection    *net.UDPConn
// }

// // SendUDP docstring
// func (crh *ClientRequestHandler) SendUDP(outcoming []byte) {
// 	crh.UDPConnection.Write(outcoming)
// }

// // ReceiveUDP docstring
// func (crh *ClientRequestHandler) ReceiveUDP() []byte {
// 	buf := make([]byte, 1024)
// 	n, _, err := crh.UDPConnection.ReadFromUDP(buf)
// 	failOnError(err, "Error reading from UDP")
// 	return buf[0:n]
// }

// // SendTCP docstring
// func (crh *ClientRequestHandler) SendTCP(outcoming []byte) {
// 	crh.TCPConnection.Write(outcoming)
// }

// // ReceiveTCP docstring
// func (crh *ClientRequestHandler) ReceiveTCP() []byte {
// 	reader := bufio.NewReader(crh.TCPConnection)
// 	ans, err := reader.ReadBytes('\n')
// 	failOnError(err, "Failed to read from TCP")
// 	return ans
// }
