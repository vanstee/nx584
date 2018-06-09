package stream

import (
	"bytes"
	"fmt"
)

func ValidateChecksum(expected []byte, input []byte) error {
	actual := Fletcher(input)
	if !bytes.Equal(expected, actual) {
		return fmt.Errorf("message checksum invalid, expected: %v, actual: %v", expected, actual)
	}
	return nil
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
