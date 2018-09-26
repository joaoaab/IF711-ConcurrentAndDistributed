package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./calculator"
	"./models"
	"./shandler"
)

// 0 for TCP
// 1 for UDP
// 2 for Middleware
const connType = 0

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

func calculate(frame models.Operation) int {
	var ans int
	switch frame.GetName() {
	case "fib":
		ans = calculator.Fib(frame.GetParam())
	case "pow":
		base := frame.GetParam()
		exp := frame.GetParam()
		ans = calculator.Pow(base, exp)
	case "mdc":
		a := frame.GetParam()
		b := frame.GetParam()
		ans = calculator.Mdc(a, b)
	case "mmc":
		a := frame.GetParam()
		b := frame.GetParam()
		ans = calculator.Mmc(a, b)
	}
	return ans
}

func sendAnswer(msg shandler.Message, frame models.Operation, ans int) {
	var ret shandler.Message
	ret.Protocol = msg.Protocol
	ret.Addr = msg.Addr
	pack := new(models.Response)
	pack.SetName(frame.GetName())
	pack.SetResult(ans)
	data, err := json.Marshal(pack)
	ret.Data = string(data) + "\n"
	if err != nil {
		fmt.Println(err)
		os.Exit(11)
	}
	//println("Package Sent : " + ret.Data)
	shandler.Reply <- ret
}

func main() {
	switch connType {
	case 0:
		go shandler.HandleTCP()
	case 1:
		go shandler.HandleUDP()
	case 2:
		go shandler.HandleMiddleware()
	}
	for {
		select {
		case msg := <-shandler.Messages:
			//fmt.Println("Package Received : " + msg.Data)
			frame := Invoke(msg.Data)
			ans := calculate(frame)
			sendAnswer(msg, frame, ans)
		}
	}
}
