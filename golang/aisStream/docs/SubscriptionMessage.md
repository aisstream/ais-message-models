# SubscriptionMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**APIKey** | **string** |  | 
**BoundingBoxes** | [**[][][]float64**]([][]float64.md) |  | 

## Methods

### NewSubscriptionMessage

`func NewSubscriptionMessage(aPIKey string, boundingBoxes [][][]float64, ) *SubscriptionMessage`

NewSubscriptionMessage instantiates a new SubscriptionMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionMessageWithDefaults

`func NewSubscriptionMessageWithDefaults() *SubscriptionMessage`

NewSubscriptionMessageWithDefaults instantiates a new SubscriptionMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAPIKey

`func (o *SubscriptionMessage) GetAPIKey() string`

GetAPIKey returns the APIKey field if non-nil, zero value otherwise.

### GetAPIKeyOk

`func (o *SubscriptionMessage) GetAPIKeyOk() (*string, bool)`

GetAPIKeyOk returns a tuple with the APIKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAPIKey

`func (o *SubscriptionMessage) SetAPIKey(v string)`

SetAPIKey sets APIKey field to given value.


### GetBoundingBoxes

`func (o *SubscriptionMessage) GetBoundingBoxes() [][][]float64`

GetBoundingBoxes returns the BoundingBoxes field if non-nil, zero value otherwise.

### GetBoundingBoxesOk

`func (o *SubscriptionMessage) GetBoundingBoxesOk() (*[][][]float64, bool)`

GetBoundingBoxesOk returns a tuple with the BoundingBoxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundingBoxes

`func (o *SubscriptionMessage) SetBoundingBoxes(v [][][]float64)`

SetBoundingBoxes sets BoundingBoxes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


