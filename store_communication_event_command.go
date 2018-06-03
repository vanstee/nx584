// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	StoreCommunicationEventCommandName              = "Store Communication Event Command"
	StoreCommunicationEventCommandLength            = 6
	StoreCommunicationEventCommandNumber            = 58
	StoreCommunicationEventCommandTransitionCapable = TransitionCapableNo
	StoreCommunicationEventCommandAcknowledged      = AcknowledgedYes
)

type StoreCommunicationEventCommand struct {
	*BaseMessage
}

func NewStoreCommunicationEventCommand(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != StoreCommunicationEventCommandLength {
		log.Printf("message length incorrect: expected %d, actual: %d", StoreCommunicationEventCommandLength, length)
	}

	message := &StoreCommunicationEventCommand{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
