// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	InterfaceConfigurationMessageName              = "Interface Configuration Message"
	InterfaceConfigurationMessageLength            = 11
	InterfaceConfigurationMessageNumber            = 1
	InterfaceConfigurationMessageTransitionCapable = TransitionCapableNo
	InterfaceConfigurationMessageAcknowledged      = AcknowledgedPossible
)

type InterfaceConfigurationMessage struct {
	*BaseMessage
}

func NewInterfaceConfigurationMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != InterfaceConfigurationMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", InterfaceConfigurationMessageLength, length)
	}

	message := &InterfaceConfigurationMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
