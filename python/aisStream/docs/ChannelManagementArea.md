# ChannelManagementArea


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**longitude1** | **float** |  | 
**latitude1** | **float** |  | 
**longitude2** | **float** |  | 
**latitude2** | **float** |  | 

## Example

```python
from aisStream.models.channel_management_area import ChannelManagementArea

# TODO update the JSON string below
json = "{}"
# create an instance of ChannelManagementArea from a JSON string
channel_management_area_instance = ChannelManagementArea.from_json(json)
# print the JSON string representation of the object
print(ChannelManagementArea.to_json())

# convert the object into a dict
channel_management_area_dict = channel_management_area_instance.to_dict()
# create an instance of ChannelManagementArea from a dict
channel_management_area_from_dict = ChannelManagementArea.from_dict(channel_management_area_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


