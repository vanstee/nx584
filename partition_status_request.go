// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"strings"
)

const (
	PartitionStatusRequestName              = "Partition Status Request"
	PartitionStatusRequestLength            = 2
	PartitionStatusRequestNumber            = 38
	PartitionStatusRequestTransitionCapable = TransitionCapableNo
	PartitionStatusRequestAcknowledged      = AcknowledgedNo
)

type PartitionStatusRequest struct {
	*BaseMessage
}

func NewPartitionStatusRequest(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PartitionStatusRequestLength {
		return nil, fmt.Errorf("message length incorrect: expected %d, actual: %d", PartitionStatusRequestLength, length)
	}

	message := &PartitionStatusRequest{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
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
