// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	LogEventRequestName              = "Log Event Request"
	LogEventRequestLength            = 2
	LogEventRequestNumber            = 42
	LogEventRequestTransitionCapable = message.TransitionCapableNo
	LogEventRequestAcknowledged      = message.AcknowledgedNo
)

type LogEventRequest struct {
	*message.Base
}

func NewLogEventRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != LogEventRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", LogEventRequestLength, length)
	}

	message := &LogEventRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
