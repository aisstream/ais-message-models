# SingleSlotBinaryMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**destination_id_valid** | **bool** |  | 
**application_id_valid** | **bool** |  | 
**destination_id** | **int** |  | 
**spare** | **int** |  | 
**application_id** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**payload** | **str** |  | 

## Example

```python
from aisStream.models.single_slot_binary_message import SingleSlotBinaryMessage

# TODO update the JSON string below
json = "{}"
# create an instance of SingleSlotBinaryMessage from a JSON string
single_slot_binary_message_instance = SingleSlotBinaryMessage.from_json(json)
# print the JSON string representation of the object
print(SingleSlotBinaryMessage.to_json())

# convert the object into a dict
single_slot_binary_message_dict = single_slot_binary_message_instance.to_dict()
# create an instance of SingleSlotBinaryMessage from a dict
single_slot_binary_message_from_dict = SingleSlotBinaryMessage.from_dict(single_slot_binary_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


