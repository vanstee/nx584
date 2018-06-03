// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZoneSnapshotMessageName              = "Zone Snapshot Message"
	ZoneSnapshotMessageLength            = 10
	ZoneSnapshotMessageNumber            = 5
	ZoneSnapshotMessageTransitionCapable = TransitionCapableYes
	ZoneSnapshotMessageAcknowledged      = AcknowledgedPossible
)

type ZoneSnapshotMessage struct {
	*BaseMessage
}

func NewZoneSnapshotMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneSnapshotMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneSnapshotMessageLength, length)
	}

	message := &ZoneSnapshotMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}

	return message, nil
}

func (m *ZoneSnapshotMessage) Number() byte {
	return ZoneSnapshotMessageNumber
}

func (m *ZoneSnapshotMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Zone Snapshot Message",
		strings.Repeat("-", len("Zone Snapshot Message")),
	)
}
