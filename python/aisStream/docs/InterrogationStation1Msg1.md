# InterrogationStation1Msg1


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**valid** | **bool** |  | 
**station_id** | **int** |  | 
**message_id** | **int** |  | 
**slot_offset** | **int** |  | 

## Example

```python
from aisStream.models.interrogation_station1_msg1 import InterrogationStation1Msg1

# TODO update the JSON string below
json = "{}"
# create an instance of InterrogationStation1Msg1 from a JSON string
interrogation_station1_msg1_instance = InterrogationStation1Msg1.from_json(json)
# print the JSON string representation of the object
print(InterrogationStation1Msg1.to_json())

# convert the object into a dict
interrogation_station1_msg1_dict = interrogation_station1_msg1_instance.to_dict()
# create an instance of InterrogationStation1Msg1 from a dict
interrogation_station1_msg1_from_dict = InterrogationStation1Msg1.from_dict(interrogation_station1_msg1_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


