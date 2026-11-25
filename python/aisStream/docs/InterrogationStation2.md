# InterrogationStation2


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**valid** | **bool** |  | 
**spare1** | **int** |  | 
**station_id** | **int** |  | 
**message_id** | **int** |  | 
**slot_offset** | **int** |  | 
**spare2** | **int** |  | 

## Example

```python
from aisStream.models.interrogation_station2 import InterrogationStation2

# TODO update the JSON string below
json = "{}"
# create an instance of InterrogationStation2 from a JSON string
interrogation_station2_instance = InterrogationStation2.from_json(json)
# print the JSON string representation of the object
print(InterrogationStation2.to_json())

# convert the object into a dict
interrogation_station2_dict = interrogation_station2_instance.to_dict()
# create an instance of InterrogationStation2 from a dict
interrogation_station2_from_dict = InterrogationStation2.from_dict(interrogation_station2_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


