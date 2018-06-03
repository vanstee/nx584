// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	SendKeypadTextMessageName              = "Send Keypad Text Message"
	SendKeypadTextMessageLength            = 12
	SendKeypadTextMessageNumber            = 43
	SendKeypadTextMessageTransitionCapable = TransitionCapableNo
	SendKeypadTextMessageAcknowledged      = AcknowledgedYes
)

type SendKeypadTextMessage struct {
	*BaseMessage
}

func NewSendKeypadTextMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SendKeypadTextMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", SendKeypadTextMessageLength, length)
	}

	message := &SendKeypadTextMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SendKeypadTextMessage) Number() byte {
	return SendKeypadTextMessageNumber
}

func (m *SendKeypadTextMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Send Keypad Text Message",
		strings.Repeat("-", len("Send Keypad Text Message")),
	)
}
