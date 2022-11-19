# StandardClassBPositionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare1** | **int32** |  | 
**Sog** | **float64** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Cog** | **float64** |  | 
**TrueHeading** | **int32** |  | 
**Timestamp** | **int32** |  | 
**Spare2** | **int32** |  | 
**ClassBUnit** | **bool** |  | 
**ClassBDisplay** | **bool** |  | 
**ClassBDsc** | **bool** |  | 
**ClassBBand** | **bool** |  | 
**ClassBMsg22** | **bool** |  | 
**AssignedMode** | **bool** |  | 
**Raim** | **bool** |  | 
**CommunicationStateIsItdma** | **bool** |  | 
**CommunicationState** | **int32** |  | 

## Methods

### NewStandardClassBPositionReport

`func NewStandardClassBPositionReport(messageID int32, repeatIndicator int32, userID int32, valid bool, spare1 int32, sog float64, positionAccuracy bool, longitude float64, latitude float64, cog float64, trueHeading int32, timestamp int32, spare2 int32, classBUnit bool, classBDisplay bool, classBDsc bool, classBBand bool, classBMsg22 bool, assignedMode bool, raim bool, communicationStateIsItdma bool, communicationState int32, ) *StandardClassBPositionReport`

NewStandardClassBPositionReport instantiates a new StandardClassBPositionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStandardClassBPositionReportWithDefaults

`func NewStandardClassBPositionReportWithDefaults() *StandardClassBPositionReport`

NewStandardClassBPositionReportWithDefaults instantiates a new StandardClassBPositionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *StandardClassBPositionReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *StandardClassBPositionReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *StandardClassBPositionReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *StandardClassBPositionReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *StandardClassBPositionReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *StandardClassBPositionReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *StandardClassBPositionReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *StandardClassBPositionReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *StandardClassBPositionReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *StandardClassBPositionReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *StandardClassBPositionReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *StandardClassBPositionReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare1

`func (o *StandardClassBPositionReport) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *StandardClassBPositionReport) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *StandardClassBPositionReport) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetSog

`func (o *StandardClassBPositionReport) GetSog() float64`

GetSog returns the Sog field if non-nil, zero value otherwise.

### GetSogOk

`func (o *StandardClassBPositionReport) GetSogOk() (*float64, bool)`

GetSogOk returns a tuple with the Sog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSog

`func (o *StandardClassBPositionReport) SetSog(v float64)`

SetSog sets Sog field to given value.


### GetPositionAccuracy

`func (o *StandardClassBPositionReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *StandardClassBPositionReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *StandardClassBPositionReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *StandardClassBPositionReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *StandardClassBPositionReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *StandardClassBPositionReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *StandardClassBPositionReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *StandardClassBPositionReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *StandardClassBPositionReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetCog

`func (o *StandardClassBPositionReport) GetCog() float64`

GetCog returns the Cog field if non-nil, zero value otherwise.

### GetCogOk

`func (o *StandardClassBPositionReport) GetCogOk() (*float64, bool)`

GetCogOk returns a tuple with the Cog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCog

`func (o *StandardClassBPositionReport) SetCog(v float64)`

SetCog sets Cog field to given value.


### GetTrueHeading

`func (o *StandardClassBPositionReport) GetTrueHeading() int32`

GetTrueHeading returns the TrueHeading field if non-nil, zero value otherwise.

### GetTrueHeadingOk

`func (o *StandardClassBPositionReport) GetTrueHeadingOk() (*int32, bool)`

GetTrueHeadingOk returns a tuple with the TrueHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueHeading

`func (o *StandardClassBPositionReport) SetTrueHeading(v int32)`

SetTrueHeading sets TrueHeading field to given value.


### GetTimestamp

`func (o *StandardClassBPositionReport) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StandardClassBPositionReport) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StandardClassBPositionReport) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSpare2

`func (o *StandardClassBPositionReport) GetSpare2() int32`

GetSpare2 returns the Spare2 field if non-nil, zero value otherwise.

### GetSpare2Ok

`func (o *StandardClassBPositionReport) GetSpare2Ok() (*int32, bool)`

GetSpare2Ok returns a tuple with the Spare2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare2

`func (o *StandardClassBPositionReport) SetSpare2(v int32)`

SetSpare2 sets Spare2 field to given value.


### GetClassBUnit

`func (o *StandardClassBPositionReport) GetClassBUnit() bool`

GetClassBUnit returns the ClassBUnit field if non-nil, zero value otherwise.

### GetClassBUnitOk

`func (o *StandardClassBPositionReport) GetClassBUnitOk() (*bool, bool)`

GetClassBUnitOk returns a tuple with the ClassBUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassBUnit

`func (o *StandardClassBPositionReport) SetClassBUnit(v bool)`

SetClassBUnit sets ClassBUnit field to given value.


### GetClassBDisplay

`func (o *StandardClassBPositionReport) GetClassBDisplay() bool`

GetClassBDisplay returns the ClassBDisplay field if non-nil, zero value otherwise.

### GetClassBDisplayOk

`func (o *StandardClassBPositionReport) GetClassBDisplayOk() (*bool, bool)`

GetClassBDisplayOk returns a tuple with the ClassBDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassBDisplay

`func (o *StandardClassBPositionReport) SetClassBDisplay(v bool)`

SetClassBDisplay sets ClassBDisplay field to given value.


### GetClassBDsc

`func (o *StandardClassBPositionReport) GetClassBDsc() bool`

GetClassBDsc returns the ClassBDsc field if non-nil, zero value otherwise.

### GetClassBDscOk

`func (o *StandardClassBPositionReport) GetClassBDscOk() (*bool, bool)`

GetClassBDscOk returns a tuple with the ClassBDsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassBDsc

`func (o *StandardClassBPositionReport) SetClassBDsc(v bool)`

SetClassBDsc sets ClassBDsc field to given value.


### GetClassBBand

`func (o *StandardClassBPositionReport) GetClassBBand() bool`

GetClassBBand returns the ClassBBand field if non-nil, zero value otherwise.

### GetClassBBandOk

`func (o *StandardClassBPositionReport) GetClassBBandOk() (*bool, bool)`

GetClassBBandOk returns a tuple with the ClassBBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassBBand

`func (o *StandardClassBPositionReport) SetClassBBand(v bool)`

SetClassBBand sets ClassBBand field to given value.


### GetClassBMsg22

`func (o *StandardClassBPositionReport) GetClassBMsg22() bool`

GetClassBMsg22 returns the ClassBMsg22 field if non-nil, zero value otherwise.

### GetClassBMsg22Ok

`func (o *StandardClassBPositionReport) GetClassBMsg22Ok() (*bool, bool)`

GetClassBMsg22Ok returns a tuple with the ClassBMsg22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassBMsg22

`func (o *StandardClassBPositionReport) SetClassBMsg22(v bool)`

SetClassBMsg22 sets ClassBMsg22 field to given value.


### GetAssignedMode

`func (o *StandardClassBPositionReport) GetAssignedMode() bool`

GetAssignedMode returns the AssignedMode field if non-nil, zero value otherwise.

### GetAssignedModeOk

`func (o *StandardClassBPositionReport) GetAssignedModeOk() (*bool, bool)`

GetAssignedModeOk returns a tuple with the AssignedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedMode

`func (o *StandardClassBPositionReport) SetAssignedMode(v bool)`

SetAssignedMode sets AssignedMode field to given value.


### GetRaim

`func (o *StandardClassBPositionReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *StandardClassBPositionReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *StandardClassBPositionReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetCommunicationStateIsItdma

`func (o *StandardClassBPositionReport) GetCommunicationStateIsItdma() bool`

GetCommunicationStateIsItdma returns the CommunicationStateIsItdma field if non-nil, zero value otherwise.

### GetCommunicationStateIsItdmaOk

`func (o *StandardClassBPositionReport) GetCommunicationStateIsItdmaOk() (*bool, bool)`

GetCommunicationStateIsItdmaOk returns a tuple with the CommunicationStateIsItdma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationStateIsItdma

`func (o *StandardClassBPositionReport) SetCommunicationStateIsItdma(v bool)`

SetCommunicationStateIsItdma sets CommunicationStateIsItdma field to given value.


### GetCommunicationState

`func (o *StandardClassBPositionReport) GetCommunicationState() int32`

GetCommunicationState returns the CommunicationState field if non-nil, zero value otherwise.

### GetCommunicationStateOk

`func (o *StandardClassBPositionReport) GetCommunicationStateOk() (*int32, bool)`

GetCommunicationStateOk returns a tuple with the CommunicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationState

`func (o *StandardClassBPositionReport) SetCommunicationState(v int32)`

SetCommunicationState sets CommunicationState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


