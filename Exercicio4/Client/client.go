package main

import (
	"fmt"

	"./algorithm"
)

// Main function
func main() {
	c := new(algorithm.Calculator)
	for i := 0; i < 1000; i++ {
		fmt.Println(c.Pow(2, 31))
	}
}
