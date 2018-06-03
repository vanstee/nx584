// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PositiveAcknowledgeName              = "Positive Acknowledge"
	PositiveAcknowledgeLength            = 1
	PositiveAcknowledgeNumber            = 29
	PositiveAcknowledgeTransitionCapable = TransitionCapableNo
	PositiveAcknowledgeAcknowledged      = AcknowledgedNo
)

type PositiveAcknowledge struct {
	*BaseMessage
}

func NewPositiveAcknowledge(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PositiveAcknowledgeLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PositiveAcknowledgeLength, length)
	}

	message := &PositiveAcknowledge{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *PositiveAcknowledge) Number() byte {
	return PositiveAcknowledgeNumber
}

func (m *PositiveAcknowledge) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Positive Acknowledge",
		strings.Repeat("-", len("Positive Acknowledge")),
	)
}
