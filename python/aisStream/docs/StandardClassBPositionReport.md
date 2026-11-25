# StandardClassBPositionReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**sog** | **float** |  | 
**position_accuracy** | **bool** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**cog** | **float** |  | 
**true_heading** | **int** |  | 
**timestamp** | **int** |  | 
**spare2** | **int** |  | 
**class_b_unit** | **bool** |  | 
**class_b_display** | **bool** |  | 
**class_b_dsc** | **bool** |  | 
**class_b_band** | **bool** |  | 
**class_b_msg22** | **bool** |  | 
**assigned_mode** | **bool** |  | 
**raim** | **bool** |  | 
**communication_state_is_itdma** | **bool** |  | 
**communication_state** | **int** |  | 

## Example

```python
from aisStream.models.standard_class_b_position_report import StandardClassBPositionReport

# TODO update the JSON string below
json = "{}"
# create an instance of StandardClassBPositionReport from a JSON string
standard_class_b_position_report_instance = StandardClassBPositionReport.from_json(json)
# print the JSON string representation of the object
print(StandardClassBPositionReport.to_json())

# convert the object into a dict
standard_class_b_position_report_dict = standard_class_b_position_report_instance.to_dict()
# create an instance of StandardClassBPositionReport from a dict
standard_class_b_position_report_from_dict = StandardClassBPositionReport.from_dict(standard_class_b_position_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


