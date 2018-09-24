package algorithm

import (
	"./requestor"
	"./requestor/models"
)

// Calculator docstring.
type Calculator struct{}

// Fib docstring.
func (c *Calculator) Fib(x int) int {
	requestor := new(requestor.Requestor)
	op := new(models.Operation)

	op.SetName("fib")
	op.AddParam(x)

	res := requestor.Invoke(op)

	return res.GetResult()
}

// Mdc docstring.
func (c *Calculator) Mdc(x int, y int) int {
	requestor := new(requestor.Requestor)
	op := new(models.Operation)

	op.SetName("mdc")
	op.AddParam(x)
	op.AddParam(y)

	res := requestor.Invoke(op)

	return res.GetResult()
}

// Mmc docstring.
func (c *Calculator) Mmc(x int, y int) int {
	requestor := new(requestor.Requestor)
	op := new(models.Operation)

	op.SetName("mmc")
	op.AddParam(x)
	op.AddParam(y)

	res := requestor.Invoke(op)

	return res.GetResult()
}

// Pow docstring.
func (c *Calculator) Pow(x int, y int) int {
	requestor := new(requestor.Requestor)
	op := new(models.Operation)

	op.SetName("pow")
	op.AddParam(x)
	op.AddParam(y)

	res := requestor.Invoke(op)

	return res.GetResult()
}
