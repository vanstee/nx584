package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/tarm/serial"
)

type Frame struct {
	Length      byte
	MessageType byte
	Data        []byte
}

func main() {
	c := &serial.Config{Name: "/Users/vanstee/Code/go/src/github.com/vanstee/nx584/dev/ttyUSB0", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Println(err)
	}

	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Opened port to serial device: /dev/ttyUSB0\n")

	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		log.Printf("Read: %v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
