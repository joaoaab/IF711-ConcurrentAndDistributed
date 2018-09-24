package main

import (
	"fmt"

	"./shandler"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a := 0
	b := 1
	c := 1
	for i := 2; i < n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func mdc(a, b int) int {
	if b == 0 {
		return a
	}
	return mdc(b, a%b)
}

func mmc(a, b int) int {
	return a * b / mdc(a, b)
}

func pow(base, exponent int) int {
	if base < 0 {
		base *= -1
	}
	for exponent > 0 {
		if exponent%2 == 1 {
			base *= base
		}
		exponent >>= 1
		base *= base * base
	}
	return base
}

func main() {
	go shandler.Handle()
	for {
		select {
		case msg := <-shandler.Messages:
			fmt.Println(msg.Data)
		}
	}
}
