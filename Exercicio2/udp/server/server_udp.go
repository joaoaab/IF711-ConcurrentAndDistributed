package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

/*CheckError A Simple function to verify error*/
func CheckError(err error) {

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	nClients := 0
	allClients := make(map[int]int)

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
		n, addr, err := ServerConn.ReadFromUDP(buf) // buf[0:n]
		if err != nil {
			continue
		}

		if _, ok := allClients[addr.Port]; !ok {
			allClients[addr.Port] = nClients
			nClients++
			fmt.Println(addr, " has just connected.")
		} else {
			fmt.Println("Received: ", string(buf[0:n]))
		}

		for port, _ := range allClients {
			if addr.Port != port {
				FullAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:"+strconv.Itoa(port))
				if err != nil {
					fmt.Println("Resolver failed with error: ", err)
					continue
				}
				_, errr := ServerConn.WriteToUDP(buf[0:n], FullAddr)
				if errr != nil {
					fmt.Println("WrtieToUDP failed with error: ", errr)
					continue
				}
			}
		}

	}
}
