// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ProgramDataCommandName              = "Program Data Command"
	ProgramDataCommandLength            = 13
	ProgramDataCommandNumber            = 49
	ProgramDataCommandTransitionCapable = message.TransitionCapableNo
	ProgramDataCommandAcknowledged      = message.AcknowledgedYes
)

type ProgramDataCommand struct {
	*message.Base
}

func NewProgramDataCommand(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ProgramDataCommandLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ProgramDataCommandLength, length)
	}

	message := &ProgramDataCommand{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
