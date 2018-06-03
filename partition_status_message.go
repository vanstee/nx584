// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PartitionStatusMessageName              = "Partition Status Message"
	PartitionStatusMessageLength            = 9
	PartitionStatusMessageNumber            = 6
	PartitionStatusMessageTransitionCapable = TransitionCapableYes
	PartitionStatusMessageAcknowledged      = AcknowledgedPossible
)

type PartitionStatusMessage struct {
	*BaseMessage
}

func NewPartitionStatusMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PartitionStatusMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PartitionStatusMessageLength, length)
	}

	message := &PartitionStatusMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *PartitionStatusMessage) Number() byte {
	return PartitionStatusMessageNumber
}

func (m *PartitionStatusMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Partition Status Message",
		strings.Repeat("-", len("Partition Status Message")),
	)
}
