package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func acceptConnections(server net.Listener, newConnections chan net.Conn) {
	for {
		c, err := server.Accept()
		if err != nil {
			fmt.Println("Error Acception:", err.Error)
			os.Exit(1)
		}
		fmt.Println("Separando a thread para comunicação")
		newConnections <- c
	}
}

func handleConnection(conn net.Conn, messages chan string, deadConnections chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		messages <- fmt.Sprintf("Messagem Recebida : %s", message)
	}
	deadConnections <- conn
}

func main() {
	nClients := 0
	allClients := make(map[net.Conn]int)
	newConnections := make(chan net.Conn)
	deadConnections := make(chan net.Conn)
	messages := make(chan string)
	// listen to incoming tcp connections
	l, err := net.Listen("tcp", "localhost:1337")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	// A common pattern is to start a loop to continously accept connections
	go acceptConnections(l, newConnections)
	for {
		select {
		case conn := <-newConnections:
			log.Printf("Aceitei Novo Cliente")
			allClients[conn] = nClients
			nClients++
			go handleConnection(conn, messages, deadConnections)

		case message := <-messages:
			for conn, _ := range allClients {
				go func(conn net.Conn, message string) {
					_, err := conn.Write([]byte(message))
					if err != nil {
						deadConnections <- conn
					}
				}(conn, message)
			}
			fmt.Println(message)
			fmt.Println("Transmitindo para todos os clientes.")

		case conn := <-deadConnections:
			fmt.Println("Cliente Desconectado")
			delete(allClients, conn)
		}
	}
}
