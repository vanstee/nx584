//go:generate go run cmd/messagegen/main.go -in data -templates templates -out messages

package nx584

import (
	"io"
	"log"
	"os"

	"github.com/tarm/serial"
	"github.com/vanstee/nx584/message"
	"github.com/vanstee/nx584/stream"
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
	logger  *log.Logger
	decoder *stream.Decoder
	encoder *stream.Encoder
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

	logger := config.Logger
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	client := &Client{
		port:    port,
		logger:  logger,
		decoder: stream.NewDecoder(port),
		encoder: stream.NewEncoder(port),
	}

	return client, nil
}

func (client *Client) ReadMessage() (message.Message, error) {
	var m message.Message
	err := client.decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (client *Client) WriteMessage(m message.Message) error {
	return client.encoder.Encode(m)
}

func (client *Client) Close() error {
	return client.port.Close()
}
