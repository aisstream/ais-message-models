# LongRangeAisBroadcastMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**position_accuracy** | **bool** |  | 
**raim** | **bool** |  | 
**navigational_status** | **int** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**sog** | **float** |  | 
**cog** | **float** |  | 
**position_latency** | **bool** |  | 
**spare** | **bool** |  | 

## Example

```python
from aisStream.models.long_range_ais_broadcast_message import LongRangeAisBroadcastMessage

# TODO update the JSON string below
json = "{}"
# create an instance of LongRangeAisBroadcastMessage from a JSON string
long_range_ais_broadcast_message_instance = LongRangeAisBroadcastMessage.from_json(json)
# print the JSON string representation of the object
print(LongRangeAisBroadcastMessage.to_json())

# convert the object into a dict
long_range_ais_broadcast_message_dict = long_range_ais_broadcast_message_instance.to_dict()
# create an instance of LongRangeAisBroadcastMessage from a dict
long_range_ais_broadcast_message_from_dict = LongRangeAisBroadcastMessage.from_dict(long_range_ais_broadcast_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


