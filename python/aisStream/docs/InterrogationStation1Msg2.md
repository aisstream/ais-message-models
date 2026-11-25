# InterrogationStation1Msg2


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**valid** | **bool** |  | 
**spare** | **int** |  | 
**message_id** | **int** |  | 
**slot_offset** | **int** |  | 

## Example

```python
from aisStream.models.interrogation_station1_msg2 import InterrogationStation1Msg2

# TODO update the JSON string below
json = "{}"
# create an instance of InterrogationStation1Msg2 from a JSON string
interrogation_station1_msg2_instance = InterrogationStation1Msg2.from_json(json)
# print the JSON string representation of the object
print(InterrogationStation1Msg2.to_json())

# convert the object into a dict
interrogation_station1_msg2_dict = interrogation_station1_msg2_instance.to_dict()
# create an instance of InterrogationStation1Msg2 from a dict
interrogation_station1_msg2_from_dict = InterrogationStation1Msg2.from_dict(interrogation_station1_msg2_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


