# ChannelManagement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare1** | **int32** |  | 
**ChannelA** | **int32** |  | 
**ChannelB** | **int32** |  | 
**TxRxMode** | **int32** |  | 
**LowPower** | **bool** |  | 
**Area** | [**ChannelManagementArea**](ChannelManagementArea.md) |  | 
**Unicast** | [**ChannelManagementUnicast**](ChannelManagementUnicast.md) |  | 
**IsAddressed** | **bool** |  | 
**BwA** | **bool** |  | 
**BwB** | **bool** |  | 
**TransitionalZoneSize** | **int32** |  | 
**Spare4** | **int32** |  | 

## Methods

### NewChannelManagement

`func NewChannelManagement(messageID int32, repeatIndicator int32, userID int32, valid bool, spare1 int32, channelA int32, channelB int32, txRxMode int32, lowPower bool, area ChannelManagementArea, unicast ChannelManagementUnicast, isAddressed bool, bwA bool, bwB bool, transitionalZoneSize int32, spare4 int32, ) *ChannelManagement`

NewChannelManagement instantiates a new ChannelManagement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelManagementWithDefaults

`func NewChannelManagementWithDefaults() *ChannelManagement`

NewChannelManagementWithDefaults instantiates a new ChannelManagement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *ChannelManagement) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *ChannelManagement) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *ChannelManagement) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *ChannelManagement) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *ChannelManagement) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *ChannelManagement) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *ChannelManagement) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ChannelManagement) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ChannelManagement) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *ChannelManagement) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ChannelManagement) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ChannelManagement) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare1

`func (o *ChannelManagement) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *ChannelManagement) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *ChannelManagement) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetChannelA

`func (o *ChannelManagement) GetChannelA() int32`

GetChannelA returns the ChannelA field if non-nil, zero value otherwise.

### GetChannelAOk

`func (o *ChannelManagement) GetChannelAOk() (*int32, bool)`

GetChannelAOk returns a tuple with the ChannelA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelA

`func (o *ChannelManagement) SetChannelA(v int32)`

SetChannelA sets ChannelA field to given value.


### GetChannelB

`func (o *ChannelManagement) GetChannelB() int32`

GetChannelB returns the ChannelB field if non-nil, zero value otherwise.

### GetChannelBOk

`func (o *ChannelManagement) GetChannelBOk() (*int32, bool)`

GetChannelBOk returns a tuple with the ChannelB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelB

`func (o *ChannelManagement) SetChannelB(v int32)`

SetChannelB sets ChannelB field to given value.


### GetTxRxMode

`func (o *ChannelManagement) GetTxRxMode() int32`

GetTxRxMode returns the TxRxMode field if non-nil, zero value otherwise.

### GetTxRxModeOk

`func (o *ChannelManagement) GetTxRxModeOk() (*int32, bool)`

GetTxRxModeOk returns a tuple with the TxRxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRxMode

`func (o *ChannelManagement) SetTxRxMode(v int32)`

SetTxRxMode sets TxRxMode field to given value.


### GetLowPower

`func (o *ChannelManagement) GetLowPower() bool`

GetLowPower returns the LowPower field if non-nil, zero value otherwise.

### GetLowPowerOk

`func (o *ChannelManagement) GetLowPowerOk() (*bool, bool)`

GetLowPowerOk returns a tuple with the LowPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowPower

`func (o *ChannelManagement) SetLowPower(v bool)`

SetLowPower sets LowPower field to given value.


### GetArea

`func (o *ChannelManagement) GetArea() ChannelManagementArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ChannelManagement) GetAreaOk() (*ChannelManagementArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ChannelManagement) SetArea(v ChannelManagementArea)`

SetArea sets Area field to given value.


### GetUnicast

`func (o *ChannelManagement) GetUnicast() ChannelManagementUnicast`

GetUnicast returns the Unicast field if non-nil, zero value otherwise.

### GetUnicastOk

`func (o *ChannelManagement) GetUnicastOk() (*ChannelManagementUnicast, bool)`

GetUnicastOk returns a tuple with the Unicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnicast

`func (o *ChannelManagement) SetUnicast(v ChannelManagementUnicast)`

SetUnicast sets Unicast field to given value.


### GetIsAddressed

`func (o *ChannelManagement) GetIsAddressed() bool`

GetIsAddressed returns the IsAddressed field if non-nil, zero value otherwise.

### GetIsAddressedOk

`func (o *ChannelManagement) GetIsAddressedOk() (*bool, bool)`

GetIsAddressedOk returns a tuple with the IsAddressed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAddressed

`func (o *ChannelManagement) SetIsAddressed(v bool)`

SetIsAddressed sets IsAddressed field to given value.


### GetBwA

`func (o *ChannelManagement) GetBwA() bool`

GetBwA returns the BwA field if non-nil, zero value otherwise.

### GetBwAOk

`func (o *ChannelManagement) GetBwAOk() (*bool, bool)`

GetBwAOk returns a tuple with the BwA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwA

`func (o *ChannelManagement) SetBwA(v bool)`

SetBwA sets BwA field to given value.


### GetBwB

`func (o *ChannelManagement) GetBwB() bool`

GetBwB returns the BwB field if non-nil, zero value otherwise.

### GetBwBOk

`func (o *ChannelManagement) GetBwBOk() (*bool, bool)`

GetBwBOk returns a tuple with the BwB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwB

`func (o *ChannelManagement) SetBwB(v bool)`

SetBwB sets BwB field to given value.


### GetTransitionalZoneSize

`func (o *ChannelManagement) GetTransitionalZoneSize() int32`

GetTransitionalZoneSize returns the TransitionalZoneSize field if non-nil, zero value otherwise.

### GetTransitionalZoneSizeOk

`func (o *ChannelManagement) GetTransitionalZoneSizeOk() (*int32, bool)`

GetTransitionalZoneSizeOk returns a tuple with the TransitionalZoneSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionalZoneSize

`func (o *ChannelManagement) SetTransitionalZoneSize(v int32)`

SetTransitionalZoneSize sets TransitionalZoneSize field to given value.


### GetSpare4

`func (o *ChannelManagement) GetSpare4() int32`

GetSpare4 returns the Spare4 field if non-nil, zero value otherwise.

### GetSpare4Ok

`func (o *ChannelManagement) GetSpare4Ok() (*int32, bool)`

GetSpare4Ok returns a tuple with the Spare4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare4

`func (o *ChannelManagement) SetSpare4(v int32)`

SetSpare4 sets Spare4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


