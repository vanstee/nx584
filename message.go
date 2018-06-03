//go:generate go run cmd/messagegen/main.go -in data -templates templates -out .

package nx584

import "fmt"

const (
	TransitionCapableNo  TransitionCapable = iota
	TransitionCapableYes TransitionCapable = iota
)

const (
	AcknowledgedNo       Acknowledged = iota
	AcknowledgedPossible Acknowledged = iota
	AcknowledgedYes      Acknowledged = iota
)

type TransitionCapable byte
type Acknowledged byte
type NewMessageFunc func(byte, bool, []byte) (Message, error)

type Message interface {
	Length() byte
	Number() byte
	AcknowledgeRequired() bool
	Data() []byte
	String() string
}

func NewMessage(length byte, number byte, acknowledgeRequired bool, data []byte) (Message, error) {
	newMessageFunc, ok := messages[number]
	if !ok {
		return nil, fmt.Errorf("message number %v unknown", number)
	}
	return newMessageFunc(length, acknowledgeRequired, data)
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

	input := bytes[0 : length+1]
	checksum := bytes[length+1 : length+3]
	err := ValidateChecksum(checksum, input)
	if err != nil {
		return nil, err
	}

	data := bytes[2 : length+1]
	return NewMessage(length, number, acknowledgeRequired, data)
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
