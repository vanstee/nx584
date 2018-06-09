// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	KeypadTerminalModeRequestName              = "Keypad Terminal Mode Request"
	KeypadTerminalModeRequestLength            = 3
	KeypadTerminalModeRequestNumber            = 44
	KeypadTerminalModeRequestTransitionCapable = message.TransitionCapableNo
	KeypadTerminalModeRequestAcknowledged      = message.AcknowledgedYes
)

type KeypadTerminalModeRequest struct {
	*message.Base
}

func NewKeypadTerminalModeRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != KeypadTerminalModeRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", KeypadTerminalModeRequestLength, length)
	}

	message := &KeypadTerminalModeRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *KeypadTerminalModeRequest) Number() byte {
	return KeypadTerminalModeRequestNumber
}

func (m *KeypadTerminalModeRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Keypad Terminal Mode Request",
		strings.Repeat("-", len("Keypad Terminal Mode Request")),
	)
}
