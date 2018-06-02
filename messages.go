package nx584

var (
	messages = map[byte]NewMessageFunc{
		ZoneStatusMessageNumber: NewZoneStatusMessage,
	}
)
