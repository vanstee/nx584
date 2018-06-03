package nx584

const (
	PositiveAcknowledgeName              = "Positive Acknowledge"
	PositiveAcknowledgeLength            = 1
	PositiveAcknowledgeNumber            = 0x1D
	PositiveAcknowledgeTransitionCapable = TransitionCapableNo
	PositiveAcknowledgeAcknowledged      = AcknowledgedNo
)

type PositiveAcknowledge struct {
	*BaseMessage
	length              byte
	acknowledgeRequired bool
	data                []byte
}

func NewPositiveAcknowledge(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	message := &PositiveAcknowledge{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},
	}
	return message, nil
}

func (m *PositiveAcknowledge) Number() byte {
	return PositiveAcknowledgeNumber
}
