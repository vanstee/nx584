// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	MessageRejectedName              = "Message Rejected"
	MessageRejectedLength            = 1
	MessageRejectedNumber            = 31
	MessageRejectedTransitionCapable = message.TransitionCapableNo
	MessageRejectedAcknowledged      = message.AcknowledgedNo
)

type MessageRejected struct {
	*message.Base
}

func NewMessageRejected(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != MessageRejectedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", MessageRejectedLength, length)
	}

	message := &MessageRejected{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *MessageRejected) Number() byte {
	return MessageRejectedNumber
}

func (m *MessageRejected) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Message Rejected",
		strings.Repeat("-", len("Message Rejected")),
	)
}
