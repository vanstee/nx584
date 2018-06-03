// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	ZoneNameRequestName              = "Zone Name Request"
	ZoneNameRequestLength            = 2
	ZoneNameRequestNumber            = 35
	ZoneNameRequestTransitionCapable = TransitionCapableNo
	ZoneNameRequestAcknowledged      = AcknowledgedNo
)

type ZoneNameRequest struct {
	*BaseMessage
	ZoneNumber byte
}

func NewZoneNameRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneNameRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZoneNameRequestLength, length)
	}

	message := &ZoneNameRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},

		ZoneNumber: data[0],
	}

	return message, nil
}

func (m *ZoneNameRequest) Number() byte {
	return ZoneNameRequestNumber
}

func (m *ZoneNameRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Number: %08b`,
		"Zone Name Request",
		strings.Repeat("-", len("Zone Name Request")),
		m.ZoneNumber,
	)
}
