package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

// Args docstring.
type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{15, 0}
	var reply int

	started := time.Now()
	err = client.Call("Arith.Fib", args, &reply)
	elapsed := time.Since(started)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: fib[%d]=%d\n", args.A, reply)
	fmt.Printf("Time: %0.3f", float64(elapsed)/float64(time.Millisecond))
}
