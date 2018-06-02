package nx584

const (
	PositiveAcknowledgeName              = "Positive Acknowledge"
	PositiveAcknowledgeLength            = 1
	PositiveAcknowledgeNumber            = 0x1D
	PositiveAcknowledgeTransitionCapable = TransitionCapableNo
	PositiveAcknowledgeAcknowledged      = AcknowledgedNo
)

type PositiveAcknowledgeType struct{}

type PositiveAcknowledge struct {
	length              byte
	acknowledgeRequired bool
	data                []byte
}

func NewPositiveAcknowledge(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	message := &PositiveAcknowledge{
		length:              length,
		acknowledgeRequired: acknowledgeRequired,
		data:                data,
	}
	return message, nil
}

func (t *PositiveAcknowledgeType) Name() string {
	return PositiveAcknowledgeName
}

func (t *PositiveAcknowledgeType) Length() byte {
	return PositiveAcknowledgeLength
}

func (t *PositiveAcknowledgeType) Number() byte {
	return PositiveAcknowledgeNumber
}

func (t *PositiveAcknowledgeType) TransitionCapable() TransitionCapable {
	return PositiveAcknowledgeTransitionCapable
}

func (t *PositiveAcknowledgeType) Acknowledged() Acknowledged {
	return PositiveAcknowledgeAcknowledged
}

func (m *PositiveAcknowledge) Type() MessageType {
	return &PositiveAcknowledgeType{}
}

func (m *PositiveAcknowledge) Length() byte {
	return m.length
}

func (m *PositiveAcknowledge) Number() byte {
	return PositiveAcknowledgeNumber
}

func (m *PositiveAcknowledge) AcknowledgeRequired() bool {
	return m.acknowledgeRequired
}

func (m *PositiveAcknowledge) Data() []byte {
	return m.data
}
