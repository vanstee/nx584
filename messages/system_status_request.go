// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SystemStatusRequestName              = "System Status Request"
	SystemStatusRequestLength            = 1
	SystemStatusRequestNumber            = 40
	SystemStatusRequestTransitionCapable = message.TransitionCapableNo
	SystemStatusRequestAcknowledged      = message.AcknowledgedNo
)

type SystemStatusRequest struct {
	*message.Base
}

func NewSystemStatusRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SystemStatusRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SystemStatusRequestLength, length)
	}

	message := &SystemStatusRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
