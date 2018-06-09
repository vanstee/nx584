// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SystemStatusMessageName              = "System Status Message"
	SystemStatusMessageLength            = 12
	SystemStatusMessageNumber            = 8
	SystemStatusMessageTransitionCapable = message.TransitionCapableYes
	SystemStatusMessageAcknowledged      = message.AcknowledgedPossible
)

type SystemStatusMessage struct {
	*message.Base
}

func NewSystemStatusMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SystemStatusMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SystemStatusMessageLength, length)
	}

	message := &SystemStatusMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
