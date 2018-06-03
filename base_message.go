package nx584

type BaseMessage struct {
	length              byte
	acknowledgeRequired bool
	data                []byte
}

func (m *BaseMessage) Length() byte {
	return m.length
}

func (m *BaseMessage) AcknowledgeRequired() bool {
	return m.acknowledgeRequired
}

func (m *BaseMessage) Data() []byte {
	return m.data
}
