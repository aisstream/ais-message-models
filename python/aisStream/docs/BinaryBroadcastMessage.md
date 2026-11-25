# BinaryBroadcastMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**application_id** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**binary_data** | **str** |  | 

## Example

```python
from aisStream.models.binary_broadcast_message import BinaryBroadcastMessage

# TODO update the JSON string below
json = "{}"
# create an instance of BinaryBroadcastMessage from a JSON string
binary_broadcast_message_instance = BinaryBroadcastMessage.from_json(json)
# print the JSON string representation of the object
print(BinaryBroadcastMessage.to_json())

# convert the object into a dict
binary_broadcast_message_dict = binary_broadcast_message_instance.to_dict()
# create an instance of BinaryBroadcastMessage from a dict
binary_broadcast_message_from_dict = BinaryBroadcastMessage.from_dict(binary_broadcast_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


