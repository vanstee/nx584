// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	UserInformationReplyName              = "User Information Reply"
	UserInformationReplyLength            = 7
	UserInformationReplyNumber            = 18
	UserInformationReplyTransitionCapable = TransitionCapableNo
	UserInformationReplyAcknowledged      = AcknowledgedNo
)

type UserInformationReply struct {
	*BaseMessage
}

func NewUserInformationReply(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != UserInformationReplyLength {
		log.Printf("message length incorrect: expected %d, actual: %d", UserInformationReplyLength, length)
	}

	message := &UserInformationReply{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
