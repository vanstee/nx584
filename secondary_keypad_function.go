// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	SecondaryKeypadFunctionName              = "Secondary Keypad Function"
	SecondaryKeypadFunctionLength            = 3
	SecondaryKeypadFunctionNumber            = 62
	SecondaryKeypadFunctionTransitionCapable = TransitionCapableNo
	SecondaryKeypadFunctionAcknowledged      = AcknowledgedYes
)

type SecondaryKeypadFunction struct {
	*BaseMessage
}

func NewSecondaryKeypadFunction(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SecondaryKeypadFunctionLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SecondaryKeypadFunctionLength, length)
	}

	message := &SecondaryKeypadFunction{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SecondaryKeypadFunction) Number() byte {
	return SecondaryKeypadFunctionNumber
}

func (m *SecondaryKeypadFunction) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Secondary Keypad Function",
		strings.Repeat("-", len("Secondary Keypad Function")),
	)
}
