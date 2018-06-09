package nx584

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"
	"testing"

	"github.com/vanstee/nx584/messages"
	"github.com/vanstee/nx584/stream"
)

type FakePort struct{}

func (port *FakePort) Read(p []byte) (int, error) {
	return 0, nil
}

func (port *FakePort) Write([]byte) (int, error) {
	return 0, nil
}

func (port *FakePort) Close() error {
	return nil
}

func TestZoneStatusMessage(t *testing.T) {
	input := "\n0A8A52B9F7000006021004B494\r"
	output := "\n011D1E1F\r"

	port := &FakePort{}
	var buffer bytes.Buffer
	client := Client{
		port,
		log.New(ioutil.Discard, "", 0),
		stream.NewDecoder(strings.NewReader(input)),
		stream.NewEncoder(&buffer),
	}

	req, err := client.ReadMessage()
	if err != nil {
		t.Errorf("client.ReadMessage() returned an error: %v", err)
	}
	if _, ok := req.(*messages.LogEventMessage); !ok {
		t.Errorf("client.ReadMessage() did not return a message of type LogEventMessage")
	}

	resp, err := messages.NewPositiveAcknowledge(1, false, []byte{})
	if err != nil {
		t.Errorf("NewPositiveAcknowledge(1, false, []byte{}) returned an error: %v", err)
	}

	err = client.WriteMessage(resp)
	if err != nil {
		t.Errorf("client.WriteMessage(resp) returned an error: %v", err)
	}

	if buffer.String() != output {
		t.Errorf("unexpected data written to port: expected %#v, actual %#v", output, buffer.String())
	}
}
