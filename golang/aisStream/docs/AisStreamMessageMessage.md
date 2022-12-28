# AisStreamMessageMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PositionReport** | Pointer to [**PositionReport**](PositionReport.md) |  | [optional] 
**UnknownMessage** | Pointer to [**UnknownMessage**](UnknownMessage.md) |  | [optional] 
**AddressedSafetyMessage** | Pointer to [**AddressedSafetyMessage**](AddressedSafetyMessage.md) |  | [optional] 
**AddressedBinaryMessage** | Pointer to [**AddressedBinaryMessage**](AddressedBinaryMessage.md) |  | [optional] 
**AidsToNavigationReport** | Pointer to [**AidsToNavigationReport**](AidsToNavigationReport.md) |  | [optional] 
**AssignedModeCommand** | Pointer to [**AssignedModeCommand**](AssignedModeCommand.md) |  | [optional] 
**BaseStationReport** | Pointer to [**BaseStationReport**](BaseStationReport.md) |  | [optional] 
**BinaryAcknowledge** | Pointer to [**BinaryAcknowledge**](BinaryAcknowledge.md) |  | [optional] 
**BinaryBroadcastMessage** | Pointer to [**BinaryBroadcastMessage**](BinaryBroadcastMessage.md) |  | [optional] 
**ChannelManagement** | Pointer to [**ChannelManagement**](ChannelManagement.md) |  | [optional] 
**CoordinatedUTCInquiry** | Pointer to [**CoordinatedUTCInquiry**](CoordinatedUTCInquiry.md) |  | [optional] 
**DataLinkManagementMessage** | Pointer to [**DataLinkManagementMessage**](DataLinkManagementMessage.md) |  | [optional] 
**DataLinkManagementMessageData** | Pointer to [**DataLinkManagementMessageData**](DataLinkManagementMessageData.md) |  | [optional] 
**ExtendedClassBPositionReport** | Pointer to [**ExtendedClassBPositionReport**](ExtendedClassBPositionReport.md) |  | [optional] 
**GnssBroadcastBinaryMessage** | Pointer to [**GnssBroadcastBinaryMessage**](GnssBroadcastBinaryMessage.md) |  | [optional] 
**GroupAssignmentCommand** | Pointer to [**GroupAssignmentCommand**](GroupAssignmentCommand.md) |  | [optional] 
**Interrogation** | Pointer to [**Interrogation**](Interrogation.md) |  | [optional] 
**LongRangeAisBroadcastMessage** | Pointer to [**LongRangeAisBroadcastMessage**](LongRangeAisBroadcastMessage.md) |  | [optional] 
**MultiSlotBinaryMessage** | Pointer to [**MultiSlotBinaryMessage**](MultiSlotBinaryMessage.md) |  | [optional] 
**SafetyBroadcastMessage** | Pointer to [**SafetyBroadcastMessage**](SafetyBroadcastMessage.md) |  | [optional] 
**ShipStaticData** | Pointer to [**ShipStaticData**](ShipStaticData.md) |  | [optional] 
**SingleSlotBinaryMessage** | Pointer to [**SingleSlotBinaryMessage**](SingleSlotBinaryMessage.md) |  | [optional] 
**StandardClassBPositionReport** | Pointer to [**StandardClassBPositionReport**](StandardClassBPositionReport.md) |  | [optional] 
**StandardSearchAndRescueAircraftReport** | Pointer to [**StandardSearchAndRescueAircraftReport**](StandardSearchAndRescueAircraftReport.md) |  | [optional] 
**StaticDataReport** | Pointer to [**StaticDataReport**](StaticDataReport.md) |  | [optional] 

## Methods

### NewAisStreamMessageMessage

`func NewAisStreamMessageMessage() *AisStreamMessageMessage`

NewAisStreamMessageMessage instantiates a new AisStreamMessageMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAisStreamMessageMessageWithDefaults

`func NewAisStreamMessageMessageWithDefaults() *AisStreamMessageMessage`

NewAisStreamMessageMessageWithDefaults instantiates a new AisStreamMessageMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPositionReport

`func (o *AisStreamMessageMessage) GetPositionReport() PositionReport`

GetPositionReport returns the PositionReport field if non-nil, zero value otherwise.

### GetPositionReportOk

`func (o *AisStreamMessageMessage) GetPositionReportOk() (*PositionReport, bool)`

GetPositionReportOk returns a tuple with the PositionReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionReport

`func (o *AisStreamMessageMessage) SetPositionReport(v PositionReport)`

SetPositionReport sets PositionReport field to given value.

### HasPositionReport

`func (o *AisStreamMessageMessage) HasPositionReport() bool`

HasPositionReport returns a boolean if a field has been set.

### GetUnknownMessage

`func (o *AisStreamMessageMessage) GetUnknownMessage() UnknownMessage`

GetUnknownMessage returns the UnknownMessage field if non-nil, zero value otherwise.

### GetUnknownMessageOk

`func (o *AisStreamMessageMessage) GetUnknownMessageOk() (*UnknownMessage, bool)`

GetUnknownMessageOk returns a tuple with the UnknownMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownMessage

`func (o *AisStreamMessageMessage) SetUnknownMessage(v UnknownMessage)`

SetUnknownMessage sets UnknownMessage field to given value.

### HasUnknownMessage

`func (o *AisStreamMessageMessage) HasUnknownMessage() bool`

HasUnknownMessage returns a boolean if a field has been set.

### GetAddressedSafetyMessage

`func (o *AisStreamMessageMessage) GetAddressedSafetyMessage() AddressedSafetyMessage`

GetAddressedSafetyMessage returns the AddressedSafetyMessage field if non-nil, zero value otherwise.

### GetAddressedSafetyMessageOk

`func (o *AisStreamMessageMessage) GetAddressedSafetyMessageOk() (*AddressedSafetyMessage, bool)`

GetAddressedSafetyMessageOk returns a tuple with the AddressedSafetyMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressedSafetyMessage

`func (o *AisStreamMessageMessage) SetAddressedSafetyMessage(v AddressedSafetyMessage)`

SetAddressedSafetyMessage sets AddressedSafetyMessage field to given value.

### HasAddressedSafetyMessage

`func (o *AisStreamMessageMessage) HasAddressedSafetyMessage() bool`

HasAddressedSafetyMessage returns a boolean if a field has been set.

### GetAddressedBinaryMessage

`func (o *AisStreamMessageMessage) GetAddressedBinaryMessage() AddressedBinaryMessage`

GetAddressedBinaryMessage returns the AddressedBinaryMessage field if non-nil, zero value otherwise.

### GetAddressedBinaryMessageOk

`func (o *AisStreamMessageMessage) GetAddressedBinaryMessageOk() (*AddressedBinaryMessage, bool)`

GetAddressedBinaryMessageOk returns a tuple with the AddressedBinaryMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressedBinaryMessage

`func (o *AisStreamMessageMessage) SetAddressedBinaryMessage(v AddressedBinaryMessage)`

SetAddressedBinaryMessage sets AddressedBinaryMessage field to given value.

### HasAddressedBinaryMessage

`func (o *AisStreamMessageMessage) HasAddressedBinaryMessage() bool`

HasAddressedBinaryMessage returns a boolean if a field has been set.

### GetAidsToNavigationReport

`func (o *AisStreamMessageMessage) GetAidsToNavigationReport() AidsToNavigationReport`

GetAidsToNavigationReport returns the AidsToNavigationReport field if non-nil, zero value otherwise.

### GetAidsToNavigationReportOk

`func (o *AisStreamMessageMessage) GetAidsToNavigationReportOk() (*AidsToNavigationReport, bool)`

GetAidsToNavigationReportOk returns a tuple with the AidsToNavigationReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAidsToNavigationReport

`func (o *AisStreamMessageMessage) SetAidsToNavigationReport(v AidsToNavigationReport)`

SetAidsToNavigationReport sets AidsToNavigationReport field to given value.

### HasAidsToNavigationReport

`func (o *AisStreamMessageMessage) HasAidsToNavigationReport() bool`

HasAidsToNavigationReport returns a boolean if a field has been set.

### GetAssignedModeCommand

`func (o *AisStreamMessageMessage) GetAssignedModeCommand() AssignedModeCommand`

GetAssignedModeCommand returns the AssignedModeCommand field if non-nil, zero value otherwise.

### GetAssignedModeCommandOk

`func (o *AisStreamMessageMessage) GetAssignedModeCommandOk() (*AssignedModeCommand, bool)`

GetAssignedModeCommandOk returns a tuple with the AssignedModeCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedModeCommand

`func (o *AisStreamMessageMessage) SetAssignedModeCommand(v AssignedModeCommand)`

SetAssignedModeCommand sets AssignedModeCommand field to given value.

### HasAssignedModeCommand

`func (o *AisStreamMessageMessage) HasAssignedModeCommand() bool`

HasAssignedModeCommand returns a boolean if a field has been set.

### GetBaseStationReport

`func (o *AisStreamMessageMessage) GetBaseStationReport() BaseStationReport`

GetBaseStationReport returns the BaseStationReport field if non-nil, zero value otherwise.

### GetBaseStationReportOk

`func (o *AisStreamMessageMessage) GetBaseStationReportOk() (*BaseStationReport, bool)`

GetBaseStationReportOk returns a tuple with the BaseStationReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStationReport

`func (o *AisStreamMessageMessage) SetBaseStationReport(v BaseStationReport)`

SetBaseStationReport sets BaseStationReport field to given value.

### HasBaseStationReport

`func (o *AisStreamMessageMessage) HasBaseStationReport() bool`

HasBaseStationReport returns a boolean if a field has been set.

### GetBinaryAcknowledge

`func (o *AisStreamMessageMessage) GetBinaryAcknowledge() BinaryAcknowledge`

GetBinaryAcknowledge returns the BinaryAcknowledge field if non-nil, zero value otherwise.

### GetBinaryAcknowledgeOk

`func (o *AisStreamMessageMessage) GetBinaryAcknowledgeOk() (*BinaryAcknowledge, bool)`

GetBinaryAcknowledgeOk returns a tuple with the BinaryAcknowledge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryAcknowledge

`func (o *AisStreamMessageMessage) SetBinaryAcknowledge(v BinaryAcknowledge)`

SetBinaryAcknowledge sets BinaryAcknowledge field to given value.

### HasBinaryAcknowledge

`func (o *AisStreamMessageMessage) HasBinaryAcknowledge() bool`

HasBinaryAcknowledge returns a boolean if a field has been set.

### GetBinaryBroadcastMessage

`func (o *AisStreamMessageMessage) GetBinaryBroadcastMessage() BinaryBroadcastMessage`

GetBinaryBroadcastMessage returns the BinaryBroadcastMessage field if non-nil, zero value otherwise.

### GetBinaryBroadcastMessageOk

`func (o *AisStreamMessageMessage) GetBinaryBroadcastMessageOk() (*BinaryBroadcastMessage, bool)`

GetBinaryBroadcastMessageOk returns a tuple with the BinaryBroadcastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryBroadcastMessage

`func (o *AisStreamMessageMessage) SetBinaryBroadcastMessage(v BinaryBroadcastMessage)`

SetBinaryBroadcastMessage sets BinaryBroadcastMessage field to given value.

### HasBinaryBroadcastMessage

`func (o *AisStreamMessageMessage) HasBinaryBroadcastMessage() bool`

HasBinaryBroadcastMessage returns a boolean if a field has been set.

### GetChannelManagement

`func (o *AisStreamMessageMessage) GetChannelManagement() ChannelManagement`

GetChannelManagement returns the ChannelManagement field if non-nil, zero value otherwise.

### GetChannelManagementOk

`func (o *AisStreamMessageMessage) GetChannelManagementOk() (*ChannelManagement, bool)`

GetChannelManagementOk returns a tuple with the ChannelManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelManagement

`func (o *AisStreamMessageMessage) SetChannelManagement(v ChannelManagement)`

SetChannelManagement sets ChannelManagement field to given value.

### HasChannelManagement

`func (o *AisStreamMessageMessage) HasChannelManagement() bool`

HasChannelManagement returns a boolean if a field has been set.

### GetCoordinatedUTCInquiry

`func (o *AisStreamMessageMessage) GetCoordinatedUTCInquiry() CoordinatedUTCInquiry`

GetCoordinatedUTCInquiry returns the CoordinatedUTCInquiry field if non-nil, zero value otherwise.

### GetCoordinatedUTCInquiryOk

`func (o *AisStreamMessageMessage) GetCoordinatedUTCInquiryOk() (*CoordinatedUTCInquiry, bool)`

GetCoordinatedUTCInquiryOk returns a tuple with the CoordinatedUTCInquiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinatedUTCInquiry

`func (o *AisStreamMessageMessage) SetCoordinatedUTCInquiry(v CoordinatedUTCInquiry)`

SetCoordinatedUTCInquiry sets CoordinatedUTCInquiry field to given value.

### HasCoordinatedUTCInquiry

`func (o *AisStreamMessageMessage) HasCoordinatedUTCInquiry() bool`

HasCoordinatedUTCInquiry returns a boolean if a field has been set.

### GetDataLinkManagementMessage

`func (o *AisStreamMessageMessage) GetDataLinkManagementMessage() DataLinkManagementMessage`

GetDataLinkManagementMessage returns the DataLinkManagementMessage field if non-nil, zero value otherwise.

### GetDataLinkManagementMessageOk

`func (o *AisStreamMessageMessage) GetDataLinkManagementMessageOk() (*DataLinkManagementMessage, bool)`

GetDataLinkManagementMessageOk returns a tuple with the DataLinkManagementMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLinkManagementMessage

`func (o *AisStreamMessageMessage) SetDataLinkManagementMessage(v DataLinkManagementMessage)`

SetDataLinkManagementMessage sets DataLinkManagementMessage field to given value.

### HasDataLinkManagementMessage

`func (o *AisStreamMessageMessage) HasDataLinkManagementMessage() bool`

HasDataLinkManagementMessage returns a boolean if a field has been set.

### GetDataLinkManagementMessageData

`func (o *AisStreamMessageMessage) GetDataLinkManagementMessageData() DataLinkManagementMessageData`

GetDataLinkManagementMessageData returns the DataLinkManagementMessageData field if non-nil, zero value otherwise.

### GetDataLinkManagementMessageDataOk

`func (o *AisStreamMessageMessage) GetDataLinkManagementMessageDataOk() (*DataLinkManagementMessageData, bool)`

GetDataLinkManagementMessageDataOk returns a tuple with the DataLinkManagementMessageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLinkManagementMessageData

`func (o *AisStreamMessageMessage) SetDataLinkManagementMessageData(v DataLinkManagementMessageData)`

SetDataLinkManagementMessageData sets DataLinkManagementMessageData field to given value.

### HasDataLinkManagementMessageData

`func (o *AisStreamMessageMessage) HasDataLinkManagementMessageData() bool`

HasDataLinkManagementMessageData returns a boolean if a field has been set.

### GetExtendedClassBPositionReport

`func (o *AisStreamMessageMessage) GetExtendedClassBPositionReport() ExtendedClassBPositionReport`

GetExtendedClassBPositionReport returns the ExtendedClassBPositionReport field if non-nil, zero value otherwise.

### GetExtendedClassBPositionReportOk

`func (o *AisStreamMessageMessage) GetExtendedClassBPositionReportOk() (*ExtendedClassBPositionReport, bool)`

GetExtendedClassBPositionReportOk returns a tuple with the ExtendedClassBPositionReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedClassBPositionReport

`func (o *AisStreamMessageMessage) SetExtendedClassBPositionReport(v ExtendedClassBPositionReport)`

SetExtendedClassBPositionReport sets ExtendedClassBPositionReport field to given value.

### HasExtendedClassBPositionReport

`func (o *AisStreamMessageMessage) HasExtendedClassBPositionReport() bool`

HasExtendedClassBPositionReport returns a boolean if a field has been set.

### GetGnssBroadcastBinaryMessage

`func (o *AisStreamMessageMessage) GetGnssBroadcastBinaryMessage() GnssBroadcastBinaryMessage`

GetGnssBroadcastBinaryMessage returns the GnssBroadcastBinaryMessage field if non-nil, zero value otherwise.

### GetGnssBroadcastBinaryMessageOk

`func (o *AisStreamMessageMessage) GetGnssBroadcastBinaryMessageOk() (*GnssBroadcastBinaryMessage, bool)`

GetGnssBroadcastBinaryMessageOk returns a tuple with the GnssBroadcastBinaryMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGnssBroadcastBinaryMessage

`func (o *AisStreamMessageMessage) SetGnssBroadcastBinaryMessage(v GnssBroadcastBinaryMessage)`

SetGnssBroadcastBinaryMessage sets GnssBroadcastBinaryMessage field to given value.

### HasGnssBroadcastBinaryMessage

`func (o *AisStreamMessageMessage) HasGnssBroadcastBinaryMessage() bool`

HasGnssBroadcastBinaryMessage returns a boolean if a field has been set.

### GetGroupAssignmentCommand

`func (o *AisStreamMessageMessage) GetGroupAssignmentCommand() GroupAssignmentCommand`

GetGroupAssignmentCommand returns the GroupAssignmentCommand field if non-nil, zero value otherwise.

### GetGroupAssignmentCommandOk

`func (o *AisStreamMessageMessage) GetGroupAssignmentCommandOk() (*GroupAssignmentCommand, bool)`

GetGroupAssignmentCommandOk returns a tuple with the GroupAssignmentCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAssignmentCommand

`func (o *AisStreamMessageMessage) SetGroupAssignmentCommand(v GroupAssignmentCommand)`

SetGroupAssignmentCommand sets GroupAssignmentCommand field to given value.

### HasGroupAssignmentCommand

`func (o *AisStreamMessageMessage) HasGroupAssignmentCommand() bool`

HasGroupAssignmentCommand returns a boolean if a field has been set.

### GetInterrogation

`func (o *AisStreamMessageMessage) GetInterrogation() Interrogation`

GetInterrogation returns the Interrogation field if non-nil, zero value otherwise.

### GetInterrogationOk

`func (o *AisStreamMessageMessage) GetInterrogationOk() (*Interrogation, bool)`

GetInterrogationOk returns a tuple with the Interrogation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterrogation

`func (o *AisStreamMessageMessage) SetInterrogation(v Interrogation)`

SetInterrogation sets Interrogation field to given value.

### HasInterrogation

`func (o *AisStreamMessageMessage) HasInterrogation() bool`

HasInterrogation returns a boolean if a field has been set.

### GetLongRangeAisBroadcastMessage

`func (o *AisStreamMessageMessage) GetLongRangeAisBroadcastMessage() LongRangeAisBroadcastMessage`

GetLongRangeAisBroadcastMessage returns the LongRangeAisBroadcastMessage field if non-nil, zero value otherwise.

### GetLongRangeAisBroadcastMessageOk

`func (o *AisStreamMessageMessage) GetLongRangeAisBroadcastMessageOk() (*LongRangeAisBroadcastMessage, bool)`

GetLongRangeAisBroadcastMessageOk returns a tuple with the LongRangeAisBroadcastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongRangeAisBroadcastMessage

`func (o *AisStreamMessageMessage) SetLongRangeAisBroadcastMessage(v LongRangeAisBroadcastMessage)`

SetLongRangeAisBroadcastMessage sets LongRangeAisBroadcastMessage field to given value.

### HasLongRangeAisBroadcastMessage

`func (o *AisStreamMessageMessage) HasLongRangeAisBroadcastMessage() bool`

HasLongRangeAisBroadcastMessage returns a boolean if a field has been set.

### GetMultiSlotBinaryMessage

`func (o *AisStreamMessageMessage) GetMultiSlotBinaryMessage() MultiSlotBinaryMessage`

GetMultiSlotBinaryMessage returns the MultiSlotBinaryMessage field if non-nil, zero value otherwise.

### GetMultiSlotBinaryMessageOk

`func (o *AisStreamMessageMessage) GetMultiSlotBinaryMessageOk() (*MultiSlotBinaryMessage, bool)`

GetMultiSlotBinaryMessageOk returns a tuple with the MultiSlotBinaryMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSlotBinaryMessage

`func (o *AisStreamMessageMessage) SetMultiSlotBinaryMessage(v MultiSlotBinaryMessage)`

SetMultiSlotBinaryMessage sets MultiSlotBinaryMessage field to given value.

### HasMultiSlotBinaryMessage

`func (o *AisStreamMessageMessage) HasMultiSlotBinaryMessage() bool`

HasMultiSlotBinaryMessage returns a boolean if a field has been set.

### GetSafetyBroadcastMessage

`func (o *AisStreamMessageMessage) GetSafetyBroadcastMessage() SafetyBroadcastMessage`

GetSafetyBroadcastMessage returns the SafetyBroadcastMessage field if non-nil, zero value otherwise.

### GetSafetyBroadcastMessageOk

`func (o *AisStreamMessageMessage) GetSafetyBroadcastMessageOk() (*SafetyBroadcastMessage, bool)`

GetSafetyBroadcastMessageOk returns a tuple with the SafetyBroadcastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyBroadcastMessage

`func (o *AisStreamMessageMessage) SetSafetyBroadcastMessage(v SafetyBroadcastMessage)`

SetSafetyBroadcastMessage sets SafetyBroadcastMessage field to given value.

### HasSafetyBroadcastMessage

`func (o *AisStreamMessageMessage) HasSafetyBroadcastMessage() bool`

HasSafetyBroadcastMessage returns a boolean if a field has been set.

### GetShipStaticData

`func (o *AisStreamMessageMessage) GetShipStaticData() ShipStaticData`

GetShipStaticData returns the ShipStaticData field if non-nil, zero value otherwise.

### GetShipStaticDataOk

`func (o *AisStreamMessageMessage) GetShipStaticDataOk() (*ShipStaticData, bool)`

GetShipStaticDataOk returns a tuple with the ShipStaticData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipStaticData

`func (o *AisStreamMessageMessage) SetShipStaticData(v ShipStaticData)`

SetShipStaticData sets ShipStaticData field to given value.

### HasShipStaticData

`func (o *AisStreamMessageMessage) HasShipStaticData() bool`

HasShipStaticData returns a boolean if a field has been set.

### GetSingleSlotBinaryMessage

`func (o *AisStreamMessageMessage) GetSingleSlotBinaryMessage() SingleSlotBinaryMessage`

GetSingleSlotBinaryMessage returns the SingleSlotBinaryMessage field if non-nil, zero value otherwise.

### GetSingleSlotBinaryMessageOk

`func (o *AisStreamMessageMessage) GetSingleSlotBinaryMessageOk() (*SingleSlotBinaryMessage, bool)`

GetSingleSlotBinaryMessageOk returns a tuple with the SingleSlotBinaryMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSlotBinaryMessage

`func (o *AisStreamMessageMessage) SetSingleSlotBinaryMessage(v SingleSlotBinaryMessage)`

SetSingleSlotBinaryMessage sets SingleSlotBinaryMessage field to given value.

### HasSingleSlotBinaryMessage

`func (o *AisStreamMessageMessage) HasSingleSlotBinaryMessage() bool`

HasSingleSlotBinaryMessage returns a boolean if a field has been set.

### GetStandardClassBPositionReport

`func (o *AisStreamMessageMessage) GetStandardClassBPositionReport() StandardClassBPositionReport`

GetStandardClassBPositionReport returns the StandardClassBPositionReport field if non-nil, zero value otherwise.

### GetStandardClassBPositionReportOk

`func (o *AisStreamMessageMessage) GetStandardClassBPositionReportOk() (*StandardClassBPositionReport, bool)`

GetStandardClassBPositionReportOk returns a tuple with the StandardClassBPositionReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardClassBPositionReport

`func (o *AisStreamMessageMessage) SetStandardClassBPositionReport(v StandardClassBPositionReport)`

SetStandardClassBPositionReport sets StandardClassBPositionReport field to given value.

### HasStandardClassBPositionReport

`func (o *AisStreamMessageMessage) HasStandardClassBPositionReport() bool`

HasStandardClassBPositionReport returns a boolean if a field has been set.

### GetStandardSearchAndRescueAircraftReport

`func (o *AisStreamMessageMessage) GetStandardSearchAndRescueAircraftReport() StandardSearchAndRescueAircraftReport`

GetStandardSearchAndRescueAircraftReport returns the StandardSearchAndRescueAircraftReport field if non-nil, zero value otherwise.

### GetStandardSearchAndRescueAircraftReportOk

`func (o *AisStreamMessageMessage) GetStandardSearchAndRescueAircraftReportOk() (*StandardSearchAndRescueAircraftReport, bool)`

GetStandardSearchAndRescueAircraftReportOk returns a tuple with the StandardSearchAndRescueAircraftReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardSearchAndRescueAircraftReport

`func (o *AisStreamMessageMessage) SetStandardSearchAndRescueAircraftReport(v StandardSearchAndRescueAircraftReport)`

SetStandardSearchAndRescueAircraftReport sets StandardSearchAndRescueAircraftReport field to given value.

### HasStandardSearchAndRescueAircraftReport

`func (o *AisStreamMessageMessage) HasStandardSearchAndRescueAircraftReport() bool`

HasStandardSearchAndRescueAircraftReport returns a boolean if a field has been set.

### GetStaticDataReport

`func (o *AisStreamMessageMessage) GetStaticDataReport() StaticDataReport`

GetStaticDataReport returns the StaticDataReport field if non-nil, zero value otherwise.

### GetStaticDataReportOk

`func (o *AisStreamMessageMessage) GetStaticDataReportOk() (*StaticDataReport, bool)`

GetStaticDataReportOk returns a tuple with the StaticDataReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDataReport

`func (o *AisStreamMessageMessage) SetStaticDataReport(v StaticDataReport)`

SetStaticDataReport sets StaticDataReport field to given value.

### HasStaticDataReport

`func (o *AisStreamMessageMessage) HasStaticDataReport() bool`

HasStaticDataReport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


