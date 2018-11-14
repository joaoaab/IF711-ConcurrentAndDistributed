package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type packet struct {
	TYPE int
	ID   int
	PORT int
}

type address struct {
	IP   string
	PORT int
}

var hash = make(map[int]address)

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

func handleConnection(conn net.Conn) {
	reader := json.NewDecoder(conn)
	var pkg packet
	for {
		err := reader.Decode(&pkg)
		if err != nil {
			fmt.Print("k")
			break
		}
		fmt.Println(pkg)
		fmt.Println(conn.RemoteAddr())
		if pkg.TYPE == 1 {
			ip := conn.RemoteAddr().String()
			port := pkg.PORT
			robj := address{IP: ip, PORT: port}
			fmt.Println(robj)
			hash[pkg.ID] = robj
		} else {
			addr := hash[pkg.ID]
			msg, _ := json.Marshal(addr)
			conn.Write(msg)
		}

	}
}

func main() {
	newConnections := make(chan net.Conn)
	l, err := net.Listen("tcp", "0.0.0.0:1337")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	go acceptConnections(l, newConnections)
	for {
		select {
		case conn := <-newConnections:
			log.Printf("New client connected")
			go handleConnection(conn)
		}
	}

}
