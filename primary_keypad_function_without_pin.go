// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PrimaryKeypadFunctionWithoutPinName              = "Primary Keypad Function without PIN"
	PrimaryKeypadFunctionWithoutPinLength            = 4
	PrimaryKeypadFunctionWithoutPinNumber            = 61
	PrimaryKeypadFunctionWithoutPinTransitionCapable = TransitionCapableNo
	PrimaryKeypadFunctionWithoutPinAcknowledged      = AcknowledgedYes
)

type PrimaryKeypadFunctionWithoutPin struct {
	*BaseMessage
}

func NewPrimaryKeypadFunctionWithoutPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PrimaryKeypadFunctionWithoutPinLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PrimaryKeypadFunctionWithoutPinLength, length)
	}

	message := &PrimaryKeypadFunctionWithoutPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
