// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	SystemStatusMessageName              = "System Status Message"
	SystemStatusMessageLength            = 12
	SystemStatusMessageNumber            = 8
	SystemStatusMessageTransitionCapable = TransitionCapableYes
	SystemStatusMessageAcknowledged      = AcknowledgedPossible
)

type SystemStatusMessage struct {
	*BaseMessage
}

func NewSystemStatusMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SystemStatusMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", SystemStatusMessageLength, length)
	}

	message := &SystemStatusMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SystemStatusMessage) Number() byte {
	return SystemStatusMessageNumber
}

func (m *SystemStatusMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"System Status Message",
		strings.Repeat("-", len("System Status Message")),
	)
}
