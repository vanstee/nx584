// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PartitionsSnapshotMessageName              = "Partitions Snapshot Message"
	PartitionsSnapshotMessageLength            = 9
	PartitionsSnapshotMessageNumber            = 7
	PartitionsSnapshotMessageTransitionCapable = message.TransitionCapableYes
	PartitionsSnapshotMessageAcknowledged      = message.AcknowledgedPossible
)

type PartitionsSnapshotMessage struct {
	*message.Base
}

func NewPartitionsSnapshotMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PartitionsSnapshotMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PartitionsSnapshotMessageLength, length)
	}

	message := &PartitionsSnapshotMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),
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
