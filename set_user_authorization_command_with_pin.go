// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	SetUserAuthorizationCommandWithPinName              = "Set User Authorization Command with PIN"
	SetUserAuthorizationCommandWithPinLength            = 7
	SetUserAuthorizationCommandWithPinNumber            = 54
	SetUserAuthorizationCommandWithPinTransitionCapable = TransitionCapableNo
	SetUserAuthorizationCommandWithPinAcknowledged      = AcknowledgedYes
)

type SetUserAuthorizationCommandWithPin struct {
	*BaseMessage
}

func NewSetUserAuthorizationCommandWithPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SetUserAuthorizationCommandWithPinLength {
		log.Printf("message length incorrect: expected %d, actual: %d", SetUserAuthorizationCommandWithPinLength, length)
	}

	message := &SetUserAuthorizationCommandWithPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SetUserAuthorizationCommandWithPin) Number() byte {
	return SetUserAuthorizationCommandWithPinNumber
}

func (m *SetUserAuthorizationCommandWithPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Set User Authorization Command with PIN",
		strings.Repeat("-", len("Set User Authorization Command with PIN")),
	)
}
