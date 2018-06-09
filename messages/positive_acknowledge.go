// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PositiveAcknowledgeName              = "Positive Acknowledge"
	PositiveAcknowledgeLength            = 1
	PositiveAcknowledgeNumber            = 29
	PositiveAcknowledgeTransitionCapable = message.TransitionCapableNo
	PositiveAcknowledgeAcknowledged      = message.AcknowledgedNo
)

type PositiveAcknowledge struct {
	*message.Base
}

func NewPositiveAcknowledge(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PositiveAcknowledgeLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PositiveAcknowledgeLength, length)
	}

	message := &PositiveAcknowledge{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
