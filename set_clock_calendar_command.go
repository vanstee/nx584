// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	SetClockCalendarCommandName              = "Set Clock / Calendar Command"
	SetClockCalendarCommandLength            = 7
	SetClockCalendarCommandNumber            = 59
	SetClockCalendarCommandTransitionCapable = TransitionCapableNo
	SetClockCalendarCommandAcknowledged      = AcknowledgedYes
)

type SetClockCalendarCommand struct {
	*BaseMessage
}

func NewSetClockCalendarCommand(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SetClockCalendarCommandLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetClockCalendarCommandLength, length)
	}

	message := &SetClockCalendarCommand{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SetClockCalendarCommand) Number() byte {
	return SetClockCalendarCommandNumber
}

func (m *SetClockCalendarCommand) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Set Clock / Calendar Command",
		strings.Repeat("-", len("Set Clock / Calendar Command")),
	)
}
