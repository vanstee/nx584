// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	SendX10RequestName              = "Send X-10 Request"
	SendX10RequestLength            = 4
	SendX10RequestNumber            = 41
	SendX10RequestTransitionCapable = message.TransitionCapableNo
	SendX10RequestAcknowledged      = message.AcknowledgedYes
)

type SendX10Request struct {
	*message.Base
}

func NewSendX10Request(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != SendX10RequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", SendX10RequestLength, length)
	}

	message := &SendX10Request{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *SendX10Request) Number() byte {
	return SendX10RequestNumber
}

func (m *SendX10Request) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Send X-10 Request",
		strings.Repeat("-", len("Send X-10 Request")),
	)
}
