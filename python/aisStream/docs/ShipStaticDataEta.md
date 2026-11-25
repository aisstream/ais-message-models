# ShipStaticDataEta


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**month** | **int** |  | 
**day** | **int** |  | 
**hour** | **int** |  | 
**minute** | **int** |  | 

## Example

```python
from aisStream.models.ship_static_data_eta import ShipStaticDataEta

# TODO update the JSON string below
json = "{}"
# create an instance of ShipStaticDataEta from a JSON string
ship_static_data_eta_instance = ShipStaticDataEta.from_json(json)
# print the JSON string representation of the object
print(ShipStaticDataEta.to_json())

# convert the object into a dict
ship_static_data_eta_dict = ship_static_data_eta_instance.to_dict()
# create an instance of ShipStaticDataEta from a dict
ship_static_data_eta_from_dict = ShipStaticDataEta.from_dict(ship_static_data_eta_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


