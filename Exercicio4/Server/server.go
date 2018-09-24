package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"./models"
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
	for i := 1; i < n; i++ {
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

//Invoke Invokes the calculations and return the json of the answer
func Invoke(data string) models.Operation {
	var res models.Operation
	err := json.Unmarshal([]byte(data), &res)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
	return res
}

func sendAnswer(msg shandler.Message, frame models.Operation, ans int) {
	var ret shandler.Message
	ret.Protocol = msg.Protocol
	ret.Addr = msg.Addr
	var pack models.Response
	pack.Name = frame.Name
	pack.Result = ans
	data, err := json.Marshal(pack)
	ret.Data = string(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(11)
	}
	println(ret.Data)
	shandler.Reply <- ret
}

func main() {
	go shandler.Handle()
	for {
		select {
		case msg := <-shandler.Messages:
			fmt.Println(msg.Data + " protocol : " + strconv.Itoa(msg.Protocol))
			frame := Invoke(msg.Data)
			if frame.GetName() == "fib" {
				ans := fib(frame.GetParam())
				sendAnswer(msg, frame, ans)
			}
		}
	}
}
