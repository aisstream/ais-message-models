# ExtendedClassBPositionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare1** | **int32** |  | 
**Sog** | **float64** |  | 
**PositionAccuracy** | **bool** |  | 
**Longitude** | **float64** |  | 
**Latitude** | **float64** |  | 
**Cog** | **float64** |  | 
**TrueHeading** | **int32** |  | 
**Timestamp** | **int32** |  | 
**Spare2** | **int32** |  | 
**Name** | **string** |  | 
**Type** | **int32** |  | 
**Dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**FixType** | **int32** |  | 
**Raim** | **bool** |  | 
**Dte** | **bool** |  | 
**AssignedMode** | **bool** |  | 
**Spare3** | **int32** |  | 

## Methods

### NewExtendedClassBPositionReport

`func NewExtendedClassBPositionReport(messageID int32, repeatIndicator int32, userID int32, valid bool, spare1 int32, sog float64, positionAccuracy bool, longitude float64, latitude float64, cog float64, trueHeading int32, timestamp int32, spare2 int32, name string, type_ int32, dimension ShipStaticDataDimension, fixType int32, raim bool, dte bool, assignedMode bool, spare3 int32, ) *ExtendedClassBPositionReport`

NewExtendedClassBPositionReport instantiates a new ExtendedClassBPositionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedClassBPositionReportWithDefaults

`func NewExtendedClassBPositionReportWithDefaults() *ExtendedClassBPositionReport`

NewExtendedClassBPositionReportWithDefaults instantiates a new ExtendedClassBPositionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *ExtendedClassBPositionReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *ExtendedClassBPositionReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *ExtendedClassBPositionReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *ExtendedClassBPositionReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *ExtendedClassBPositionReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *ExtendedClassBPositionReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *ExtendedClassBPositionReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ExtendedClassBPositionReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ExtendedClassBPositionReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *ExtendedClassBPositionReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ExtendedClassBPositionReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ExtendedClassBPositionReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare1

`func (o *ExtendedClassBPositionReport) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *ExtendedClassBPositionReport) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *ExtendedClassBPositionReport) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetSog

`func (o *ExtendedClassBPositionReport) GetSog() float64`

GetSog returns the Sog field if non-nil, zero value otherwise.

### GetSogOk

`func (o *ExtendedClassBPositionReport) GetSogOk() (*float64, bool)`

GetSogOk returns a tuple with the Sog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSog

`func (o *ExtendedClassBPositionReport) SetSog(v float64)`

SetSog sets Sog field to given value.


### GetPositionAccuracy

`func (o *ExtendedClassBPositionReport) GetPositionAccuracy() bool`

GetPositionAccuracy returns the PositionAccuracy field if non-nil, zero value otherwise.

### GetPositionAccuracyOk

`func (o *ExtendedClassBPositionReport) GetPositionAccuracyOk() (*bool, bool)`

GetPositionAccuracyOk returns a tuple with the PositionAccuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionAccuracy

`func (o *ExtendedClassBPositionReport) SetPositionAccuracy(v bool)`

SetPositionAccuracy sets PositionAccuracy field to given value.


### GetLongitude

`func (o *ExtendedClassBPositionReport) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *ExtendedClassBPositionReport) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *ExtendedClassBPositionReport) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetLatitude

`func (o *ExtendedClassBPositionReport) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *ExtendedClassBPositionReport) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *ExtendedClassBPositionReport) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetCog

`func (o *ExtendedClassBPositionReport) GetCog() float64`

GetCog returns the Cog field if non-nil, zero value otherwise.

### GetCogOk

`func (o *ExtendedClassBPositionReport) GetCogOk() (*float64, bool)`

GetCogOk returns a tuple with the Cog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCog

`func (o *ExtendedClassBPositionReport) SetCog(v float64)`

SetCog sets Cog field to given value.


### GetTrueHeading

`func (o *ExtendedClassBPositionReport) GetTrueHeading() int32`

GetTrueHeading returns the TrueHeading field if non-nil, zero value otherwise.

### GetTrueHeadingOk

`func (o *ExtendedClassBPositionReport) GetTrueHeadingOk() (*int32, bool)`

GetTrueHeadingOk returns a tuple with the TrueHeading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueHeading

`func (o *ExtendedClassBPositionReport) SetTrueHeading(v int32)`

SetTrueHeading sets TrueHeading field to given value.


### GetTimestamp

`func (o *ExtendedClassBPositionReport) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtendedClassBPositionReport) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtendedClassBPositionReport) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetSpare2

`func (o *ExtendedClassBPositionReport) GetSpare2() int32`

GetSpare2 returns the Spare2 field if non-nil, zero value otherwise.

### GetSpare2Ok

`func (o *ExtendedClassBPositionReport) GetSpare2Ok() (*int32, bool)`

GetSpare2Ok returns a tuple with the Spare2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare2

`func (o *ExtendedClassBPositionReport) SetSpare2(v int32)`

SetSpare2 sets Spare2 field to given value.


### GetName

`func (o *ExtendedClassBPositionReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtendedClassBPositionReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtendedClassBPositionReport) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ExtendedClassBPositionReport) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedClassBPositionReport) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedClassBPositionReport) SetType(v int32)`

SetType sets Type field to given value.


### GetDimension

`func (o *ExtendedClassBPositionReport) GetDimension() ShipStaticDataDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ExtendedClassBPositionReport) GetDimensionOk() (*ShipStaticDataDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ExtendedClassBPositionReport) SetDimension(v ShipStaticDataDimension)`

SetDimension sets Dimension field to given value.


### GetFixType

`func (o *ExtendedClassBPositionReport) GetFixType() int32`

GetFixType returns the FixType field if non-nil, zero value otherwise.

### GetFixTypeOk

`func (o *ExtendedClassBPositionReport) GetFixTypeOk() (*int32, bool)`

GetFixTypeOk returns a tuple with the FixType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixType

`func (o *ExtendedClassBPositionReport) SetFixType(v int32)`

SetFixType sets FixType field to given value.


### GetRaim

`func (o *ExtendedClassBPositionReport) GetRaim() bool`

GetRaim returns the Raim field if non-nil, zero value otherwise.

### GetRaimOk

`func (o *ExtendedClassBPositionReport) GetRaimOk() (*bool, bool)`

GetRaimOk returns a tuple with the Raim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaim

`func (o *ExtendedClassBPositionReport) SetRaim(v bool)`

SetRaim sets Raim field to given value.


### GetDte

`func (o *ExtendedClassBPositionReport) GetDte() bool`

GetDte returns the Dte field if non-nil, zero value otherwise.

### GetDteOk

`func (o *ExtendedClassBPositionReport) GetDteOk() (*bool, bool)`

GetDteOk returns a tuple with the Dte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDte

`func (o *ExtendedClassBPositionReport) SetDte(v bool)`

SetDte sets Dte field to given value.


### GetAssignedMode

`func (o *ExtendedClassBPositionReport) GetAssignedMode() bool`

GetAssignedMode returns the AssignedMode field if non-nil, zero value otherwise.

### GetAssignedModeOk

`func (o *ExtendedClassBPositionReport) GetAssignedModeOk() (*bool, bool)`

GetAssignedModeOk returns a tuple with the AssignedMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedMode

`func (o *ExtendedClassBPositionReport) SetAssignedMode(v bool)`

SetAssignedMode sets AssignedMode field to given value.


### GetSpare3

`func (o *ExtendedClassBPositionReport) GetSpare3() int32`

GetSpare3 returns the Spare3 field if non-nil, zero value otherwise.

### GetSpare3Ok

`func (o *ExtendedClassBPositionReport) GetSpare3Ok() (*int32, bool)`

GetSpare3Ok returns a tuple with the Spare3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare3

`func (o *ExtendedClassBPositionReport) SetSpare3(v int32)`

SetSpare3 sets Spare3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


