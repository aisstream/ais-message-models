# BaseStationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**UtcYear** | **int32** |  | 
**UtcMonth** | **int32** |  | 
**UtcDay** | **int32** |  | 
**UtcHour** | **int32** |  | 
**UtcMinute** | **int32** |  | 
**UtcSecond** | **int32** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**FixType** | **int32** |  | 
**LongRangeEnable** | **bool** |  | 
**Spare** | **int32** |  | 
**Raim** | **bool** |  | 
**CommunicationState** | **int32** |  | 

## Methods

### NewBaseStationReport

`func NewBaseStationReport(messageID int32, repeatIndicator int32, userID int32, valid bool, utcYear int32, utcMonth int32, utcDay int32, utcHour int32, utcMinute int32, utcSecond int32, positionAccuracy bool, longitude float64, latitude float64, fixType int32, longRangeEnable bool, spare int32, raim bool, communicationState int32, ) *BaseStationReport`

NewBaseStationReport instantiates a new BaseStationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseStationReportWithDefaults

`func NewBaseStationReportWithDefaults() *BaseStationReport`

NewBaseStationReportWithDefaults instantiates a new BaseStationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *BaseStationReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *BaseStationReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *BaseStationReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *BaseStationReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *BaseStationReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *BaseStationReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *BaseStationReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *BaseStationReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *BaseStationReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *BaseStationReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *BaseStationReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *BaseStationReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetUtcYear

`func (o *BaseStationReport) GetUtcYear() int32`

GetUtcYear returns the UtcYear field if non-nil, zero value otherwise.

### GetUtcYearOk

`func (o *BaseStationReport) GetUtcYearOk() (*int32, bool)`

GetUtcYearOk returns a tuple with the UtcYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcYear

`func (o *BaseStationReport) SetUtcYear(v int32)`

SetUtcYear sets UtcYear field to given value.


### GetUtcMonth

`func (o *BaseStationReport) GetUtcMonth() int32`

GetUtcMonth returns the UtcMonth field if non-nil, zero value otherwise.

### GetUtcMonthOk

`func (o *BaseStationReport) GetUtcMonthOk() (*int32, bool)`

GetUtcMonthOk returns a tuple with the UtcMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcMonth

`func (o *BaseStationReport) SetUtcMonth(v int32)`

SetUtcMonth sets UtcMonth field to given value.


### GetUtcDay

`func (o *BaseStationReport) GetUtcDay() int32`

GetUtcDay returns the UtcDay field if non-nil, zero value otherwise.

### GetUtcDayOk

`func (o *BaseStationReport) GetUtcDayOk() (*int32, bool)`

GetUtcDayOk returns a tuple with the UtcDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcDay

`func (o *BaseStationReport) SetUtcDay(v int32)`

SetUtcDay sets UtcDay field to given value.


### GetUtcHour

`func (o *BaseStationReport) GetUtcHour() int32`

GetUtcHour returns the UtcHour field if non-nil, zero value otherwise.

### GetUtcHourOk

`func (o *BaseStationReport) GetUtcHourOk() (*int32, bool)`

GetUtcHourOk returns a tuple with the UtcHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcHour

`func (o *BaseStationReport) SetUtcHour(v int32)`

SetUtcHour sets UtcHour field to given value.


### GetUtcMinute

`func (o *BaseStationReport) GetUtcMinute() int32`

GetUtcMinute returns the UtcMinute field if non-nil, zero value otherwise.

### GetUtcMinuteOk

`func (o *BaseStationReport) GetUtcMinuteOk() (*int32, bool)`

GetUtcMinuteOk returns a tuple with the UtcMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcMinute

`func (o *BaseStationReport) SetUtcMinute(v int32)`

SetUtcMinute sets UtcMinute field to given value.


### GetUtcSecond

`func (o *BaseStationReport) GetUtcSecond() int32`

GetUtcSecond returns the UtcSecond field if non-nil, zero value otherwise.

### GetUtcSecondOk

`func (o *BaseStationReport) GetUtcSecondOk() (*int32, bool)`

GetUtcSecondOk returns a tuple with the UtcSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtcSecond

`func (o *BaseStationReport) SetUtcSecond(v int32)`

SetUtcSecond sets UtcSecond field to given value.


### GetPositionAccuracy

`func (o *BaseStationReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *BaseStationReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *BaseStationReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *BaseStationReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *BaseStationReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *BaseStationReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *BaseStationReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *BaseStationReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *BaseStationReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetFixType

`func (o *BaseStationReport) GetFixType() int32`

GetFixType returns the FixType field if non-nil, zero value otherwise.

### GetFixTypeOk

`func (o *BaseStationReport) GetFixTypeOk() (*int32, bool)`

GetFixTypeOk returns a tuple with the FixType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixType

`func (o *BaseStationReport) SetFixType(v int32)`

SetFixType sets FixType field to given value.


### GetLongRangeEnable

`func (o *BaseStationReport) GetLongRangeEnable() bool`

GetLongRangeEnable returns the LongRangeEnable field if non-nil, zero value otherwise.

### GetLongRangeEnableOk

`func (o *BaseStationReport) GetLongRangeEnableOk() (*bool, bool)`

GetLongRangeEnableOk returns a tuple with the LongRangeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongRangeEnable

`func (o *BaseStationReport) SetLongRangeEnable(v bool)`

SetLongRangeEnable sets LongRangeEnable field to given value.


### GetSpare

`func (o *BaseStationReport) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *BaseStationReport) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *BaseStationReport) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetRaim

`func (o *BaseStationReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *BaseStationReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *BaseStationReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetCommunicationState

`func (o *BaseStationReport) GetCommunicationState() int32`

GetCommunicationState returns the CommunicationState field if non-nil, zero value otherwise.

### GetCommunicationStateOk

`func (o *BaseStationReport) GetCommunicationStateOk() (*int32, bool)`

GetCommunicationStateOk returns a tuple with the CommunicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationState

`func (o *BaseStationReport) SetCommunicationState(v int32)`

SetCommunicationState sets CommunicationState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


