package crh

import (
	"fmt"
)

// ClientRequestHandler test.
type ClientRequestHandler struct {
	Host           string
	Port           int
	sentMsgSize    int
	receiveMsgSize int
}

// Send test.
func (crh *ClientRequestHandler) Send() {
	fmt.Printf("%s:%d -> %d\n", crh.Host, crh.Port, crh.sentMsgSize)
}

// Receive test.
func (crh *ClientRequestHandler) Receive() {
	fmt.Printf("%s:%d <- %d\n", crh.Host, crh.Port, crh.receiveMsgSize)
}
