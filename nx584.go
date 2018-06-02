package nx584

import (
	"bufio"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/tarm/serial"
)

const (
	defaultBaud = 9600
)

type Config struct {
	Path string
	Baud int
}

type Client struct {
	port    io.ReadWriteCloser
	scanner *bufio.Scanner
}

func Open(config *Config) (*Client, error) {
	baud := config.Baud
	if baud == 0 {
		baud = defaultBaud
	}

	port, err := serial.OpenPort(
		&serial.Config{
			Name: config.Path,
			Baud: baud,
		},
	)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(port)
	scanner.Split(ScanMessages)

	client := &Client{
		port:    port,
		scanner: scanner,
	}

	return client, nil
}

func (client *Client) ReadMessage() (Message, error) {
	more := client.scanner.Scan()
	if !more {
		return nil, io.EOF
	}

	ascii := client.scanner.Text()
	if len(ascii) == 0 {
		return nil, errors.New("message was empty")
	}

	bytes, err := hex.DecodeString(ascii)
	if err != nil {
		return nil, err
	}

	log.Printf("decoding message: %#v", bytes)

	message, err := DecodeMessage(bytes)
	if err != nil {
		return nil, err
	}

	if err := client.scanner.Err(); err != nil {
		return nil, err
	}

	return message, err
}

func (client *Client) WriteMessage(message Message) error {
	bytes := EncodeMessage(message)
	log.Printf("encoding message: %#v", bytes)

	ascii := strings.ToUpper(hex.EncodeToString(bytes))
	_, err := client.port.Write([]byte(ascii))
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Close() error {
	return client.port.Close()
}
