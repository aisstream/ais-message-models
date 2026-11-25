# GnssBroadcastBinaryMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**spare2** | **int** |  | 
**data** | **str** |  | 

## Example

```python
from aisStream.models.gnss_broadcast_binary_message import GnssBroadcastBinaryMessage

# TODO update the JSON string below
json = "{}"
# create an instance of GnssBroadcastBinaryMessage from a JSON string
gnss_broadcast_binary_message_instance = GnssBroadcastBinaryMessage.from_json(json)
# print the JSON string representation of the object
print(GnssBroadcastBinaryMessage.to_json())

# convert the object into a dict
gnss_broadcast_binary_message_dict = gnss_broadcast_binary_message_instance.to_dict()
# create an instance of GnssBroadcastBinaryMessage from a dict
gnss_broadcast_binary_message_from_dict = GnssBroadcastBinaryMessage.from_dict(gnss_broadcast_binary_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


