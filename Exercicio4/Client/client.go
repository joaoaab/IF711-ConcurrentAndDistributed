package main

import (
	"fmt"

	"./algorithm"
)

// Main function
func main() {
	c := new(algorithm.Calculator)

	fmt.Printf("%d", c.Fib(10))
}
