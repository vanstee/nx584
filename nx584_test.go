package nx584

import (
	"bufio"
	"strings"
	"testing"
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
	port := &FakePort{}
	input := strings.NewReader("0A8405015853E001000200249B")
	scanner := bufio.NewScanner(input)
	client := Client{port, scanner}

	req, err := client.ReadMessage()
	if err != nil {
		t.Errorf("client.ReadMessage() returned an error: %v", err)
	}
	if _, ok := req.(*ZoneStatusMessage); !ok {
		t.Errorf("client.ReadMessage() did not return a message of type ZoneStatusMessage")
	}

	resp, err := NewPositiveAcknowledge(1, false, []byte{})
	if err != nil {
		t.Errorf("NewPositiveAcknowledge(1, false, []byte{}) returned an error: %v", err)
	}

	err = client.WriteMessage(resp)
	if err != nil {
		t.Errorf("client.WriteMessage(resp) returned an error: %v", err)
	}
}
