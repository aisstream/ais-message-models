# StaticDataReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**reserved** | **int** |  | 
**part_number** | **bool** |  | 
**report_a** | [**StaticDataReportReportA**](StaticDataReportReportA.md) |  | 
**report_b** | [**StaticDataReportReportB**](StaticDataReportReportB.md) |  | 

## Example

```python
from aisStream.models.static_data_report import StaticDataReport

# TODO update the JSON string below
json = "{}"
# create an instance of StaticDataReport from a JSON string
static_data_report_instance = StaticDataReport.from_json(json)
# print the JSON string representation of the object
print(StaticDataReport.to_json())

# convert the object into a dict
static_data_report_dict = static_data_report_instance.to_dict()
# create an instance of StaticDataReport from a dict
static_data_report_from_dict = StaticDataReport.from_dict(static_data_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


