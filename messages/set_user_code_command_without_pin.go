// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SetUserCodeCommandWithoutPinName              = "Set User Code Command without PIN"
	SetUserCodeCommandWithoutPinLength            = 5
	SetUserCodeCommandWithoutPinNumber            = 53
	SetUserCodeCommandWithoutPinTransitionCapable = message.TransitionCapableNo
	SetUserCodeCommandWithoutPinAcknowledged      = message.AcknowledgedYes
)

type SetUserCodeCommandWithoutPin struct {
	*message.Base
}

func NewSetUserCodeCommandWithoutPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SetUserCodeCommandWithoutPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetUserCodeCommandWithoutPinLength, length)
	}

	message := &SetUserCodeCommandWithoutPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
