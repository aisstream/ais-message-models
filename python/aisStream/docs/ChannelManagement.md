# ChannelManagement


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**channel_a** | **int** |  | 
**channel_b** | **int** |  | 
**tx_rx_mode** | **int** |  | 
**low_power** | **bool** |  | 
**area** | [**ChannelManagementArea**](ChannelManagementArea.md) |  | 
**unicast** | [**ChannelManagementUnicast**](ChannelManagementUnicast.md) |  | 
**is_addressed** | **bool** |  | 
**bw_a** | **bool** |  | 
**bw_b** | **bool** |  | 
**transitional_zone_size** | **int** |  | 
**spare4** | **int** |  | 

## Example

```python
from aisStream.models.channel_management import ChannelManagement

# TODO update the JSON string below
json = "{}"
# create an instance of ChannelManagement from a JSON string
channel_management_instance = ChannelManagement.from_json(json)
# print the JSON string representation of the object
print(ChannelManagement.to_json())

# convert the object into a dict
channel_management_dict = channel_management_instance.to_dict()
# create an instance of ChannelManagement from a dict
channel_management_from_dict = ChannelManagement.from_dict(channel_management_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


