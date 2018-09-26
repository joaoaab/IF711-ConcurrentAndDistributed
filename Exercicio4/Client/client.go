package main

import (
	"fmt"
	"os"
	"time"

	"./algorithm"
)

// Main function
func main() {
	c := new(algorithm.Calculator)
	err := c.Setup()
	if err != nil {
		os.Exit(1)
	}
	defer c.Close()

	for i := 0; i < 10; i++ {
		fmt.Printf("Result: %d\n", c.Pow(2, 31))
		time.Sleep(10 * time.Second) // Tempo pra respirar
	}
}
