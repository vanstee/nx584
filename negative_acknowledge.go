// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	NegativeAcknowledgeName              = "Negative Acknowledge"
	NegativeAcknowledgeLength            = 1
	NegativeAcknowledgeNumber            = 30
	NegativeAcknowledgeTransitionCapable = TransitionCapableNo
	NegativeAcknowledgeAcknowledged      = AcknowledgedNo
)

type NegativeAcknowledge struct {
	*BaseMessage
}

func NewNegativeAcknowledge(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != NegativeAcknowledgeLength {
		log.Printf("message length incorrect: expected %d, actual: %d", NegativeAcknowledgeLength, length)
	}

	message := &NegativeAcknowledge{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *NegativeAcknowledge) Number() byte {
	return NegativeAcknowledgeNumber
}

func (m *NegativeAcknowledge) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Negative Acknowledge",
		strings.Repeat("-", len("Negative Acknowledge")),
	)
}
