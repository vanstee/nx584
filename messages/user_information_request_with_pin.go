// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	UserInformationRequestWithPinName              = "User Information Request with PIN"
	UserInformationRequestWithPinLength            = 5
	UserInformationRequestWithPinNumber            = 50
	UserInformationRequestWithPinTransitionCapable = message.TransitionCapableNo
	UserInformationRequestWithPinAcknowledged      = message.AcknowledgedNo
)

type UserInformationRequestWithPin struct {
	*message.Base
}

func NewUserInformationRequestWithPin(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != UserInformationRequestWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", UserInformationRequestWithPinLength, length)
	}

	message := &UserInformationRequestWithPin{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *UserInformationRequestWithPin) Number() byte {
	return UserInformationRequestWithPinNumber
}

func (m *UserInformationRequestWithPin) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"User Information Request with PIN",
		strings.Repeat("-", len("User Information Request with PIN")),
	)
}
