package algorithm

import (
	"./requestor"
	"./requestor/models"
)

// Calculator docstring.
type Calculator struct {
	requestor *requestor.Requestor
}

// Setup for Calculator.
func (c *Calculator) Setup() error {
	if c.requestor == nil {
		c.requestor = new(requestor.Requestor)
		return c.requestor.Setup(0) // 0=tcp, 1=udp, 2=rabbitmq.rpc
	}
	return nil
}

// Close Calculator.
func (c *Calculator) Close() {
	c.requestor.Close()
}

// Fib docstring.
func (c *Calculator) Fib(x int) int {
	op := new(models.Operation)

	op.SetName("fib")
	op.AddParam(x)

	res := c.requestor.Invoke(op)

	return res.GetResult()
}

// Mdc docstring.
func (c *Calculator) Mdc(x int, y int) int {
	op := new(models.Operation)

	op.SetName("mdc")
	op.AddParam(x)
	op.AddParam(y)

	res := c.requestor.Invoke(op)

	return res.GetResult()
}

// Mmc docstring.
func (c *Calculator) Mmc(x int, y int) int {
	op := new(models.Operation)

	op.SetName("mmc")
	op.AddParam(x)
	op.AddParam(y)

	res := c.requestor.Invoke(op)

	return res.GetResult()
}

// Pow docstring.
func (c *Calculator) Pow(x int, y int) int {
	op := new(models.Operation)

	op.SetName("pow")
	op.AddParam(x)
	op.AddParam(y)

	res := c.requestor.Invoke(op)

	return res.GetResult()
}
