# LongRangeAisBroadcastMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**PositionAccuracy** | **bool** |  | 
**Raim** | **bool** |  | 
**NavigationalStatus** | **int32** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Sog** | **float64** |  | 
**Cog** | **float64** |  | 
**PositionLatency** | **bool** |  | 
**Spare** | **bool** |  | 

## Methods

### NewLongRangeAisBroadcastMessage

`func NewLongRangeAisBroadcastMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, positionAccuracy bool, raim bool, navigationalStatus int32, longitude float64, latitude float64, sog float64, cog float64, positionLatency bool, spare bool, ) *LongRangeAisBroadcastMessage`

NewLongRangeAisBroadcastMessage instantiates a new LongRangeAisBroadcastMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongRangeAisBroadcastMessageWithDefaults

`func NewLongRangeAisBroadcastMessageWithDefaults() *LongRangeAisBroadcastMessage`

NewLongRangeAisBroadcastMessageWithDefaults instantiates a new LongRangeAisBroadcastMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *LongRangeAisBroadcastMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *LongRangeAisBroadcastMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *LongRangeAisBroadcastMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *LongRangeAisBroadcastMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *LongRangeAisBroadcastMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *LongRangeAisBroadcastMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *LongRangeAisBroadcastMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *LongRangeAisBroadcastMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *LongRangeAisBroadcastMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *LongRangeAisBroadcastMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *LongRangeAisBroadcastMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *LongRangeAisBroadcastMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetPositionAccuracy

`func (o *LongRangeAisBroadcastMessage) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *LongRangeAisBroadcastMessage) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *LongRangeAisBroadcastMessage) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetRaim

`func (o *LongRangeAisBroadcastMessage) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *LongRangeAisBroadcastMessage) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *LongRangeAisBroadcastMessage) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetNavigationalStatus

`func (o *LongRangeAisBroadcastMessage) GetNavigationalStatus() int32`

GetNavigationalStatus returns the NavigationalStatus field if non-nil, zero value otherwise.

### GetNavigationalStatusOk

`func (o *LongRangeAisBroadcastMessage) GetNavigationalStatusOk() (*int32, bool)`

GetNavigationalStatusOk returns a tuple with the NavigationalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigationalStatus

`func (o *LongRangeAisBroadcastMessage) SetNavigationalStatus(v int32)`

SetNavigationalStatus sets NavigationalStatus field to given value.


### GetLongitude

`func (o *LongRangeAisBroadcastMessage) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *LongRangeAisBroadcastMessage) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *LongRangeAisBroadcastMessage) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *LongRangeAisBroadcastMessage) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *LongRangeAisBroadcastMessage) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *LongRangeAisBroadcastMessage) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetSog

`func (o *LongRangeAisBroadcastMessage) GetSog() float64`

GetSog returns the Sog field if non-nil, zero value otherwise.

### GetSogOk

`func (o *LongRangeAisBroadcastMessage) GetSogOk() (*float64, bool)`

GetSogOk returns a tuple with the Sog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSog

`func (o *LongRangeAisBroadcastMessage) SetSog(v float64)`

SetSog sets Sog field to given value.


### GetCog

`func (o *LongRangeAisBroadcastMessage) GetCog() float64`

GetCog returns the Cog field if non-nil, zero value otherwise.

### GetCogOk

`func (o *LongRangeAisBroadcastMessage) GetCogOk() (*float64, bool)`

GetCogOk returns a tuple with the Cog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCog

`func (o *LongRangeAisBroadcastMessage) SetCog(v float64)`

SetCog sets Cog field to given value.


### GetPositionLatency

`func (o *LongRangeAisBroadcastMessage) GetPositionLatency() bool`

GetPositionLatency returns the PositionLatency field if non-nil, zero value otherwise.

### GetPositionLatencyOk

`func (o *LongRangeAisBroadcastMessage) GetPositionLatencyOk() (*bool, bool)`

GetPositionLatencyOk returns a tuple with the PositionLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionLatency

`func (o *LongRangeAisBroadcastMessage) SetPositionLatency(v bool)`

SetPositionLatency sets PositionLatency field to given value.


### GetSpare

`func (o *LongRangeAisBroadcastMessage) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *LongRangeAisBroadcastMessage) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *LongRangeAisBroadcastMessage) SetSpare(v bool)`

SetSpare sets Spare field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


