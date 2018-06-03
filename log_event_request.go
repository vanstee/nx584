// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	LogEventRequestName              = "Log Event Request"
	LogEventRequestLength            = 2
	LogEventRequestNumber            = 42
	LogEventRequestTransitionCapable = TransitionCapableNo
	LogEventRequestAcknowledged      = AcknowledgedNo
)

type LogEventRequest struct {
	*BaseMessage
}

func NewLogEventRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != LogEventRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", LogEventRequestLength, length)
	}

	message := &LogEventRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *LogEventRequest) Number() byte {
	return LogEventRequestNumber
}

func (m *LogEventRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Log Event Request",
		strings.Repeat("-", len("Log Event Request")),
	)
}
