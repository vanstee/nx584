// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SetUserAuthorizationCommandWithPinName              = "Set User Authorization Command with PIN"
	SetUserAuthorizationCommandWithPinLength            = 7
	SetUserAuthorizationCommandWithPinNumber            = 54
	SetUserAuthorizationCommandWithPinTransitionCapable = message.TransitionCapableNo
	SetUserAuthorizationCommandWithPinAcknowledged      = message.AcknowledgedYes
)

type SetUserAuthorizationCommandWithPin struct {
	*message.Base
}

func NewSetUserAuthorizationCommandWithPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SetUserAuthorizationCommandWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetUserAuthorizationCommandWithPinLength, length)
	}

	message := &SetUserAuthorizationCommandWithPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
