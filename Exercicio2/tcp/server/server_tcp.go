package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// listen to incoming tcp connections
	l, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	// A common pattern is to start a loop to continously accept connections
	for {
		//accept connections using Listener.Accept()
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error Accepting:", err.Error())
			os.Exit(1)
		}
		//It's common to handle accepted connection on different goroutines
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buffer := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buffer)
	reqLen = reqLen + 1 - 1
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}
