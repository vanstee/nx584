// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	UserInformationReplyName              = "User Information Reply"
	UserInformationReplyLength            = 7
	UserInformationReplyNumber            = 18
	UserInformationReplyTransitionCapable = message.TransitionCapableNo
	UserInformationReplyAcknowledged      = message.AcknowledgedNo
)

type UserInformationReply struct {
	*message.Base
}

func NewUserInformationReply(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != UserInformationReplyLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", UserInformationReplyLength, length)
	}

	message := &UserInformationReply{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *UserInformationReply) Number() byte {
	return UserInformationReplyNumber
}

func (m *UserInformationReply) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"User Information Reply",
		strings.Repeat("-", len("User Information Reply")),
	)
}
