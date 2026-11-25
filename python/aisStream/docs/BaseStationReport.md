# BaseStationReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**utc_year** | **int** |  | 
**utc_month** | **int** |  | 
**utc_day** | **int** |  | 
**utc_hour** | **int** |  | 
**utc_minute** | **int** |  | 
**utc_second** | **int** |  | 
**position_accuracy** | **bool** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**fix_type** | **int** |  | 
**long_range_enable** | **bool** |  | 
**spare** | **int** |  | 
**raim** | **bool** |  | 
**communication_state** | **int** |  | 

## Example

```python
from aisStream.models.base_station_report import BaseStationReport

# TODO update the JSON string below
json = "{}"
# create an instance of BaseStationReport from a JSON string
base_station_report_instance = BaseStationReport.from_json(json)
# print the JSON string representation of the object
print(BaseStationReport.to_json())

# convert the object into a dict
base_station_report_dict = base_station_report_instance.to_dict()
# create an instance of BaseStationReport from a dict
base_station_report_from_dict = BaseStationReport.from_dict(base_station_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


