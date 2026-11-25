# DataLinkManagementMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**data** | [**DataLinkManagementMessageData**](DataLinkManagementMessageData.md) |  | 

## Example

```python
from aisStream.models.data_link_management_message import DataLinkManagementMessage

# TODO update the JSON string below
json = "{}"
# create an instance of DataLinkManagementMessage from a JSON string
data_link_management_message_instance = DataLinkManagementMessage.from_json(json)
# print the JSON string representation of the object
print(DataLinkManagementMessage.to_json())

# convert the object into a dict
data_link_management_message_dict = data_link_management_message_instance.to_dict()
# create an instance of DataLinkManagementMessage from a dict
data_link_management_message_from_dict = DataLinkManagementMessage.from_dict(data_link_management_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


