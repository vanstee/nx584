// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	X10MessageReceivedName              = "X-10 Message Received"
	X10MessageReceivedLength            = 4
	X10MessageReceivedNumber            = 9
	X10MessageReceivedTransitionCapable = TransitionCapableYes
	X10MessageReceivedAcknowledged      = AcknowledgedYes
)

type X10MessageReceived struct {
	*BaseMessage
}

func NewX10MessageReceived(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != X10MessageReceivedLength {
		log.Printf("message length incorrect: expected %d, actual: %d", X10MessageReceivedLength, length)
	}

	message := &X10MessageReceived{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *X10MessageReceived) Number() byte {
	return X10MessageReceivedNumber
}

func (m *X10MessageReceived) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"X-10 Message Received",
		strings.Repeat("-", len("X-10 Message Received")),
	)
}
