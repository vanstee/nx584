// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZoneNameMessageName              = "Zone Name Message"
	ZoneNameMessageLength            = 18
	ZoneNameMessageNumber            = 3
	ZoneNameMessageTransitionCapable = TransitionCapableNo
	ZoneNameMessageAcknowledged      = AcknowledgedNo
)

type ZoneNameMessage struct {
	*BaseMessage
}

func NewZoneNameMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneNameMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneNameMessageLength, length)
	}

	message := &ZoneNameMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZoneNameMessage) Number() byte {
	return ZoneNameMessageNumber
}

func (m *ZoneNameMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zone Name Message",
		strings.Repeat("-", len("Zone Name Message")),
	)
}
