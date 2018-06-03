// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	ProgramDataRequestName              = "Program Data Request"
	ProgramDataRequestLength            = 4
	ProgramDataRequestNumber            = 48
	ProgramDataRequestTransitionCapable = TransitionCapableNo
	ProgramDataRequestAcknowledged      = AcknowledgedNo
)

type ProgramDataRequest struct {
	*BaseMessage
}

func NewProgramDataRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ProgramDataRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ProgramDataRequestLength, length)
	}

	message := &ProgramDataRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ProgramDataRequest) Number() byte {
	return ProgramDataRequestNumber
}

func (m *ProgramDataRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Program Data Request",
		strings.Repeat("-", len("Program Data Request")),
	)
}
