# SafetyBroadcastMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**text** | **str** |  | 

## Example

```python
from aisStream.models.safety_broadcast_message import SafetyBroadcastMessage

# TODO update the JSON string below
json = "{}"
# create an instance of SafetyBroadcastMessage from a JSON string
safety_broadcast_message_instance = SafetyBroadcastMessage.from_json(json)
# print the JSON string representation of the object
print(SafetyBroadcastMessage.to_json())

# convert the object into a dict
safety_broadcast_message_dict = safety_broadcast_message_instance.to_dict()
# create an instance of SafetyBroadcastMessage from a dict
safety_broadcast_message_from_dict = SafetyBroadcastMessage.from_dict(safety_broadcast_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


