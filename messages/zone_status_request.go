// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ZoneStatusRequestName              = "Zone Status Request"
	ZoneStatusRequestLength            = 2
	ZoneStatusRequestNumber            = 36
	ZoneStatusRequestTransitionCapable = message.TransitionCapableNo
	ZoneStatusRequestAcknowledged      = message.AcknowledgedNo
)

type ZoneStatusRequest struct {
	*message.Base
	ZoneNumber byte
}

func NewZoneStatusRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ZoneStatusRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZoneStatusRequestLength, length)
	}

	message := &ZoneStatusRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),

		ZoneNumber: data[0],
	}

	return message, nil
}

func (m *ZoneStatusRequest) Number() byte {
	return ZoneStatusRequestNumber
}

func (m *ZoneStatusRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Number: %08b`,
		"Zone Status Request",
		strings.Repeat("-", len("Zone Status Request")),
		m.ZoneNumber,
	)
}
