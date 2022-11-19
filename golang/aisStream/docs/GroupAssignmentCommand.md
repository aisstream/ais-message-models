# GroupAssignmentCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Spare1** | **int32** |  | 
**Longitude1** | **float64** |  | 
**Latitude1** | **float64** |  | 
**Longitude2** | **float64** |  | 
**Latitude2** | **float64** |  | 
**StationType** | **int32** |  | 
**ShipType** | **int32** |  | 
**Spare2** | **int32** |  | 
**TxRxMode** | **int32** |  | 
**ReportingInterval** | **int32** |  | 
**QuietTime** | **int32** |  | 
**Spare3** | **int32** |  | 

## Methods

### NewGroupAssignmentCommand

`func NewGroupAssignmentCommand(messageID int32, repeatIndicator int32, userID int32, valid bool, spare1 int32, longitude1 float64, latitude1 float64, longitude2 float64, latitude2 float64, stationType int32, shipType int32, spare2 int32, txRxMode int32, reportingInterval int32, quietTime int32, spare3 int32, ) *GroupAssignmentCommand`

NewGroupAssignmentCommand instantiates a new GroupAssignmentCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupAssignmentCommandWithDefaults

`func NewGroupAssignmentCommandWithDefaults() *GroupAssignmentCommand`

NewGroupAssignmentCommandWithDefaults instantiates a new GroupAssignmentCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *GroupAssignmentCommand) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *GroupAssignmentCommand) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *GroupAssignmentCommand) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *GroupAssignmentCommand) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *GroupAssignmentCommand) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *GroupAssignmentCommand) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *GroupAssignmentCommand) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *GroupAssignmentCommand) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *GroupAssignmentCommand) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *GroupAssignmentCommand) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *GroupAssignmentCommand) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *GroupAssignmentCommand) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetSpare1

`func (o *GroupAssignmentCommand) GetSpare1() int32`

GetSpare1 returns the Spare1 field if non-nil, zero value otherwise.

### GetSpare1Ok

`func (o *GroupAssignmentCommand) GetSpare1Ok() (*int32, bool)`

GetSpare1Ok returns a tuple with the Spare1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare1

`func (o *GroupAssignmentCommand) SetSpare1(v int32)`

SetSpare1 sets Spare1 field to given value.


### GetLongitude1

`func (o *GroupAssignmentCommand) GetLongitude1() float64`

GetLongitude1 returns the Longitude1 field if non-nil, zero value otherwise.

### GetLongitude1Ok

`func (o *GroupAssignmentCommand) GetLongitude1Ok() (*float64, bool)`

GetLongitude1Ok returns a tuple with the Longitude1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude1

`func (o *GroupAssignmentCommand) SetLongitude1(v float64)`

SetLongitude1 sets Longitude1 field to given value.


### GetLatitude1

`func (o *GroupAssignmentCommand) GetLatitude1() float64`

GetLatitude1 returns the Latitude1 field if non-nil, zero value otherwise.

### GetLatitude1Ok

`func (o *GroupAssignmentCommand) GetLatitude1Ok() (*float64, bool)`

GetLatitude1Ok returns a tuple with the Latitude1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude1

`func (o *GroupAssignmentCommand) SetLatitude1(v float64)`

SetLatitude1 sets Latitude1 field to given value.


### GetLongitude2

`func (o *GroupAssignmentCommand) GetLongitude2() float64`

GetLongitude2 returns the Longitude2 field if non-nil, zero value otherwise.

### GetLongitude2Ok

`func (o *GroupAssignmentCommand) GetLongitude2Ok() (*float64, bool)`

GetLongitude2Ok returns a tuple with the Longitude2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude2

`func (o *GroupAssignmentCommand) SetLongitude2(v float64)`

SetLongitude2 sets Longitude2 field to given value.


### GetLatitude2

`func (o *GroupAssignmentCommand) GetLatitude2() float64`

GetLatitude2 returns the Latitude2 field if non-nil, zero value otherwise.

### GetLatitude2Ok

`func (o *GroupAssignmentCommand) GetLatitude2Ok() (*float64, bool)`

GetLatitude2Ok returns a tuple with the Latitude2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude2

`func (o *GroupAssignmentCommand) SetLatitude2(v float64)`

SetLatitude2 sets Latitude2 field to given value.


### GetStationType

`func (o *GroupAssignmentCommand) GetStationType() int32`

GetStationType returns the StationType field if non-nil, zero value otherwise.

### GetStationTypeOk

`func (o *GroupAssignmentCommand) GetStationTypeOk() (*int32, bool)`

GetStationTypeOk returns a tuple with the StationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationType

`func (o *GroupAssignmentCommand) SetStationType(v int32)`

SetStationType sets StationType field to given value.


### GetShipType

`func (o *GroupAssignmentCommand) GetShipType() int32`

GetShipType returns the ShipType field if non-nil, zero value otherwise.

### GetShipTypeOk

`func (o *GroupAssignmentCommand) GetShipTypeOk() (*int32, bool)`

GetShipTypeOk returns a tuple with the ShipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipType

`func (o *GroupAssignmentCommand) SetShipType(v int32)`

SetShipType sets ShipType field to given value.


### GetSpare2

`func (o *GroupAssignmentCommand) GetSpare2() int32`

GetSpare2 returns the Spare2 field if non-nil, zero value otherwise.

### GetSpare2Ok

`func (o *GroupAssignmentCommand) GetSpare2Ok() (*int32, bool)`

GetSpare2Ok returns a tuple with the Spare2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare2

`func (o *GroupAssignmentCommand) SetSpare2(v int32)`

SetSpare2 sets Spare2 field to given value.


### GetTxRxMode

`func (o *GroupAssignmentCommand) GetTxRxMode() int32`

GetTxRxMode returns the TxRxMode field if non-nil, zero value otherwise.

### GetTxRxModeOk

`func (o *GroupAssignmentCommand) GetTxRxModeOk() (*int32, bool)`

GetTxRxModeOk returns a tuple with the TxRxMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxRxMode

`func (o *GroupAssignmentCommand) SetTxRxMode(v int32)`

SetTxRxMode sets TxRxMode field to given value.


### GetReportingInterval

`func (o *GroupAssignmentCommand) GetReportingInterval() int32`

GetReportingInterval returns the ReportingInterval field if non-nil, zero value otherwise.

### GetReportingIntervalOk

`func (o *GroupAssignmentCommand) GetReportingIntervalOk() (*int32, bool)`

GetReportingIntervalOk returns a tuple with the ReportingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportingInterval

`func (o *GroupAssignmentCommand) SetReportingInterval(v int32)`

SetReportingInterval sets ReportingInterval field to given value.


### GetQuietTime

`func (o *GroupAssignmentCommand) GetQuietTime() int32`

GetQuietTime returns the QuietTime field if non-nil, zero value otherwise.

### GetQuietTimeOk

`func (o *GroupAssignmentCommand) GetQuietTimeOk() (*int32, bool)`

GetQuietTimeOk returns a tuple with the QuietTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietTime

`func (o *GroupAssignmentCommand) SetQuietTime(v int32)`

SetQuietTime sets QuietTime field to given value.


### GetSpare3

`func (o *GroupAssignmentCommand) GetSpare3() int32`

GetSpare3 returns the Spare3 field if non-nil, zero value otherwise.

### GetSpare3Ok

`func (o *GroupAssignmentCommand) GetSpare3Ok() (*int32, bool)`

GetSpare3Ok returns a tuple with the Spare3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare3

`func (o *GroupAssignmentCommand) SetSpare3(v int32)`

SetSpare3 sets Spare3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


