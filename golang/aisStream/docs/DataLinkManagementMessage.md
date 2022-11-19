# DataLinkManagementMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare** | **int32** |  | 
**Data** | [**DataLinkManagementMessageData**](DataLinkManagementMessageData.md) |  | 

## Methods

### NewDataLinkManagementMessage

`func NewDataLinkManagementMessage(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, data DataLinkManagementMessageData, ) *DataLinkManagementMessage`

NewDataLinkManagementMessage instantiates a new DataLinkManagementMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLinkManagementMessageWithDefaults

`func NewDataLinkManagementMessageWithDefaults() *DataLinkManagementMessage`

NewDataLinkManagementMessageWithDefaults instantiates a new DataLinkManagementMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *DataLinkManagementMessage) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *DataLinkManagementMessage) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *DataLinkManagementMessage) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *DataLinkManagementMessage) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *DataLinkManagementMessage) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *DataLinkManagementMessage) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *DataLinkManagementMessage) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *DataLinkManagementMessage) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *DataLinkManagementMessage) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *DataLinkManagementMessage) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *DataLinkManagementMessage) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *DataLinkManagementMessage) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare

`func (o *DataLinkManagementMessage) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *DataLinkManagementMessage) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *DataLinkManagementMessage) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetData

`func (o *DataLinkManagementMessage) GetData() DataLinkManagementMessageData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataLinkManagementMessage) GetDataOk() (*DataLinkManagementMessageData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataLinkManagementMessage) SetData(v DataLinkManagementMessageData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


