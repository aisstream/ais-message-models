# PositionReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**navigational_status** | **int** |  | 
**rate_of_turn** | **int** |  | 
**sog** | **float** |  | 
**position_accuracy** | **bool** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**cog** | **float** |  | 
**true_heading** | **int** |  | 
**timestamp** | **int** |  | 
**special_manoeuvre_indicator** | **int** |  | 
**spare** | **int** |  | 
**raim** | **bool** |  | 
**communication_state** | **int** |  | 

## Example

```python
from aisStream.models.position_report import PositionReport

# TODO update the JSON string below
json = "{}"
# create an instance of PositionReport from a JSON string
position_report_instance = PositionReport.from_json(json)
# print the JSON string representation of the object
print(PositionReport.to_json())

# convert the object into a dict
position_report_dict = position_report_instance.to_dict()
# create an instance of PositionReport from a dict
position_report_from_dict = PositionReport.from_dict(position_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


