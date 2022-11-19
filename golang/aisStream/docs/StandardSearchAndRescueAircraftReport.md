# StandardSearchAndRescueAircraftReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Altitude** | **int32** |  | 
**Sog** | **float64** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Cog** | **float64** |  | 
**Timestamp** | **int32** |  | 
**AltFromBaro** | **bool** |  | 
**Spare1** | **int32** |  | 
**Dte** | **bool** |  | 
**Spare2** | **int32** |  | 
**AssignedMode** | **bool** |  | 
**Raim** | **bool** |  | 
**CommunicationStateIsItdma** | **bool** |  | 
**CommunicationState** | **int32** |  | 

## Methods

### NewStandardSearchAndRescueAircraftReport

`func NewStandardSearchAndRescueAircraftReport(messageID int32, repeatIndicator int32, userID int32, valid bool, altitude int32, sog float64, positionAccuracy bool, longitude float64, latitude float64, cog float64, timestamp int32, altFromBaro bool, spare1 int32, dte bool, spare2 int32, assignedMode bool, raim bool, communicationStateIsItdma bool, communicationState int32, ) *StandardSearchAndRescueAircraftReport`

NewStandardSearchAndRescueAircraftReport instantiates a new StandardSearchAndRescueAircraftReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardSearchAndRescueAircraftReportWithDefaults

`func NewStandardSearchAndRescueAircraftReportWithDefaults() *StandardSearchAndRescueAircraftReport`

NewStandardSearchAndRescueAircraftReportWithDefaults instantiates a new StandardSearchAndRescueAircraftReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *StandardSearchAndRescueAircraftReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *StandardSearchAndRescueAircraftReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *StandardSearchAndRescueAircraftReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *StandardSearchAndRescueAircraftReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *StandardSearchAndRescueAircraftReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *StandardSearchAndRescueAircraftReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *StandardSearchAndRescueAircraftReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *StandardSearchAndRescueAircraftReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *StandardSearchAndRescueAircraftReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *StandardSearchAndRescueAircraftReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *StandardSearchAndRescueAircraftReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *StandardSearchAndRescueAircraftReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetAltitude

`func (o *StandardSearchAndRescueAircraftReport) GetAltitude() int32`

GetAltitude returns the Altitude field if non-nil, zero value otherwise.

### GetAltitudeOk

`func (o *StandardSearchAndRescueAircraftReport) GetAltitudeOk() (*int32, bool)`

GetAltitudeOk returns a tuple with the Altitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltitude

`func (o *StandardSearchAndRescueAircraftReport) SetAltitude(v int32)`

SetAltitude sets Altitude field to given value.


### GetSog

`func (o *StandardSearchAndRescueAircraftReport) GetSog() float64`

GetSog returns the Sog field if non-nil, zero value otherwise.

### GetSogOk

`func (o *StandardSearchAndRescueAircraftReport) GetSogOk() (*float64, bool)`

GetSogOk returns a tuple with the Sog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSog

`func (o *StandardSearchAndRescueAircraftReport) SetSog(v float64)`

SetSog sets Sog field to given value.


### GetPositionAccuracy

`func (o *StandardSearchAndRescueAircraftReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *StandardSearchAndRescueAircraftReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *StandardSearchAndRescueAircraftReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *StandardSearchAndRescueAircraftReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *StandardSearchAndRescueAircraftReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *StandardSearchAndRescueAircraftReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *StandardSearchAndRescueAircraftReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *StandardSearchAndRescueAircraftReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *StandardSearchAndRescueAircraftReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetCog

`func (o *StandardSearchAndRescueAircraftReport) GetCog() float64`

GetCog returns the Cog field if non-nil, zero value otherwise.

### GetCogOk

`func (o *StandardSearchAndRescueAircraftReport) GetCogOk() (*float64, bool)`

GetCogOk returns a tuple with the Cog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCog

`func (o *StandardSearchAndRescueAircraftReport) SetCog(v float64)`

SetCog sets Cog field to given value.


### GetTimestamp

`func (o *StandardSearchAndRescueAircraftReport) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StandardSearchAndRescueAircraftReport) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StandardSearchAndRescueAircraftReport) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetAltFromBaro

`func (o *StandardSearchAndRescueAircraftReport) GetAltFromBaro() bool`

GetAltFromBaro returns the AltFromBaro field if non-nil, zero value otherwise.

### GetAltFromBaroOk

`func (o *StandardSearchAndRescueAircraftReport) GetAltFromBaroOk() (*bool, bool)`

GetAltFromBaroOk returns a tuple with the AltFromBaro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltFromBaro

`func (o *StandardSearchAndRescueAircraftReport) SetAltFromBaro(v bool)`

SetAltFromBaro sets AltFromBaro field to given value.


### GetSpare1

`func (o *StandardSearchAndRescueAircraftReport) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *StandardSearchAndRescueAircraftReport) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *StandardSearchAndRescueAircraftReport) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetDte

`func (o *StandardSearchAndRescueAircraftReport) GetDte() bool`

GetDte returns the Dte field if non-nil, zero value otherwise.

### GetDteOk

`func (o *StandardSearchAndRescueAircraftReport) GetDteOk() (*bool, bool)`

GetDteOk returns a tuple with the Dte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDte

`func (o *StandardSearchAndRescueAircraftReport) SetDte(v bool)`

SetDte sets Dte field to given value.


### GetSpare2

`func (o *StandardSearchAndRescueAircraftReport) GetSpare2() int32`

GetSpare2 returns the Spare2 field if non-nil, zero value otherwise.

### GetSpare2Ok

`func (o *StandardSearchAndRescueAircraftReport) GetSpare2Ok() (*int32, bool)`

GetSpare2Ok returns a tuple with the Spare2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare2

`func (o *StandardSearchAndRescueAircraftReport) SetSpare2(v int32)`

SetSpare2 sets Spare2 field to given value.


### GetAssignedMode

`func (o *StandardSearchAndRescueAircraftReport) GetAssignedMode() bool`

GetAssignedMode returns the AssignedMode field if non-nil, zero value otherwise.

### GetAssignedModeOk

`func (o *StandardSearchAndRescueAircraftReport) GetAssignedModeOk() (*bool, bool)`

GetAssignedModeOk returns a tuple with the AssignedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedMode

`func (o *StandardSearchAndRescueAircraftReport) SetAssignedMode(v bool)`

SetAssignedMode sets AssignedMode field to given value.


### GetRaim

`func (o *StandardSearchAndRescueAircraftReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *StandardSearchAndRescueAircraftReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *StandardSearchAndRescueAircraftReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetCommunicationStateIsItdma

`func (o *StandardSearchAndRescueAircraftReport) GetCommunicationStateIsItdma() bool`

GetCommunicationStateIsItdma returns the CommunicationStateIsItdma field if non-nil, zero value otherwise.

### GetCommunicationStateIsItdmaOk

`func (o *StandardSearchAndRescueAircraftReport) GetCommunicationStateIsItdmaOk() (*bool, bool)`

GetCommunicationStateIsItdmaOk returns a tuple with the CommunicationStateIsItdma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationStateIsItdma

`func (o *StandardSearchAndRescueAircraftReport) SetCommunicationStateIsItdma(v bool)`

SetCommunicationStateIsItdma sets CommunicationStateIsItdma field to given value.


### GetCommunicationState

`func (o *StandardSearchAndRescueAircraftReport) GetCommunicationState() int32`

GetCommunicationState returns the CommunicationState field if non-nil, zero value otherwise.

### GetCommunicationStateOk

`func (o *StandardSearchAndRescueAircraftReport) GetCommunicationStateOk() (*int32, bool)`

GetCommunicationStateOk returns a tuple with the CommunicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationState

`func (o *StandardSearchAndRescueAircraftReport) SetCommunicationState(v int32)`

SetCommunicationState sets CommunicationState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


