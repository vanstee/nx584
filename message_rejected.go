// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	MessageRejectedName              = "Message Rejected"
	MessageRejectedLength            = 1
	MessageRejectedNumber            = 31
	MessageRejectedTransitionCapable = TransitionCapableNo
	MessageRejectedAcknowledged      = AcknowledgedNo
)

type MessageRejected struct {
	*BaseMessage
}

func NewMessageRejected(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != MessageRejectedLength {
		log.Printf("message length incorrect: expected %d, actual: %d", MessageRejectedLength, length)
	}

	message := &MessageRejected{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
