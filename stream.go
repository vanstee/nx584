package nx584

import (
	"bytes"
	"log"
)

func ScanMessages(data []byte, atEOF bool) (int, []byte, error) {
	start := bytes.IndexByte(data, '\n')
	if start < 0 {
		return 0, nil, nil
	}

	end := bytes.IndexByte(data[start+1:], '\r')
	if end < 0 {
		return 0, nil, nil
	}

	if start > 0 {
		log.Printf("skipping unrecognized stream: %#v", data[0:start])
	}

	return start + end + 2, data[start+1 : start+end+1], nil
}
