// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	InterfaceConfigurationRequestName              = "Interface Configuration Request"
	InterfaceConfigurationRequestLength            = 1
	InterfaceConfigurationRequestNumber            = 33
	InterfaceConfigurationRequestTransitionCapable = message.TransitionCapableNo
	InterfaceConfigurationRequestAcknowledged      = message.AcknowledgedNo
)

type InterfaceConfigurationRequest struct {
	*message.Base
}

func NewInterfaceConfigurationRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != InterfaceConfigurationRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", InterfaceConfigurationRequestLength, length)
	}

	message := &InterfaceConfigurationRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
