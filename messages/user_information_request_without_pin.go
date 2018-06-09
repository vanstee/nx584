// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	UserInformationRequestWithoutPinName              = "User Information Request without PIN"
	UserInformationRequestWithoutPinLength            = 2
	UserInformationRequestWithoutPinNumber            = 51
	UserInformationRequestWithoutPinTransitionCapable = message.TransitionCapableNo
	UserInformationRequestWithoutPinAcknowledged      = message.AcknowledgedNo
)

type UserInformationRequestWithoutPin struct {
	*message.Base
}

func NewUserInformationRequestWithoutPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != UserInformationRequestWithoutPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", UserInformationRequestWithoutPinLength, length)
	}

	message := &UserInformationRequestWithoutPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *UserInformationRequestWithoutPin) Number() byte {
	return UserInformationRequestWithoutPinNumber
}

func (m *UserInformationRequestWithoutPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"User Information Request without PIN",
		strings.Repeat("-", len("User Information Request without PIN")),
	)
}
