// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SetUserAuthorizationCommandWithoutPinName              = "Set User Authorization Command without PIN"
	SetUserAuthorizationCommandWithoutPinLength            = 4
	SetUserAuthorizationCommandWithoutPinNumber            = 55
	SetUserAuthorizationCommandWithoutPinTransitionCapable = message.TransitionCapableNo
	SetUserAuthorizationCommandWithoutPinAcknowledged      = message.AcknowledgedYes
)

type SetUserAuthorizationCommandWithoutPin struct {
	*message.Base
}

func NewSetUserAuthorizationCommandWithoutPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SetUserAuthorizationCommandWithoutPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetUserAuthorizationCommandWithoutPinLength, length)
	}

	message := &SetUserAuthorizationCommandWithoutPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
