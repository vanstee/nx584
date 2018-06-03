// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	ProgramDataCommandName              = "Program Data Command"
	ProgramDataCommandLength            = 13
	ProgramDataCommandNumber            = 49
	ProgramDataCommandTransitionCapable = TransitionCapableNo
	ProgramDataCommandAcknowledged      = AcknowledgedYes
)

type ProgramDataCommand struct {
	*BaseMessage
}

func NewProgramDataCommand(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ProgramDataCommandLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ProgramDataCommandLength, length)
	}

	message := &ProgramDataCommand{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ProgramDataCommand) Number() byte {
	return ProgramDataCommandNumber
}

func (m *ProgramDataCommand) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Program Data Command",
		strings.Repeat("-", len("Program Data Command")),
	)
}
