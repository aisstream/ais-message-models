# ChannelManagementUnicast


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address_station1** | **int** |  | 
**spare2** | **int** |  | 
**address_station2** | **int** |  | 
**spare3** | **int** |  | 

## Example

```python
from aisStream.models.channel_management_unicast import ChannelManagementUnicast

# TODO update the JSON string below
json = "{}"
# create an instance of ChannelManagementUnicast from a JSON string
channel_management_unicast_instance = ChannelManagementUnicast.from_json(json)
# print the JSON string representation of the object
print(ChannelManagementUnicast.to_json())

# convert the object into a dict
channel_management_unicast_dict = channel_management_unicast_instance.to_dict()
# create an instance of ChannelManagementUnicast from a dict
channel_management_unicast_from_dict = ChannelManagementUnicast.from_dict(channel_management_unicast_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


