# AssignedModeCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare** | **int32** |  | 
**Commands** | [**AssignedModeCommandCommands**](AssignedModeCommandCommands.md) |  | 

## Methods

### NewAssignedModeCommand

`func NewAssignedModeCommand(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, commands AssignedModeCommandCommands, ) *AssignedModeCommand`

NewAssignedModeCommand instantiates a new AssignedModeCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignedModeCommandWithDefaults

`func NewAssignedModeCommandWithDefaults() *AssignedModeCommand`

NewAssignedModeCommandWithDefaults instantiates a new AssignedModeCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *AssignedModeCommand) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *AssignedModeCommand) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *AssignedModeCommand) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *AssignedModeCommand) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *AssignedModeCommand) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *AssignedModeCommand) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *AssignedModeCommand) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AssignedModeCommand) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AssignedModeCommand) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *AssignedModeCommand) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AssignedModeCommand) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AssignedModeCommand) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare

`func (o *AssignedModeCommand) GetSpare() int32`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *AssignedModeCommand) GetSpareOk() (*int32, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *AssignedModeCommand) SetSpare(v int32)`

SetSpare sets Spare field to given value.


### GetCommands

`func (o *AssignedModeCommand) GetCommands() AssignedModeCommandCommands`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *AssignedModeCommand) GetCommandsOk() (*AssignedModeCommandCommands, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *AssignedModeCommand) SetCommands(v AssignedModeCommandCommands)`

SetCommands sets Commands field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


