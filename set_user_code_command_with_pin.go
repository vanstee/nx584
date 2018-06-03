// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	SetUserCodeCommandWithPinName              = "Set User Code Command with PIN"
	SetUserCodeCommandWithPinLength            = 8
	SetUserCodeCommandWithPinNumber            = 52
	SetUserCodeCommandWithPinTransitionCapable = TransitionCapableNo
	SetUserCodeCommandWithPinAcknowledged      = AcknowledgedYes
)

type SetUserCodeCommandWithPin struct {
	*BaseMessage
}

func NewSetUserCodeCommandWithPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SetUserCodeCommandWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetUserCodeCommandWithPinLength, length)
	}

	message := &SetUserCodeCommandWithPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
