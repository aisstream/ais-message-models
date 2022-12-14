/*
Ais-Stream WebsocketObjects

OpenAPI 3.0 definitions for the data models used by aisstream.io.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// InterrogationStation1Msg1 struct for InterrogationStation1Msg1
type InterrogationStation1Msg1 struct {
	Valid bool `json:"Valid"`
	StationID int32 `json:"StationID"`
	MessageID int32 `json:"MessageID"`
	SlotOffset int32 `json:"SlotOffset"`
}

// NewInterrogationStation1Msg1 instantiates a new InterrogationStation1Msg1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterrogationStation1Msg1(valid bool, stationID int32, messageID int32, slotOffset int32) *InterrogationStation1Msg1 {
	this := InterrogationStation1Msg1{}
	this.Valid = valid
	this.StationID = stationID
	this.MessageID = messageID
	this.SlotOffset = slotOffset
	return &this
}

// NewInterrogationStation1Msg1WithDefaults instantiates a new InterrogationStation1Msg1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterrogationStation1Msg1WithDefaults() *InterrogationStation1Msg1 {
	this := InterrogationStation1Msg1{}
	return &this
}

// GetValid returns the Valid field value
func (o *InterrogationStation1Msg1) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg1) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *InterrogationStation1Msg1) SetValid(v bool) {
	o.Valid = v
}

// GetStationID returns the StationID field value
func (o *InterrogationStation1Msg1) GetStationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StationID
}

// GetStationIDOk returns a tuple with the StationID field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg1) GetStationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StationID, true
}

// SetStationID sets field value
func (o *InterrogationStation1Msg1) SetStationID(v int32) {
	o.StationID = v
}

// GetMessageID returns the MessageID field value
func (o *InterrogationStation1Msg1) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg1) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *InterrogationStation1Msg1) SetMessageID(v int32) {
	o.MessageID = v
}

// GetSlotOffset returns the SlotOffset field value
func (o *InterrogationStation1Msg1) GetSlotOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotOffset
}

// GetSlotOffsetOk returns a tuple with the SlotOffset field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg1) GetSlotOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotOffset, true
}

// SetSlotOffset sets field value
func (o *InterrogationStation1Msg1) SetSlotOffset(v int32) {
	o.SlotOffset = v
}

func (o InterrogationStation1Msg1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["StationID"] = o.StationID
	}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["SlotOffset"] = o.SlotOffset
	}
	return json.Marshal(toSerialize)
}

type NullableInterrogationStation1Msg1 struct {
	value *InterrogationStation1Msg1
	isSet bool
}

func (v NullableInterrogationStation1Msg1) Get() *InterrogationStation1Msg1 {
	return v.value
}

func (v *NullableInterrogationStation1Msg1) Set(val *InterrogationStation1Msg1) {
	v.value = val
	v.isSet = true
}

func (v NullableInterrogationStation1Msg1) IsSet() bool {
	return v.isSet
}

func (v *NullableInterrogationStation1Msg1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterrogationStation1Msg1(val *InterrogationStation1Msg1) *NullableInterrogationStation1Msg1 {
	return &NullableInterrogationStation1Msg1{value: val, isSet: true}
}

func (v NullableInterrogationStation1Msg1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterrogationStation1Msg1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


