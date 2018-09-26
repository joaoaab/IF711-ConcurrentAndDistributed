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
		r.handler = new(crh.TCPRequestHandler)
		port = 6969
	case 1: // UDP
		r.handler = new(crh.UDPRequestHandler)
		port = 1111
	case 2: // RabbitMQ
		r.handler = new(crh.MiddlewareRequestHandler)
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
	// fmt.Println("Teste1")
	aux := r.handler.Receive()
	// fmt.Println("Teste2")
	elapsed := time.Since(start)
	fmt.Printf("%.3f\n", float64(elapsed)/float64(time.Millisecond))

	// Deserealização
	err = json.Unmarshal(aux, &res)
	if err != nil {
		return models.Response{Name: "err", Result: 500}
	}

	return res
}
