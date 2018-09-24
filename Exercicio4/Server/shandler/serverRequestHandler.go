package shandler

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// Messages Channel
var Messages = make(chan Message)

// Reply Channel
var Reply = make(chan Message)

// Message type
type Message struct {
	Data     string
	Addr     net.Addr
	Protocol int
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func acceptTCPConnections(server net.Listener, newConnections chan net.Conn) {
	for {
		c, err := server.Accept()
		if err != nil {
			fmt.Println("Error Acception:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Separando a thread para comunicação")
		newConnections <- c
	}
}

func handleTCPConnection(conn net.Conn, TCPMessages chan Message, deadTCPConnections chan net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		m := Message{message, conn.RemoteAddr(), 0}
		//fmt.Println("li do socket " + m.Data)
		TCPMessages <- m
	}
	deadTCPConnections <- conn
}

func handleUDPMessages(conn *net.UDPConn, allClients map[int]int, udpMessages chan Message) {
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf) // buf[0:n]
		if err != nil {
			continue
		}
		m := Message{string(buf[0:n]), addr, 1}
		udpMessages <- m
	}
}

// HandleTCP handles tcp connections for the server
func HandleTCP() {
	TCPConnections := make(map[net.Addr]net.Conn)
	newTCPConnections := make(chan net.Conn)
	deadTCPConnections := make(chan net.Conn)
	TCPMessages := make(chan Message)
	s, err := net.Listen("tcp", "localhost:6900")
	checkError(err)
	go acceptTCPConnections(s, newTCPConnections)
	defer s.Close()

	for {
		select {
		case conn := <-newTCPConnections:
			TCPConnections[conn.RemoteAddr()] = conn
			go handleTCPConnection(conn, TCPMessages, deadTCPConnections)
			fmt.Println(conn.RemoteAddr())
		case conn := <-deadTCPConnections:
			delete(TCPConnections, conn.RemoteAddr())
		case msg := <-TCPMessages:
			Messages <- msg
		case ret := <-Reply:
			address := ret.Addr
			data := ret.Data
			TCPConnections[address].Write([]byte(data))
		}
	}
}

// HandleUDP handles udp connections for the server
func HandleUDP() {
	// Setting up UDP
	udpClients := make(map[int]int)
	udpMessages := make(chan Message)
	serverAddr, err := net.ResolveUDPAddr("udp", ":1111")
	checkError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	checkError(err)
	defer serverConn.Close()

	go handleUDPMessages(serverConn, udpClients, udpMessages)

	for {
		select {
		case msg := <-udpMessages:
			Messages <- msg
		case ret := <-Reply:
			address := ret.Addr
			data := ret.Data
			fullAddr, err := net.ResolveUDPAddr("udp", address.String())
			if err != nil {
				fmt.Println("Resolver Failed with error : ", err)
				continue
			}
			_, err = serverConn.WriteToUDP([]byte(data), fullAddr)
		}
	}
}
