# Interrogation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare** | **int32** |  | 
**Station1Msg1** | [**InterrogationStation1Msg1**](InterrogationStation1Msg1.md) |  | 
**Station1Msg2** | [**InterrogationStation1Msg2**](InterrogationStation1Msg2.md) |  | 
**Station2** | [**InterrogationStation2**](InterrogationStation2.md) |  | 

## Methods

### NewInterrogation

`func NewInterrogation(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, station1Msg1 InterrogationStation1Msg1, station1Msg2 InterrogationStation1Msg2, station2 InterrogationStation2, ) *Interrogation`

NewInterrogation instantiates a new Interrogation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterrogationWithDefaults

`func NewInterrogationWithDefaults() *Interrogation`

NewInterrogationWithDefaults instantiates a new Interrogation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *Interrogation) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *Interrogation) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *Interrogation) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *Interrogation) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *Interrogation) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *Interrogation) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *Interrogation) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Interrogation) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Interrogation) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *Interrogation) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *Interrogation) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *Interrogation) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare

`func (o *Interrogation) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *Interrogation) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *Interrogation) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetStation1Msg1

`func (o *Interrogation) GetStation1Msg1() InterrogationStation1Msg1`

GetStation1Msg1 returns the Station1Msg1 field if non-nil, zero value otherwise.

### GetStation1Msg1Ok

`func (o *Interrogation) GetStation1Msg1Ok() (*InterrogationStation1Msg1, bool)`

GetStation1Msg1Ok returns a tuple with the Station1Msg1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStation1Msg1

`func (o *Interrogation) SetStation1Msg1(v InterrogationStation1Msg1)`

SetStation1Msg1 sets Station1Msg1 field to given value.


### GetStation1Msg2

`func (o *Interrogation) GetStation1Msg2() InterrogationStation1Msg2`

GetStation1Msg2 returns the Station1Msg2 field if non-nil, zero value otherwise.

### GetStation1Msg2Ok

`func (o *Interrogation) GetStation1Msg2Ok() (*InterrogationStation1Msg2, bool)`

GetStation1Msg2Ok returns a tuple with the Station1Msg2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStation1Msg2

`func (o *Interrogation) SetStation1Msg2(v InterrogationStation1Msg2)`

SetStation1Msg2 sets Station1Msg2 field to given value.


### GetStation2

`func (o *Interrogation) GetStation2() InterrogationStation2`

GetStation2 returns the Station2 field if non-nil, zero value otherwise.

### GetStation2Ok

`func (o *Interrogation) GetStation2Ok() (*InterrogationStation2, bool)`

GetStation2Ok returns a tuple with the Station2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStation2

`func (o *Interrogation) SetStation2(v InterrogationStation2)`

SetStation2 sets Station2 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


