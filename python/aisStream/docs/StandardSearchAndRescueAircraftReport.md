# StandardSearchAndRescueAircraftReport


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**altitude** | **int** |  | 
**sog** | **float** |  | 
**position_accuracy** | **bool** |  | 
**longitude** | **float** |  | 
**latitude** | **float** |  | 
**cog** | **float** |  | 
**timestamp** | **int** |  | 
**alt_from_baro** | **bool** |  | 
**spare1** | **int** |  | 
**dte** | **bool** |  | 
**spare2** | **int** |  | 
**assigned_mode** | **bool** |  | 
**raim** | **bool** |  | 
**communication_state_is_itdma** | **bool** |  | 
**communication_state** | **int** |  | 

## Example

```python
from aisStream.models.standard_search_and_rescue_aircraft_report import StandardSearchAndRescueAircraftReport

# TODO update the JSON string below
json = "{}"
# create an instance of StandardSearchAndRescueAircraftReport from a JSON string
standard_search_and_rescue_aircraft_report_instance = StandardSearchAndRescueAircraftReport.from_json(json)
# print the JSON string representation of the object
print(StandardSearchAndRescueAircraftReport.to_json())

# convert the object into a dict
standard_search_and_rescue_aircraft_report_dict = standard_search_and_rescue_aircraft_report_instance.to_dict()
# create an instance of StandardSearchAndRescueAircraftReport from a dict
standard_search_and_rescue_aircraft_report_from_dict = StandardSearchAndRescueAircraftReport.from_dict(standard_search_and_rescue_aircraft_report_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


