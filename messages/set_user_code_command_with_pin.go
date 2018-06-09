// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SetUserCodeCommandWithPinName              = "Set User Code Command with PIN"
	SetUserCodeCommandWithPinLength            = 8
	SetUserCodeCommandWithPinNumber            = 52
	SetUserCodeCommandWithPinTransitionCapable = message.TransitionCapableNo
	SetUserCodeCommandWithPinAcknowledged      = message.AcknowledgedYes
)

type SetUserCodeCommandWithPin struct {
	*message.Base
}

func NewSetUserCodeCommandWithPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SetUserCodeCommandWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetUserCodeCommandWithPinLength, length)
	}

	message := &SetUserCodeCommandWithPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *SetUserCodeCommandWithPin) Number() byte {
	return SetUserCodeCommandWithPinNumber
}

func (m *SetUserCodeCommandWithPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Set User Code Command with PIN",
		strings.Repeat("-", len("Set User Code Command with PIN")),
	)
}
