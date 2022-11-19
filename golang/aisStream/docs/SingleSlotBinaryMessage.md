# SingleSlotBinaryMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**DestinationIDValid** | **bool** |  | 
**ApplicationIDValid** | **bool** |  | 
**DestinationID** | **int32** |  | 
**Spare** | **int32** |  | 
**ApplicationID** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**Payload** | **string** |  | 

## Methods

### NewSingleSlotBinaryMessage

`func NewSingleSlotBinaryMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, destinationIDValid bool, applicationIDValid bool, destinationID int32, spare int32, applicationID AddressedBinaryMessageApplicationID, payload string, ) *SingleSlotBinaryMessage`

NewSingleSlotBinaryMessage instantiates a new SingleSlotBinaryMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleSlotBinaryMessageWithDefaults

`func NewSingleSlotBinaryMessageWithDefaults() *SingleSlotBinaryMessage`

NewSingleSlotBinaryMessageWithDefaults instantiates a new SingleSlotBinaryMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *SingleSlotBinaryMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *SingleSlotBinaryMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *SingleSlotBinaryMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *SingleSlotBinaryMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *SingleSlotBinaryMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *SingleSlotBinaryMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *SingleSlotBinaryMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *SingleSlotBinaryMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *SingleSlotBinaryMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *SingleSlotBinaryMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SingleSlotBinaryMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SingleSlotBinaryMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetDestinationIDValid

`func (o *SingleSlotBinaryMessage) GetDestinationIDValid() bool`

GetDestinationIDValid returns the DestinationIDValid field if non-nil, zero value otherwise.

### GetDestinationIDValidOk

`func (o *SingleSlotBinaryMessage) GetDestinationIDValidOk() (*bool, bool)`

GetDestinationIDValidOk returns a tuple with the DestinationIDValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIDValid

`func (o *SingleSlotBinaryMessage) SetDestinationIDValid(v bool)`

SetDestinationIDValid sets DestinationIDValid field to given value.


### GetApplicationIDValid

`func (o *SingleSlotBinaryMessage) GetApplicationIDValid() bool`

GetApplicationIDValid returns the ApplicationIDValid field if non-nil, zero value otherwise.

### GetApplicationIDValidOk

`func (o *SingleSlotBinaryMessage) GetApplicationIDValidOk() (*bool, bool)`

GetApplicationIDValidOk returns a tuple with the ApplicationIDValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIDValid

`func (o *SingleSlotBinaryMessage) SetApplicationIDValid(v bool)`

SetApplicationIDValid sets ApplicationIDValid field to given value.


### GetDestinationID

`func (o *SingleSlotBinaryMessage) GetDestinationID() int32`

GetDestinationID returns the DestinationID field if non-nil, zero value otherwise.

### GetDestinationIDOk

`func (o *SingleSlotBinaryMessage) GetDestinationIDOk() (*int32, bool)`

GetDestinationIDOk returns a tuple with the DestinationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationID

`func (o *SingleSlotBinaryMessage) SetDestinationID(v int32)`

SetDestinationID sets DestinationID field to given value.


### GetSpare

`func (o *SingleSlotBinaryMessage) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *SingleSlotBinaryMessage) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *SingleSlotBinaryMessage) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetApplicationID

`func (o *SingleSlotBinaryMessage) GetApplicationID() AddressedBinaryMessageApplicationID`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *SingleSlotBinaryMessage) GetApplicationIDOk() (*AddressedBinaryMessageApplicationID, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *SingleSlotBinaryMessage) SetApplicationID(v AddressedBinaryMessageApplicationID)`

SetApplicationID sets ApplicationID field to given value.


### GetPayload

`func (o *SingleSlotBinaryMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SingleSlotBinaryMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SingleSlotBinaryMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


