// DO NOT EDIT -- This file was generated from templates/message.tmpl

package messages

import (
	"fmt"
	"strings"

	"github.com/vanstee/nx584/message"
)

const (
	PartitionStatusRequestName              = "Partition Status Request"
	PartitionStatusRequestLength            = 2
	PartitionStatusRequestNumber            = 38
	PartitionStatusRequestTransitionCapable = message.TransitionCapableNo
	PartitionStatusRequestAcknowledged      = message.AcknowledgedNo
)

type PartitionStatusRequest struct {
	*message.Base
}

func NewPartitionStatusRequest(length byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	if length != PartitionStatusRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PartitionStatusRequestLength, length)
	}

	message := &PartitionStatusRequest{
		Base: message.NewBase(length, acknowledgeRequired, data),
	}

	return message, nil
}

func (m *PartitionStatusRequest) Number() byte {
	return PartitionStatusRequestNumber
}

func (m *PartitionStatusRequest) String() string {
	return fmt.Sprintf(
		`%s
                    %s`,
		"Partition Status Request",
		strings.Repeat("-", len("Partition Status Request")),
	)
}
