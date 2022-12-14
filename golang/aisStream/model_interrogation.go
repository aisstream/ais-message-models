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

// Interrogation struct for Interrogation
type Interrogation struct {
	MessageID int32 `json:"MessageID"`
	RepeatIndicator int32 `json:"RepeatIndicator"`
	UserID int32 `json:"UserID"`
	Valid bool `json:"Valid"`
	Spare int32 `json:"Spare"`
	Station1Msg1 InterrogationStation1Msg1 `json:"Station1Msg1"`
	Station1Msg2 InterrogationStation1Msg2 `json:"Station1Msg2"`
	Station2 InterrogationStation2 `json:"Station2"`
}

// NewInterrogation instantiates a new Interrogation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterrogation(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, station1Msg1 InterrogationStation1Msg1, station1Msg2 InterrogationStation1Msg2, station2 InterrogationStation2) *Interrogation {
	this := Interrogation{}
	this.MessageID = messageID
	this.RepeatIndicator = repeatIndicator
	this.UserID = userID
	this.Valid = valid
	this.Spare = spare
	this.Station1Msg1 = station1Msg1
	this.Station1Msg2 = station1Msg2
	this.Station2 = station2
	return &this
}

// NewInterrogationWithDefaults instantiates a new Interrogation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterrogationWithDefaults() *Interrogation {
	this := Interrogation{}
	return &this
}

// GetMessageID returns the MessageID field value
func (o *Interrogation) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *Interrogation) SetMessageID(v int32) {
	o.MessageID = v
}

// GetRepeatIndicator returns the RepeatIndicator field value
func (o *Interrogation) GetRepeatIndicator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIndicator
}

// GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetRepeatIndicatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatIndicator, true
}

// SetRepeatIndicator sets field value
func (o *Interrogation) SetRepeatIndicator(v int32) {
	o.RepeatIndicator = v
}

// GetUserID returns the UserID field value
func (o *Interrogation) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *Interrogation) SetUserID(v int32) {
	o.UserID = v
}

// GetValid returns the Valid field value
func (o *Interrogation) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *Interrogation) SetValid(v bool) {
	o.Valid = v
}

// GetSpare returns the Spare field value
func (o *Interrogation) GetSpare() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare
}

// GetSpareOk returns a tuple with the Spare field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetSpareOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare, true
}

// SetSpare sets field value
func (o *Interrogation) SetSpare(v int32) {
	o.Spare = v
}

// GetStation1Msg1 returns the Station1Msg1 field value
func (o *Interrogation) GetStation1Msg1() InterrogationStation1Msg1 {
	if o == nil {
		var ret InterrogationStation1Msg1
		return ret
	}

	return o.Station1Msg1
}

// GetStation1Msg1Ok returns a tuple with the Station1Msg1 field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetStation1Msg1Ok() (*InterrogationStation1Msg1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Station1Msg1, true
}

// SetStation1Msg1 sets field value
func (o *Interrogation) SetStation1Msg1(v InterrogationStation1Msg1) {
	o.Station1Msg1 = v
}

// GetStation1Msg2 returns the Station1Msg2 field value
func (o *Interrogation) GetStation1Msg2() InterrogationStation1Msg2 {
	if o == nil {
		var ret InterrogationStation1Msg2
		return ret
	}

	return o.Station1Msg2
}

// GetStation1Msg2Ok returns a tuple with the Station1Msg2 field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetStation1Msg2Ok() (*InterrogationStation1Msg2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Station1Msg2, true
}

// SetStation1Msg2 sets field value
func (o *Interrogation) SetStation1Msg2(v InterrogationStation1Msg2) {
	o.Station1Msg2 = v
}

// GetStation2 returns the Station2 field value
func (o *Interrogation) GetStation2() InterrogationStation2 {
	if o == nil {
		var ret InterrogationStation2
		return ret
	}

	return o.Station2
}

// GetStation2Ok returns a tuple with the Station2 field value
// and a boolean to check if the value has been set.
func (o *Interrogation) GetStation2Ok() (*InterrogationStation2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Station2, true
}

// SetStation2 sets field value
func (o *Interrogation) SetStation2(v InterrogationStation2) {
	o.Station2 = v
}

func (o Interrogation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["RepeatIndicator"] = o.RepeatIndicator
	}
	if true {
		toSerialize["UserID"] = o.UserID
	}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["Spare"] = o.Spare
	}
	if true {
		toSerialize["Station1Msg1"] = o.Station1Msg1
	}
	if true {
		toSerialize["Station1Msg2"] = o.Station1Msg2
	}
	if true {
		toSerialize["Station2"] = o.Station2
	}
	return json.Marshal(toSerialize)
}

type NullableInterrogation struct {
	value *Interrogation
	isSet bool
}

func (v NullableInterrogation) Get() *Interrogation {
	return v.value
}

func (v *NullableInterrogation) Set(val *Interrogation) {
	v.value = val
	v.isSet = true
}

func (v NullableInterrogation) IsSet() bool {
	return v.isSet
}

func (v *NullableInterrogation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterrogation(val *Interrogation) *NullableInterrogation {
	return &NullableInterrogation{value: val, isSet: true}
}

func (v NullableInterrogation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterrogation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


