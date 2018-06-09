// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ZoneBypassToggleName              = "Zone Bypass Toggle"
	ZoneBypassToggleLength            = 2
	ZoneBypassToggleNumber            = 63
	ZoneBypassToggleTransitionCapable = message.TransitionCapableNo
	ZoneBypassToggleAcknowledged      = message.AcknowledgedYes
)

type ZoneBypassToggle struct {
	*message.Base
}

func NewZoneBypassToggle(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ZoneBypassToggleLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZoneBypassToggleLength, length)
	}

	message := &ZoneBypassToggle{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
