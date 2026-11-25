# AddressedBinaryMessageApplicationID


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**valid** | **bool** |  | 
**designated_area_code** | **int** |  | 
**function_identifier** | **int** |  | 

## Example

```python
from aisStream.models.addressed_binary_message_application_id import AddressedBinaryMessageApplicationID

# TODO update the JSON string below
json = "{}"
# create an instance of AddressedBinaryMessageApplicationID from a JSON string
addressed_binary_message_application_id_instance = AddressedBinaryMessageApplicationID.from_json(json)
# print the JSON string representation of the object
print(AddressedBinaryMessageApplicationID.to_json())

# convert the object into a dict
addressed_binary_message_application_id_dict = addressed_binary_message_application_id_instance.to_dict()
# create an instance of AddressedBinaryMessageApplicationID from a dict
addressed_binary_message_application_id_from_dict = AddressedBinaryMessageApplicationID.from_dict(addressed_binary_message_application_id_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


