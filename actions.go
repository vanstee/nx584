package nx584

import (
	"errors"
	"time"

	"github.com/vanstee/nx584/message"
	"github.com/vanstee/nx584/messages"
)

const (
	ZoneCount = 16
)

const (
	ZoneStateNormal ZoneState = iota
	ZoneStateFault  ZoneState = iota
)

type ZoneState byte

type Zone struct {
	Number byte
	Name   string
	State  ZoneState
}

func (client *Client) ListZones() ([]*Zone, error) {
	zones := make([]*Zone, ZoneCount)
	for i := 0; i < ZoneCount; i++ {
		req, err := messages.NewZoneNameRequest(2, false, []byte{byte(i)})
		if err != nil {
			return nil, err
		}

		err = client.WriteMessage(req)
		if err != nil {
			return nil, err
		}

		resp, err := client.ReadMessage()
		if err != nil {
			return nil, err
		}
		if _, ok := resp.(*messages.ZoneNameMessage); !ok {
			return nil, errors.New("expected zone name message")
		}

		data := resp.Data()
		zone := &Zone{
			Number: data[0],
			Name:   string(data[1:]),
		}

		resp, err = messages.NewPositiveAcknowledge(1, false, []byte{})
		if err != nil {
			return nil, err
		}

		err = client.WriteMessage(resp)
		if err != nil {
			return nil, err
		}

		req, err = messages.NewZoneStatusRequest(2, false, []byte{byte(i)})
		if err != nil {
			return nil, err
		}

		err = client.WriteMessage(req)
		if err != nil {
			return nil, err
		}

		resp, err = client.ReadMessage()
		if err != nil {
			return nil, err
		}

		message, ok := resp.(*messages.ZoneStatusMessage)
		if !ok {
			return nil, errors.New("expected zone name message")
		}

		if message.Faulted == 0x1 {
			zone.State = ZoneStateFault
		} else {
			zone.State = ZoneStateNormal
		}

		resp, err = messages.NewPositiveAcknowledge(1, false, []byte{})
		if err != nil {
			return nil, err
		}

		err = client.WriteMessage(resp)
		if err != nil {
			return nil, err
		}

		zones[i] = zone
	}

	return zones, nil
}

func (client *Client) ZoneChanges() (<-chan *Zone, <-chan error) {
	messagesc := make(chan *Zone)
	errc := make(chan error)

	go func() {
		for {
			resp, err := client.ReadMessage()
			if err != nil {
				errc <- err
				break
			}

			if message, ok := resp.(*messages.ZoneStatusMessage); ok {
				zone := &Zone{Number: message.ZoneNumber}

				if message.Faulted == 0x1 {
					zone.State = ZoneStateFault
				} else {
					zone.State = ZoneStateNormal
				}

				messagesc <- zone
			}

			resp, err = messages.NewPositiveAcknowledge(1, false, []byte{})
			if err != nil {
				errc <- err
				break
			}

			err = client.WriteMessage(resp)
			if err != nil {
				errc <- err
				break
			}
		}

		close(messagesc)
		close(errc)
	}()

	return messagesc, errc
}

func (client *Client) Clear(timeout time.Duration) error {
	messagec := make(chan message.Message, 0)
	errc := make(chan error, 1)

	go func() {
		for {
			message, err := client.ReadMessage()
			if err != nil {
				errc <- err
				return
			}

			messagec <- message
		}
	}()

	handleStale := true
	for handleStale {
		select {
		case req := <-messagec:
			if req.AcknowledgeRequired() {
				resp, err := messages.NewPositiveAcknowledge(1, false, []byte{})
				if err != nil {
					return err
				}

				err = client.WriteMessage(resp)
				if err != nil {
					return err
				}
			}
		case err := <-errc:
			return err
		case <-time.After(timeout):
			handleStale = false
		}
	}

	return nil
}
