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
func (crh *ClientRequestHandler) Send(outcoming []byte) {
	crh.sentMsgSize = len(outcoming)
	fmt.Printf("%s:%d -> %d\n", crh.Host, crh.Port, crh.sentMsgSize)
}

// Receive test.
func (crh *ClientRequestHandler) Receive() []byte {
	var jsonBlob = []byte("{\"Name\": \"Fib\", \"Result\": 45}")

	fmt.Printf("%s:%d <- %d\n", crh.Host, crh.Port, crh.receiveMsgSize)

	return jsonBlob
}
