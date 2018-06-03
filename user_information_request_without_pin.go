// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	UserInformationRequestWithoutPinName              = "User Information Request without PIN"
	UserInformationRequestWithoutPinLength            = 2
	UserInformationRequestWithoutPinNumber            = 51
	UserInformationRequestWithoutPinTransitionCapable = TransitionCapableNo
	UserInformationRequestWithoutPinAcknowledged      = AcknowledgedNo
)

type UserInformationRequestWithoutPin struct {
	*BaseMessage
}

func NewUserInformationRequestWithoutPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != UserInformationRequestWithoutPinLength {
		log.Printf("message length incorrect: expected %d, actual: %d", UserInformationRequestWithoutPinLength, length)
	}

	message := &UserInformationRequestWithoutPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
