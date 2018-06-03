// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	KeypadTerminalModeRequestName              = "Keypad Terminal Mode Request"
	KeypadTerminalModeRequestLength            = 3
	KeypadTerminalModeRequestNumber            = 44
	KeypadTerminalModeRequestTransitionCapable = TransitionCapableNo
	KeypadTerminalModeRequestAcknowledged      = AcknowledgedYes
)

type KeypadTerminalModeRequest struct {
	*BaseMessage
}

func NewKeypadTerminalModeRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != KeypadTerminalModeRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", KeypadTerminalModeRequestLength, length)
	}

	message := &KeypadTerminalModeRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
