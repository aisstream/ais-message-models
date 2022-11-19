# AidsToNavigationReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Type** | **int32** |  | 
**Name** | **string** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**Fixtype** | **int32** |  | 
**Timestamp** | **int32** |  | 
**OffPosition** | **bool** |  | 
**AtoN** | **int32** |  | 
**Raim** | **bool** |  | 
**VirtualAtoN** | **bool** |  | 
**AssignedMode** | **bool** |  | 
**Spare** | **bool** |  | 
**NameExtension** | **string** |  | 

## Methods

### NewAidsToNavigationReport

`func NewAidsToNavigationReport(messageID int32, repeatIndicator int32, userID int32, valid bool, type_ int32, name string, positionAccuracy bool, longitude float64, latitude float64, dimension ShipStaticDataDimension, fixtype int32, timestamp int32, offPosition bool, atoN int32, raim bool, virtualAtoN bool, assignedMode bool, spare bool, nameExtension string, ) *AidsToNavigationReport`

NewAidsToNavigationReport instantiates a new AidsToNavigationReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAidsToNavigationReportWithDefaults

`func NewAidsToNavigationReportWithDefaults() *AidsToNavigationReport`

NewAidsToNavigationReportWithDefaults instantiates a new AidsToNavigationReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *AidsToNavigationReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *AidsToNavigationReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *AidsToNavigationReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *AidsToNavigationReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *AidsToNavigationReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *AidsToNavigationReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *AidsToNavigationReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *AidsToNavigationReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *AidsToNavigationReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *AidsToNavigationReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AidsToNavigationReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AidsToNavigationReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetType

`func (o *AidsToNavigationReport) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AidsToNavigationReport) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AidsToNavigationReport) SetType(v int32)`

SetType sets Type field to given value.


### GetName

`func (o *AidsToNavigationReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AidsToNavigationReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AidsToNavigationReport) SetName(v string)`

SetName sets Name field to given value.


### GetPositionAccuracy

`func (o *AidsToNavigationReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *AidsToNavigationReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *AidsToNavigationReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *AidsToNavigationReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AidsToNavigationReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AidsToNavigationReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *AidsToNavigationReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AidsToNavigationReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AidsToNavigationReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetDimension

`func (o *AidsToNavigationReport) GetDimension() ShipStaticDataDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *AidsToNavigationReport) GetDimensionOk() (*ShipStaticDataDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *AidsToNavigationReport) SetDimension(v ShipStaticDataDimension)`

SetDimension sets Dimension field to given value.


### GetFixtype

`func (o *AidsToNavigationReport) GetFixtype() int32`

GetFixtype returns the Fixtype field if non-nil, zero value otherwise.

### GetFixtypeOk

`func (o *AidsToNavigationReport) GetFixtypeOk() (*int32, bool)`

GetFixtypeOk returns a tuple with the Fixtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixtype

`func (o *AidsToNavigationReport) SetFixtype(v int32)`

SetFixtype sets Fixtype field to given value.


### GetTimestamp

`func (o *AidsToNavigationReport) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AidsToNavigationReport) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AidsToNavigationReport) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetOffPosition

`func (o *AidsToNavigationReport) GetOffPosition() bool`

GetOffPosition returns the OffPosition field if non-nil, zero value otherwise.

### GetOffPositionOk

`func (o *AidsToNavigationReport) GetOffPositionOk() (*bool, bool)`

GetOffPositionOk returns a tuple with the OffPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffPosition

`func (o *AidsToNavigationReport) SetOffPosition(v bool)`

SetOffPosition sets OffPosition field to given value.


### GetAtoN

`func (o *AidsToNavigationReport) GetAtoN() int32`

GetAtoN returns the AtoN field if non-nil, zero value otherwise.

### GetAtoNOk

`func (o *AidsToNavigationReport) GetAtoNOk() (*int32, bool)`

GetAtoNOk returns a tuple with the AtoN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtoN

`func (o *AidsToNavigationReport) SetAtoN(v int32)`

SetAtoN sets AtoN field to given value.


### GetRaim

`func (o *AidsToNavigationReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *AidsToNavigationReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *AidsToNavigationReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetVirtualAtoN

`func (o *AidsToNavigationReport) GetVirtualAtoN() bool`

GetVirtualAtoN returns the VirtualAtoN field if non-nil, zero value otherwise.

### GetVirtualAtoNOk

`func (o *AidsToNavigationReport) GetVirtualAtoNOk() (*bool, bool)`

GetVirtualAtoNOk returns a tuple with the VirtualAtoN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAtoN

`func (o *AidsToNavigationReport) SetVirtualAtoN(v bool)`

SetVirtualAtoN sets VirtualAtoN field to given value.


### GetAssignedMode

`func (o *AidsToNavigationReport) GetAssignedMode() bool`

GetAssignedMode returns the AssignedMode field if non-nil, zero value otherwise.

### GetAssignedModeOk

`func (o *AidsToNavigationReport) GetAssignedModeOk() (*bool, bool)`

GetAssignedModeOk returns a tuple with the AssignedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedMode

`func (o *AidsToNavigationReport) SetAssignedMode(v bool)`

SetAssignedMode sets AssignedMode field to given value.


### GetSpare

`func (o *AidsToNavigationReport) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *AidsToNavigationReport) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *AidsToNavigationReport) SetSpare(v bool)`

SetSpare sets Spare field to given value.


### GetNameExtension

`func (o *AidsToNavigationReport) GetNameExtension() string`

GetNameExtension returns the NameExtension field if non-nil, zero value otherwise.

### GetNameExtensionOk

`func (o *AidsToNavigationReport) GetNameExtensionOk() (*string, bool)`

GetNameExtensionOk returns a tuple with the NameExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameExtension

`func (o *AidsToNavigationReport) SetNameExtension(v string)`

SetNameExtension sets NameExtension field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


