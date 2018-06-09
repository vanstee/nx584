// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SendKeypadTextMessageName              = "Send Keypad Text Message"
	SendKeypadTextMessageLength            = 12
	SendKeypadTextMessageNumber            = 43
	SendKeypadTextMessageTransitionCapable = message.TransitionCapableNo
	SendKeypadTextMessageAcknowledged      = message.AcknowledgedYes
)

type SendKeypadTextMessage struct {
	*message.Base
}

func NewSendKeypadTextMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SendKeypadTextMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SendKeypadTextMessageLength, length)
	}

	message := &SendKeypadTextMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
