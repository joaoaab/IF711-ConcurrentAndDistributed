package algorithm

import (
	"./requestor"
	"./requestor/models"
)

// Calculator test.
type Calculator struct{}

// Fib test.
func (c *Calculator) Fib(x int) int {
	requestor := new(requestor.Requestor)
	op := new(models.Operation)

	op.SetName("fib")
	op.AddParam(x)

	res := requestor.Invoke(op)

	return res.GetResult()
}

// Mdc test.
func (c *Calculator) Mdc(x int, y int) int {
	op := new(models.Operation)

	op.SetName("mdc")
	op.AddParam(x)
	op.AddParam(y)

	return x
}

// Mmc test.
func (c *Calculator) Mmc(x int, y int) int {
	op := new(models.Operation)

	op.SetName("mmc")
	op.AddParam(x)
	op.AddParam(y)

	return x
}

// Pow test.
func (c *Calculator) Pow(x int, y int) int {
	op := new(models.Operation)

	op.SetName("pow")
	op.AddParam(x)
	op.AddParam(y)

	return x
}
