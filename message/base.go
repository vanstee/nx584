package message

type Base struct {
	length              byte
	acknowledgeRequired bool
	data                []byte
}

func NewBase(length byte, acknowledgeRequired bool, data []byte) *Base {
	return &Base{
		length:              length,
		acknowledgeRequired: acknowledgeRequired,
		data:                data,
	}
}

func (m *Base) Length() byte {
	return m.length
}

func (m *Base) AcknowledgeRequired() bool {
	return m.acknowledgeRequired
}

func (m *Base) Data() []byte {
	return m.data
}
