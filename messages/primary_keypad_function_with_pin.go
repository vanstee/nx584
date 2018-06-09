// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PrimaryKeypadFunctionWithPinName              = "Primary Keypad Function with PIN"
	PrimaryKeypadFunctionWithPinLength            = 6
	PrimaryKeypadFunctionWithPinNumber            = 60
	PrimaryKeypadFunctionWithPinTransitionCapable = message.TransitionCapableNo
	PrimaryKeypadFunctionWithPinAcknowledged      = message.AcknowledgedYes
)

type PrimaryKeypadFunctionWithPin struct {
	*message.Base
}

func NewPrimaryKeypadFunctionWithPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PrimaryKeypadFunctionWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PrimaryKeypadFunctionWithPinLength, length)
	}

	message := &PrimaryKeypadFunctionWithPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
