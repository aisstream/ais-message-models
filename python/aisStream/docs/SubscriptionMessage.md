# SubscriptionMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**api_key** | **str** |  | 
**bounding_boxes** | **List[List[List[float]]]** |  | 
**filters_ship_mmsi** | **List[str]** |  | [optional] 
**filter_message_types** | [**List[AisMessageTypes]**](AisMessageTypes.md) |  | [optional] 

## Example

```python
from aisStream.models.subscription_message import SubscriptionMessage

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionMessage from a JSON string
subscription_message_instance = SubscriptionMessage.from_json(json)
# print the JSON string representation of the object
print(SubscriptionMessage.to_json())

# convert the object into a dict
subscription_message_dict = subscription_message_instance.to_dict()
# create an instance of SubscriptionMessage from a dict
subscription_message_from_dict = SubscriptionMessage.from_dict(subscription_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


