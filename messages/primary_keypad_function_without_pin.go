// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PrimaryKeypadFunctionWithoutPinName              = "Primary Keypad Function without PIN"
	PrimaryKeypadFunctionWithoutPinLength            = 4
	PrimaryKeypadFunctionWithoutPinNumber            = 61
	PrimaryKeypadFunctionWithoutPinTransitionCapable = message.TransitionCapableNo
	PrimaryKeypadFunctionWithoutPinAcknowledged      = message.AcknowledgedYes
)

type PrimaryKeypadFunctionWithoutPin struct {
	*message.Base
}

func NewPrimaryKeypadFunctionWithoutPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PrimaryKeypadFunctionWithoutPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PrimaryKeypadFunctionWithoutPinLength, length)
	}

	message := &PrimaryKeypadFunctionWithoutPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *PrimaryKeypadFunctionWithoutPin) Number() byte {
	return PrimaryKeypadFunctionWithoutPinNumber
}

func (m *PrimaryKeypadFunctionWithoutPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Primary Keypad Function without PIN",
		strings.Repeat("-", len("Primary Keypad Function without PIN")),
	)
}
