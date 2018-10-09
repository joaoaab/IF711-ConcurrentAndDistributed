package main

import (
	"context"
	"log"
	"net"

	"../protocol"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Calc(ctx context.Context, in *fib.Number) (*fib.Result, error) {
	n := in.Value
	if n == 0 {
		return &fib.Result{Value: int64(0)}, nil
	} else if n == 1 {
		return &fib.Result{Value: int64(1)}, nil
	}
	a := 0
	b := 1
	c := 1
	for i := 1; i < int(n); i++ {
		c = a + b
		a = b
		b = c
	}
	return &fib.Result{Value: int64(c)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("failed %v", err)
	}
	grpcServer := grpc.NewServer()
	fib.RegisterCalculatorServer(grpcServer, &server{})

	reflection.Register(grpcServer)

	log.Println("Server is listening on port 50000")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("ih rapaz, deu ruim, meldels %v", err)
	}
}
