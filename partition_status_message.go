// DO NOT EDIT -- This file was generated from templates/message.tmpl

package nx584

import (
	"fmt"
	"log"
	"strings"
)

const (
	PartitionStatusMessageName              = "Partition Status Message"
	PartitionStatusMessageLength            = 9
	PartitionStatusMessageNumber            = 6
	PartitionStatusMessageTransitionCapable = TransitionCapableYes
	PartitionStatusMessageAcknowledged      = AcknowledgedPossible
)

type PartitionStatusMessage struct {
	*BaseMessage
	PartitionNumber            byte
	PartitionConditionFlags1   byte
	PartitionConditionFlags2   byte
	PartitionConditionFlags3   byte
	PartitionConditionFlags4   byte
	LastUserNumber             byte
	PartitionConditionFlags5   byte
	PartitionConditionFlags6   byte
	BypassCodeRequired         byte
	FireTrouble                byte
	Fire                       byte
	PulsingBuzzer              byte
	TlmFaultMemory             byte
	Armed                      byte
	Instant                    byte
	PreviousAlarm              byte
	SirenOn                    byte
	SteadySirenOn              byte
	AlarmMemory                byte
	Tamper                     byte
	CancelCommandEntered       byte
	CodeEntered                byte
	CancelPending              byte
	SilentExitEnabled          byte
	Entryguard                 byte
	ChimeModeOn                byte
	Entry                      byte
	DelayExpirationWarning     byte
	Exit1                      byte
	Exit2                      byte
	LedExtinguish              byte
	CrossTiming                byte
	RecentClosingBeingTimed    byte
	ExitErrorTriggered         byte
	AutoHomeInhibited          byte
	SensorLowBattery           byte
	SensorLostSupervision      byte
	ZoneBypassed               byte
	ForceArmTriggeredByAutoArm byte
	ReadyToArm                 byte
	ReadyToForceArm            byte
	ValidPinAccepted           byte
	ChimeOn                    byte
	ErrorBeep                  byte
	ToneOn                     byte
	Entry1                     byte
	OpenPeriod                 byte
	AlarmSentUsingPhoneNumber1 byte
	AlarmSentUsingPhoneNumber2 byte
	AlarmSentUsingPhoneNumber3 byte
	CancelReportIsInTheStack   byte
	KeyswitchArmed             byte
	DelayTripInProgress        byte
}

func NewPartitionStatusMessage(length byte, acknowledgeRequired bool, data []byte) (Message, error) {
	if length != PartitionStatusMessageLength {
		log.Printf("message length incorrect: expected %d, actual: %d", PartitionStatusMessageLength, length)
	}

	message := &PartitionStatusMessage{
		BaseMessage: &BaseMessage{
			length:              length,
			acknowledgeRequired: acknowledgeRequired,
			data:                data,
		},

		PartitionNumber: data[0],

		PartitionConditionFlags1: data[1],
		BypassCodeRequired:       (data[1] >> 0) & 0x1,
		FireTrouble:              (data[1] >> 1) & 0x1,
		Fire:                     (data[1] >> 2) & 0x1,
		PulsingBuzzer:            (data[1] >> 3) & 0x1,
		TlmFaultMemory:           (data[1] >> 4) & 0x1,
		Armed:                    (data[1] >> 6) & 0x1,
		Instant:                  (data[1] >> 7) & 0x1,

		PartitionConditionFlags2: data[2],
		PreviousAlarm:            (data[2] >> 0) & 0x1,
		SirenOn:                  (data[2] >> 1) & 0x1,
		SteadySirenOn:            (data[2] >> 2) & 0x1,
		AlarmMemory:              (data[2] >> 3) & 0x1,
		Tamper:                   (data[2] >> 4) & 0x1,
		CancelCommandEntered:     (data[2] >> 5) & 0x1,
		CodeEntered:              (data[2] >> 6) & 0x1,
		CancelPending:            (data[2] >> 7) & 0x1,

		PartitionConditionFlags3: data[3],
		SilentExitEnabled:        (data[3] >> 1) & 0x1,
		Entryguard:               (data[3] >> 2) & 0x1,
		ChimeModeOn:              (data[3] >> 3) & 0x1,
		Entry:                    (data[3] >> 4) & 0x1,
		DelayExpirationWarning: (data[3] >> 5) & 0x1,
		Exit1: (data[3] >> 6) & 0x1,
		Exit2: (data[3] >> 7) & 0x1,

		PartitionConditionFlags4: data[4],
		LedExtinguish:            (data[4] >> 0) & 0x1,
		CrossTiming:              (data[4] >> 1) & 0x1,
		RecentClosingBeingTimed:  (data[4] >> 2) & 0x1,
		ExitErrorTriggered:       (data[4] >> 4) & 0x1,
		AutoHomeInhibited:        (data[4] >> 5) & 0x1,
		SensorLowBattery:         (data[4] >> 6) & 0x1,
		SensorLostSupervision:    (data[4] >> 7) & 0x1,

		LastUserNumber: data[5],

		PartitionConditionFlags5:   data[6],
		ZoneBypassed:               (data[6] >> 0) & 0x1,
		ForceArmTriggeredByAutoArm: (data[6] >> 1) & 0x1,
		ReadyToArm:                 (data[6] >> 2) & 0x1,
		ReadyToForceArm:            (data[6] >> 3) & 0x1,
		ValidPinAccepted:           (data[6] >> 4) & 0x1,
		ChimeOn:                    (data[6] >> 5) & 0x1,
		ErrorBeep:                  (data[6] >> 6) & 0x1,
		ToneOn:                     (data[6] >> 7) & 0x1,

		PartitionConditionFlags6: data[7],
		Entry1:                     (data[7] >> 0) & 0x1,
		OpenPeriod:                 (data[7] >> 1) & 0x1,
		AlarmSentUsingPhoneNumber1: (data[7] >> 2) & 0x1,
		AlarmSentUsingPhoneNumber2: (data[7] >> 3) & 0x1,
		AlarmSentUsingPhoneNumber3: (data[7] >> 4) & 0x1,
		CancelReportIsInTheStack:   (data[7] >> 5) & 0x1,
		KeyswitchArmed:             (data[7] >> 6) & 0x1,
		DelayTripInProgress:        (data[7] >> 7) & 0x1,
	}

	return message, nil
}

func (m *PartitionStatusMessage) Number() byte {
	return PartitionStatusMessageNumber
}

func (m *PartitionStatusMessage) String() string {
	return fmt.Sprintf(
		`%s
                    %s
                    Partition Number: %08b
                    Partition Condition Flags 1: %08b
                      Bypass Code Required: %01b
                      Fire Trouble: %01b
                      Fire: %01b
                      Pulsing Buzzer: %01b
                      TLM Fault Memory: %01b
                      Armed: %01b
                      Instant: %01b
                    Partition Condition Flags 2: %08b
                      Previous Alarm: %01b
                      Siren On: %01b
                      Steady Siren On: %01b
                      Alarm Memory: %01b
                      Tamper: %01b
                      Cancel Command Entered: %01b
                      Code Entered: %01b
                      Cancel Pending: %01b
                    Partition Condition Flags 3: %08b
                      Silent Exit Enabled: %01b
                      Entryguard: %01b
                      Chime Mode On: %01b
                      Entry: %01b
                      Delay Expiration Warning: %01b
                      Exit 1: %01b
                      Exit 2: %01b
                    Partition Condition Flags 4: %08b
                      LED Extinguish: %01b
                      Cross Timing: %01b
                      Recent Closing Being Timed: %01b
                      Exit Error Triggered: %01b
                      Auto Home Inhibited: %01b
                      Sensor Low Battery: %01b
                      Sensor Lost Supervision: %01b
                    Last User Number: %08b
                    Partition Condition Flags 5: %08b
                      Zone Bypassed: %01b
                      Force Arm Triggered By Auto Arm: %01b
                      Ready To Arm: %01b
                      Ready To Force Arm: %01b
                      Valid PIN Accepted: %01b
                      Chime On: %01b
                      Error Beep: %01b
                      Tone On: %01b
                    Partition Condition Flags 6: %08b
                      Entry 1: %01b
                      Open Period: %01b
                      Alarm Sent Using Phone Number 1: %01b
                      Alarm Sent Using Phone Number 2: %01b
                      Alarm Sent Using Phone Number 3: %01b
                      Cancel Report Is In the Stack: %01b
                      Keyswitch Armed: %01b
                      Delay Trip In Progress: %01b`,
		"Partition Status Message",
		strings.Repeat("-", len("Partition Status Message")),
		m.PartitionNumber,
		m.PartitionConditionFlags1,
		m.BypassCodeRequired,
		m.FireTrouble,
		m.Fire,
		m.PulsingBuzzer,
		m.TlmFaultMemory,
		m.Armed,
		m.Instant,
		m.PartitionConditionFlags2,
		m.PreviousAlarm,
		m.SirenOn,
		m.SteadySirenOn,
		m.AlarmMemory,
		m.Tamper,
		m.CancelCommandEntered,
		m.CodeEntered,
		m.CancelPending,
		m.PartitionConditionFlags3,
		m.SilentExitEnabled,
		m.Entryguard,
		m.ChimeModeOn,
		m.Entry,
		m.DelayExpirationWarning,
		m.Exit1,
		m.Exit2,
		m.PartitionConditionFlags4,
		m.LedExtinguish,
		m.CrossTiming,
		m.RecentClosingBeingTimed,
		m.ExitErrorTriggered,
		m.AutoHomeInhibited,
		m.SensorLowBattery,
		m.SensorLostSupervision,
		m.LastUserNumber,
		m.PartitionConditionFlags5,
		m.ZoneBypassed,
		m.ForceArmTriggeredByAutoArm,
		m.ReadyToArm,
		m.ReadyToForceArm,
		m.ValidPinAccepted,
		m.ChimeOn,
		m.ErrorBeep,
		m.ToneOn,
		m.PartitionConditionFlags6,
		m.Entry1,
		m.OpenPeriod,
		m.AlarmSentUsingPhoneNumber1,
		m.AlarmSentUsingPhoneNumber2,
		m.AlarmSentUsingPhoneNumber3,
		m.CancelReportIsInTheStack,
		m.KeyswitchArmed,
		m.DelayTripInProgress,
	)
}
