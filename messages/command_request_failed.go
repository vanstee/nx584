// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	CommandRequestFailedName              = "Command / Request Failed"
	CommandRequestFailedLength            = 1
	CommandRequestFailedNumber            = 28
	CommandRequestFailedTransitionCapable = message.TransitionCapableNo
	CommandRequestFailedAcknowledged      = message.AcknowledgedNo
)

type CommandRequestFailed struct {
	*message.Base
}

func NewCommandRequestFailed(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != CommandRequestFailedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", CommandRequestFailedLength, length)
	}

	message := &CommandRequestFailed{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *CommandRequestFailed) Number() byte {
	return CommandRequestFailedNumber
}

func (m *CommandRequestFailed) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Command / Request Failed",
		strings.Repeat("-", len("Command / Request Failed")),
	)
}
