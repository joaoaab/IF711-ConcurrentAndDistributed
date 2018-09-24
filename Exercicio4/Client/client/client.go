package main

import (
	"fmt"
)

// IntQueue test.
type IntQueue struct {
	elements []int
}

func (iq *IntQueue) enqueue(x int) {
	iq.elements = append(iq.elements, x)
}

func (iq *IntQueue) dequeue(x int) int {
	r := iq.elements[0]
	iq.elements = iq.elements[1:]
	return r
}

// Operation data to send.
type Operation struct {
	name   string
	params IntQueue
}

// ClientRequestHandler test.
type ClientRequestHandler struct {
	host           string
	port           int
	sentMsgSize    int
	receiveMsgSize int
}

func (crh *ClientRequestHandler) send() {
	fmt.Printf("%s:%d -> %d\n", crh.host, crh.port, crh.sentMsgSize)
}

func (crh *ClientRequestHandler) receive() {
	fmt.Printf("%s:%d <- %d\n", crh.host, crh.port, crh.receiveMsgSize)
}

// Requestor test.
type Requestor struct {
	crh ClientRequestHandler
}

func (r *Requestor) invoke() {
	r.crh = ClientRequestHandler{"localhost", 8080, 256, 512}
	r.crh.send()
	r.crh.receive()
}

// Calculator test.
type Calculator struct{}

func (c *Calculator) fib(x int) int {
	op := new(Operation)

	op.name = "fib"
	op.params.enqueue(x)

	return x
}

func (c *Calculator) mdc(x int, y int) int {
	return 3
}

func (c *Calculator) mmc(x int, y int) int {
	return 2
}

func (c *Calculator) pow(x int, y int) int {
	return 1
}

// Main function
func main() {
	r := new(Requestor)
	r.invoke()
}
