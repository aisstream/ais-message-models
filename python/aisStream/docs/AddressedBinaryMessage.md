# AddressedBinaryMessage


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
**application_id** | [**AddressedBinaryMessageApplicationID**](AddressedBinaryMessageApplicationID.md) |  | 
**binary_data** | **str** |  | 

## Example

```python
from aisStream.models.addressed_binary_message import AddressedBinaryMessage

# TODO update the JSON string below
json = "{}"
# create an instance of AddressedBinaryMessage from a JSON string
addressed_binary_message_instance = AddressedBinaryMessage.from_json(json)
# print the JSON string representation of the object
print(AddressedBinaryMessage.to_json())

# convert the object into a dict
addressed_binary_message_dict = addressed_binary_message_instance.to_dict()
# create an instance of AddressedBinaryMessage from a dict
addressed_binary_message_from_dict = AddressedBinaryMessage.from_dict(addressed_binary_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


