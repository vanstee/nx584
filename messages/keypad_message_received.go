// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	KeypadMessageReceivedName              = "Keypad Message Received"
	KeypadMessageReceivedLength            = 3
	KeypadMessageReceivedNumber            = 11
	KeypadMessageReceivedTransitionCapable = message.TransitionCapableYes
	KeypadMessageReceivedAcknowledged      = message.AcknowledgedYes
)

type KeypadMessageReceived struct {
	*message.Base
}

func NewKeypadMessageReceived(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != KeypadMessageReceivedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", KeypadMessageReceivedLength, length)
	}

	message := &KeypadMessageReceived{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *KeypadMessageReceived) Number() byte {
	return KeypadMessageReceivedNumber
}

func (m *KeypadMessageReceived) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Keypad Message Received",
		strings.Repeat("-", len("Keypad Message Received")),
	)
}
