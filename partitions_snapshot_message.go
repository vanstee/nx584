// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PartitionsSnapshotMessageName              = "Partitions Snapshot Message"
	PartitionsSnapshotMessageLength            = 9
	PartitionsSnapshotMessageNumber            = 7
	PartitionsSnapshotMessageTransitionCapable = TransitionCapableYes
	PartitionsSnapshotMessageAcknowledged      = AcknowledgedPossible
)

type PartitionsSnapshotMessage struct {
	*BaseMessage
}

func NewPartitionsSnapshotMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PartitionsSnapshotMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PartitionsSnapshotMessageLength, length)
	}

	message := &PartitionsSnapshotMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *PartitionsSnapshotMessage) Number() byte {
	return PartitionsSnapshotMessageNumber
}

func (m *PartitionsSnapshotMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Partitions Snapshot Message",
		strings.Repeat("-", len("Partitions Snapshot Message")),
	)
}
