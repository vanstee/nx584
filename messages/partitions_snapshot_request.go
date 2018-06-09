// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PartitionsSnapshotRequestName              = "Partitions Snapshot Request"
	PartitionsSnapshotRequestLength            = 1
	PartitionsSnapshotRequestNumber            = 39
	PartitionsSnapshotRequestTransitionCapable = message.TransitionCapableNo
	PartitionsSnapshotRequestAcknowledged      = message.AcknowledgedNo
)

type PartitionsSnapshotRequest struct {
	*message.Base
}

func NewPartitionsSnapshotRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PartitionsSnapshotRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PartitionsSnapshotRequestLength, length)
	}

	message := &PartitionsSnapshotRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
