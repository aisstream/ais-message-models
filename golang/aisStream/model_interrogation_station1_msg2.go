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

// InterrogationStation1Msg2 struct for InterrogationStation1Msg2
type InterrogationStation1Msg2 struct {
	Valid bool `json:"Valid"`
	Spare int32 `json:"Spare"`
	MessageID int32 `json:"MessageID"`
	SlotOffset int32 `json:"SlotOffset"`
}

// NewInterrogationStation1Msg2 instantiates a new InterrogationStation1Msg2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterrogationStation1Msg2(valid bool, spare int32, messageID int32, slotOffset int32) *InterrogationStation1Msg2 {
	this := InterrogationStation1Msg2{}
	this.Valid = valid
	this.Spare = spare
	this.MessageID = messageID
	this.SlotOffset = slotOffset
	return &this
}

// NewInterrogationStation1Msg2WithDefaults instantiates a new InterrogationStation1Msg2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterrogationStation1Msg2WithDefaults() *InterrogationStation1Msg2 {
	this := InterrogationStation1Msg2{}
	return &this
}

// GetValid returns the Valid field value
func (o *InterrogationStation1Msg2) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg2) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *InterrogationStation1Msg2) SetValid(v bool) {
	o.Valid = v
}

// GetSpare returns the Spare field value
func (o *InterrogationStation1Msg2) GetSpare() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare
}

// GetSpareOk returns a tuple with the Spare field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg2) GetSpareOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare, true
}

// SetSpare sets field value
func (o *InterrogationStation1Msg2) SetSpare(v int32) {
	o.Spare = v
}

// GetMessageID returns the MessageID field value
func (o *InterrogationStation1Msg2) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg2) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *InterrogationStation1Msg2) SetMessageID(v int32) {
	o.MessageID = v
}

// GetSlotOffset returns the SlotOffset field value
func (o *InterrogationStation1Msg2) GetSlotOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotOffset
}

// GetSlotOffsetOk returns a tuple with the SlotOffset field value
// and a boolean to check if the value has been set.
func (o *InterrogationStation1Msg2) GetSlotOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotOffset, true
}

// SetSlotOffset sets field value
func (o *InterrogationStation1Msg2) SetSlotOffset(v int32) {
	o.SlotOffset = v
}

func (o InterrogationStation1Msg2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["Spare"] = o.Spare
	}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["SlotOffset"] = o.SlotOffset
	}
	return json.Marshal(toSerialize)
}

type NullableInterrogationStation1Msg2 struct {
	value *InterrogationStation1Msg2
	isSet bool
}

func (v NullableInterrogationStation1Msg2) Get() *InterrogationStation1Msg2 {
	return v.value
}

func (v *NullableInterrogationStation1Msg2) Set(val *InterrogationStation1Msg2) {
	v.value = val
	v.isSet = true
}

func (v NullableInterrogationStation1Msg2) IsSet() bool {
	return v.isSet
}

func (v *NullableInterrogationStation1Msg2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterrogationStation1Msg2(val *InterrogationStation1Msg2) *NullableInterrogationStation1Msg2 {
	return &NullableInterrogationStation1Msg2{value: val, isSet: true}
}

func (v NullableInterrogationStation1Msg2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterrogationStation1Msg2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


