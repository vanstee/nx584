// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	ZoneNameMessageName              = "Zone Name Message"
	ZoneNameMessageLength            = 18
	ZoneNameMessageNumber            = 3
	ZoneNameMessageTransitionCapable = message.TransitionCapableNo
	ZoneNameMessageAcknowledged      = message.AcknowledgedNo
)

type ZoneNameMessage struct {
	*message.Base
	ZoneNumber          byte
	ZoneNameCharacter1  byte
	ZoneNameCharacter2  byte
	ZoneNameCharacter3  byte
	ZoneNameCharacter4  byte
	ZoneNameCharacter5  byte
	ZoneNameCharacter6  byte
	ZoneNameCharacter7  byte
	ZoneNameCharacter8  byte
	ZoneNameCharacter9  byte
	ZoneNameCharacter10 byte
	ZoneNameCharacter11 byte
	ZoneNameCharacter12 byte
	ZoneNameCharacter13 byte
	ZoneNameCharacter14 byte
	ZoneNameCharacter15 byte
	ZoneNameCharacter16 byte
}

func NewZoneNameMessage(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != ZoneNameMessageLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", ZoneNameMessageLength, length)
	}

	message := &ZoneNameMessage{
		Base: message.NewBase(length, acknowledgeRequired, data),

		ZoneNumber: data[0],

		ZoneNameCharacter1: data[1],

		ZoneNameCharacter2: data[2],

		ZoneNameCharacter3: data[3],

		ZoneNameCharacter4: data[4],

		ZoneNameCharacter5: data[5],

		ZoneNameCharacter6: data[6],

		ZoneNameCharacter7: data[7],

		ZoneNameCharacter8: data[8],

		ZoneNameCharacter9: data[9],

		ZoneNameCharacter10: data[10],

		ZoneNameCharacter11: data[11],

		ZoneNameCharacter12: data[12],

		ZoneNameCharacter13: data[13],

		ZoneNameCharacter14: data[14],

		ZoneNameCharacter15: data[15],

		ZoneNameCharacter16: data[16],
	}

	return message, nil
}

func (m *ZoneNameMessage) Number() byte {
	return ZoneNameMessageNumber
}

func (m *ZoneNameMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Zone Number: %08b
                    Zone Name Character 1: %08b
                    Zone Name Character 2: %08b
                    Zone Name Character 3: %08b
                    Zone Name Character 4: %08b
                    Zone Name Character 5: %08b
                    Zone Name Character 6: %08b
                    Zone Name Character 7: %08b
                    Zone Name Character 8: %08b
                    Zone Name Character 9: %08b
                    Zone Name Character 10: %08b
                    Zone Name Character 11: %08b
                    Zone Name Character 12: %08b
                    Zone Name Character 13: %08b
                    Zone Name Character 14: %08b
                    Zone Name Character 15: %08b
                    Zone Name Character 16: %08b`,
		"Zone Name Message",
		strings.Repeat("-", len("Zone Name Message")),
		m.ZoneNumber,
		m.ZoneNameCharacter1,
		m.ZoneNameCharacter2,
		m.ZoneNameCharacter3,
		m.ZoneNameCharacter4,
		m.ZoneNameCharacter5,
		m.ZoneNameCharacter6,
		m.ZoneNameCharacter7,
		m.ZoneNameCharacter8,
		m.ZoneNameCharacter9,
		m.ZoneNameCharacter10,
		m.ZoneNameCharacter11,
		m.ZoneNameCharacter12,
		m.ZoneNameCharacter13,
		m.ZoneNameCharacter14,
		m.ZoneNameCharacter15,
		m.ZoneNameCharacter16,
	)
}
