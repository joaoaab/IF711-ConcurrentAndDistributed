package clientrequesthandler

import (
	"fmt"
)

// ClientRequestHandler test.
type ClientRequestHandler struct {
	host           string
	port           int
	sentMsgSize    int
	receiveMsgSize int
}

func (crh *ClientRequestHandler) send() {
	fmt.Printf("%s:%d -> %d\n", crh.host, crh.port, crh.sentMsgSize)
}

func (crh *ClientRequestHandler) receive() {
	fmt.Printf("%s:%d <- %d\n", crh.host, crh.port, crh.receiveMsgSize)
}
