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
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newConnections <- conn
	}
}

func handleConnections(conn net.Conn, clientID int, messages chan string, deadConnections chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		incoming, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		messages <- fmt.Sprintf("Client %d > %s", clientID, incoming)

	}
	// if there's an error reading from the buffer
	// that means there's a lost connection
	deadConnections <- conn
}

func main() {
	numberOfClients := 0
	allClients := make(map[net.Conn]int)
	newConnections := make(chan net.Conn)
	deadConnections := make(chan net.Conn)
	messages := make(chan string)

	// Listen to connections
	server, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Goroutine to listen to accept connections and send to the main routine
	// on the newConnections channel
	go acceptConnections(server, newConnections)

	// main GoRoutine
	for {
		select {
		// If a new connection is signaled on the channel
		case conn := <-newConnections:
			log.Printf("Accepted new client")
			allClients[conn] = numberOfClients
			numberOfClients++
			// New routine to read from the client
			go handleConnections(conn, allClients[conn], messages, deadConnections)
		// If there's a new message on the channel
		case message := <-messages:
			// New routine to broadcast to all clients
			for conn, _ := range allClients {
				go func(conn net.Conn, message string) {
					_, err := conn.Write([]byte(message))

					if err != nil {
						deadConnections <- conn
					}
				}(conn, message)
			}
			log.Printf("New Message: %s ", message)
			log.Printf("Broadcast to all %d Clients", len(allClients))
		// if a connection dropped on the channel
		case conn := <-deadConnections:
			log.Printf("Client %d disconnected", allClients[conn])
			delete(allClients, conn)
		}
	}
}
