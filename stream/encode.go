package stream

import (
	"encoding/hex"
	"fmt"
	"io"
	"strings"

	"github.com/vanstee/nx584/message"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer: writer,
	}
}

func (encoder *Encoder) Encode(m message.Message) error {
	length := m.Length()
	bytes := make([]byte, 0, int(length+3))
	bytes = append(bytes, length)

	number := m.Number()
	acknowledgeRequired := message.AcknowledgedNo
	if m.AcknowledgeRequired() {
		acknowledgeRequired = message.AcknowledgedYes
	}
	bytes = append(bytes, (number&0x3F)+(byte(acknowledgeRequired)<<7))
	bytes = append(bytes, m.Data()[0:length-1]...)

	checksum := Fletcher(bytes)
	bytes = append(bytes, checksum...)

	ascii := strings.ToUpper(hex.EncodeToString(bytes))
	line := fmt.Sprintf("\n%s\r", ascii)
	encoder.writer.Write([]byte(line))

	return nil
}
