package crh

import (
	"fmt"
)

// ClientRequestHandler test.
type ClientRequestHandler struct {
	Host           string
	Port           int
	SentMsgSize    int
	ReceiveMsgSize int
}

// Send test.
func (crh *ClientRequestHandler) Send() {
	fmt.Printf("%s:%d -> %d\n", crh.Host, crh.Port, crh.SentMsgSize)
}

// Receive test.
func (crh *ClientRequestHandler) Receive() {
	fmt.Printf("%s:%d <- %d\n", crh.Host, crh.Port, crh.ReceiveMsgSize)
}
