package requestor

import (
	"encoding/json"
	"fmt"

	"./crh"
	"./models"
)

// Requestor test.
type Requestor struct {
	handler crh.ClientRequestHandler
}

// Invoke test.
func (r *Requestor) Invoke(op *models.Operation) models.Response {
	r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 5672} // RabbitMQ
	var res models.Response

	msg, err := json.Marshal(op)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 404}
	}

	r.handler.Send(msg)
	err = json.Unmarshal(r.handler.Receive(), &res)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 500}
	}

	return res
}
