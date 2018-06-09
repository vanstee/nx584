// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ZonesSnapshotRequestName              = "Zones Snapshot Request"
	ZonesSnapshotRequestLength            = 2
	ZonesSnapshotRequestNumber            = 37
	ZonesSnapshotRequestTransitionCapable = message.TransitionCapableNo
	ZonesSnapshotRequestAcknowledged      = message.AcknowledgedNo
)

type ZonesSnapshotRequest struct {
	*message.Base
	ZoneOffsetNumber byte
}

func NewZonesSnapshotRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ZonesSnapshotRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZonesSnapshotRequestLength, length)
	}

	message := &ZonesSnapshotRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),

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
