# ExtendedClassBPositionReport


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
**name** | **str** |  | 
**type** | **int** |  | 
**dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**fix_type** | **int** |  | 
**raim** | **bool** |  | 
**dte** | **bool** |  | 
**assigned_mode** | **bool** |  | 
**spare3** | **int** |  | 

## Example

```python
from aisStream.models.extended_class_b_position_report import ExtendedClassBPositionReport

# TODO update the JSON string below
json = "{}"
# create an instance of ExtendedClassBPositionReport from a JSON string
extended_class_b_position_report_instance = ExtendedClassBPositionReport.from_json(json)
# print the JSON string representation of the object
print(ExtendedClassBPositionReport.to_json())

# convert the object into a dict
extended_class_b_position_report_dict = extended_class_b_position_report_instance.to_dict()
# create an instance of ExtendedClassBPositionReport from a dict
extended_class_b_position_report_from_dict = ExtendedClassBPositionReport.from_dict(extended_class_b_position_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


