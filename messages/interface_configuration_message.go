// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	InterfaceConfigurationMessageName              = "Interface Configuration Message"
	InterfaceConfigurationMessageLength            = 11
	InterfaceConfigurationMessageNumber            = 1
	InterfaceConfigurationMessageTransitionCapable = message.TransitionCapableNo
	InterfaceConfigurationMessageAcknowledged      = message.AcknowledgedPossible
)

type InterfaceConfigurationMessage struct {
	*message.Base
}

func NewInterfaceConfigurationMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != InterfaceConfigurationMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", InterfaceConfigurationMessageLength, length)
	}

	message := &InterfaceConfigurationMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *InterfaceConfigurationMessage) Number() byte {
	return InterfaceConfigurationMessageNumber
}

func (m *InterfaceConfigurationMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Interface Configuration Message",
		strings.Repeat("-", len("Interface Configuration Message")),
	)
}
