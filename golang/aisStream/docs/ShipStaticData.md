# ShipStaticData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageID** | **int32** |  | 
**RepeatIndicator** | **int32** |  | 
**UserID** | **int32** |  | 
**Valid** | **bool** |  | 
**AisVersion** | **int32** |  | 
**ImoNumber** | **int32** |  | 
**CallSign** | **string** |  | 
**Name** | **string** |  | 
**Type** | **int32** |  | 
**Dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**FixType** | **int32** |  | 
**Eta** | [**ShipStaticDataEta**](ShipStaticDataEta.md) |  | 
**MaximumStaticDraught** | **float64** |  | 
**Destination** | **string** |  | 
**Dte** | **bool** |  | 
**Spare** | **bool** |  | 

## Methods

### NewShipStaticData

`func NewShipStaticData(messageID int32, repeatIndicator int32, userID int32, valid bool, aisVersion int32, imoNumber int32, callSign string, name string, type_ int32, dimension ShipStaticDataDimension, fixType int32, eta ShipStaticDataEta, maximumStaticDraught float64, destination string, dte bool, spare bool, ) *ShipStaticData`

NewShipStaticData instantiates a new ShipStaticData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipStaticDataWithDefaults

`func NewShipStaticDataWithDefaults() *ShipStaticData`

NewShipStaticDataWithDefaults instantiates a new ShipStaticData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageID

`func (o *ShipStaticData) GetMessageID() int32`

GetMessageID returns the MessageID field if non-nil, zero value otherwise.

### GetMessageIDOk

`func (o *ShipStaticData) GetMessageIDOk() (*int32, bool)`

GetMessageIDOk returns a tuple with the MessageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageID

`func (o *ShipStaticData) SetMessageID(v int32)`

SetMessageID sets MessageID field to given value.


### GetRepeatIndicator

`func (o *ShipStaticData) GetRepeatIndicator() int32`

GetRepeatIndicator returns the RepeatIndicator field if non-nil, zero value otherwise.

### GetRepeatIndicatorOk

`func (o *ShipStaticData) GetRepeatIndicatorOk() (*int32, bool)`

GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatIndicator

`func (o *ShipStaticData) SetRepeatIndicator(v int32)`

SetRepeatIndicator sets RepeatIndicator field to given value.


### GetUserID

`func (o *ShipStaticData) GetUserID() int32`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *ShipStaticData) GetUserIDOk() (*int32, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *ShipStaticData) SetUserID(v int32)`

SetUserID sets UserID field to given value.


### GetValid

`func (o *ShipStaticData) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ShipStaticData) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ShipStaticData) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetAisVersion

`func (o *ShipStaticData) GetAisVersion() int32`

GetAisVersion returns the AisVersion field if non-nil, zero value otherwise.

### GetAisVersionOk

`func (o *ShipStaticData) GetAisVersionOk() (*int32, bool)`

GetAisVersionOk returns a tuple with the AisVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAisVersion

`func (o *ShipStaticData) SetAisVersion(v int32)`

SetAisVersion sets AisVersion field to given value.


### GetImoNumber

`func (o *ShipStaticData) GetImoNumber() int32`

GetImoNumber returns the ImoNumber field if non-nil, zero value otherwise.

### GetImoNumberOk

`func (o *ShipStaticData) GetImoNumberOk() (*int32, bool)`

GetImoNumberOk returns a tuple with the ImoNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImoNumber

`func (o *ShipStaticData) SetImoNumber(v int32)`

SetImoNumber sets ImoNumber field to given value.


### GetCallSign

`func (o *ShipStaticData) GetCallSign() string`

GetCallSign returns the CallSign field if non-nil, zero value otherwise.

### GetCallSignOk

`func (o *ShipStaticData) GetCallSignOk() (*string, bool)`

GetCallSignOk returns a tuple with the CallSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSign

`func (o *ShipStaticData) SetCallSign(v string)`

SetCallSign sets CallSign field to given value.


### GetName

`func (o *ShipStaticData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShipStaticData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShipStaticData) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ShipStaticData) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShipStaticData) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShipStaticData) SetType(v int32)`

SetType sets Type field to given value.


### GetDimension

`func (o *ShipStaticData) GetDimension() ShipStaticDataDimension`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *ShipStaticData) GetDimensionOk() (*ShipStaticDataDimension, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *ShipStaticData) SetDimension(v ShipStaticDataDimension)`

SetDimension sets Dimension field to given value.


### GetFixType

`func (o *ShipStaticData) GetFixType() int32`

GetFixType returns the FixType field if non-nil, zero value otherwise.

### GetFixTypeOk

`func (o *ShipStaticData) GetFixTypeOk() (*int32, bool)`

GetFixTypeOk returns a tuple with the FixType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixType

`func (o *ShipStaticData) SetFixType(v int32)`

SetFixType sets FixType field to given value.


### GetEta

`func (o *ShipStaticData) GetEta() ShipStaticDataEta`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *ShipStaticData) GetEtaOk() (*ShipStaticDataEta, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *ShipStaticData) SetEta(v ShipStaticDataEta)`

SetEta sets Eta field to given value.


### GetMaximumStaticDraught

`func (o *ShipStaticData) GetMaximumStaticDraught() float64`

GetMaximumStaticDraught returns the MaximumStaticDraught field if non-nil, zero value otherwise.

### GetMaximumStaticDraughtOk

`func (o *ShipStaticData) GetMaximumStaticDraughtOk() (*float64, bool)`

GetMaximumStaticDraughtOk returns a tuple with the MaximumStaticDraught field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumStaticDraught

`func (o *ShipStaticData) SetMaximumStaticDraught(v float64)`

SetMaximumStaticDraught sets MaximumStaticDraught field to given value.


### GetDestination

`func (o *ShipStaticData) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ShipStaticData) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ShipStaticData) SetDestination(v string)`

SetDestination sets Destination field to given value.


### GetDte

`func (o *ShipStaticData) GetDte() bool`

GetDte returns the Dte field if non-nil, zero value otherwise.

### GetDteOk

`func (o *ShipStaticData) GetDteOk() (*bool, bool)`

GetDteOk returns a tuple with the Dte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDte

`func (o *ShipStaticData) SetDte(v bool)`

SetDte sets Dte field to given value.


### GetSpare

`func (o *ShipStaticData) GetSpare() bool`

GetSpare returns the Spare field if non-nil, zero value otherwise.

### GetSpareOk

`func (o *ShipStaticData) GetSpareOk() (*bool, bool)`

GetSpareOk returns a tuple with the Spare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpare

`func (o *ShipStaticData) SetSpare(v bool)`

SetSpare sets Spare field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


