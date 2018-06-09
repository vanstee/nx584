// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SetClockCalendarCommandName              = "Set Clock / Calendar Command"
	SetClockCalendarCommandLength            = 7
	SetClockCalendarCommandNumber            = 59
	SetClockCalendarCommandTransitionCapable = message.TransitionCapableNo
	SetClockCalendarCommandAcknowledged      = message.AcknowledgedYes
)

type SetClockCalendarCommand struct {
	*message.Base
}

func NewSetClockCalendarCommand(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SetClockCalendarCommandLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SetClockCalendarCommandLength, length)
	}

	message := &SetClockCalendarCommand{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
