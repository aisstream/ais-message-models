# AssignedModeCommand


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**commands** | [**AssignedModeCommandCommands**](AssignedModeCommandCommands.md) |  | 

## Example

```python
from aisStream.models.assigned_mode_command import AssignedModeCommand

# TODO update the JSON string below
json = "{}"
# create an instance of AssignedModeCommand from a JSON string
assigned_mode_command_instance = AssignedModeCommand.from_json(json)
# print the JSON string representation of the object
print(AssignedModeCommand.to_json())

# convert the object into a dict
assigned_mode_command_dict = assigned_mode_command_instance.to_dict()
# create an instance of AssignedModeCommand from a dict
assigned_mode_command_from_dict = AssignedModeCommand.from_dict(assigned_mode_command_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


