# SubscriptionConfirmationMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compression_enabled** | **bool** |  | 

## Example

```python
from aisStream.models.subscription_confirmation_message import SubscriptionConfirmationMessage

# TODO update the JSON string below
json = "{}"
# create an instance of SubscriptionConfirmationMessage from a JSON string
subscription_confirmation_message_instance = SubscriptionConfirmationMessage.from_json(json)
# print the JSON string representation of the object
print(SubscriptionConfirmationMessage.to_json())

# convert the object into a dict
subscription_confirmation_message_dict = subscription_confirmation_message_instance.to_dict()
# create an instance of SubscriptionConfirmationMessage from a dict
subscription_confirmation_message_from_dict = SubscriptionConfirmationMessage.from_dict(subscription_confirmation_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


