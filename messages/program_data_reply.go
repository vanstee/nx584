// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ProgramDataReplyName              = "Program Data Reply"
	ProgramDataReplyLength            = 13
	ProgramDataReplyNumber            = 16
	ProgramDataReplyTransitionCapable = message.TransitionCapableNo
	ProgramDataReplyAcknowledged      = message.AcknowledgedNo
)

type ProgramDataReply struct {
	*message.Base
}

func NewProgramDataReply(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ProgramDataReplyLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ProgramDataReplyLength, length)
	}

	message := &ProgramDataReply{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *ProgramDataReply) Number() byte {
	return ProgramDataReplyNumber
}

func (m *ProgramDataReply) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Program Data Reply",
		strings.Repeat("-", len("Program Data Reply")),
	)
}
