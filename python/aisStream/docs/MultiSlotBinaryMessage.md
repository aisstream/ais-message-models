# MultiSlotBinaryMessage


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
**spare1** | **int** |  | 
**application_id** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**payload** | **str** |  | 
**spare2** | **int** |  | 
**communication_state_is_itdma** | **bool** |  | 
**communication_state** | **int** |  | 

## Example

```python
from aisStream.models.multi_slot_binary_message import MultiSlotBinaryMessage

# TODO update the JSON string below
json = "{}"
# create an instance of MultiSlotBinaryMessage from a JSON string
multi_slot_binary_message_instance = MultiSlotBinaryMessage.from_json(json)
# print the JSON string representation of the object
print(MultiSlotBinaryMessage.to_json())

# convert the object into a dict
multi_slot_binary_message_dict = multi_slot_binary_message_instance.to_dict()
# create an instance of MultiSlotBinaryMessage from a dict
multi_slot_binary_message_from_dict = MultiSlotBinaryMessage.from_dict(multi_slot_binary_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


