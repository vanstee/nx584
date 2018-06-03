package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vanstee/nx584"
)

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
		Path: path,
		Baud: baud,
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

	log.Printf("handling stale messages")
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
			log.Printf("timed out waiting for messages")
			handleStale = false
		}
	}
	log.Printf("done handling stale messages")

	req, err := nx584.NewZonesSnapshotRequest(2, false, []byte{0x0})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("sending zones snapshot request")
	err = client.WriteMessage(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("waiting for zones snapshot message")
	select {
	case resp := <-messages:
		if _, ok := resp.(*nx584.ZonesSnapshotMessage); !ok {
			log.Fatal("expected zones snapshot message")
		}

		log.Printf(resp.String())
	case err := <-errc:
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < 16; i++ {
		req, err := nx584.NewZoneNameRequest(2, false, []byte{byte(i)})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("sending zone name request")
		err = client.WriteMessage(req)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("waiting for zone name message")
		select {
		case resp := <-messages:
			if _, ok := resp.(*nx584.ZoneNameMessage); !ok {
				log.Fatal("expected zone name message")
			}

			data := resp.Data()
			number := data[0]
			name := data[1:]

			log.Printf(
				`zone %d
                    ------
                    %v
                    %v`,
				number,
				string(name),
				name,
			)
		case err := <-errc:
			if err != nil {
				log.Fatal(err)
			}
		}

		req, err = nx584.NewZoneStatusRequest(2, false, []byte{byte(i)})
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("sending zone status request")
		err = client.WriteMessage(req)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("waiting for zone status message")
		select {
		case resp := <-messages:
			if _, ok := resp.(*nx584.ZoneStatusMessage); !ok {
				log.Fatal("expected zone status message")
			}

			log.Printf(resp.String())
		case err := <-errc:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
