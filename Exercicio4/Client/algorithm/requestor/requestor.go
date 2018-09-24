package requestor

import (
	"encoding/json"
	"fmt"

	"./crh"
	"./models"
)

// Requestor docstring.
type Requestor struct {
	handler crh.ClientRequestHandler
}

// Invoke docstring.
func (r *Requestor) Invoke(op *models.Operation) models.Response {
	r.handler = crh.ClientRequestHandler{Host: "localhost", Port: 5672} // RabbitMQ
	var res models.Response

	// Serialização
	msg, err := json.Marshal(op)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 404}
	}

	// Begin time evaluation
	r.handler.Send(msg)
	aux := r.handler.Receive()
	// End time evaluation

	// Deserialização
	err = json.Unmarshal(aux, &res)
	if err != nil {
		fmt.Println(err)
		return models.Response{Name: "err", Result: 500}
	}

	return res
}
