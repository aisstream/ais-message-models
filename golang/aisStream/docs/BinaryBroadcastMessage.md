# BinaryBroadcastMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare** | **int32** |  | 
**ApplicationID** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**BinaryData** | **string** |  | 

## Methods

### NewBinaryBroadcastMessage

`func NewBinaryBroadcastMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, applicationID AddressedBinaryMessageApplicationID, binaryData string, ) *BinaryBroadcastMessage`

NewBinaryBroadcastMessage instantiates a new BinaryBroadcastMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryBroadcastMessageWithDefaults

`func NewBinaryBroadcastMessageWithDefaults() *BinaryBroadcastMessage`

NewBinaryBroadcastMessageWithDefaults instantiates a new BinaryBroadcastMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *BinaryBroadcastMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *BinaryBroadcastMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *BinaryBroadcastMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *BinaryBroadcastMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *BinaryBroadcastMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *BinaryBroadcastMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *BinaryBroadcastMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *BinaryBroadcastMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *BinaryBroadcastMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *BinaryBroadcastMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *BinaryBroadcastMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *BinaryBroadcastMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare

`func (o *BinaryBroadcastMessage) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *BinaryBroadcastMessage) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *BinaryBroadcastMessage) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetApplicationID

`func (o *BinaryBroadcastMessage) GetApplicationID() AddressedBinaryMessageApplicationID`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *BinaryBroadcastMessage) GetApplicationIDOk() (*AddressedBinaryMessageApplicationID, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *BinaryBroadcastMessage) SetApplicationID(v AddressedBinaryMessageApplicationID)`

SetApplicationID sets ApplicationID field to given value.


### GetBinaryData

`func (o *BinaryBroadcastMessage) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *BinaryBroadcastMessage) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *BinaryBroadcastMessage) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


