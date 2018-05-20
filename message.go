package nx584

import (
	"bytes"
	"errors"
	"fmt"
)

type TransitionCapable byte
type Acknowledged byte

const (
	TransitionCapableNo  TransitionCapable = iota
	TransitionCapableYes TransitionCapable = iota

	AcknowledgedNo       Acknowledged = iota
	AcknowledgedPossible Acknowledged = iota
	AcknowledgedYes      Acknowledged = iota

	PartitionStatusMessageName              = "Parition Status Message"
	PartitionStatusMessageNumber            = 6
	PartitionStatusMessageLength            = 9
	PartitionStatusMessageTransitionCapable = TransitionCapableYes
	PartitionStatusMessageAcknowledged      = AcknowledgedPossible

	PositiveAcknowledgeName              = "Positive Acknowledge"
	PositiveAcknowledgeNumber            = 29
	PositiveAcknowledgeLength            = 1
	PositiveAcknowledgeTransitionCapable = TransitionCapableNo
	PositiveAcknowledgeAcknowledged      = AcknowledgedNo
)

type Message interface {
	Length() byte
	Number() byte
	AcknowledgeRequired() bool
	Data() []byte
}

type BaseMessage struct {
	Length              byte
	Number              byte
	AcknowledgeRequired bool
	Data                []byte
}

type PartitionStatusMessage struct {
	*BaseMessage
}

type PositiveAcknowledge struct {
	*BaseMessage
}

func (message *PartitionStatusMessage) Length() byte {
	return PartitionStatusMessageLength
}

func (message *PartitionStatusMessage) Number() byte {
	return PartitionStatusMessageNumber
}

func (message *PartitionStatusMessage) AcknowledgeRequired() bool {
	return message.BaseMessage.AcknowledgeRequired
}

func (message *PartitionStatusMessage) Data() []byte {
	return message.BaseMessage.Data
}

func (message *PositiveAcknowledge) Length() byte {
	return PositiveAcknowledgeLength
}

func (message *PositiveAcknowledge) Number() byte {
	return PositiveAcknowledgeNumber
}

func (message *PositiveAcknowledge) AcknowledgeRequired() bool {
	return message.BaseMessage.AcknowledgeRequired
}

func (message *PositiveAcknowledge) Data() []byte {
	return []byte{}
}

func DecodeMessage(bytes []byte) (Message, error) {
	if len(bytes) < 2 {
		return nil, fmt.Errorf("message data malformed: message shorter than minimum length, expected: 2, actual: %v", len(bytes))
	}

	length := bytes[0]
	number := bytes[1] & 0x3F
	acknowledgeRequired := bytes[1]>>7 == 0x1
	if len(bytes) < int(length+3) {
		return nil, fmt.Errorf("message data malformed: message shorter than length specified, expected: %v, actual: %v", int(length+3), len(bytes))
	}

	data := bytes[0 : length+1]
	checksum := bytes[length+1 : length+3]
	if !IsValidChecksum(checksum, data) {
		return nil, fmt.Errorf("message checksum invalid, expected: %v, actual: %v", checksum, Fletcher(data))
	}

	base := &BaseMessage{
		Length:              length,
		Number:              number,
		AcknowledgeRequired: acknowledgeRequired,
		Data:                data[2 : length+1],
	}

	switch number {
	case PartitionStatusMessageNumber:
		if length != PartitionStatusMessageLength {
			return nil, errors.New("message length invalid")
		}
		if acknowledgeRequired && PartitionStatusMessageAcknowledged == AcknowledgedNo {
			return nil, errors.New("message acknowledge required invalid")
		}
		if !acknowledgeRequired && PartitionStatusMessageAcknowledged == AcknowledgedYes {
			return nil, errors.New("message acknowledge required invalid")
		}
		return &PartitionStatusMessage{base}, nil
	default:
		return nil, errors.New("message number unknown")
	}
}

func EncodeMessage(message Message) []byte {
	length := message.Length()
	bytes := make([]byte, 0, int(length+3))
	bytes = append(bytes, length)

	number := message.Number()
	acknowledgeRequired := AcknowledgedNo
	if message.AcknowledgeRequired() {
		acknowledgeRequired = AcknowledgedYes
	}
	bytes = append(bytes, (number&0x3F)+(byte(acknowledgeRequired)<<7))
	bytes = append(bytes, message.Data()[0:length-1]...)

	checksum := Fletcher(bytes)
	bytes = append(bytes, checksum...)

	return bytes
}

func IsValidChecksum(checksum []byte, data []byte) bool {
	return bytes.Equal(checksum, Fletcher(data))
}

func Fletcher(bytes []byte) []byte {
	var sum1, sum2 byte
	for _, b := range bytes {
		if 0xFF-sum1 < b {
			sum1 += 1
		}
		sum1 += b
		if sum1 == 0xFF {
			sum1 = 0
		}
		if 0xFF-sum2 < sum1 {
			sum2 += 1
		}
		sum2 += sum1
		if sum2 == 0xFF {
			sum2 = 0
		}
	}
	return []byte{sum1, sum2}
}
