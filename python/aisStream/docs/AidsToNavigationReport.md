# AidsToNavigationReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**type** | **int** |  | 
**name** | **str** |  | 
**position_accuracy** | **bool** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**fixtype** | **int** |  | 
**timestamp** | **int** |  | 
**off_position** | **bool** |  | 
**ato_n** | **int** |  | 
**raim** | **bool** |  | 
**virtual_ato_n** | **bool** |  | 
**assigned_mode** | **bool** |  | 
**spare** | **bool** |  | 
**name_extension** | **str** |  | 

## Example

```python
from aisStream.models.aids_to_navigation_report import AidsToNavigationReport

# TODO update the JSON string below
json = "{}"
# create an instance of AidsToNavigationReport from a JSON string
aids_to_navigation_report_instance = AidsToNavigationReport.from_json(json)
# print the JSON string representation of the object
print(AidsToNavigationReport.to_json())

# convert the object into a dict
aids_to_navigation_report_dict = aids_to_navigation_report_instance.to_dict()
# create an instance of AidsToNavigationReport from a dict
aids_to_navigation_report_from_dict = AidsToNavigationReport.from_dict(aids_to_navigation_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


