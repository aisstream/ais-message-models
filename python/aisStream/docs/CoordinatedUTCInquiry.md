# CoordinatedUTCInquiry


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**destination_id** | **int** |  | 
**spare2** | **int** |  | 

## Example

```python
from aisStream.models.coordinated_utc_inquiry import CoordinatedUTCInquiry

# TODO update the JSON string below
json = "{}"
# create an instance of CoordinatedUTCInquiry from a JSON string
coordinated_utc_inquiry_instance = CoordinatedUTCInquiry.from_json(json)
# print the JSON string representation of the object
print(CoordinatedUTCInquiry.to_json())

# convert the object into a dict
coordinated_utc_inquiry_dict = coordinated_utc_inquiry_instance.to_dict()
# create an instance of CoordinatedUTCInquiry from a dict
coordinated_utc_inquiry_from_dict = CoordinatedUTCInquiry.from_dict(coordinated_utc_inquiry_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


