package messages

import (
	"fmt"

	"github.com/vanstee/nx584/message"
)

var (
	messages = map[byte]message.NewMessageFunc{
		InterfaceConfigurationMessageNumber:         NewInterfaceConfigurationMessage,
		ZoneNameMessageNumber:                       NewZoneNameMessage,
		ZoneStatusMessageNumber:                     NewZoneStatusMessage,
		ZonesSnapshotMessageNumber:                  NewZonesSnapshotMessage,
		PartitionStatusMessageNumber:                NewPartitionStatusMessage,
		PartitionsSnapshotMessageNumber:             NewPartitionsSnapshotMessage,
		SystemStatusMessageNumber:                   NewSystemStatusMessage,
		X10MessageReceivedNumber:                    NewX10MessageReceived,
		LogEventMessageNumber:                       NewLogEventMessage,
		KeypadMessageReceivedNumber:                 NewKeypadMessageReceived,
		ProgramDataReplyNumber:                      NewProgramDataReply,
		UserInformationReplyNumber:                  NewUserInformationReply,
		CommandRequestFailedNumber:                  NewCommandRequestFailed,
		PositiveAcknowledgeNumber:                   NewPositiveAcknowledge,
		NegativeAcknowledgeNumber:                   NewNegativeAcknowledge,
		MessageRejectedNumber:                       NewMessageRejected,
		InterfaceConfigurationRequestNumber:         NewInterfaceConfigurationRequest,
		ZoneNameRequestNumber:                       NewZoneNameRequest,
		ZoneStatusRequestNumber:                     NewZoneStatusRequest,
		ZonesSnapshotRequestNumber:                  NewZonesSnapshotRequest,
		PartitionStatusRequestNumber:                NewPartitionStatusRequest,
		PartitionsSnapshotRequestNumber:             NewPartitionsSnapshotRequest,
		SystemStatusRequestNumber:                   NewSystemStatusRequest,
		SendX10RequestNumber:                        NewSendX10Request,
		LogEventRequestNumber:                       NewLogEventRequest,
		SendKeypadTextMessageNumber:                 NewSendKeypadTextMessage,
		KeypadTerminalModeRequestNumber:             NewKeypadTerminalModeRequest,
		ProgramDataRequestNumber:                    NewProgramDataRequest,
		ProgramDataCommandNumber:                    NewProgramDataCommand,
		UserInformationRequestWithPinNumber:         NewUserInformationRequestWithPin,
		UserInformationRequestWithoutPinNumber:      NewUserInformationRequestWithoutPin,
		SetUserCodeCommandWithPinNumber:             NewSetUserCodeCommandWithPin,
		SetUserCodeCommandWithoutPinNumber:          NewSetUserCodeCommandWithoutPin,
		SetUserAuthorizationCommandWithPinNumber:    NewSetUserAuthorizationCommandWithPin,
		SetUserAuthorizationCommandWithoutPinNumber: NewSetUserAuthorizationCommandWithoutPin,
		StoreCommunicationEventCommandNumber:        NewStoreCommunicationEventCommand,
		SetClockCalendarCommandNumber:               NewSetClockCalendarCommand,
		PrimaryKeypadFunctionWithPinNumber:          NewPrimaryKeypadFunctionWithPin,
		PrimaryKeypadFunctionWithoutPinNumber:       NewPrimaryKeypadFunctionWithoutPin,
		SecondaryKeypadFunctionNumber:               NewSecondaryKeypadFunction,
		ZoneBypassToggleNumber:                      NewZoneBypassToggle,
	}
)

func New(length byte, number byte, acknowledgeRequired bool, data []byte) (message.Message, error) {
	newMessageFunc, ok := messages[number]
	if !ok {
		return nil, fmt.Errorf("message number %v unknown", number)
	}
	return newMessageFunc(length, acknowledgeRequired, data)
}
