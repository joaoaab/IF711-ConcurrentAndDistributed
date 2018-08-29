package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

/*CheckError A Simple function to verify error*/
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func sendMessage(messages chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Digite uma mensagem para enviar : ")
		sendMessage, _ := reader.ReadString('\n')
		messages <- fmt.Sprintf("%s", sendMessage[:len(sendMessage)-1])
	}
}

func receiveMessage(conn *net.UDPConn, messages chan string) {
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		CheckError(err)
		messages <- string(buf[0:n])
	}
}

func main() {
	incomingMessages := make(chan string)
	outgoingMessages := make(chan string)
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	start := []byte("starting")
	Conn.Write(start)
	CheckError(err)

	go sendMessage(outgoingMessages)
	go receiveMessage(Conn, incomingMessages)
	defer Conn.Close()
	for {
		select {
		case message := <-incomingMessages:
			fmt.Println("Server -> " + message)

		case message := <-outgoingMessages:
			buf := []byte(message)
			Conn.Write(buf)
		}
	}
}
