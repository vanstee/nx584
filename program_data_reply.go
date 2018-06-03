// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ProgramDataReplyName              = "Program Data Reply"
	ProgramDataReplyLength            = 13
	ProgramDataReplyNumber            = 16
	ProgramDataReplyTransitionCapable = TransitionCapableNo
	ProgramDataReplyAcknowledged      = AcknowledgedNo
)

type ProgramDataReply struct {
	*BaseMessage
}

func NewProgramDataReply(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ProgramDataReplyLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ProgramDataReplyLength, length)
	}

	message := &ProgramDataReply{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
