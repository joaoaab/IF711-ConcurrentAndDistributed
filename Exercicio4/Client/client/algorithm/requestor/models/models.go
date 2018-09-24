package models

import "fmt"

// Operation data to send.
type Operation struct {
	Name   string
	Params []int
}

// AddParam to Operation.
func (op *Operation) AddParam(x int) {
	op.Params = append(op.Params, x)
}

// GetParam and remove it from Operation.
func (op *Operation) GetParam() int {
	r := op.Params[0]
	op.Params = op.Params[1:]
	return r
}

// SetName test.
func (op *Operation) SetName(name string) {
	op.Name = name
}

// Print test.
func (op *Operation) Print() {
	fmt.Println(op.Name)
	for _, param := range op.Params {
		fmt.Println(param)
	}
}
