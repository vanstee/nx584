// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	SetUserAuthorizationCommandWithoutPinName              = "Set User Authorization Command without PIN"
	SetUserAuthorizationCommandWithoutPinLength            = 4
	SetUserAuthorizationCommandWithoutPinNumber            = 55
	SetUserAuthorizationCommandWithoutPinTransitionCapable = TransitionCapableNo
	SetUserAuthorizationCommandWithoutPinAcknowledged      = AcknowledgedYes
)

type SetUserAuthorizationCommandWithoutPin struct {
	*BaseMessage
}

func NewSetUserAuthorizationCommandWithoutPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SetUserAuthorizationCommandWithoutPinLength {
		log.Printf("message length incorrect: expected %d, actual: %d", SetUserAuthorizationCommandWithoutPinLength, length)
	}

	message := &SetUserAuthorizationCommandWithoutPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SetUserAuthorizationCommandWithoutPin) Number() byte {
	return SetUserAuthorizationCommandWithoutPinNumber
}

func (m *SetUserAuthorizationCommandWithoutPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Set User Authorization Command without PIN",
		strings.Repeat("-", len("Set User Authorization Command without PIN")),
	)
}
