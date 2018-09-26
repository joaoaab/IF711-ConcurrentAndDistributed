package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"./calculator"
)

// Args docstring.
type Args struct {
	A, B int
}

// Arith docstring.
type Arith int

// Fib docstrig.
func (t *Arith) Fib(args *Args, reply *int) error {
	*reply = calculator.Fib(args.A)
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	forever := make(chan bool)
	<-forever
}
