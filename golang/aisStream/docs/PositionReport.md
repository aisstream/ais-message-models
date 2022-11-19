# PositionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**NavigationalStatus** | **int32** |  | 
**RateOfTurn** | **int32** |  | 
**Sog** | **float64** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Cog** | **float64** |  | 
**TrueHeading** | **int32** |  | 
**Timestamp** | **int32** |  | 
**SpecialManoeuvreIndicator** | **int32** |  | 
**Spare** | **int32** |  | 
**Raim** | **bool** |  | 
**CommunicationState** | **int32** |  | 

## Methods

### NewPositionReport

`func NewPositionReport(messageID int32, repeatIndicator int32, userID int32, valid bool, navigationalStatus int32, rateOfTurn int32, sog float64, positionAccuracy bool, longitude float64, latitude float64, cog float64, trueHeading int32, timestamp int32, specialManoeuvreIndicator int32, spare int32, raim bool, communicationState int32, ) *PositionReport`

NewPositionReport instantiates a new PositionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionReportWithDefaults

`func NewPositionReportWithDefaults() *PositionReport`

NewPositionReportWithDefaults instantiates a new PositionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *PositionReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *PositionReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *PositionReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *PositionReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *PositionReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *PositionReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *PositionReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *PositionReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *PositionReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *PositionReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *PositionReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *PositionReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetNavigationalStatus

`func (o *PositionReport) GetNavigationalStatus() int32`

GetNavigationalStatus returns the NavigationalStatus field if non-nil, zero value otherwise.

### GetNavigationalStatusOk

`func (o *PositionReport) GetNavigationalStatusOk() (*int32, bool)`

GetNavigationalStatusOk returns a tuple with the NavigationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationalStatus

`func (o *PositionReport) SetNavigationalStatus(v int32)`

SetNavigationalStatus sets NavigationalStatus field to given value.


### GetRateOfTurn

`func (o *PositionReport) GetRateOfTurn() int32`

GetRateOfTurn returns the RateOfTurn field if non-nil, zero value otherwise.

### GetRateOfTurnOk

`func (o *PositionReport) GetRateOfTurnOk() (*int32, bool)`

GetRateOfTurnOk returns a tuple with the RateOfTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateOfTurn

`func (o *PositionReport) SetRateOfTurn(v int32)`

SetRateOfTurn sets RateOfTurn field to given value.


### GetSog

`func (o *PositionReport) GetSog() float64`

GetSog returns the Sog field if non-nil, zero value otherwise.

### GetSogOk

`func (o *PositionReport) GetSogOk() (*float64, bool)`

GetSogOk returns a tuple with the Sog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSog

`func (o *PositionReport) SetSog(v float64)`

SetSog sets Sog field to given value.


### GetPositionAccuracy

`func (o *PositionReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *PositionReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *PositionReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *PositionReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *PositionReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *PositionReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *PositionReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *PositionReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *PositionReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetCog

`func (o *PositionReport) GetCog() float64`

GetCog returns the Cog field if non-nil, zero value otherwise.

### GetCogOk

`func (o *PositionReport) GetCogOk() (*float64, bool)`

GetCogOk returns a tuple with the Cog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCog

`func (o *PositionReport) SetCog(v float64)`

SetCog sets Cog field to given value.


### GetTrueHeading

`func (o *PositionReport) GetTrueHeading() int32`

GetTrueHeading returns the TrueHeading field if non-nil, zero value otherwise.

### GetTrueHeadingOk

`func (o *PositionReport) GetTrueHeadingOk() (*int32, bool)`

GetTrueHeadingOk returns a tuple with the TrueHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueHeading

`func (o *PositionReport) SetTrueHeading(v int32)`

SetTrueHeading sets TrueHeading field to given value.


### GetTimestamp

`func (o *PositionReport) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PositionReport) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PositionReport) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSpecialManoeuvreIndicator

`func (o *PositionReport) GetSpecialManoeuvreIndicator() int32`

GetSpecialManoeuvreIndicator returns the SpecialManoeuvreIndicator field if non-nil, zero value otherwise.

### GetSpecialManoeuvreIndicatorOk

`func (o *PositionReport) GetSpecialManoeuvreIndicatorOk() (*int32, bool)`

GetSpecialManoeuvreIndicatorOk returns a tuple with the SpecialManoeuvreIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialManoeuvreIndicator

`func (o *PositionReport) SetSpecialManoeuvreIndicator(v int32)`

SetSpecialManoeuvreIndicator sets SpecialManoeuvreIndicator field to given value.


### GetSpare

`func (o *PositionReport) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *PositionReport) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *PositionReport) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetRaim

`func (o *PositionReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *PositionReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *PositionReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetCommunicationState

`func (o *PositionReport) GetCommunicationState() int32`

GetCommunicationState returns the CommunicationState field if non-nil, zero value otherwise.

### GetCommunicationStateOk

`func (o *PositionReport) GetCommunicationStateOk() (*int32, bool)`

GetCommunicationStateOk returns a tuple with the CommunicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationState

`func (o *PositionReport) SetCommunicationState(v int32)`

SetCommunicationState sets CommunicationState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


