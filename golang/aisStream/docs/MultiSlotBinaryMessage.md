# MultiSlotBinaryMessage

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
**Spare1** | **int32** |  | 
**ApplicationID** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**Payload** | **string** |  | 
**Spare2** | **int32** |  | 
**CommunicationStateIsItdma** | **bool** |  | 
**CommunicationState** | **int32** |  | 

## Methods

### NewMultiSlotBinaryMessage

`func NewMultiSlotBinaryMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, destinationIDValid bool, applicationIDValid bool, destinationID int32, spare1 int32, applicationID AddressedBinaryMessageApplicationID, payload string, spare2 int32, communicationStateIsItdma bool, communicationState int32, ) *MultiSlotBinaryMessage`

NewMultiSlotBinaryMessage instantiates a new MultiSlotBinaryMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiSlotBinaryMessageWithDefaults

`func NewMultiSlotBinaryMessageWithDefaults() *MultiSlotBinaryMessage`

NewMultiSlotBinaryMessageWithDefaults instantiates a new MultiSlotBinaryMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *MultiSlotBinaryMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *MultiSlotBinaryMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *MultiSlotBinaryMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *MultiSlotBinaryMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *MultiSlotBinaryMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *MultiSlotBinaryMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *MultiSlotBinaryMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *MultiSlotBinaryMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *MultiSlotBinaryMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *MultiSlotBinaryMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *MultiSlotBinaryMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *MultiSlotBinaryMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetDestinationIDValid

`func (o *MultiSlotBinaryMessage) GetDestinationIDValid() bool`

GetDestinationIDValid returns the DestinationIDValid field if non-nil, zero value otherwise.

### GetDestinationIDValidOk

`func (o *MultiSlotBinaryMessage) GetDestinationIDValidOk() (*bool, bool)`

GetDestinationIDValidOk returns a tuple with the DestinationIDValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationIDValid

`func (o *MultiSlotBinaryMessage) SetDestinationIDValid(v bool)`

SetDestinationIDValid sets DestinationIDValid field to given value.


### GetApplicationIDValid

`func (o *MultiSlotBinaryMessage) GetApplicationIDValid() bool`

GetApplicationIDValid returns the ApplicationIDValid field if non-nil, zero value otherwise.

### GetApplicationIDValidOk

`func (o *MultiSlotBinaryMessage) GetApplicationIDValidOk() (*bool, bool)`

GetApplicationIDValidOk returns a tuple with the ApplicationIDValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIDValid

`func (o *MultiSlotBinaryMessage) SetApplicationIDValid(v bool)`

SetApplicationIDValid sets ApplicationIDValid field to given value.


### GetDestinationID

`func (o *MultiSlotBinaryMessage) GetDestinationID() int32`

GetDestinationID returns the DestinationID field if non-nil, zero value otherwise.

### GetDestinationIDOk

`func (o *MultiSlotBinaryMessage) GetDestinationIDOk() (*int32, bool)`

GetDestinationIDOk returns a tuple with the DestinationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationID

`func (o *MultiSlotBinaryMessage) SetDestinationID(v int32)`

SetDestinationID sets DestinationID field to given value.


### GetSpare1

`func (o *MultiSlotBinaryMessage) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *MultiSlotBinaryMessage) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *MultiSlotBinaryMessage) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetApplicationID

`func (o *MultiSlotBinaryMessage) GetApplicationID() AddressedBinaryMessageApplicationID`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *MultiSlotBinaryMessage) GetApplicationIDOk() (*AddressedBinaryMessageApplicationID, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *MultiSlotBinaryMessage) SetApplicationID(v AddressedBinaryMessageApplicationID)`

SetApplicationID sets ApplicationID field to given value.


### GetPayload

`func (o *MultiSlotBinaryMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *MultiSlotBinaryMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *MultiSlotBinaryMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetSpare2

`func (o *MultiSlotBinaryMessage) GetSpare2() int32`

GetSpare2 returns the Spare2 field if non-nil, zero value otherwise.

### GetSpare2Ok

`func (o *MultiSlotBinaryMessage) GetSpare2Ok() (*int32, bool)`

GetSpare2Ok returns a tuple with the Spare2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare2

`func (o *MultiSlotBinaryMessage) SetSpare2(v int32)`

SetSpare2 sets Spare2 field to given value.


### GetCommunicationStateIsItdma

`func (o *MultiSlotBinaryMessage) GetCommunicationStateIsItdma() bool`

GetCommunicationStateIsItdma returns the CommunicationStateIsItdma field if non-nil, zero value otherwise.

### GetCommunicationStateIsItdmaOk

`func (o *MultiSlotBinaryMessage) GetCommunicationStateIsItdmaOk() (*bool, bool)`

GetCommunicationStateIsItdmaOk returns a tuple with the CommunicationStateIsItdma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationStateIsItdma

`func (o *MultiSlotBinaryMessage) SetCommunicationStateIsItdma(v bool)`

SetCommunicationStateIsItdma sets CommunicationStateIsItdma field to given value.


### GetCommunicationState

`func (o *MultiSlotBinaryMessage) GetCommunicationState() int32`

GetCommunicationState returns the CommunicationState field if non-nil, zero value otherwise.

### GetCommunicationStateOk

`func (o *MultiSlotBinaryMessage) GetCommunicationStateOk() (*int32, bool)`

GetCommunicationStateOk returns a tuple with the CommunicationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicationState

`func (o *MultiSlotBinaryMessage) SetCommunicationState(v int32)`

SetCommunicationState sets CommunicationState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


