// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	InterfaceConfigurationRequestName              = "Interface Configuration Request"
	InterfaceConfigurationRequestLength            = 1
	InterfaceConfigurationRequestNumber            = 33
	InterfaceConfigurationRequestTransitionCapable = TransitionCapableNo
	InterfaceConfigurationRequestAcknowledged      = AcknowledgedNo
)

type InterfaceConfigurationRequest struct {
	*BaseMessage
}

func NewInterfaceConfigurationRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != InterfaceConfigurationRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", InterfaceConfigurationRequestLength, length)
	}

	message := &InterfaceConfigurationRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *InterfaceConfigurationRequest) Number() byte {
	return InterfaceConfigurationRequestNumber
}

func (m *InterfaceConfigurationRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Interface Configuration Request",
		strings.Repeat("-", len("Interface Configuration Request")),
	)
}
