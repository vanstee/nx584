// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
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
}

func NewZoneNameRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneNameRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneNameRequestLength, length)
	}

	message := &ZoneNameRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZoneNameRequest) Number() byte {
	return ZoneNameRequestNumber
}

func (m *ZoneNameRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zone Name Request",
		strings.Repeat("-", len("Zone Name Request")),
	)
}
