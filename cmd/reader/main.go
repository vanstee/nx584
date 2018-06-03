package main

import (
	"io"
	"log"
	"os"
	"strconv"

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

	for {
		req, err := client.ReadMessage()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("error reading message: %v", err)
		}

		switch req.(type) {
		case *nx584.ZoneStatusMessage:
			log.Printf("%s", req)

			resp, err := nx584.NewPositiveAcknowledge(1, false, []byte{})
			if err != nil {
				log.Fatal(err)
			}

			err = client.WriteMessage(resp)
			if err != nil {
				log.Fatal(err)
			}
		default:
			log.Fatal("message type unknown")
		}
	}
}
