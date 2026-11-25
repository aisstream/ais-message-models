# ShipStaticData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message_id** | **int** |  | 
**repeat_indicator** | **int** |  | 
**user_id** | **int** |  | 
**valid** | **bool** |  | 
**ais_version** | **int** |  | 
**imo_number** | **int** |  | 
**call_sign** | **str** |  | 
**name** | **str** |  | 
**type** | **int** |  | 
**dimension** | [**ShipStaticDataDimension**](ShipStaticDataDimension.md) |  | 
**fix_type** | **int** |  | 
**eta** | [**ShipStaticDataEta**](ShipStaticDataEta.md) |  | 
**maximum_static_draught** | **float** |  | 
**destination** | **str** |  | 
**dte** | **bool** |  | 
**spare** | **bool** |  | 

## Example

```python
from aisStream.models.ship_static_data import ShipStaticData

# TODO update the JSON string below
json = "{}"
# create an instance of ShipStaticData from a JSON string
ship_static_data_instance = ShipStaticData.from_json(json)
# print the JSON string representation of the object
print(ShipStaticData.to_json())

# convert the object into a dict
ship_static_data_dict = ship_static_data_instance.to_dict()
# create an instance of ShipStaticData from a dict
ship_static_data_from_dict = ShipStaticData.from_dict(ship_static_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


