# AisStreamMessageMessage


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**position_report** | [**PositionReport**](PositionReport.md) |  | [optional] 
**unknown_message** | [**UnknownMessage**](UnknownMessage.md) |  | [optional] 
**addressed_safety_message** | [**AddressedSafetyMessage**](AddressedSafetyMessage.md) |  | [optional] 
**addressed_binary_message** | [**AddressedBinaryMessage**](AddressedBinaryMessage.md) |  | [optional] 
**aids_to_navigation_report** | [**AidsToNavigationReport**](AidsToNavigationReport.md) |  | [optional] 
**assigned_mode_command** | [**AssignedModeCommand**](AssignedModeCommand.md) |  | [optional] 
**base_station_report** | [**BaseStationReport**](BaseStationReport.md) |  | [optional] 
**binary_acknowledge** | [**BinaryAcknowledge**](BinaryAcknowledge.md) |  | [optional] 
**binary_broadcast_message** | [**BinaryBroadcastMessage**](BinaryBroadcastMessage.md) |  | [optional] 
**channel_management** | [**ChannelManagement**](ChannelManagement.md) |  | [optional] 
**coordinated_utc_inquiry** | [**CoordinatedUTCInquiry**](CoordinatedUTCInquiry.md) |  | [optional] 
**data_link_management_message** | [**DataLinkManagementMessage**](DataLinkManagementMessage.md) |  | [optional] 
**data_link_management_message_data** | [**DataLinkManagementMessageData**](DataLinkManagementMessageData.md) |  | [optional] 
**extended_class_b_position_report** | [**ExtendedClassBPositionReport**](ExtendedClassBPositionReport.md) |  | [optional] 
**gnss_broadcast_binary_message** | [**GnssBroadcastBinaryMessage**](GnssBroadcastBinaryMessage.md) |  | [optional] 
**group_assignment_command** | [**GroupAssignmentCommand**](GroupAssignmentCommand.md) |  | [optional] 
**interrogation** | [**Interrogation**](Interrogation.md) |  | [optional] 
**long_range_ais_broadcast_message** | [**LongRangeAisBroadcastMessage**](LongRangeAisBroadcastMessage.md) |  | [optional] 
**multi_slot_binary_message** | [**MultiSlotBinaryMessage**](MultiSlotBinaryMessage.md) |  | [optional] 
**safety_broadcast_message** | [**SafetyBroadcastMessage**](SafetyBroadcastMessage.md) |  | [optional] 
**ship_static_data** | [**ShipStaticData**](ShipStaticData.md) |  | [optional] 
**single_slot_binary_message** | [**SingleSlotBinaryMessage**](SingleSlotBinaryMessage.md) |  | [optional] 
**standard_class_b_position_report** | [**StandardClassBPositionReport**](StandardClassBPositionReport.md) |  | [optional] 
**standard_search_and_rescue_aircraft_report** | [**StandardSearchAndRescueAircraftReport**](StandardSearchAndRescueAircraftReport.md) |  | [optional] 
**static_data_report** | [**StaticDataReport**](StaticDataReport.md) |  | [optional] 

## Example

```python
from aisStream.models.ais_stream_message_message import AisStreamMessageMessage

# TODO update the JSON string below
json = "{}"
# create an instance of AisStreamMessageMessage from a JSON string
ais_stream_message_message_instance = AisStreamMessageMessage.from_json(json)
# print the JSON string representation of the object
print(AisStreamMessageMessage.to_json())

# convert the object into a dict
ais_stream_message_message_dict = ais_stream_message_message_instance.to_dict()
# create an instance of AisStreamMessageMessage from a dict
ais_stream_message_message_from_dict = AisStreamMessageMessage.from_dict(ais_stream_message_message_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


