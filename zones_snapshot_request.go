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
	ZoneOffsetNumber byte
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

		ZoneOffsetNumber: data[0],
	}

	return message, nil
}

func (m *ZonesSnapshotRequest) Number() byte {
	return ZonesSnapshotRequestNumber
}

func (m *ZonesSnapshotRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Offset Number: %08b`,
		"Zones Snapshot Request",
		strings.Repeat("-", len("Zones Snapshot Request")),
		m.ZoneOffsetNumber,
	)
}
