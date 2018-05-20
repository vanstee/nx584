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
			log.Printf("error reading message: %v", err)
		}

		var resp nx584.Message
		switch message := req.(type) {
		case *nx584.PartitionStatusMessage:
			resp = &nx584.PositiveAcknowledge{
				&nx584.BaseMessage{
					Length: 1,
					Number: message.Number(),
				},
			}
		default:
			log.Fatal("message type unknown")
		}

		err = client.WriteMessage(resp)
		if err != nil {
			log.Fatal(err)
		}
	}
}
