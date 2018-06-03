// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZoneStatusMessageName              = "Zone Status Message"
	ZoneStatusMessageLength            = 8
	ZoneStatusMessageNumber            = 4
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
	KeySwitch           byte
	Follower            byte
	EntryExitDelay1     byte
	EntryExitDelay2     byte
	Interior            byte
	LocalOnly           byte
	KeypadSounder       byte
	YelpingSiren        byte
	SteadySiren         byte
	Chime               byte
	Bypassable          byte
	GroupBypassable     byte
	ForceArmable        byte
	EntryGuard          byte
	FastLoopResponse    byte
	DoubleEolTamper     byte
	Trouble1            byte
	CrossZone           byte
	DialerDelay         byte
	SwingerShutdown     byte
	Restorable          byte
	ListenIn            byte
	Faulted             byte
	Tampered            byte
	Trouble2            byte
	Bypassed            byte
	Inhibited           byte
	LowBattery          byte
	LossOfSupervision   byte
	AlarmMemory         byte
	BypassMemory        byte
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

		ZoneNumber: data[0],

		PartitionMask:    data[1],
		Partition1Enable: (data[1] >> 0) & 0x1,
		Partition2Enable: (data[1] >> 1) & 0x1,
		Partition3Enable: (data[1] >> 2) & 0x1,
		Partition4Enable: (data[1] >> 3) & 0x1,
		Partition5Enable: (data[1] >> 4) & 0x1,
		Partition6Enable: (data[1] >> 5) & 0x1,
		Partition7Enable: (data[1] >> 6) & 0x1,
		Partition8Enable: (data[1] >> 7) & 0x1,

		ZoneTypeFlags1:  data[2],
		Fire:            (data[2] >> 0) & 0x1,
		N24Hour:         (data[2] >> 1) & 0x1,
		KeySwitch:       (data[2] >> 2) & 0x1,
		Follower:        (data[2] >> 3) & 0x1,
		EntryExitDelay1: (data[2] >> 4) & 0x1,
		EntryExitDelay2: (data[2] >> 5) & 0x1,
		Interior:        (data[2] >> 6) & 0x1,
		LocalOnly:       (data[2] >> 7) & 0x1,

		ZoneTypeFlags2:  data[3],
		KeypadSounder:   (data[3] >> 0) & 0x1,
		YelpingSiren:    (data[3] >> 1) & 0x1,
		SteadySiren:     (data[3] >> 2) & 0x1,
		Chime:           (data[3] >> 3) & 0x1,
		Bypassable:      (data[3] >> 4) & 0x1,
		GroupBypassable: (data[3] >> 5) & 0x1,
		ForceArmable:    (data[3] >> 6) & 0x1,
		EntryGuard:      (data[3] >> 7) & 0x1,

		ZoneTypeFlags3:   data[4],
		FastLoopResponse: (data[4] >> 0) & 0x1,
		DoubleEolTamper:  (data[4] >> 1) & 0x1,
		Trouble1:         (data[4] >> 2) & 0x1,
		CrossZone:        (data[4] >> 3) & 0x1,
		DialerDelay:      (data[4] >> 4) & 0x1,
		SwingerShutdown:  (data[4] >> 5) & 0x1,
		Restorable:       (data[4] >> 6) & 0x1,
		ListenIn:         (data[4] >> 7) & 0x1,

		ZoneConditionFlags1: data[5],
		Faulted:             (data[5] >> 0) & 0x1,
		Tampered:            (data[5] >> 1) & 0x1,
		Trouble2:            (data[5] >> 2) & 0x1,
		Bypassed:            (data[5] >> 3) & 0x1,
		Inhibited:           (data[5] >> 4) & 0x1,
		LowBattery:          (data[5] >> 5) & 0x1,
		LossOfSupervision:   (data[5] >> 6) & 0x1,

		ZoneConditionFlags2: data[6],
		AlarmMemory:         (data[6] >> 0) & 0x1,
		BypassMemory:        (data[6] >> 1) & 0x1,
	}

	return message, nil
}

func (m *ZoneStatusMessage) Number() byte {
	return ZoneStatusMessageNumber
}

func (m *ZoneStatusMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Number: %08b
                    Partition Mask: %08b
                      Partition 1 Enable: %01b
                      Partition 2 Enable: %01b
                      Partition 3 Enable: %01b
                      Partition 4 Enable: %01b
                      Partition 5 Enable: %01b
                      Partition 6 Enable: %01b
                      Partition 7 Enable: %01b
                      Partition 8 Enable: %01b
                    Zone Type Flags 1: %08b
                      Fire: %01b
                      24 Hour: %01b
                      Key-switch: %01b
                      Follower: %01b
                      Entry / Exit Delay 1: %01b
                      Entry / Exit Delay 2: %01b
                      Interior: %01b
                      Local Only: %01b
                    Zone Type Flags 2: %08b
                      Keypad Sounder: %01b
                      Yelping Siren: %01b
                      Steady Siren: %01b
                      Chime: %01b
                      Bypassable: %01b
                      Group Bypassable: %01b
                      Force Armable: %01b
                      Entry Guard: %01b
                    Zone Type Flags 3: %08b
                      Fast Loop Response: %01b
                      Double EOL Tamper: %01b
                      Trouble 1: %01b
                      Cross Zone: %01b
                      Dialer Delay: %01b
                      Swinger Shutdown: %01b
                      Restorable: %01b
                      Listen In: %01b
                    Zone Condition Flags 1: %08b
                      Faulted: %01b
                      Tampered: %01b
                      Trouble 2: %01b
                      Bypassed: %01b
                      Inhibited: %01b
                      Low Battery: %01b
                      Loss of Supervision: %01b
                    Zone Condition Flags 2: %08b
                      Alarm Memory: %01b
                      Bypass Memory: %01b`,
		"Zone Status Message",
		strings.Repeat("-", len("Zone Status Message")),
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
		m.KeySwitch,
		m.Follower,
		m.EntryExitDelay1,
		m.EntryExitDelay2,
		m.Interior,
		m.LocalOnly,
		m.ZoneTypeFlags2,
		m.KeypadSounder,
		m.YelpingSiren,
		m.SteadySiren,
		m.Chime,
		m.Bypassable,
		m.GroupBypassable,
		m.ForceArmable,
		m.EntryGuard,
		m.ZoneTypeFlags3,
		m.FastLoopResponse,
		m.DoubleEolTamper,
		m.Trouble1,
		m.CrossZone,
		m.DialerDelay,
		m.SwingerShutdown,
		m.Restorable,
		m.ListenIn,
		m.ZoneConditionFlags1,
		m.Faulted,
		m.Tampered,
		m.Trouble2,
		m.Bypassed,
		m.Inhibited,
		m.LowBattery,
		m.LossOfSupervision,
		m.ZoneConditionFlags2,
		m.AlarmMemory,
		m.BypassMemory,
	)
}
