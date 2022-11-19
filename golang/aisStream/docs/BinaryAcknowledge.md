# BinaryAcknowledge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare** | **int32** |  | 
**Destinations** | [**BinaryAcknowledgeDestinations**](BinaryAcknowledgeDestinations.md) |  | 

## Methods

### NewBinaryAcknowledge

`func NewBinaryAcknowledge(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, destinations BinaryAcknowledgeDestinations, ) *BinaryAcknowledge`

NewBinaryAcknowledge instantiates a new BinaryAcknowledge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryAcknowledgeWithDefaults

`func NewBinaryAcknowledgeWithDefaults() *BinaryAcknowledge`

NewBinaryAcknowledgeWithDefaults instantiates a new BinaryAcknowledge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *BinaryAcknowledge) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *BinaryAcknowledge) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *BinaryAcknowledge) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *BinaryAcknowledge) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *BinaryAcknowledge) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *BinaryAcknowledge) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *BinaryAcknowledge) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *BinaryAcknowledge) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *BinaryAcknowledge) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *BinaryAcknowledge) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *BinaryAcknowledge) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *BinaryAcknowledge) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare

`func (o *BinaryAcknowledge) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *BinaryAcknowledge) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *BinaryAcknowledge) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetDestinations

`func (o *BinaryAcknowledge) GetDestinations() BinaryAcknowledgeDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *BinaryAcknowledge) GetDestinationsOk() (*BinaryAcknowledgeDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *BinaryAcknowledge) SetDestinations(v BinaryAcknowledgeDestinations)`

SetDestinations sets Destinations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


