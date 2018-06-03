package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/vanstee/nx584"
)

const (
	MaxZones = 16

	StateUnknown   = "UNKNOWN"
	StateNormal    = "NORMAL"
	StateFaulted   = "FAULT"
	StateTampered  = "TAMPER"
	StateBypassed  = "BYPASS"
	StateInhibited = "INHIBIT"
)

type Zone struct {
	Name  string
	State string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("path required")
	}
	path := os.Args[1]

	baud := 9600
	if len(os.Args) > 2 {
		var err error
		baud, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	config := &nx584.Config{
		Path:   path,
		Baud:   baud,
		Logger: log.New(ioutil.Discard, "", log.LstdFlags),
	}

	client, err := nx584.Open(config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	messages := make(chan nx584.Message, 0)
	errc := make(chan error, 1)

	go func() {
		for {
			message, err := client.ReadMessage()
			if err != nil {
				errc <- err
				return
			}

			messages <- message
		}
	}()

	handleStale := true
	for handleStale {
		select {
		case req := <-messages:
			if req.AcknowledgeRequired() {
				resp, err := nx584.NewPositiveAcknowledge(1, false, []byte{})
				if err != nil {
					log.Fatal(err)
				}

				err = client.WriteMessage(resp)
				if err != nil {
					log.Fatal(err)
				}
			}
		case err := <-errc:
			log.Fatal(err)
		case <-time.After(3 * time.Second):
			handleStale = false
		}
	}

	zones := make(map[byte]*Zone)
	for i := 0; i < MaxZones; i++ {
		req, err := nx584.NewZoneNameRequest(2, false, []byte{byte(i)})
		if err != nil {
			log.Fatal(err)
		}

		err = client.WriteMessage(req)
		if err != nil {
			log.Fatal(err)
		}

		select {
		case resp := <-messages:
			if _, ok := resp.(*nx584.ZoneNameMessage); !ok {
				log.Fatal("expected zone name message")
			}

			zones[byte(i)] = &Zone{
				Name:  strings.TrimSpace(string(resp.Data()[1:])),
				State: StateUnknown,
			}
		case err := <-errc:
			if err != nil {
				log.Fatal(err)
			}
		}

		req, err = nx584.NewZoneStatusRequest(2, false, []byte{byte(i)})
		if err != nil {
			log.Fatal(err)
		}

		err = client.WriteMessage(req)
		if err != nil {
			log.Fatal(err)
		}

		select {
		case resp := <-messages:
			status, ok := resp.(*nx584.ZoneStatusMessage)
			if !ok {
				log.Fatal("expected zone status message")
			}

			zone := zones[status.ZoneNumber]

			if status.Faulted == 0x1 {
				zone.State = StateFaulted
			} else if status.Tampered == 0x1 {
				zone.State = StateTampered
			} else if status.Bypassed == 0x1 {
				zone.State = StateBypassed
			} else if status.Inhibited == 0x1 {
				zone.State = StateInhibited
			} else {
				zone.State = StateNormal
			}
		case err := <-errc:
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	for _, zone := range zones {
		log.Printf("%17s %s", fmt.Sprintf("%s:", zone.Name), zone.State)
	}

	for {
		select {
		case req := <-messages:
			switch message := req.(type) {
			case *nx584.ZoneStatusMessage:
				zone := zones[message.ZoneNumber]

				if message.Faulted == 0x1 {
					zone.State = StateFaulted
				} else if message.Tampered == 0x1 {
					zone.State = StateTampered
				} else if message.Bypassed == 0x1 {
					zone.State = StateBypassed
				} else if message.Inhibited == 0x1 {
					zone.State = StateInhibited
				} else {
					zone.State = StateNormal
				}

				log.Printf("%17s %s", fmt.Sprintf("%s:", zone.Name), zone.State)
			case *nx584.PartitionStatusMessage:
				// ignore
			default:
				log.Printf("unexpected message number %s", req.Number())
			}

			if req.AcknowledgeRequired() {
				resp, err := nx584.NewPositiveAcknowledge(1, false, []byte{})
				if err != nil {
					log.Fatal(err)
				}

				err = client.WriteMessage(resp)
				if err != nil {
					log.Fatal(err)
				}
			}
		case err := <-errc:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
