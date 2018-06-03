// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	PrimaryKeypadFunctionWithPinName              = "Primary Keypad Function with PIN"
	PrimaryKeypadFunctionWithPinLength            = 6
	PrimaryKeypadFunctionWithPinNumber            = 60
	PrimaryKeypadFunctionWithPinTransitionCapable = TransitionCapableNo
	PrimaryKeypadFunctionWithPinAcknowledged      = AcknowledgedYes
)

type PrimaryKeypadFunctionWithPin struct {
	*BaseMessage
}

func NewPrimaryKeypadFunctionWithPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PrimaryKeypadFunctionWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PrimaryKeypadFunctionWithPinLength, length)
	}

	message := &PrimaryKeypadFunctionWithPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *PrimaryKeypadFunctionWithPin) Number() byte {
	return PrimaryKeypadFunctionWithPinNumber
}

func (m *PrimaryKeypadFunctionWithPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Primary Keypad Function with PIN",
		strings.Repeat("-", len("Primary Keypad Function with PIN")),
	)
}
