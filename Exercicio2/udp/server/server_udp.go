package main

import (
	"fmt"
	"net"
	"os"
)

/*CheckError A Simple function to verify error*/
func CheckError(err error) {

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	/* Lets prepare a address at any address at port 10001*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":10001")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	fmt.Println("Server just started!")
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		ServerConn.WriteToUDP(buf[0:n], addr) // buf[0:n]
		if err != nil {
			continue
		}

	}
}
