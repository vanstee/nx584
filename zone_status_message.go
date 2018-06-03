package nx584

import (
	"fmt"
	"log"
)

const (
	ZoneStatusMessageName              = "Zone Status Message"
	ZoneStatusMessageLength            = 8
	ZoneStatusMessageNumber            = 0x04
	ZoneStatusMessageTransitionCapable = TransitionCapableYes
	ZoneStatusMessageAcknowledged      = AcknowledgedPossible
)

type ZoneStatusMessage struct {
	*BaseMessage
	ZoneNumber          byte
	PartitionMask       byte
	ZoneTypeFlags1      byte
	ZoneTypeFlags2      byte
	ZoneTypeFlags3      byte
	ZoneConditionFlags1 byte
	ZoneConditionFlags2 byte
	Partition1Enable    byte
	Partition2Enable    byte
	Partition3Enable    byte
	Partition4Enable    byte
	Partition5Enable    byte
	Partition6Enable    byte
	Partition7Enable    byte
	Partition8Enable    byte
	Fire                byte
	N24Hour             byte
	Keyswitch           byte
	Follower            byte
	EntryExitDelay1     byte
	EntryExitDelay2     byte
	Interior            byte
	LocalOnly           byte
}

func NewZoneStatusMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZoneStatusMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZoneStatusMessageLength, length)
	}

	message := &ZoneStatusMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
		ZoneNumber:          data[0],
		PartitionMask:       data[1],
		ZoneTypeFlags1:      data[2],
		ZoneTypeFlags2:      data[3],
		ZoneTypeFlags3:      data[4],
		ZoneConditionFlags1: data[5],
		ZoneConditionFlags2: data[6],
		Partition1Enable:    (data[1] >> 7) & 0x1,
		Partition2Enable:    (data[1] >> 6) & 0x1,
		Partition3Enable:    (data[1] >> 5) & 0x1,
		Partition4Enable:    (data[1] >> 4) & 0x1,
		Partition5Enable:    (data[1] >> 3) & 0x1,
		Partition6Enable:    (data[1] >> 2) & 0x1,
		Partition7Enable:    (data[1] >> 1) & 0x1,
		Partition8Enable:    (data[1] >> 0) & 0x1,
		Fire:                (data[2] >> 7) & 0x1,
		N24Hour:             (data[2] >> 6) & 0x1,
		Keyswitch:           (data[2] >> 5) & 0x1,
		Follower:            (data[2] >> 4) & 0x1,
		EntryExitDelay1:     (data[2] >> 3) & 0x1,
		EntryExitDelay2:     (data[2] >> 2) & 0x1,
		Interior:            (data[2] >> 1) & 0x1,
		LocalOnly:           (data[2] >> 0) & 0x1,
	}

	return message, nil
}

func (m *ZoneStatusMessage) Number() byte {
	return ZoneStatusMessageNumber
}

func (m *ZoneStatusMessage) String() string {
	return fmt.Sprintf(`%s
                    -------------------
                    Zone Number: %d
                    Partition Mask: %b
                      Partition 1 Enable: %b
                      Partition 2 Enable: %b
                      Partition 3 Enable: %b
                      Partition 4 Enable: %b
                      Partition 5 Enable: %b
                      Partition 6 Enable: %b
                      Partition 7 Enable: %b
                      Partition 8 Enable: %b
                    Zone Type Flags 1: %b
                      Fire: %b
                      24 Hour: %b
                      Key-switch: %b
                      Follower: %b
                      Entry Exit Delay 1: %b
                      Entry Exit Delay 2: %b
                      Interior: %b
                      Local Only: %b`,
		ZoneStatusMessageName,
		m.ZoneNumber,
		m.PartitionMask,
		m.Partition1Enable,
		m.Partition2Enable,
		m.Partition3Enable,
		m.Partition4Enable,
		m.Partition5Enable,
		m.Partition6Enable,
		m.Partition7Enable,
		m.Partition8Enable,
		m.ZoneTypeFlags1,
		m.Fire,
		m.N24Hour,
		m.Keyswitch,
		m.Follower,
		m.EntryExitDelay1,
		m.EntryExitDelay2,
		m.Interior,
		m.LocalOnly,
	)
}
