// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PartitionsSnapshotRequestName              = "Partitions Snapshot Request"
	PartitionsSnapshotRequestLength            = 1
	PartitionsSnapshotRequestNumber            = 39
	PartitionsSnapshotRequestTransitionCapable = TransitionCapableNo
	PartitionsSnapshotRequestAcknowledged      = AcknowledgedNo
)

type PartitionsSnapshotRequest struct {
	*BaseMessage
}

func NewPartitionsSnapshotRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PartitionsSnapshotRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PartitionsSnapshotRequestLength, length)
	}

	message := &PartitionsSnapshotRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *PartitionsSnapshotRequest) Number() byte {
	return PartitionsSnapshotRequestNumber
}

func (m *PartitionsSnapshotRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Partitions Snapshot Request",
		strings.Repeat("-", len("Partitions Snapshot Request")),
	)
}
