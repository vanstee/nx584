// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	X10MessageReceivedName              = "X-10 Message Received"
	X10MessageReceivedLength            = 4
	X10MessageReceivedNumber            = 9
	X10MessageReceivedTransitionCapable = message.TransitionCapableYes
	X10MessageReceivedAcknowledged      = message.AcknowledgedYes
)

type X10MessageReceived struct {
	*message.Base
}

func NewX10MessageReceived(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != X10MessageReceivedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", X10MessageReceivedLength, length)
	}

	message := &X10MessageReceived{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
