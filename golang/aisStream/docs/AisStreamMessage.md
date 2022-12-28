# AisStreamMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaData** | **map[string]interface{}** |  | 
**MessageType** | [**AisMessageTypes**](AisMessageTypes.md) |  | 
**Message** | [**AisStreamMessageMessage**](AisStreamMessageMessage.md) |  | 

## Methods

### NewAisStreamMessage

`func NewAisStreamMessage(metaData map[string]interface{}, messageType AisMessageTypes, message AisStreamMessageMessage, ) *AisStreamMessage`

NewAisStreamMessage instantiates a new AisStreamMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAisStreamMessageWithDefaults

`func NewAisStreamMessageWithDefaults() *AisStreamMessage`

NewAisStreamMessageWithDefaults instantiates a new AisStreamMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaData

`func (o *AisStreamMessage) GetMetaData() map[string]interface{}`

GetMetaData returns the MetaData field if non-nil, zero value otherwise.

### GetMetaDataOk

`func (o *AisStreamMessage) GetMetaDataOk() (*map[string]interface{}, bool)`

GetMetaDataOk returns a tuple with the MetaData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaData

`func (o *AisStreamMessage) SetMetaData(v map[string]interface{})`

SetMetaData sets MetaData field to given value.


### GetMessageType

`func (o *AisStreamMessage) GetMessageType() AisMessageTypes`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *AisStreamMessage) GetMessageTypeOk() (*AisMessageTypes, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *AisStreamMessage) SetMessageType(v AisMessageTypes)`

SetMessageType sets MessageType field to given value.


### GetMessage

`func (o *AisStreamMessage) GetMessage() AisStreamMessageMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AisStreamMessage) GetMessageOk() (*AisStreamMessageMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AisStreamMessage) SetMessage(v AisStreamMessageMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


