package main

import (
	"fmt"

	"./algorithm"
)

// Main function
func main() {
	c := new(algorithm.Calculator)
	fmt.Println(c.Fib(25))
	fmt.Println(c.Pow(2, 8))
	fmt.Println(c.Mdc(4, 6))
	fmt.Println(c.Mmc(4, 2))
}
