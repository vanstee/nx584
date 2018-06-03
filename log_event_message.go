// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	LogEventMessageName              = "Log Event Message"
	LogEventMessageLength            = 10
	LogEventMessageNumber            = 10
	LogEventMessageTransitionCapable = TransitionCapableYes
	LogEventMessageAcknowledged      = AcknowledgedPossible
)

type LogEventMessage struct {
	*BaseMessage
}

func NewLogEventMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != LogEventMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", LogEventMessageLength, length)
	}

	message := &LogEventMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *LogEventMessage) Number() byte {
	return LogEventMessageNumber
}

func (m *LogEventMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Log Event Message",
		strings.Repeat("-", len("Log Event Message")),
	)
}
