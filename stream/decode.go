package stream

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/vanstee/nx584/message"
	"github.com/vanstee/nx584/messages"
)

type Decoder struct {
	reader  io.Reader
	scanner *bufio.Scanner
}

func NewDecoder(reader io.Reader) *Decoder {
	scanner := bufio.NewScanner(reader)
	scanner.Split(ScanMessages)

	return &Decoder{
		reader:  reader,
		scanner: scanner,
	}
}

func (decoder *Decoder) Decode(m *message.Message) error {
	more := decoder.scanner.Scan()
	if !more {
		return io.EOF
	}

	ascii := decoder.scanner.Text()
	if ascii == "" {
		return errors.New("message was empty")
	}

	bytes, err := hex.DecodeString(ascii)
	if err != nil {
		return err
	}

	if len(bytes) < 2 {
		return fmt.Errorf("message data malformed: message shorter than minimum length, expected: 2, actual: %v", len(bytes))
	}

	length := bytes[0]
	number := bytes[1] & 0x3F
	acknowledgeRequired := bytes[1]>>7 == 0x1
	if len(bytes) < int(length+3) {
		return fmt.Errorf("message data malformed: message shorter than length specified, expected: %v, actual: %v", int(length+3), len(bytes))
	}

	input := bytes[0 : length+1]
	checksum := bytes[length+1 : length+3]
	err = ValidateChecksum(checksum, input)
	if err != nil {
		return err
	}

	data := bytes[2 : length+1]
	decoded, err := messages.New(length, number, acknowledgeRequired, data)
	if err != nil {
		return err
	}

	*m = decoded

	if err := decoder.scanner.Err(); err != nil {
		return err
	}

	return nil
}
