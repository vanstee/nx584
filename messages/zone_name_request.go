// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ZoneNameRequestName              = "Zone Name Request"
	ZoneNameRequestLength            = 2
	ZoneNameRequestNumber            = 35
	ZoneNameRequestTransitionCapable = message.TransitionCapableNo
	ZoneNameRequestAcknowledged      = message.AcknowledgedNo
)

type ZoneNameRequest struct {
	*message.Base
	ZoneNumber byte
}

func NewZoneNameRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ZoneNameRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZoneNameRequestLength, length)
	}

	message := &ZoneNameRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),

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
