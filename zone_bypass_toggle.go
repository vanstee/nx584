// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZoneBypassToggleName              = "Zone Bypass Toggle"
	ZoneBypassToggleLength            = 2
	ZoneBypassToggleNumber            = 63
	ZoneBypassToggleTransitionCapable = TransitionCapableNo
	ZoneBypassToggleAcknowledged      = AcknowledgedYes
)

type ZoneBypassToggle struct {
	*BaseMessage
}

func NewZoneBypassToggle(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneBypassToggleLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneBypassToggleLength, length)
	}

	message := &ZoneBypassToggle{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZoneBypassToggle) Number() byte {
	return ZoneBypassToggleNumber
}

func (m *ZoneBypassToggle) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zone Bypass Toggle",
		strings.Repeat("-", len("Zone Bypass Toggle")),
	)
}
