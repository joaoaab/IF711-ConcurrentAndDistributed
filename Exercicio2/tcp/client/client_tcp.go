package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func sendMessage(out chan string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Digite para enviar:")
		sendMessage, _ := reader.ReadString('\n')
		fmt.Println("digitado : " + sendMessage)
		out <- fmt.Sprintf("%s", sendMessage[:len(sendMessage)-1])
	}
}

func receiveMessage(server net.Conn, in chan string) {
	reader := bufio.NewReader(server)
	for {
		message, _ := reader.ReadString('\n')
		in <- message

	}
}

func main() {
	//Connect TCP
	conn, err := net.Dial("tcp", "localhost:1337")
	incomingMessages := make(chan string)
	outgoingMessages := make(chan string)
	fmt.Println("--Tentando conexao--")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//efer conn.Close()
	fmt.Println("--conexao feita--")
	go sendMessage(outgoingMessages)
	go receiveMessage(conn, incomingMessages)
	for {
		select {
		case text := <-incomingMessages:
			fmt.Println("Server -> " + text)

		case text := <-outgoingMessages:
			fmt.Printf("Enviando : %s \n", text)
			conn.Write([]byte(text + "\n"))
		}
	}

}
