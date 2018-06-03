// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZoneStatusRequestName              = "Zone Status Request"
	ZoneStatusRequestLength            = 2
	ZoneStatusRequestNumber            = 36
	ZoneStatusRequestTransitionCapable = TransitionCapableNo
	ZoneStatusRequestAcknowledged      = AcknowledgedNo
)

type ZoneStatusRequest struct {
	*BaseMessage
}

func NewZoneStatusRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneStatusRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneStatusRequestLength, length)
	}

	message := &ZoneStatusRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZoneStatusRequest) Number() byte {
	return ZoneStatusRequestNumber
}

func (m *ZoneStatusRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zone Status Request",
		strings.Repeat("-", len("Zone Status Request")),
	)
}
