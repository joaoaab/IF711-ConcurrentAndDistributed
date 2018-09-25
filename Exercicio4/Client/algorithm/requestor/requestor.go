package requestor

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"./crh"
	"./models"
)

// TCP = 0
// UDP = 1
// RabbitMQ = 2
const connType = 1

// Requestor docstring.
type Requestor struct {
	handler crh.ClientRequestHandler
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Invoke docstring.
func (r *Requestor) Invoke(op *models.Operation) models.Response {
	var res models.Response

	// Serialização
	msg, err := json.Marshal(op)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 404}
	}
	// Create Structures to connect to server
	switch connType {
	case 0:
		r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 6969}
		r.handler.TCPConnection, err = net.Dial("tcp", r.handler.Host+":"+strconv.Itoa(r.handler.Port))
		failOnError(err, "Couldn't Create TCP'Connection")
	case 1:
		r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 1111}
		ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1111")
		LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
		r.handler.UDPConnection, _ = net.DialUDP("udp", LocalAddr, ServerAddr)
	case 2:
		r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 5672} // RabbitMQ
	}
	var aux []byte
	//	Send and Receive
	switch connType {
	case 0:
		start := time.Now()
		r.handler.SendTCP(msg)
		fmt.Println("Sent" + string(msg))
		aux = r.handler.ReceiveTCP()
		elapsed := time.Since(start)
		fmt.Printf("%.0f\n", float64(elapsed)/float64(time.Millisecond))
	case 1:
		start := time.Now()
		r.handler.SendUDP(msg)
		aux = r.handler.ReceiveUDP()
		elapsed := time.Since(start)
		fmt.Printf("%.0f\n", float64(elapsed)/float64(time.Millisecond))
	case 2:
		start := time.Now()
		r.handler.SendMiddleware(msg)
		aux = r.handler.ReceiveMiddleware()
		elapsed := time.Since(start)
		fmt.Printf("%.0f\n", float64(elapsed)/float64(time.Millisecond))
	}
	/*// Begin time evaluation
	start := time.Now()
	r.handler.SendMiddleware(msg)
	aux := r.handler.ReceiveMiddleware()
	elapsed := time.Since(start)
	fmt.Printf("%.0f\n", float64(elapsed)/float64(time.Millisecond))
	// End time evaluation
	*/
	// Deserialização
	err = json.Unmarshal(aux, &res)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 500}
	}

	return res
}
