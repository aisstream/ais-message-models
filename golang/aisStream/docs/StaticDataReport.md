# StaticDataReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**Reserved** | **int32** |  | 
**PartNumber** | **bool** |  | 
**ReportA** | [**StaticDataReportReportA**](StaticDataReportReportA.md) |  | 
**ReportB** | [**StaticDataReportReportB**](StaticDataReportReportB.md) |  | 

## Methods

### NewStaticDataReport

`func NewStaticDataReport(messageID int32, repeatIndicator int32, userID int32, valid bool, reserved int32, partNumber bool, reportA StaticDataReportReportA, reportB StaticDataReportReportB, ) *StaticDataReport`

NewStaticDataReport instantiates a new StaticDataReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticDataReportWithDefaults

`func NewStaticDataReportWithDefaults() *StaticDataReport`

NewStaticDataReportWithDefaults instantiates a new StaticDataReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *StaticDataReport) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *StaticDataReport) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *StaticDataReport) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *StaticDataReport) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *StaticDataReport) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *StaticDataReport) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *StaticDataReport) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *StaticDataReport) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *StaticDataReport) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *StaticDataReport) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *StaticDataReport) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *StaticDataReport) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetReserved

`func (o *StaticDataReport) GetReserved() int32`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *StaticDataReport) GetReservedOk() (*int32, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *StaticDataReport) SetReserved(v int32)`

SetReserved sets Reserved field to given value.


### GetPartNumber

`func (o *StaticDataReport) GetPartNumber() bool`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *StaticDataReport) GetPartNumberOk() (*bool, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *StaticDataReport) SetPartNumber(v bool)`

SetPartNumber sets PartNumber field to given value.


### GetReportA

`func (o *StaticDataReport) GetReportA() StaticDataReportReportA`

GetReportA returns the ReportA field if non-nil, zero value otherwise.

### GetReportAOk

`func (o *StaticDataReport) GetReportAOk() (*StaticDataReportReportA, bool)`

GetReportAOk returns a tuple with the ReportA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportA

`func (o *StaticDataReport) SetReportA(v StaticDataReportReportA)`

SetReportA sets ReportA field to given value.


### GetReportB

`func (o *StaticDataReport) GetReportB() StaticDataReportReportB`

GetReportB returns the ReportB field if non-nil, zero value otherwise.

### GetReportBOk

`func (o *StaticDataReport) GetReportBOk() (*StaticDataReportReportB, bool)`

GetReportBOk returns a tuple with the ReportB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportB

`func (o *StaticDataReport) SetReportB(v StaticDataReportReportB)`

SetReportB sets ReportB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


