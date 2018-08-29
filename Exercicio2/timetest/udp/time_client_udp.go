package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	ServerAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	LocalAddr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	Conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)
	Conn.Write([]byte("dale"))
	buf := make([]byte, 1024)
	for i := 0; i < 50; i++ {
		start := time.Now()
		for j := 0; j < 1000; j++ {
			Conn.Write([]byte(string(j) + "\n"))
			Conn.ReadFromUDP(buf)
		}
		elapsed := time.Since(start)
		fmt.Println(float64(elapsed) / float64(time.Millisecond))
	}
}
