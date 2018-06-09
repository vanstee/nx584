// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	NegativeAcknowledgeName              = "Negative Acknowledge"
	NegativeAcknowledgeLength            = 1
	NegativeAcknowledgeNumber            = 30
	NegativeAcknowledgeTransitionCapable = message.TransitionCapableNo
	NegativeAcknowledgeAcknowledged      = message.AcknowledgedNo
)

type NegativeAcknowledge struct {
	*message.Base
}

func NewNegativeAcknowledge(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != NegativeAcknowledgeLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", NegativeAcknowledgeLength, length)
	}

	message := &NegativeAcknowledge{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
