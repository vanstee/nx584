package nx584

import (
	"bufio"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/tarm/serial"
)

const (
	defaultBaud = 9600
)

type Config struct {
	Path   string
	Baud   int
	Logger *log.Logger
}

type Client struct {
	port    io.ReadWriteCloser
	scanner *bufio.Scanner
	logger  *log.Logger
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

	logger := config.Logger
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	client := &Client{
		port:    port,
		scanner: scanner,
		logger:  logger,
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

	client.logger.Printf("decoding message: %#v", bytes)

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
	client.logger.Printf("encoding message: %#v", bytes)

	ascii := strings.ToUpper(hex.EncodeToString(bytes))
	line := fmt.Sprintf("\n%s\r", ascii)
	_, err := client.port.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Close() error {
	return client.port.Close()
}
