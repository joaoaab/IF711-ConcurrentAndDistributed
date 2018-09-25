package shandler

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"github.com/streadway/amqp"
)

// Messages Channel
var Messages = make(chan Message)

// Reply Channel
var Reply = make(chan Message)

// Message type
type Message struct {
	Data     string
	Addr     net.Addr
	Protocol int
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func acceptTCPConnections(server net.Listener, newConnections chan net.Conn) {
	for {
		c, err := server.Accept()
		if err != nil {
			fmt.Println("Error Acception:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Separando a thread para comunicação")
		newConnections <- c
	}
}

func handleTCPConnection(conn net.Conn, TCPMessages chan Message, deadTCPConnections chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		m := Message{message, conn.RemoteAddr(), 0}
		//fmt.Println("li do socket " + m.Data)
		TCPMessages <- m
	}
	deadTCPConnections <- conn
}

func handleUDPMessages(conn *net.UDPConn, allClients map[int]int, udpMessages chan Message) {
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf) // buf[0:n]
		if err != nil {
			continue
		}
		m := Message{string(buf[0:n]), addr, 1}
		udpMessages <- m
	}
}

// HandleTCP handles tcp connections for the server
func HandleTCP() {
	TCPConnections := make(map[net.Addr]net.Conn)
	newTCPConnections := make(chan net.Conn)
	deadTCPConnections := make(chan net.Conn)
	TCPMessages := make(chan Message)
	s, err := net.Listen("tcp", "localhost:6900")
	checkError(err)
	go acceptTCPConnections(s, newTCPConnections)
	defer s.Close()

	for {
		select {
		case conn := <-newTCPConnections:
			TCPConnections[conn.RemoteAddr()] = conn
			go handleTCPConnection(conn, TCPMessages, deadTCPConnections)
			fmt.Println(conn.RemoteAddr())
		case conn := <-deadTCPConnections:
			delete(TCPConnections, conn.RemoteAddr())
		case msg := <-TCPMessages:
			Messages <- msg
		case ret := <-Reply:
			address := ret.Addr
			data := ret.Data
			TCPConnections[address].Write([]byte(data))
		}
	}
}

// HandleUDP handles udp connections for the server
func HandleUDP() {
	// Setting up UDP
	udpClients := make(map[int]int)
	udpMessages := make(chan Message)
	serverAddr, err := net.ResolveUDPAddr("udp", ":1111")
	checkError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	checkError(err)
	defer serverConn.Close()

	go handleUDPMessages(serverConn, udpClients, udpMessages)

	for {
		select {
		case msg := <-udpMessages:
			Messages <- msg
		case ret := <-Reply:
			address := ret.Addr
			data := ret.Data
			fullAddr, err := net.ResolveUDPAddr("udp", address.String())
			if err != nil {
				fmt.Println("Resolver Failed with error : ", err)
				continue
			}
			_, err = serverConn.WriteToUDP([]byte(data), fullAddr)
		}
	}
}

// HandleMiddleware handles middleware connections using rabbitmq
func HandleMiddleware() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			rec := string(d.Body)

			fmt.Println("Received " + string(rec))
			Messages <- Message{Data: rec, Addr: nil, Protocol: 2}

			ret := <-Reply

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(ret.Data),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf("Awaiting Requests with RabbitMQ")
	<-forever
}
