// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	ZonesSnapshotMessageName              = "Zones Snapshot Message"
	ZonesSnapshotMessageLength            = 10
	ZonesSnapshotMessageNumber            = 5
	ZonesSnapshotMessageTransitionCapable = TransitionCapableYes
	ZonesSnapshotMessageAcknowledged      = AcknowledgedPossible
)

type ZonesSnapshotMessage struct {
	*BaseMessage
	ZoneOffset        byte
	Zone12            byte
	Zone34            byte
	Zone56            byte
	Zone78            byte
	Zone910           byte
	Zone1112          byte
	Zone1314          byte
	Zone1516          byte
	Zone1Faulted      byte
	Zone1Bypass       byte
	Zone1Trouble      byte
	Zone1AlarmMemory  byte
	Zone2Faulted      byte
	Zone2Bypass       byte
	Zone2Trouble      byte
	Zone2AlarmMemory  byte
	Zone3Faulted      byte
	Zone3Bypass       byte
	Zone3Trouble      byte
	Zone3AlarmMemory  byte
	Zone4Faulted      byte
	Zone4Bypass       byte
	Zone4Trouble      byte
	Zone4AlarmMemory  byte
	Zone5Faulted      byte
	Zone5Bypass       byte
	Zone5Trouble      byte
	Zone5AlarmMemory  byte
	Zone6Faulted      byte
	Zone6Bypass       byte
	Zone6Trouble      byte
	Zone6AlarmMemory  byte
	Zone7Faulted      byte
	Zone7Bypass       byte
	Zone7Trouble      byte
	Zone7AlarmMemory  byte
	Zone8Faulted      byte
	Zone8Bypass       byte
	Zone8Trouble      byte
	Zone8AlarmMemory  byte
	Zone9Faulted      byte
	Zone9Bypass       byte
	Zone9Trouble      byte
	Zone9AlarmMemory  byte
	Zone10Faulted     byte
	Zone10Bypass      byte
	Zone10Trouble     byte
	Zone10AlarmMemory byte
	Zone11Faulted     byte
	Zone11Bypass      byte
	Zone11Trouble     byte
	Zone11AlarmMemory byte
	Zone12Faulted     byte
	Zone12Bypass      byte
	Zone12Trouble     byte
	Zone12AlarmMemory byte
	Zone13Faulted     byte
	Zone13Bypass      byte
	Zone13Trouble     byte
	Zone13AlarmMemory byte
	Zone14Faulted     byte
	Zone14Bypass      byte
	Zone14Trouble     byte
	Zone14AlarmMemory byte
	Zone15Faulted     byte
	Zone15Bypass      byte
	Zone15Trouble     byte
	Zone15AlarmMemory byte
	Zone16Faulted     byte
	Zone16Bypass      byte
	Zone16Trouble     byte
	Zone16AlarmMemory byte
}

func NewZonesSnapshotMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != ZonesSnapshotMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", ZonesSnapshotMessageLength, length)
	}

	message := &ZonesSnapshotMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},

		ZoneOffset: data[0],

		Zone12:           data[1],
		Zone1Faulted:     (data[1] >> 0) & 0x1,
		Zone1Bypass:      (data[1] >> 1) & 0x1,
		Zone1Trouble:     (data[1] >> 2) & 0x1,
		Zone1AlarmMemory: (data[1] >> 3) & 0x1,
		Zone2Faulted:     (data[1] >> 4) & 0x1,
		Zone2Bypass:      (data[1] >> 5) & 0x1,
		Zone2Trouble:     (data[1] >> 6) & 0x1,
		Zone2AlarmMemory: (data[1] >> 7) & 0x1,

		Zone34:           data[2],
		Zone3Faulted:     (data[2] >> 0) & 0x1,
		Zone3Bypass:      (data[2] >> 1) & 0x1,
		Zone3Trouble:     (data[2] >> 2) & 0x1,
		Zone3AlarmMemory: (data[2] >> 3) & 0x1,
		Zone4Faulted:     (data[2] >> 4) & 0x1,
		Zone4Bypass:      (data[2] >> 5) & 0x1,
		Zone4Trouble:     (data[2] >> 6) & 0x1,
		Zone4AlarmMemory: (data[2] >> 7) & 0x1,

		Zone56:           data[3],
		Zone5Faulted:     (data[3] >> 0) & 0x1,
		Zone5Bypass:      (data[3] >> 1) & 0x1,
		Zone5Trouble:     (data[3] >> 2) & 0x1,
		Zone5AlarmMemory: (data[3] >> 3) & 0x1,
		Zone6Faulted:     (data[3] >> 4) & 0x1,
		Zone6Bypass:      (data[3] >> 5) & 0x1,
		Zone6Trouble:     (data[3] >> 6) & 0x1,
		Zone6AlarmMemory: (data[3] >> 7) & 0x1,

		Zone78:           data[4],
		Zone7Faulted:     (data[4] >> 0) & 0x1,
		Zone7Bypass:      (data[4] >> 1) & 0x1,
		Zone7Trouble:     (data[4] >> 2) & 0x1,
		Zone7AlarmMemory: (data[4] >> 3) & 0x1,
		Zone8Faulted:     (data[4] >> 4) & 0x1,
		Zone8Bypass:      (data[4] >> 5) & 0x1,
		Zone8Trouble:     (data[4] >> 6) & 0x1,
		Zone8AlarmMemory: (data[4] >> 7) & 0x1,

		Zone910:           data[5],
		Zone9Faulted:      (data[5] >> 0) & 0x1,
		Zone9Bypass:       (data[5] >> 1) & 0x1,
		Zone9Trouble:      (data[5] >> 2) & 0x1,
		Zone9AlarmMemory:  (data[5] >> 3) & 0x1,
		Zone10Faulted:     (data[5] >> 4) & 0x1,
		Zone10Bypass:      (data[5] >> 5) & 0x1,
		Zone10Trouble:     (data[5] >> 6) & 0x1,
		Zone10AlarmMemory: (data[5] >> 7) & 0x1,

		Zone1112:          data[6],
		Zone11Faulted:     (data[6] >> 0) & 0x1,
		Zone11Bypass:      (data[6] >> 1) & 0x1,
		Zone11Trouble:     (data[6] >> 2) & 0x1,
		Zone11AlarmMemory: (data[6] >> 3) & 0x1,
		Zone12Faulted:     (data[6] >> 4) & 0x1,
		Zone12Bypass:      (data[6] >> 5) & 0x1,
		Zone12Trouble:     (data[6] >> 6) & 0x1,
		Zone12AlarmMemory: (data[6] >> 7) & 0x1,

		Zone1314:          data[7],
		Zone13Faulted:     (data[7] >> 0) & 0x1,
		Zone13Bypass:      (data[7] >> 1) & 0x1,
		Zone13Trouble:     (data[7] >> 2) & 0x1,
		Zone13AlarmMemory: (data[7] >> 3) & 0x1,
		Zone14Faulted:     (data[7] >> 4) & 0x1,
		Zone14Bypass:      (data[7] >> 5) & 0x1,
		Zone14Trouble:     (data[7] >> 6) & 0x1,
		Zone14AlarmMemory: (data[7] >> 7) & 0x1,

		Zone1516:          data[8],
		Zone15Faulted:     (data[8] >> 0) & 0x1,
		Zone15Bypass:      (data[8] >> 1) & 0x1,
		Zone15Trouble:     (data[8] >> 2) & 0x1,
		Zone15AlarmMemory: (data[8] >> 3) & 0x1,
		Zone16Faulted:     (data[8] >> 4) & 0x1,
		Zone16Bypass:      (data[8] >> 5) & 0x1,
		Zone16Trouble:     (data[8] >> 6) & 0x1,
		Zone16AlarmMemory: (data[8] >> 7) & 0x1,
	}

	return message, nil
}

func (m *ZonesSnapshotMessage) Number() byte {
	return ZonesSnapshotMessageNumber
}

func (m *ZonesSnapshotMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Offset: %08b
                    Zone 1 &amp; 2: %08b
                      Zone 1 Faulted: %01b
                      Zone 1 Bypass: %01b
                      Zone 1 Trouble: %01b
                      Zone 1 Alarm Memory: %01b
                      Zone 2 Faulted: %01b
                      Zone 2 Bypass: %01b
                      Zone 2 Trouble: %01b
                      Zone 2 Alarm Memory: %01b
                    Zone 3 &amp; 4: %08b
                      Zone 3 Faulted: %01b
                      Zone 3 Bypass: %01b
                      Zone 3 Trouble: %01b
                      Zone 3 Alarm Memory: %01b
                      Zone 4 Faulted: %01b
                      Zone 4 Bypass: %01b
                      Zone 4 Trouble: %01b
                      Zone 4 Alarm Memory: %01b
                    Zone 5 &amp; 6: %08b
                      Zone 5 Faulted: %01b
                      Zone 5 Bypass: %01b
                      Zone 5 Trouble: %01b
                      Zone 5 Alarm Memory: %01b
                      Zone 6 Faulted: %01b
                      Zone 6 Bypass: %01b
                      Zone 6 Trouble: %01b
                      Zone 6 Alarm Memory: %01b
                    Zone 7 &amp; 8: %08b
                      Zone 7 Faulted: %01b
                      Zone 7 Bypass: %01b
                      Zone 7 Trouble: %01b
                      Zone 7 Alarm Memory: %01b
                      Zone 8 Faulted: %01b
                      Zone 8 Bypass: %01b
                      Zone 8 Trouble: %01b
                      Zone 8 Alarm Memory: %01b
                    Zone 9 &amp; 10: %08b
                      Zone 9 Faulted: %01b
                      Zone 9 Bypass: %01b
                      Zone 9 Trouble: %01b
                      Zone 9 Alarm Memory: %01b
                      Zone 10 Faulted: %01b
                      Zone 10 Bypass: %01b
                      Zone 10 Trouble: %01b
                      Zone 10 Alarm Memory: %01b
                    Zone 11 &amp; 12: %08b
                      Zone 11 Faulted: %01b
                      Zone 11 Bypass: %01b
                      Zone 11 Trouble: %01b
                      Zone 11 Alarm Memory: %01b
                      Zone 12 Faulted: %01b
                      Zone 12 Bypass: %01b
                      Zone 12 Trouble: %01b
                      Zone 12 Alarm Memory: %01b
                    Zone 13 &amp; 14: %08b
                      Zone 13 Faulted: %01b
                      Zone 13 Bypass: %01b
                      Zone 13 Trouble: %01b
                      Zone 13 Alarm Memory: %01b
                      Zone 14 Faulted: %01b
                      Zone 14 Bypass: %01b
                      Zone 14 Trouble: %01b
                      Zone 14 Alarm Memory: %01b
                    Zone 15 &amp; 16: %08b
                      Zone 15 Faulted: %01b
                      Zone 15 Bypass: %01b
                      Zone 15 Trouble: %01b
                      Zone 15 Alarm Memory: %01b
                      Zone 16 Faulted: %01b
                      Zone 16 Bypass: %01b
                      Zone 16 Trouble: %01b
                      Zone 16 Alarm Memory: %01b`,
		"Zones Snapshot Message",
		strings.Repeat("-", len("Zones Snapshot Message")),
		m.ZoneOffset,
		m.Zone12,
		m.Zone1Faulted,
		m.Zone1Bypass,
		m.Zone1Trouble,
		m.Zone1AlarmMemory,
		m.Zone2Faulted,
		m.Zone2Bypass,
		m.Zone2Trouble,
		m.Zone2AlarmMemory,
		m.Zone34,
		m.Zone3Faulted,
		m.Zone3Bypass,
		m.Zone3Trouble,
		m.Zone3AlarmMemory,
		m.Zone4Faulted,
		m.Zone4Bypass,
		m.Zone4Trouble,
		m.Zone4AlarmMemory,
		m.Zone56,
		m.Zone5Faulted,
		m.Zone5Bypass,
		m.Zone5Trouble,
		m.Zone5AlarmMemory,
		m.Zone6Faulted,
		m.Zone6Bypass,
		m.Zone6Trouble,
		m.Zone6AlarmMemory,
		m.Zone78,
		m.Zone7Faulted,
		m.Zone7Bypass,
		m.Zone7Trouble,
		m.Zone7AlarmMemory,
		m.Zone8Faulted,
		m.Zone8Bypass,
		m.Zone8Trouble,
		m.Zone8AlarmMemory,
		m.Zone910,
		m.Zone9Faulted,
		m.Zone9Bypass,
		m.Zone9Trouble,
		m.Zone9AlarmMemory,
		m.Zone10Faulted,
		m.Zone10Bypass,
		m.Zone10Trouble,
		m.Zone10AlarmMemory,
		m.Zone1112,
		m.Zone11Faulted,
		m.Zone11Bypass,
		m.Zone11Trouble,
		m.Zone11AlarmMemory,
		m.Zone12Faulted,
		m.Zone12Bypass,
		m.Zone12Trouble,
		m.Zone12AlarmMemory,
		m.Zone1314,
		m.Zone13Faulted,
		m.Zone13Bypass,
		m.Zone13Trouble,
		m.Zone13AlarmMemory,
		m.Zone14Faulted,
		m.Zone14Bypass,
		m.Zone14Trouble,
		m.Zone14AlarmMemory,
		m.Zone1516,
		m.Zone15Faulted,
		m.Zone15Bypass,
		m.Zone15Trouble,
		m.Zone15AlarmMemory,
		m.Zone16Faulted,
		m.Zone16Bypass,
		m.Zone16Trouble,
		m.Zone16AlarmMemory,
	)
}
