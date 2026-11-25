# SubscriptionConfirmation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageType** | **string** |  | 
**Message** | [**SubscriptionConfirmationMessage**](SubscriptionConfirmationMessage.md) |  | 

## Methods

### NewSubscriptionConfirmation

`func NewSubscriptionConfirmation(messageType string, message SubscriptionConfirmationMessage, ) *SubscriptionConfirmation`

NewSubscriptionConfirmation instantiates a new SubscriptionConfirmation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionConfirmationWithDefaults

`func NewSubscriptionConfirmationWithDefaults() *SubscriptionConfirmation`

NewSubscriptionConfirmationWithDefaults instantiates a new SubscriptionConfirmation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageType

`func (o *SubscriptionConfirmation) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SubscriptionConfirmation) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SubscriptionConfirmation) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetMessage

`func (o *SubscriptionConfirmation) GetMessage() SubscriptionConfirmationMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SubscriptionConfirmation) GetMessageOk() (*SubscriptionConfirmationMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SubscriptionConfirmation) SetMessage(v SubscriptionConfirmationMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


