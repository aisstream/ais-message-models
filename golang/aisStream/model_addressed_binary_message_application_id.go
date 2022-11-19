/*
Ais-Stream WebsocketObjects

A sample API to illustrate OpenAPI concepts

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// AddressedBinaryMessageApplicationID struct for AddressedBinaryMessageApplicationID
type AddressedBinaryMessageApplicationID struct {
	Valid bool `json:"Valid"`
	DesignatedAreaCode int32 `json:"DesignatedAreaCode"`
	FunctionIdentifier int32 `json:"FunctionIdentifier"`
}

// NewAddressedBinaryMessageApplicationID instantiates a new AddressedBinaryMessageApplicationID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressedBinaryMessageApplicationID(valid bool, designatedAreaCode int32, functionIdentifier int32) *AddressedBinaryMessageApplicationID {
	this := AddressedBinaryMessageApplicationID{}
	this.Valid = valid
	this.DesignatedAreaCode = designatedAreaCode
	this.FunctionIdentifier = functionIdentifier
	return &this
}

// NewAddressedBinaryMessageApplicationIDWithDefaults instantiates a new AddressedBinaryMessageApplicationID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressedBinaryMessageApplicationIDWithDefaults() *AddressedBinaryMessageApplicationID {
	this := AddressedBinaryMessageApplicationID{}
	return &this
}

// GetValid returns the Valid field value
func (o *AddressedBinaryMessageApplicationID) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *AddressedBinaryMessageApplicationID) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *AddressedBinaryMessageApplicationID) SetValid(v bool) {
	o.Valid = v
}

// GetDesignatedAreaCode returns the DesignatedAreaCode field value
func (o *AddressedBinaryMessageApplicationID) GetDesignatedAreaCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DesignatedAreaCode
}

// GetDesignatedAreaCodeOk returns a tuple with the DesignatedAreaCode field value
// and a boolean to check if the value has been set.
func (o *AddressedBinaryMessageApplicationID) GetDesignatedAreaCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesignatedAreaCode, true
}

// SetDesignatedAreaCode sets field value
func (o *AddressedBinaryMessageApplicationID) SetDesignatedAreaCode(v int32) {
	o.DesignatedAreaCode = v
}

// GetFunctionIdentifier returns the FunctionIdentifier field value
func (o *AddressedBinaryMessageApplicationID) GetFunctionIdentifier() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FunctionIdentifier
}

// GetFunctionIdentifierOk returns a tuple with the FunctionIdentifier field value
// and a boolean to check if the value has been set.
func (o *AddressedBinaryMessageApplicationID) GetFunctionIdentifierOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FunctionIdentifier, true
}

// SetFunctionIdentifier sets field value
func (o *AddressedBinaryMessageApplicationID) SetFunctionIdentifier(v int32) {
	o.FunctionIdentifier = v
}

func (o AddressedBinaryMessageApplicationID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["DesignatedAreaCode"] = o.DesignatedAreaCode
	}
	if true {
		toSerialize["FunctionIdentifier"] = o.FunctionIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableAddressedBinaryMessageApplicationID struct {
	value *AddressedBinaryMessageApplicationID
	isSet bool
}

func (v NullableAddressedBinaryMessageApplicationID) Get() *AddressedBinaryMessageApplicationID {
	return v.value
}

func (v *NullableAddressedBinaryMessageApplicationID) Set(val *AddressedBinaryMessageApplicationID) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressedBinaryMessageApplicationID) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressedBinaryMessageApplicationID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressedBinaryMessageApplicationID(val *AddressedBinaryMessageApplicationID) *NullableAddressedBinaryMessageApplicationID {
	return &NullableAddressedBinaryMessageApplicationID{value: val, isSet: true}
}

func (v NullableAddressedBinaryMessageApplicationID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressedBinaryMessageApplicationID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

