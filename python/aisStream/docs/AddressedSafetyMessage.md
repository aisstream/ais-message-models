# AddressedSafetyMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**sequenceinteger** | **int** |  | 
**destination_id** | **int** |  | 
**retransmission** | **bool** |  | 
**spare** | **bool** |  | 
**text** | **str** |  | 

## Example

```python
from aisStream.models.addressed_safety_message import AddressedSafetyMessage

# TODO update the JSON string below
json = "{}"
# create an instance of AddressedSafetyMessage from a JSON string
addressed_safety_message_instance = AddressedSafetyMessage.from_json(json)
# print the JSON string representation of the object
print(AddressedSafetyMessage.to_json())

# convert the object into a dict
addressed_safety_message_dict = addressed_safety_message_instance.to_dict()
# create an instance of AddressedSafetyMessage from a dict
addressed_safety_message_from_dict = AddressedSafetyMessage.from_dict(addressed_safety_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


