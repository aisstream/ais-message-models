# AddressedBinaryMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Sequenceinteger** | **int32** |  | 
**DestinationID** | **int32** |  | 
**Retransmission** | **bool** |  | 
**Spare** | **bool** |  | 
**ApplicationID** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**BinaryData** | **string** |  | 

## Methods

### NewAddressedBinaryMessage

`func NewAddressedBinaryMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, sequenceinteger int32, destinationID int32, retransmission bool, spare bool, applicationID AddressedBinaryMessageApplicationID, binaryData string, ) *AddressedBinaryMessage`

NewAddressedBinaryMessage instantiates a new AddressedBinaryMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressedBinaryMessageWithDefaults

`func NewAddressedBinaryMessageWithDefaults() *AddressedBinaryMessage`

NewAddressedBinaryMessageWithDefaults instantiates a new AddressedBinaryMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *AddressedBinaryMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *AddressedBinaryMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *AddressedBinaryMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *AddressedBinaryMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *AddressedBinaryMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *AddressedBinaryMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *AddressedBinaryMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AddressedBinaryMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AddressedBinaryMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *AddressedBinaryMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AddressedBinaryMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AddressedBinaryMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSequenceinteger

`func (o *AddressedBinaryMessage) GetSequenceinteger() int32`

GetSequenceinteger returns the Sequenceinteger field if non-nil, zero value otherwise.

### GetSequenceintegerOk

`func (o *AddressedBinaryMessage) GetSequenceintegerOk() (*int32, bool)`

GetSequenceintegerOk returns a tuple with the Sequenceinteger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceinteger

`func (o *AddressedBinaryMessage) SetSequenceinteger(v int32)`

SetSequenceinteger sets Sequenceinteger field to given value.


### GetDestinationID

`func (o *AddressedBinaryMessage) GetDestinationID() int32`

GetDestinationID returns the DestinationID field if non-nil, zero value otherwise.

### GetDestinationIDOk

`func (o *AddressedBinaryMessage) GetDestinationIDOk() (*int32, bool)`

GetDestinationIDOk returns a tuple with the DestinationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationID

`func (o *AddressedBinaryMessage) SetDestinationID(v int32)`

SetDestinationID sets DestinationID field to given value.


### GetRetransmission

`func (o *AddressedBinaryMessage) GetRetransmission() bool`

GetRetransmission returns the Retransmission field if non-nil, zero value otherwise.

### GetRetransmissionOk

`func (o *AddressedBinaryMessage) GetRetransmissionOk() (*bool, bool)`

GetRetransmissionOk returns a tuple with the Retransmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmission

`func (o *AddressedBinaryMessage) SetRetransmission(v bool)`

SetRetransmission sets Retransmission field to given value.


### GetSpare

`func (o *AddressedBinaryMessage) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *AddressedBinaryMessage) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *AddressedBinaryMessage) SetSpare(v bool)`

SetSpare sets Spare field to given value.


### GetApplicationID

`func (o *AddressedBinaryMessage) GetApplicationID() AddressedBinaryMessageApplicationID`

GetApplicationID returns the ApplicationID field if non-nil, zero value otherwise.

### GetApplicationIDOk

`func (o *AddressedBinaryMessage) GetApplicationIDOk() (*AddressedBinaryMessageApplicationID, bool)`

GetApplicationIDOk returns a tuple with the ApplicationID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationID

`func (o *AddressedBinaryMessage) SetApplicationID(v AddressedBinaryMessageApplicationID)`

SetApplicationID sets ApplicationID field to given value.


### GetBinaryData

`func (o *AddressedBinaryMessage) GetBinaryData() string`

GetBinaryData returns the BinaryData field if non-nil, zero value otherwise.

### GetBinaryDataOk

`func (o *AddressedBinaryMessage) GetBinaryDataOk() (*string, bool)`

GetBinaryDataOk returns a tuple with the BinaryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryData

`func (o *AddressedBinaryMessage) SetBinaryData(v string)`

SetBinaryData sets BinaryData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


