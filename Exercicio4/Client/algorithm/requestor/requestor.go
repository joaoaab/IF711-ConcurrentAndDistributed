package requestor

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"./crh"
	"./models"
)

// Requestor docstring.
type Requestor struct {
	handler crh.ClientRequestHandler
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Setup for Requestor.
func (r *Requestor) Setup(connType int) error {
	var port int
	switch connType {
	case 0: // TCP
		r.handler = new(crh.MiddlewareRequestHandler) // {Host: "localhost", Port: 6969}
		port = 6969
		// r.handler.TCPConnection, err = net.Dial("tcp", r.handler.Host+":"+strconv.Itoa(r.handler.Port))
	case 1: // UDP
		r.handler = new(crh.MiddlewareRequestHandler) // {Host: "localhost", Port: 1111}
		port = 1111
		// ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:1111")
		// LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
		// r.handler.UDPConnection, _ = net.DialUDP("udp", LocalAddr, ServerAddr)
	case 2: // RabbitMQ
		r.handler = new(crh.MiddlewareRequestHandler) // {Host: "localhost", Port: 5672}
		port = 5672
	}

	return r.handler.Setup("localhost", port)
}

// Close Requestor.
func (r *Requestor) Close() {
	r.handler.Close()
}

// Invoke docstring.
func (r *Requestor) Invoke(op *models.Operation) models.Response {
	var res models.Response

	// Serialização
	msg, err := json.Marshal(op)
	if err != nil {
		return models.Response{Name: "err", Result: 404}
	}

	//	Send and Receive
	start := time.Now()
	r.handler.Send(msg)
	aux := r.handler.Receive()
	elapsed := time.Since(start)
	fmt.Printf("%.2f\n", float64(elapsed)/float64(time.Millisecond))

	// Deserealização
	err = json.Unmarshal(aux, &res)
	if err != nil {
		return models.Response{Name: "err", Result: 500}
	}

	return res
}
