# GroupAssignmentCommand


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**longitude1** | **float** |  | 
**latitude1** | **float** |  | 
**longitude2** | **float** |  | 
**latitude2** | **float** |  | 
**station_type** | **int** |  | 
**ship_type** | **int** |  | 
**spare2** | **int** |  | 
**tx_rx_mode** | **int** |  | 
**reporting_interval** | **int** |  | 
**quiet_time** | **int** |  | 
**spare3** | **int** |  | 

## Example

```python
from aisStream.models.group_assignment_command import GroupAssignmentCommand

# TODO update the JSON string below
json = "{}"
# create an instance of GroupAssignmentCommand from a JSON string
group_assignment_command_instance = GroupAssignmentCommand.from_json(json)
# print the JSON string representation of the object
print(GroupAssignmentCommand.to_json())

# convert the object into a dict
group_assignment_command_dict = group_assignment_command_instance.to_dict()
# create an instance of GroupAssignmentCommand from a dict
group_assignment_command_from_dict = GroupAssignmentCommand.from_dict(group_assignment_command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


