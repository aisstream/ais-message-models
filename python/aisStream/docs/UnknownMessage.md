# UnknownMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**error** | **str** |  | [optional] 

## Example

```python
from aisStream.models.unknown_message import UnknownMessage

# TODO update the JSON string below
json = "{}"
# create an instance of UnknownMessage from a JSON string
unknown_message_instance = UnknownMessage.from_json(json)
# print the JSON string representation of the object
print(UnknownMessage.to_json())

# convert the object into a dict
unknown_message_dict = unknown_message_instance.to_dict()
# create an instance of UnknownMessage from a dict
unknown_message_from_dict = UnknownMessage.from_dict(unknown_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


