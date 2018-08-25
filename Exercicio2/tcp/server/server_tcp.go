package main
g:syntastic_quiet_messages
import (
	"fmt"
	"net"
	"os"
)

func main() {
	// listen to incoming tcp connections
	l, err := net.Listen("tcp", "localhost:1313")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer l.Close()
	// A common pattern is to start a loop to continously accept connections
	for {
		//accept connections using Listener.Accept()
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//It's common to handle accepted connection on different goroutines
		go handleConnection(c)
	}
}

func handleConnection(conn net.Conn) {

}
