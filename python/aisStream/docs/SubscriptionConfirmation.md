# SubscriptionConfirmation


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_type** | **str** |  | 
**message** | [**SubscriptionConfirmationMessage**](SubscriptionConfirmationMessage.md) |  | 

## Example

```python
from aisStream.models.subscription_confirmation import SubscriptionConfirmation

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionConfirmation from a JSON string
subscription_confirmation_instance = SubscriptionConfirmation.from_json(json)
# print the JSON string representation of the object
print(SubscriptionConfirmation.to_json())

# convert the object into a dict
subscription_confirmation_dict = subscription_confirmation_instance.to_dict()
# create an instance of SubscriptionConfirmation from a dict
subscription_confirmation_from_dict = SubscriptionConfirmation.from_dict(subscription_confirmation_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


