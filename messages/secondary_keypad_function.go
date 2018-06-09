// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SecondaryKeypadFunctionName              = "Secondary Keypad Function"
	SecondaryKeypadFunctionLength            = 3
	SecondaryKeypadFunctionNumber            = 62
	SecondaryKeypadFunctionTransitionCapable = message.TransitionCapableNo
	SecondaryKeypadFunctionAcknowledged      = message.AcknowledgedYes
)

type SecondaryKeypadFunction struct {
	*message.Base
}

func NewSecondaryKeypadFunction(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SecondaryKeypadFunctionLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SecondaryKeypadFunctionLength, length)
	}

	message := &SecondaryKeypadFunction{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
