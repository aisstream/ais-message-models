# Interrogation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**spare** | **int** |  | 
**station1_msg1** | [**InterrogationStation1Msg1**](InterrogationStation1Msg1.md) |  | 
**station1_msg2** | [**InterrogationStation1Msg2**](InterrogationStation1Msg2.md) |  | 
**station2** | [**InterrogationStation2**](InterrogationStation2.md) |  | 

## Example

```python
from aisStream.models.interrogation import Interrogation

# TODO update the JSON string below
json = "{}"
# create an instance of Interrogation from a JSON string
interrogation_instance = Interrogation.from_json(json)
# print the JSON string representation of the object
print(Interrogation.to_json())

# convert the object into a dict
interrogation_dict = interrogation_instance.to_dict()
# create an instance of Interrogation from a dict
interrogation_from_dict = Interrogation.from_dict(interrogation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


