// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	LogEventMessageName              = "Log Event Message"
	LogEventMessageLength            = 10
	LogEventMessageNumber            = 10
	LogEventMessageTransitionCapable = message.TransitionCapableYes
	LogEventMessageAcknowledged      = message.AcknowledgedPossible
)

type LogEventMessage struct {
	*message.Base
}

func NewLogEventMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != LogEventMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", LogEventMessageLength, length)
	}

	message := &LogEventMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
