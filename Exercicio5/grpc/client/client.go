package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"../protocol"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to create connection %v", err)
	}

	defer conn.Close()

	c := fib.NewCalculatorClient(conn)
	for i := 0; i < 1000; i++ {
		start := time.Now()
		_, err := c.Calc(context.Background(), &fib.Number{Value: 15})
		if err != nil {
			log.Fatalf("Ih rapaz, deu ruim %v", err)
		}
		elapsed := time.Since(start)
		fmt.Printf("%.2f\n", float64(elapsed)/float64(time.Millisecond))
	}

}
