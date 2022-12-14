/*
Ais-Stream WebsocketObjects

OpenAPI 3.0 definitions for the data models used by aisstream.io.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// AisStreamMessageMessage struct for AisStreamMessageMessage
type AisStreamMessageMessage struct {
	PositionReport *PositionReport `json:"PositionReport,omitempty"`
	UnknownMessage *UnknownMessage `json:"UnknownMessage,omitempty"`
	AddressedSafetyMessage *AddressedSafetyMessage `json:"AddressedSafetyMessage,omitempty"`
	AddressedBinaryMessage *AddressedBinaryMessage `json:"AddressedBinaryMessage,omitempty"`
	AidsToNavigationReport *AidsToNavigationReport `json:"AidsToNavigationReport,omitempty"`
	AssignedModeCommand *AssignedModeCommand `json:"AssignedModeCommand,omitempty"`
	BaseStationReport *BaseStationReport `json:"BaseStationReport,omitempty"`
	BinaryAcknowledge *BinaryAcknowledge `json:"BinaryAcknowledge,omitempty"`
	BinaryBroadcastMessage *BinaryBroadcastMessage `json:"BinaryBroadcastMessage,omitempty"`
	ChannelManagement *ChannelManagement `json:"ChannelManagement,omitempty"`
	CoordinatedUTCInquiry *CoordinatedUTCInquiry `json:"CoordinatedUTCInquiry,omitempty"`
	DataLinkManagementMessage *DataLinkManagementMessage `json:"DataLinkManagementMessage,omitempty"`
	DataLinkManagementMessageData *DataLinkManagementMessageData `json:"DataLinkManagementMessageData,omitempty"`
	ExtendedClassBPositionReport *ExtendedClassBPositionReport `json:"ExtendedClassBPositionReport,omitempty"`
	GnssBroadcastBinaryMessage *GnssBroadcastBinaryMessage `json:"GnssBroadcastBinaryMessage,omitempty"`
	GroupAssignmentCommand *GroupAssignmentCommand `json:"GroupAssignmentCommand,omitempty"`
	Interrogation *Interrogation `json:"Interrogation,omitempty"`
	LongRangeAisBroadcastMessage *LongRangeAisBroadcastMessage `json:"LongRangeAisBroadcastMessage,omitempty"`
	MultiSlotBinaryMessage *MultiSlotBinaryMessage `json:"MultiSlotBinaryMessage,omitempty"`
	SafetyBroadcastMessage *SafetyBroadcastMessage `json:"SafetyBroadcastMessage,omitempty"`
	ShipStaticData *ShipStaticData `json:"ShipStaticData,omitempty"`
	SingleSlotBinaryMessage *SingleSlotBinaryMessage `json:"SingleSlotBinaryMessage,omitempty"`
	StandardClassBPositionReport *StandardClassBPositionReport `json:"StandardClassBPositionReport,omitempty"`
	StandardSearchAndRescueAircraftReport *StandardSearchAndRescueAircraftReport `json:"StandardSearchAndRescueAircraftReport,omitempty"`
	StaticDataReport *StaticDataReport `json:"StaticDataReport,omitempty"`
}

// NewAisStreamMessageMessage instantiates a new AisStreamMessageMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAisStreamMessageMessage() *AisStreamMessageMessage {
	this := AisStreamMessageMessage{}
	return &this
}

// NewAisStreamMessageMessageWithDefaults instantiates a new AisStreamMessageMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAisStreamMessageMessageWithDefaults() *AisStreamMessageMessage {
	this := AisStreamMessageMessage{}
	return &this
}

// GetPositionReport returns the PositionReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetPositionReport() PositionReport {
	if o == nil || o.PositionReport == nil {
		var ret PositionReport
		return ret
	}
	return *o.PositionReport
}

// GetPositionReportOk returns a tuple with the PositionReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetPositionReportOk() (*PositionReport, bool) {
	if o == nil || o.PositionReport == nil {
		return nil, false
	}
	return o.PositionReport, true
}

// HasPositionReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasPositionReport() bool {
	if o != nil && o.PositionReport != nil {
		return true
	}

	return false
}

// SetPositionReport gets a reference to the given PositionReport and assigns it to the PositionReport field.
func (o *AisStreamMessageMessage) SetPositionReport(v PositionReport) {
	o.PositionReport = &v
}

// GetUnknownMessage returns the UnknownMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetUnknownMessage() UnknownMessage {
	if o == nil || o.UnknownMessage == nil {
		var ret UnknownMessage
		return ret
	}
	return *o.UnknownMessage
}

// GetUnknownMessageOk returns a tuple with the UnknownMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetUnknownMessageOk() (*UnknownMessage, bool) {
	if o == nil || o.UnknownMessage == nil {
		return nil, false
	}
	return o.UnknownMessage, true
}

// HasUnknownMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasUnknownMessage() bool {
	if o != nil && o.UnknownMessage != nil {
		return true
	}

	return false
}

// SetUnknownMessage gets a reference to the given UnknownMessage and assigns it to the UnknownMessage field.
func (o *AisStreamMessageMessage) SetUnknownMessage(v UnknownMessage) {
	o.UnknownMessage = &v
}

// GetAddressedSafetyMessage returns the AddressedSafetyMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetAddressedSafetyMessage() AddressedSafetyMessage {
	if o == nil || o.AddressedSafetyMessage == nil {
		var ret AddressedSafetyMessage
		return ret
	}
	return *o.AddressedSafetyMessage
}

// GetAddressedSafetyMessageOk returns a tuple with the AddressedSafetyMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetAddressedSafetyMessageOk() (*AddressedSafetyMessage, bool) {
	if o == nil || o.AddressedSafetyMessage == nil {
		return nil, false
	}
	return o.AddressedSafetyMessage, true
}

// HasAddressedSafetyMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasAddressedSafetyMessage() bool {
	if o != nil && o.AddressedSafetyMessage != nil {
		return true
	}

	return false
}

// SetAddressedSafetyMessage gets a reference to the given AddressedSafetyMessage and assigns it to the AddressedSafetyMessage field.
func (o *AisStreamMessageMessage) SetAddressedSafetyMessage(v AddressedSafetyMessage) {
	o.AddressedSafetyMessage = &v
}

// GetAddressedBinaryMessage returns the AddressedBinaryMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetAddressedBinaryMessage() AddressedBinaryMessage {
	if o == nil || o.AddressedBinaryMessage == nil {
		var ret AddressedBinaryMessage
		return ret
	}
	return *o.AddressedBinaryMessage
}

// GetAddressedBinaryMessageOk returns a tuple with the AddressedBinaryMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetAddressedBinaryMessageOk() (*AddressedBinaryMessage, bool) {
	if o == nil || o.AddressedBinaryMessage == nil {
		return nil, false
	}
	return o.AddressedBinaryMessage, true
}

// HasAddressedBinaryMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasAddressedBinaryMessage() bool {
	if o != nil && o.AddressedBinaryMessage != nil {
		return true
	}

	return false
}

// SetAddressedBinaryMessage gets a reference to the given AddressedBinaryMessage and assigns it to the AddressedBinaryMessage field.
func (o *AisStreamMessageMessage) SetAddressedBinaryMessage(v AddressedBinaryMessage) {
	o.AddressedBinaryMessage = &v
}

// GetAidsToNavigationReport returns the AidsToNavigationReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetAidsToNavigationReport() AidsToNavigationReport {
	if o == nil || o.AidsToNavigationReport == nil {
		var ret AidsToNavigationReport
		return ret
	}
	return *o.AidsToNavigationReport
}

// GetAidsToNavigationReportOk returns a tuple with the AidsToNavigationReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetAidsToNavigationReportOk() (*AidsToNavigationReport, bool) {
	if o == nil || o.AidsToNavigationReport == nil {
		return nil, false
	}
	return o.AidsToNavigationReport, true
}

// HasAidsToNavigationReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasAidsToNavigationReport() bool {
	if o != nil && o.AidsToNavigationReport != nil {
		return true
	}

	return false
}

// SetAidsToNavigationReport gets a reference to the given AidsToNavigationReport and assigns it to the AidsToNavigationReport field.
func (o *AisStreamMessageMessage) SetAidsToNavigationReport(v AidsToNavigationReport) {
	o.AidsToNavigationReport = &v
}

// GetAssignedModeCommand returns the AssignedModeCommand field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetAssignedModeCommand() AssignedModeCommand {
	if o == nil || o.AssignedModeCommand == nil {
		var ret AssignedModeCommand
		return ret
	}
	return *o.AssignedModeCommand
}

// GetAssignedModeCommandOk returns a tuple with the AssignedModeCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetAssignedModeCommandOk() (*AssignedModeCommand, bool) {
	if o == nil || o.AssignedModeCommand == nil {
		return nil, false
	}
	return o.AssignedModeCommand, true
}

// HasAssignedModeCommand returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasAssignedModeCommand() bool {
	if o != nil && o.AssignedModeCommand != nil {
		return true
	}

	return false
}

// SetAssignedModeCommand gets a reference to the given AssignedModeCommand and assigns it to the AssignedModeCommand field.
func (o *AisStreamMessageMessage) SetAssignedModeCommand(v AssignedModeCommand) {
	o.AssignedModeCommand = &v
}

// GetBaseStationReport returns the BaseStationReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetBaseStationReport() BaseStationReport {
	if o == nil || o.BaseStationReport == nil {
		var ret BaseStationReport
		return ret
	}
	return *o.BaseStationReport
}

// GetBaseStationReportOk returns a tuple with the BaseStationReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetBaseStationReportOk() (*BaseStationReport, bool) {
	if o == nil || o.BaseStationReport == nil {
		return nil, false
	}
	return o.BaseStationReport, true
}

// HasBaseStationReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasBaseStationReport() bool {
	if o != nil && o.BaseStationReport != nil {
		return true
	}

	return false
}

// SetBaseStationReport gets a reference to the given BaseStationReport and assigns it to the BaseStationReport field.
func (o *AisStreamMessageMessage) SetBaseStationReport(v BaseStationReport) {
	o.BaseStationReport = &v
}

// GetBinaryAcknowledge returns the BinaryAcknowledge field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetBinaryAcknowledge() BinaryAcknowledge {
	if o == nil || o.BinaryAcknowledge == nil {
		var ret BinaryAcknowledge
		return ret
	}
	return *o.BinaryAcknowledge
}

// GetBinaryAcknowledgeOk returns a tuple with the BinaryAcknowledge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetBinaryAcknowledgeOk() (*BinaryAcknowledge, bool) {
	if o == nil || o.BinaryAcknowledge == nil {
		return nil, false
	}
	return o.BinaryAcknowledge, true
}

// HasBinaryAcknowledge returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasBinaryAcknowledge() bool {
	if o != nil && o.BinaryAcknowledge != nil {
		return true
	}

	return false
}

// SetBinaryAcknowledge gets a reference to the given BinaryAcknowledge and assigns it to the BinaryAcknowledge field.
func (o *AisStreamMessageMessage) SetBinaryAcknowledge(v BinaryAcknowledge) {
	o.BinaryAcknowledge = &v
}

// GetBinaryBroadcastMessage returns the BinaryBroadcastMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetBinaryBroadcastMessage() BinaryBroadcastMessage {
	if o == nil || o.BinaryBroadcastMessage == nil {
		var ret BinaryBroadcastMessage
		return ret
	}
	return *o.BinaryBroadcastMessage
}

// GetBinaryBroadcastMessageOk returns a tuple with the BinaryBroadcastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetBinaryBroadcastMessageOk() (*BinaryBroadcastMessage, bool) {
	if o == nil || o.BinaryBroadcastMessage == nil {
		return nil, false
	}
	return o.BinaryBroadcastMessage, true
}

// HasBinaryBroadcastMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasBinaryBroadcastMessage() bool {
	if o != nil && o.BinaryBroadcastMessage != nil {
		return true
	}

	return false
}

// SetBinaryBroadcastMessage gets a reference to the given BinaryBroadcastMessage and assigns it to the BinaryBroadcastMessage field.
func (o *AisStreamMessageMessage) SetBinaryBroadcastMessage(v BinaryBroadcastMessage) {
	o.BinaryBroadcastMessage = &v
}

// GetChannelManagement returns the ChannelManagement field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetChannelManagement() ChannelManagement {
	if o == nil || o.ChannelManagement == nil {
		var ret ChannelManagement
		return ret
	}
	return *o.ChannelManagement
}

// GetChannelManagementOk returns a tuple with the ChannelManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetChannelManagementOk() (*ChannelManagement, bool) {
	if o == nil || o.ChannelManagement == nil {
		return nil, false
	}
	return o.ChannelManagement, true
}

// HasChannelManagement returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasChannelManagement() bool {
	if o != nil && o.ChannelManagement != nil {
		return true
	}

	return false
}

// SetChannelManagement gets a reference to the given ChannelManagement and assigns it to the ChannelManagement field.
func (o *AisStreamMessageMessage) SetChannelManagement(v ChannelManagement) {
	o.ChannelManagement = &v
}

// GetCoordinatedUTCInquiry returns the CoordinatedUTCInquiry field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetCoordinatedUTCInquiry() CoordinatedUTCInquiry {
	if o == nil || o.CoordinatedUTCInquiry == nil {
		var ret CoordinatedUTCInquiry
		return ret
	}
	return *o.CoordinatedUTCInquiry
}

// GetCoordinatedUTCInquiryOk returns a tuple with the CoordinatedUTCInquiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetCoordinatedUTCInquiryOk() (*CoordinatedUTCInquiry, bool) {
	if o == nil || o.CoordinatedUTCInquiry == nil {
		return nil, false
	}
	return o.CoordinatedUTCInquiry, true
}

// HasCoordinatedUTCInquiry returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasCoordinatedUTCInquiry() bool {
	if o != nil && o.CoordinatedUTCInquiry != nil {
		return true
	}

	return false
}

// SetCoordinatedUTCInquiry gets a reference to the given CoordinatedUTCInquiry and assigns it to the CoordinatedUTCInquiry field.
func (o *AisStreamMessageMessage) SetCoordinatedUTCInquiry(v CoordinatedUTCInquiry) {
	o.CoordinatedUTCInquiry = &v
}

// GetDataLinkManagementMessage returns the DataLinkManagementMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetDataLinkManagementMessage() DataLinkManagementMessage {
	if o == nil || o.DataLinkManagementMessage == nil {
		var ret DataLinkManagementMessage
		return ret
	}
	return *o.DataLinkManagementMessage
}

// GetDataLinkManagementMessageOk returns a tuple with the DataLinkManagementMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetDataLinkManagementMessageOk() (*DataLinkManagementMessage, bool) {
	if o == nil || o.DataLinkManagementMessage == nil {
		return nil, false
	}
	return o.DataLinkManagementMessage, true
}

// HasDataLinkManagementMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasDataLinkManagementMessage() bool {
	if o != nil && o.DataLinkManagementMessage != nil {
		return true
	}

	return false
}

// SetDataLinkManagementMessage gets a reference to the given DataLinkManagementMessage and assigns it to the DataLinkManagementMessage field.
func (o *AisStreamMessageMessage) SetDataLinkManagementMessage(v DataLinkManagementMessage) {
	o.DataLinkManagementMessage = &v
}

// GetDataLinkManagementMessageData returns the DataLinkManagementMessageData field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetDataLinkManagementMessageData() DataLinkManagementMessageData {
	if o == nil || o.DataLinkManagementMessageData == nil {
		var ret DataLinkManagementMessageData
		return ret
	}
	return *o.DataLinkManagementMessageData
}

// GetDataLinkManagementMessageDataOk returns a tuple with the DataLinkManagementMessageData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetDataLinkManagementMessageDataOk() (*DataLinkManagementMessageData, bool) {
	if o == nil || o.DataLinkManagementMessageData == nil {
		return nil, false
	}
	return o.DataLinkManagementMessageData, true
}

// HasDataLinkManagementMessageData returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasDataLinkManagementMessageData() bool {
	if o != nil && o.DataLinkManagementMessageData != nil {
		return true
	}

	return false
}

// SetDataLinkManagementMessageData gets a reference to the given DataLinkManagementMessageData and assigns it to the DataLinkManagementMessageData field.
func (o *AisStreamMessageMessage) SetDataLinkManagementMessageData(v DataLinkManagementMessageData) {
	o.DataLinkManagementMessageData = &v
}

// GetExtendedClassBPositionReport returns the ExtendedClassBPositionReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetExtendedClassBPositionReport() ExtendedClassBPositionReport {
	if o == nil || o.ExtendedClassBPositionReport == nil {
		var ret ExtendedClassBPositionReport
		return ret
	}
	return *o.ExtendedClassBPositionReport
}

// GetExtendedClassBPositionReportOk returns a tuple with the ExtendedClassBPositionReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetExtendedClassBPositionReportOk() (*ExtendedClassBPositionReport, bool) {
	if o == nil || o.ExtendedClassBPositionReport == nil {
		return nil, false
	}
	return o.ExtendedClassBPositionReport, true
}

// HasExtendedClassBPositionReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasExtendedClassBPositionReport() bool {
	if o != nil && o.ExtendedClassBPositionReport != nil {
		return true
	}

	return false
}

// SetExtendedClassBPositionReport gets a reference to the given ExtendedClassBPositionReport and assigns it to the ExtendedClassBPositionReport field.
func (o *AisStreamMessageMessage) SetExtendedClassBPositionReport(v ExtendedClassBPositionReport) {
	o.ExtendedClassBPositionReport = &v
}

// GetGnssBroadcastBinaryMessage returns the GnssBroadcastBinaryMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetGnssBroadcastBinaryMessage() GnssBroadcastBinaryMessage {
	if o == nil || o.GnssBroadcastBinaryMessage == nil {
		var ret GnssBroadcastBinaryMessage
		return ret
	}
	return *o.GnssBroadcastBinaryMessage
}

// GetGnssBroadcastBinaryMessageOk returns a tuple with the GnssBroadcastBinaryMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetGnssBroadcastBinaryMessageOk() (*GnssBroadcastBinaryMessage, bool) {
	if o == nil || o.GnssBroadcastBinaryMessage == nil {
		return nil, false
	}
	return o.GnssBroadcastBinaryMessage, true
}

// HasGnssBroadcastBinaryMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasGnssBroadcastBinaryMessage() bool {
	if o != nil && o.GnssBroadcastBinaryMessage != nil {
		return true
	}

	return false
}

// SetGnssBroadcastBinaryMessage gets a reference to the given GnssBroadcastBinaryMessage and assigns it to the GnssBroadcastBinaryMessage field.
func (o *AisStreamMessageMessage) SetGnssBroadcastBinaryMessage(v GnssBroadcastBinaryMessage) {
	o.GnssBroadcastBinaryMessage = &v
}

// GetGroupAssignmentCommand returns the GroupAssignmentCommand field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetGroupAssignmentCommand() GroupAssignmentCommand {
	if o == nil || o.GroupAssignmentCommand == nil {
		var ret GroupAssignmentCommand
		return ret
	}
	return *o.GroupAssignmentCommand
}

// GetGroupAssignmentCommandOk returns a tuple with the GroupAssignmentCommand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetGroupAssignmentCommandOk() (*GroupAssignmentCommand, bool) {
	if o == nil || o.GroupAssignmentCommand == nil {
		return nil, false
	}
	return o.GroupAssignmentCommand, true
}

// HasGroupAssignmentCommand returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasGroupAssignmentCommand() bool {
	if o != nil && o.GroupAssignmentCommand != nil {
		return true
	}

	return false
}

// SetGroupAssignmentCommand gets a reference to the given GroupAssignmentCommand and assigns it to the GroupAssignmentCommand field.
func (o *AisStreamMessageMessage) SetGroupAssignmentCommand(v GroupAssignmentCommand) {
	o.GroupAssignmentCommand = &v
}

// GetInterrogation returns the Interrogation field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetInterrogation() Interrogation {
	if o == nil || o.Interrogation == nil {
		var ret Interrogation
		return ret
	}
	return *o.Interrogation
}

// GetInterrogationOk returns a tuple with the Interrogation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetInterrogationOk() (*Interrogation, bool) {
	if o == nil || o.Interrogation == nil {
		return nil, false
	}
	return o.Interrogation, true
}

// HasInterrogation returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasInterrogation() bool {
	if o != nil && o.Interrogation != nil {
		return true
	}

	return false
}

// SetInterrogation gets a reference to the given Interrogation and assigns it to the Interrogation field.
func (o *AisStreamMessageMessage) SetInterrogation(v Interrogation) {
	o.Interrogation = &v
}

// GetLongRangeAisBroadcastMessage returns the LongRangeAisBroadcastMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetLongRangeAisBroadcastMessage() LongRangeAisBroadcastMessage {
	if o == nil || o.LongRangeAisBroadcastMessage == nil {
		var ret LongRangeAisBroadcastMessage
		return ret
	}
	return *o.LongRangeAisBroadcastMessage
}

// GetLongRangeAisBroadcastMessageOk returns a tuple with the LongRangeAisBroadcastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetLongRangeAisBroadcastMessageOk() (*LongRangeAisBroadcastMessage, bool) {
	if o == nil || o.LongRangeAisBroadcastMessage == nil {
		return nil, false
	}
	return o.LongRangeAisBroadcastMessage, true
}

// HasLongRangeAisBroadcastMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasLongRangeAisBroadcastMessage() bool {
	if o != nil && o.LongRangeAisBroadcastMessage != nil {
		return true
	}

	return false
}

// SetLongRangeAisBroadcastMessage gets a reference to the given LongRangeAisBroadcastMessage and assigns it to the LongRangeAisBroadcastMessage field.
func (o *AisStreamMessageMessage) SetLongRangeAisBroadcastMessage(v LongRangeAisBroadcastMessage) {
	o.LongRangeAisBroadcastMessage = &v
}

// GetMultiSlotBinaryMessage returns the MultiSlotBinaryMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetMultiSlotBinaryMessage() MultiSlotBinaryMessage {
	if o == nil || o.MultiSlotBinaryMessage == nil {
		var ret MultiSlotBinaryMessage
		return ret
	}
	return *o.MultiSlotBinaryMessage
}

// GetMultiSlotBinaryMessageOk returns a tuple with the MultiSlotBinaryMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetMultiSlotBinaryMessageOk() (*MultiSlotBinaryMessage, bool) {
	if o == nil || o.MultiSlotBinaryMessage == nil {
		return nil, false
	}
	return o.MultiSlotBinaryMessage, true
}

// HasMultiSlotBinaryMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasMultiSlotBinaryMessage() bool {
	if o != nil && o.MultiSlotBinaryMessage != nil {
		return true
	}

	return false
}

// SetMultiSlotBinaryMessage gets a reference to the given MultiSlotBinaryMessage and assigns it to the MultiSlotBinaryMessage field.
func (o *AisStreamMessageMessage) SetMultiSlotBinaryMessage(v MultiSlotBinaryMessage) {
	o.MultiSlotBinaryMessage = &v
}

// GetSafetyBroadcastMessage returns the SafetyBroadcastMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetSafetyBroadcastMessage() SafetyBroadcastMessage {
	if o == nil || o.SafetyBroadcastMessage == nil {
		var ret SafetyBroadcastMessage
		return ret
	}
	return *o.SafetyBroadcastMessage
}

// GetSafetyBroadcastMessageOk returns a tuple with the SafetyBroadcastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetSafetyBroadcastMessageOk() (*SafetyBroadcastMessage, bool) {
	if o == nil || o.SafetyBroadcastMessage == nil {
		return nil, false
	}
	return o.SafetyBroadcastMessage, true
}

// HasSafetyBroadcastMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasSafetyBroadcastMessage() bool {
	if o != nil && o.SafetyBroadcastMessage != nil {
		return true
	}

	return false
}

// SetSafetyBroadcastMessage gets a reference to the given SafetyBroadcastMessage and assigns it to the SafetyBroadcastMessage field.
func (o *AisStreamMessageMessage) SetSafetyBroadcastMessage(v SafetyBroadcastMessage) {
	o.SafetyBroadcastMessage = &v
}

// GetShipStaticData returns the ShipStaticData field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetShipStaticData() ShipStaticData {
	if o == nil || o.ShipStaticData == nil {
		var ret ShipStaticData
		return ret
	}
	return *o.ShipStaticData
}

// GetShipStaticDataOk returns a tuple with the ShipStaticData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetShipStaticDataOk() (*ShipStaticData, bool) {
	if o == nil || o.ShipStaticData == nil {
		return nil, false
	}
	return o.ShipStaticData, true
}

// HasShipStaticData returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasShipStaticData() bool {
	if o != nil && o.ShipStaticData != nil {
		return true
	}

	return false
}

// SetShipStaticData gets a reference to the given ShipStaticData and assigns it to the ShipStaticData field.
func (o *AisStreamMessageMessage) SetShipStaticData(v ShipStaticData) {
	o.ShipStaticData = &v
}

// GetSingleSlotBinaryMessage returns the SingleSlotBinaryMessage field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetSingleSlotBinaryMessage() SingleSlotBinaryMessage {
	if o == nil || o.SingleSlotBinaryMessage == nil {
		var ret SingleSlotBinaryMessage
		return ret
	}
	return *o.SingleSlotBinaryMessage
}

// GetSingleSlotBinaryMessageOk returns a tuple with the SingleSlotBinaryMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetSingleSlotBinaryMessageOk() (*SingleSlotBinaryMessage, bool) {
	if o == nil || o.SingleSlotBinaryMessage == nil {
		return nil, false
	}
	return o.SingleSlotBinaryMessage, true
}

// HasSingleSlotBinaryMessage returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasSingleSlotBinaryMessage() bool {
	if o != nil && o.SingleSlotBinaryMessage != nil {
		return true
	}

	return false
}

// SetSingleSlotBinaryMessage gets a reference to the given SingleSlotBinaryMessage and assigns it to the SingleSlotBinaryMessage field.
func (o *AisStreamMessageMessage) SetSingleSlotBinaryMessage(v SingleSlotBinaryMessage) {
	o.SingleSlotBinaryMessage = &v
}

// GetStandardClassBPositionReport returns the StandardClassBPositionReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetStandardClassBPositionReport() StandardClassBPositionReport {
	if o == nil || o.StandardClassBPositionReport == nil {
		var ret StandardClassBPositionReport
		return ret
	}
	return *o.StandardClassBPositionReport
}

// GetStandardClassBPositionReportOk returns a tuple with the StandardClassBPositionReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetStandardClassBPositionReportOk() (*StandardClassBPositionReport, bool) {
	if o == nil || o.StandardClassBPositionReport == nil {
		return nil, false
	}
	return o.StandardClassBPositionReport, true
}

// HasStandardClassBPositionReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasStandardClassBPositionReport() bool {
	if o != nil && o.StandardClassBPositionReport != nil {
		return true
	}

	return false
}

// SetStandardClassBPositionReport gets a reference to the given StandardClassBPositionReport and assigns it to the StandardClassBPositionReport field.
func (o *AisStreamMessageMessage) SetStandardClassBPositionReport(v StandardClassBPositionReport) {
	o.StandardClassBPositionReport = &v
}

// GetStandardSearchAndRescueAircraftReport returns the StandardSearchAndRescueAircraftReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetStandardSearchAndRescueAircraftReport() StandardSearchAndRescueAircraftReport {
	if o == nil || o.StandardSearchAndRescueAircraftReport == nil {
		var ret StandardSearchAndRescueAircraftReport
		return ret
	}
	return *o.StandardSearchAndRescueAircraftReport
}

// GetStandardSearchAndRescueAircraftReportOk returns a tuple with the StandardSearchAndRescueAircraftReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetStandardSearchAndRescueAircraftReportOk() (*StandardSearchAndRescueAircraftReport, bool) {
	if o == nil || o.StandardSearchAndRescueAircraftReport == nil {
		return nil, false
	}
	return o.StandardSearchAndRescueAircraftReport, true
}

// HasStandardSearchAndRescueAircraftReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasStandardSearchAndRescueAircraftReport() bool {
	if o != nil && o.StandardSearchAndRescueAircraftReport != nil {
		return true
	}

	return false
}

// SetStandardSearchAndRescueAircraftReport gets a reference to the given StandardSearchAndRescueAircraftReport and assigns it to the StandardSearchAndRescueAircraftReport field.
func (o *AisStreamMessageMessage) SetStandardSearchAndRescueAircraftReport(v StandardSearchAndRescueAircraftReport) {
	o.StandardSearchAndRescueAircraftReport = &v
}

// GetStaticDataReport returns the StaticDataReport field value if set, zero value otherwise.
func (o *AisStreamMessageMessage) GetStaticDataReport() StaticDataReport {
	if o == nil || o.StaticDataReport == nil {
		var ret StaticDataReport
		return ret
	}
	return *o.StaticDataReport
}

// GetStaticDataReportOk returns a tuple with the StaticDataReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AisStreamMessageMessage) GetStaticDataReportOk() (*StaticDataReport, bool) {
	if o == nil || o.StaticDataReport == nil {
		return nil, false
	}
	return o.StaticDataReport, true
}

// HasStaticDataReport returns a boolean if a field has been set.
func (o *AisStreamMessageMessage) HasStaticDataReport() bool {
	if o != nil && o.StaticDataReport != nil {
		return true
	}

	return false
}

// SetStaticDataReport gets a reference to the given StaticDataReport and assigns it to the StaticDataReport field.
func (o *AisStreamMessageMessage) SetStaticDataReport(v StaticDataReport) {
	o.StaticDataReport = &v
}

func (o AisStreamMessageMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PositionReport != nil {
		toSerialize["PositionReport"] = o.PositionReport
	}
	if o.UnknownMessage != nil {
		toSerialize["UnknownMessage"] = o.UnknownMessage
	}
	if o.AddressedSafetyMessage != nil {
		toSerialize["AddressedSafetyMessage"] = o.AddressedSafetyMessage
	}
	if o.AddressedBinaryMessage != nil {
		toSerialize["AddressedBinaryMessage"] = o.AddressedBinaryMessage
	}
	if o.AidsToNavigationReport != nil {
		toSerialize["AidsToNavigationReport"] = o.AidsToNavigationReport
	}
	if o.AssignedModeCommand != nil {
		toSerialize["AssignedModeCommand"] = o.AssignedModeCommand
	}
	if o.BaseStationReport != nil {
		toSerialize["BaseStationReport"] = o.BaseStationReport
	}
	if o.BinaryAcknowledge != nil {
		toSerialize["BinaryAcknowledge"] = o.BinaryAcknowledge
	}
	if o.BinaryBroadcastMessage != nil {
		toSerialize["BinaryBroadcastMessage"] = o.BinaryBroadcastMessage
	}
	if o.ChannelManagement != nil {
		toSerialize["ChannelManagement"] = o.ChannelManagement
	}
	if o.CoordinatedUTCInquiry != nil {
		toSerialize["CoordinatedUTCInquiry"] = o.CoordinatedUTCInquiry
	}
	if o.DataLinkManagementMessage != nil {
		toSerialize["DataLinkManagementMessage"] = o.DataLinkManagementMessage
	}
	if o.DataLinkManagementMessageData != nil {
		toSerialize["DataLinkManagementMessageData"] = o.DataLinkManagementMessageData
	}
	if o.ExtendedClassBPositionReport != nil {
		toSerialize["ExtendedClassBPositionReport"] = o.ExtendedClassBPositionReport
	}
	if o.GnssBroadcastBinaryMessage != nil {
		toSerialize["GnssBroadcastBinaryMessage"] = o.GnssBroadcastBinaryMessage
	}
	if o.GroupAssignmentCommand != nil {
		toSerialize["GroupAssignmentCommand"] = o.GroupAssignmentCommand
	}
	if o.Interrogation != nil {
		toSerialize["Interrogation"] = o.Interrogation
	}
	if o.LongRangeAisBroadcastMessage != nil {
		toSerialize["LongRangeAisBroadcastMessage"] = o.LongRangeAisBroadcastMessage
	}
	if o.MultiSlotBinaryMessage != nil {
		toSerialize["MultiSlotBinaryMessage"] = o.MultiSlotBinaryMessage
	}
	if o.SafetyBroadcastMessage != nil {
		toSerialize["SafetyBroadcastMessage"] = o.SafetyBroadcastMessage
	}
	if o.ShipStaticData != nil {
		toSerialize["ShipStaticData"] = o.ShipStaticData
	}
	if o.SingleSlotBinaryMessage != nil {
		toSerialize["SingleSlotBinaryMessage"] = o.SingleSlotBinaryMessage
	}
	if o.StandardClassBPositionReport != nil {
		toSerialize["StandardClassBPositionReport"] = o.StandardClassBPositionReport
	}
	if o.StandardSearchAndRescueAircraftReport != nil {
		toSerialize["StandardSearchAndRescueAircraftReport"] = o.StandardSearchAndRescueAircraftReport
	}
	if o.StaticDataReport != nil {
		toSerialize["StaticDataReport"] = o.StaticDataReport
	}
	return json.Marshal(toSerialize)
}

type NullableAisStreamMessageMessage struct {
	value *AisStreamMessageMessage
	isSet bool
}

func (v NullableAisStreamMessageMessage) Get() *AisStreamMessageMessage {
	return v.value
}

func (v *NullableAisStreamMessageMessage) Set(val *AisStreamMessageMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAisStreamMessageMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAisStreamMessageMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAisStreamMessageMessage(val *AisStreamMessageMessage) *NullableAisStreamMessageMessage {
	return &NullableAisStreamMessageMessage{value: val, isSet: true}
}

func (v NullableAisStreamMessageMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAisStreamMessageMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


