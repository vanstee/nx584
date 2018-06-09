// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	StoreCommunicationEventCommandName              = "Store Communication Event Command"
	StoreCommunicationEventCommandLength            = 6
	StoreCommunicationEventCommandNumber            = 58
	StoreCommunicationEventCommandTransitionCapable = message.TransitionCapableNo
	StoreCommunicationEventCommandAcknowledged      = message.AcknowledgedYes
)

type StoreCommunicationEventCommand struct {
	*message.Base
}

func NewStoreCommunicationEventCommand(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != StoreCommunicationEventCommandLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", StoreCommunicationEventCommandLength, length)
	}

	message := &StoreCommunicationEventCommand{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *StoreCommunicationEventCommand) Number() byte {
	return StoreCommunicationEventCommandNumber
}

func (m *StoreCommunicationEventCommand) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Store Communication Event Command",
		strings.Repeat("-", len("Store Communication Event Command")),
	)
}
