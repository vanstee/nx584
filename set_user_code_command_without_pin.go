// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	SetUserCodeCommandWithoutPinName              = "Set User Code Command without PIN"
	SetUserCodeCommandWithoutPinLength            = 5
	SetUserCodeCommandWithoutPinNumber            = 53
	SetUserCodeCommandWithoutPinTransitionCapable = TransitionCapableNo
	SetUserCodeCommandWithoutPinAcknowledged      = AcknowledgedYes
)

type SetUserCodeCommandWithoutPin struct {
	*BaseMessage
}

func NewSetUserCodeCommandWithoutPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SetUserCodeCommandWithoutPinLength {
		log.Printf("message length incorrect: expected %d, actual: %d", SetUserCodeCommandWithoutPinLength, length)
	}

	message := &SetUserCodeCommandWithoutPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SetUserCodeCommandWithoutPin) Number() byte {
	return SetUserCodeCommandWithoutPinNumber
}

func (m *SetUserCodeCommandWithoutPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Set User Code Command without PIN",
		strings.Repeat("-", len("Set User Code Command without PIN")),
	)
}
