package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Connect TCP
	conn, err := net.Dial("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	//simple Read
	buffer := make([]byte, 1024)
	conn.Read(buffer)

	//simple write
	conn.Write([]byte("Hello from client"))
}
