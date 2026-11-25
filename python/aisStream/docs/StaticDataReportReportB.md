# StaticDataReportReportB


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**valid** | **bool** |  | 
**ship_type** | **int** |  | 
**vendor_id_name** | **str** |  | 
**vender_id_model** | **int** |  | 
**vender_id_serial** | **int** |  | 
**call_sign** | **str** |  | 
**dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**fix_type** | **int** |  | 
**spare** | **int** |  | 

## Example

```python
from aisStream.models.static_data_report_report_b import StaticDataReportReportB

# TODO update the JSON string below
json = "{}"
# create an instance of StaticDataReportReportB from a JSON string
static_data_report_report_b_instance = StaticDataReportReportB.from_json(json)
# print the JSON string representation of the object
print(StaticDataReportReportB.to_json())

# convert the object into a dict
static_data_report_report_b_dict = static_data_report_report_b_instance.to_dict()
# create an instance of StaticDataReportReportB from a dict
static_data_report_report_b_from_dict = StaticDataReportReportB.from_dict(static_data_report_report_b_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


