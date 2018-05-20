package nx584

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"

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

	client := &Client{
		port:    port,
		scanner: bufio.NewScanner(port),
	}

	return client, nil
}

func (client *Client) ReadMessage() (Message, error) {
	more := client.scanner.Scan()
	if !more {
		return nil, io.EOF
	}

	ascii := client.scanner.Bytes()
	if len(ascii) == 0 {
		return nil, errors.New("message was empty")
	}

	bytes, err := asciiHexToBytes(ascii)
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

	ascii, err := bytesToAsciiHex(bytes)
	if err != nil {
		return err
	}

	_, err = client.port.Write(ascii)
	if err != nil {
		return err
	}

	return nil
}

func (client *Client) Close() error {
	return client.port.Close()
}

func asciiHexToBytes(ascii []byte) ([]byte, error) {
	bytes := make([]byte, len(ascii)/2)
	for n, _ := range bytes {
		i := n * 2
		higher, err := asciiHexToByte(ascii[i])
		if err != nil {
			return nil, err
		}
		lower, err := asciiHexToByte(ascii[i+1])
		if err != nil {
			return nil, err
		}
		bytes[n] = (higher << 4) + lower
	}
	return bytes, nil
}

func asciiHexToByte(ascii byte) (byte, error) {
	if ascii >= 48 && ascii <= 57 {
		return ascii - 48, nil
	}
	if ascii >= 65 && ascii <= 70 {
		return ascii - 65 + 10, nil
	}
	return 0, fmt.Errorf("byte %#v not in ascii hex range", ascii)
}

func bytesToAsciiHex(bytes []byte) ([]byte, error) {
	ascii := make([]byte, len(bytes)*2)
	for i, b := range bytes {
		higher, err := byteToAsciiHex(b >> 4)
		if err != nil {
			return nil, err
		}
		ascii[i*2] = higher
		lower, err := byteToAsciiHex(b & 0x0F)
		if err != nil {
			return nil, err
		}
		ascii[(i*2)+1] = lower
	}
	return ascii, nil
}

func byteToAsciiHex(b byte) (byte, error) {
	if b >= 0 && b <= 9 {
		return b + 48, nil
	}
	if b >= 10 && b <= 15 {
		return b + 65, nil
	}
	return 0, fmt.Errorf("byte %#v not in ascii hex range", b)
}
