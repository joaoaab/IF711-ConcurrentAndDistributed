package main

import (
	"./algorithm"
)

// Main function
func main() {
	c := new(algorithm.Calculator)
	for i := 0; i < 1000; i++ {
		c.Fib(15)
	}
}
