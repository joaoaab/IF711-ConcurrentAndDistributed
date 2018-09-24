package models

// Operation data to send.
type Operation struct {
	name   string
	params []int
}

// AddParam to Operation.
func (op *Operation) AddParam(x int) {
	op.params = append(op.params, x)
}

// GetParam and remove it from Operation.
func (op *Operation) GetParam() int {
	r := op.params[0]
	op.params = op.params[1:]
	return r
}

// SetName test.
func (op *Operation) SetName(name string) {
	op.name = name
}
