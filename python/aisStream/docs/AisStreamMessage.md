# AisStreamMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta_data** | **object** |  | 
**message_type** | [**AisMessageTypes**](AisMessageTypes.md) |  | 
**message** | [**AisStreamMessageMessage**](AisStreamMessageMessage.md) |  | 

## Example

```python
from aisStream.models.ais_stream_message import AisStreamMessage

# TODO update the JSON string below
json = "{}"
# create an instance of AisStreamMessage from a JSON string
ais_stream_message_instance = AisStreamMessage.from_json(json)
# print the JSON string representation of the object
print(AisStreamMessage.to_json())

# convert the object into a dict
ais_stream_message_dict = ais_stream_message_instance.to_dict()
# create an instance of AisStreamMessage from a dict
ais_stream_message_from_dict = AisStreamMessage.from_dict(ais_stream_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


