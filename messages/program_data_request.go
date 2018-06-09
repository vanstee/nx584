// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ProgramDataRequestName              = "Program Data Request"
	ProgramDataRequestLength            = 4
	ProgramDataRequestNumber            = 48
	ProgramDataRequestTransitionCapable = message.TransitionCapableNo
	ProgramDataRequestAcknowledged      = message.AcknowledgedNo
)

type ProgramDataRequest struct {
	*message.Base
}

func NewProgramDataRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ProgramDataRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ProgramDataRequestLength, length)
	}

	message := &ProgramDataRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
