// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	KeypadMessageReceivedName              = "Keypad Message Received"
	KeypadMessageReceivedLength            = 3
	KeypadMessageReceivedNumber            = 11
	KeypadMessageReceivedTransitionCapable = TransitionCapableYes
	KeypadMessageReceivedAcknowledged      = AcknowledgedYes
)

type KeypadMessageReceived struct {
	*BaseMessage
}

func NewKeypadMessageReceived(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != KeypadMessageReceivedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", KeypadMessageReceivedLength, length)
	}

	message := &KeypadMessageReceived{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
