package nx584

const (
	ZoneStatusMessageName              = "Zone Status Message"
	ZoneStatusMessageLength            = 10
	ZoneStatusMessageNumber            = 0x04
	ZoneStatusMessageTransitionCapable = TransitionCapableYes
	ZoneStatusMessageAcknowledged      = AcknowledgedPossible
)

type ZoneStatusMessageType struct{}

type ZoneStatusMessage struct {
	length              byte
	acknowledgeRequired bool
	data                []byte
}

func NewZoneStatusMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	message := &ZoneStatusMessage{
		length:              length,
		acknowledgeRequired: acknowledgeRequired,
		data:                data,
	}
	return message, nil
}

func (t *ZoneStatusMessageType) Name() string {
	return ZoneStatusMessageName
}

func (t *ZoneStatusMessageType) Length() byte {
	return ZoneStatusMessageLength
}

func (t *ZoneStatusMessageType) Number() byte {
	return ZoneStatusMessageNumber
}

func (t *ZoneStatusMessageType) TransitionCapable() TransitionCapable {
	return ZoneStatusMessageTransitionCapable
}

func (t *ZoneStatusMessageType) Acknowledged() Acknowledged {
	return ZoneStatusMessageAcknowledged
}

func (m *ZoneStatusMessage) Type() MessageType {
	return &ZoneStatusMessageType{}
}

func (m *ZoneStatusMessage) Length() byte {
	return m.length
}

func (m *ZoneStatusMessage) Number() byte {
	return ZoneStatusMessageNumber
}

func (m *ZoneStatusMessage) AcknowledgeRequired() bool {
	return m.acknowledgeRequired
}

func (m *ZoneStatusMessage) Data() []byte {
	return m.data
}
