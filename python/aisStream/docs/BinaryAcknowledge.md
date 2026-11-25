# BinaryAcknowledge


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**destinations** | [**BinaryAcknowledgeDestinations**](BinaryAcknowledgeDestinations.md) |  | 

## Example

```python
from aisStream.models.binary_acknowledge import BinaryAcknowledge

# TODO update the JSON string below
json = "{}"
# create an instance of BinaryAcknowledge from a JSON string
binary_acknowledge_instance = BinaryAcknowledge.from_json(json)
# print the JSON string representation of the object
print(BinaryAcknowledge.to_json())

# convert the object into a dict
binary_acknowledge_dict = binary_acknowledge_instance.to_dict()
# create an instance of BinaryAcknowledge from a dict
binary_acknowledge_from_dict = BinaryAcknowledge.from_dict(binary_acknowledge_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


