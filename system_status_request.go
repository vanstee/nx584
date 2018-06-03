// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	SystemStatusRequestName              = "System Status Request"
	SystemStatusRequestLength            = 1
	SystemStatusRequestNumber            = 40
	SystemStatusRequestTransitionCapable = TransitionCapableNo
	SystemStatusRequestAcknowledged      = AcknowledgedNo
)

type SystemStatusRequest struct {
	*BaseMessage
}

func NewSystemStatusRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != SystemStatusRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SystemStatusRequestLength, length)
	}

	message := &SystemStatusRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *SystemStatusRequest) Number() byte {
	return SystemStatusRequestNumber
}

func (m *SystemStatusRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"System Status Request",
		strings.Repeat("-", len("System Status Request")),
	)
}
