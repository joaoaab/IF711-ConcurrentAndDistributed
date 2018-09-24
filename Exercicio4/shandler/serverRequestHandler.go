package shandler

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type TCPMessage struct{
	data string
	addr net.Addr
}

type UDPMessage struct{
	data string
	addr net.UDPAddr
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func acceptTcpConnections(server net.Listener, newConnections chan net.Conn) {
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

func handleTcpConnection(conn net.Conn, tcpMessages chan TCPMessage, deadTcpConnections chan net.Conn) {
	reader := bufio.NewReader()
	for {
		message, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		m := TCPMessage{message, conn.RemoteAddr}
		tcpMessages <- m
	}
	deadTcpConnections <- conn
}

func handleUdpMessages(conn net.UDPConn, allClients map[int]int ,udpMessages chan UDPMessage, *nClients int){
	for {
		n, addr, err := conn.ReadFromUDP(buf) // buf[0:n]
		if err != nil {
			continue
		}

		if _, ok := allClients[addr.Port]; !ok {
			allClients[addr.Port] = nClients
			nClients++
			fmt.Println(addr, " has just connected in UDP.")
		} else {
			fmt.Println("Received:", string(buf[0:n]))
		}
		m := UDPMessage{string(buf[0:n]), conn.RemoteAddr}
		udpMessages <- m

	}
}



func handleConnections() {
	tcpConnections := make(map[net.Conn]int)
	newTcpConnections := make(chan net.Conn)
	deadTcpConnections := make(chan net.Conn)
	tcpMessages := make(chan TCPMessage)

	s, err := net.Listen("tcp", "localhost:6969")
	CheckError(err)
	go acceptTcpConnections(s, newTcpConnections)
	defer s.Close()

	udpClients := make(map[int]int)
	udpMessages := make(chan string)
	serverAddr, err := net.ResolveUDPAddr("udp", ":1111")
	CheckError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()

	buf := make([]byte, 1024)
	go handleUdpMessages(serverConn, udpClients, udpMessages, &nClients)
	
	for{
		select{
			case message := <-tcpMessages:

		}
	}
}