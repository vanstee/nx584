//go:generate go run cmd/messagegen/main.go -in data -templates templates -out messages

package message

const (
	TransitionCapableNo  TransitionCapable = iota
	TransitionCapableYes TransitionCapable = iota
)

const (
	AcknowledgedNo       Acknowledged = iota
	AcknowledgedPossible Acknowledged = iota
	AcknowledgedYes      Acknowledged = iota
)

type TransitionCapable byte
type Acknowledged byte
type NewMessageFunc func(byte, bool, []byte) (Message, error)

type Message interface {
	Length() byte
	Number() byte
	AcknowledgeRequired() bool
	Data() []byte
	String() string
}
