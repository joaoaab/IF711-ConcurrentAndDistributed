package requestor

import (
	"./crh"
)

// Requestor test.
type Requestor struct {
	handler crh.ClientRequestHandler
}

// Invoke test.
func (r *Requestor) Invoke() {
	r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 8080}
	r.handler.Send()
	r.handler.Receive()
}
