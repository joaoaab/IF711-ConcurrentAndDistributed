package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Tentando Conexão")
	fmt.Println("Conexão Feita")
	reader := bufio.NewReader(conn)
	for i := 0; i < 50; i++ {
		start := time.Now()
		for j := 0; j < 10000; j++ {
			conn.Write([]byte(string(j) + "\n"))
			reader.ReadString('\n')
		}
		elapsed := time.Since(start)
		fmt.Println(float64(elapsed) / float64(time.Millisecond))
	}
}
