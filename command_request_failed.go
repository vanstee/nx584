// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	CommandRequestFailedName              = "Command / Request Failed"
	CommandRequestFailedLength            = 1
	CommandRequestFailedNumber            = 28
	CommandRequestFailedTransitionCapable = TransitionCapableNo
	CommandRequestFailedAcknowledged      = AcknowledgedNo
)

type CommandRequestFailed struct {
	*BaseMessage
}

func NewCommandRequestFailed(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != CommandRequestFailedLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", CommandRequestFailedLength, length)
	}

	message := &CommandRequestFailed{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
