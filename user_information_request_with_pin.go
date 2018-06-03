// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	UserInformationRequestWithPinName              = "User Information Request with PIN"
	UserInformationRequestWithPinLength            = 5
	UserInformationRequestWithPinNumber            = 50
	UserInformationRequestWithPinTransitionCapable = TransitionCapableNo
	UserInformationRequestWithPinAcknowledged      = AcknowledgedNo
)

type UserInformationRequestWithPin struct {
	*BaseMessage
}

func NewUserInformationRequestWithPin(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != UserInformationRequestWithPinLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", UserInformationRequestWithPinLength, length)
	}

	message := &UserInformationRequestWithPin{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
