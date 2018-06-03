// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZonesSnapshotRequestName              = "Zones Snapshot Request"
	ZonesSnapshotRequestLength            = 2
	ZonesSnapshotRequestNumber            = 37
	ZonesSnapshotRequestTransitionCapable = TransitionCapableNo
	ZonesSnapshotRequestAcknowledged      = AcknowledgedNo
)

type ZonesSnapshotRequest struct {
	*BaseMessage
}

func NewZonesSnapshotRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZonesSnapshotRequestLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZonesSnapshotRequestLength, length)
	}

	message := &ZonesSnapshotRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZonesSnapshotRequest) Number() byte {
	return ZonesSnapshotRequestNumber
}

func (m *ZonesSnapshotRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zones Snapshot Request",
		strings.Repeat("-", len("Zones Snapshot Request")),
	)
}
