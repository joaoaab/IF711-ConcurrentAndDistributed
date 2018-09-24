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

// SetName docstring.
func (op *Operation) SetName(name string) {
	op.Name = name
}

// GetName docstrong
func (op *Operation) GetName() string {
	return op.Name
}

// Print docstring.
func (op *Operation) Print() {
	fmt.Println(op.Name)
	for _, param := range op.Params {
		fmt.Println(param)
	}
}

// Response docstring.
type Response struct {
	Name   string
	Result int
}

// GetResult docstring.
func (r *Response) GetResult() int {
	return r.Result
}

// GetName docstring
func (r *Response) GetName() string {
	return r.Name
}

// SetName docstring
func (r *Response) SetName(name string) {
	r.Name = name
}

// SetResult docstring
func (r *Response) SetResult(result int) {
	r.Result = result
}

// Print for test purposes.
func (r *Response) Print() {
	fmt.Println(r.Name, r.GetResult())
}
