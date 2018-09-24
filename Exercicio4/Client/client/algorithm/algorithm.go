package algorithm

import (
	"encoding/json"
	"fmt"
	"os"

	"./requestor/models"
)

// Calculator test.
type Calculator struct{}

// Fib test.
func (c *Calculator) Fib(x int) int {
	op := new(models.Operation)

	op.SetName("fib")
	op.AddParam(x)

	msg, err := json.Marshal(op)
	if (err) != nil {
		fmt.Println("error:", err)
	} else {
		os.Stdout.Write(msg)
	}

	return x
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
