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

// AssignedModeCommand struct for AssignedModeCommand
type AssignedModeCommand struct {
	MessageID int32 `json:"MessageID"`
	RepeatIndicator int32 `json:"RepeatIndicator"`
	UserID int32 `json:"UserID"`
	Valid bool `json:"Valid"`
	Spare int32 `json:"Spare"`
	Commands AssignedModeCommandCommands `json:"Commands"`
}

// NewAssignedModeCommand instantiates a new AssignedModeCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignedModeCommand(messageID int32, repeatIndicator int32, userID int32, valid bool, spare int32, commands AssignedModeCommandCommands) *AssignedModeCommand {
	this := AssignedModeCommand{}
	this.MessageID = messageID
	this.RepeatIndicator = repeatIndicator
	this.UserID = userID
	this.Valid = valid
	this.Spare = spare
	this.Commands = commands
	return &this
}

// NewAssignedModeCommandWithDefaults instantiates a new AssignedModeCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignedModeCommandWithDefaults() *AssignedModeCommand {
	this := AssignedModeCommand{}
	return &this
}

// GetMessageID returns the MessageID field value
func (o *AssignedModeCommand) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *AssignedModeCommand) SetMessageID(v int32) {
	o.MessageID = v
}

// GetRepeatIndicator returns the RepeatIndicator field value
func (o *AssignedModeCommand) GetRepeatIndicator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIndicator
}

// GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetRepeatIndicatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatIndicator, true
}

// SetRepeatIndicator sets field value
func (o *AssignedModeCommand) SetRepeatIndicator(v int32) {
	o.RepeatIndicator = v
}

// GetUserID returns the UserID field value
func (o *AssignedModeCommand) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *AssignedModeCommand) SetUserID(v int32) {
	o.UserID = v
}

// GetValid returns the Valid field value
func (o *AssignedModeCommand) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *AssignedModeCommand) SetValid(v bool) {
	o.Valid = v
}

// GetSpare returns the Spare field value
func (o *AssignedModeCommand) GetSpare() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare
}

// GetSpareOk returns a tuple with the Spare field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetSpareOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare, true
}

// SetSpare sets field value
func (o *AssignedModeCommand) SetSpare(v int32) {
	o.Spare = v
}

// GetCommands returns the Commands field value
func (o *AssignedModeCommand) GetCommands() AssignedModeCommandCommands {
	if o == nil {
		var ret AssignedModeCommandCommands
		return ret
	}

	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
func (o *AssignedModeCommand) GetCommandsOk() (*AssignedModeCommandCommands, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Commands, true
}

// SetCommands sets field value
func (o *AssignedModeCommand) SetCommands(v AssignedModeCommandCommands) {
	o.Commands = v
}

func (o AssignedModeCommand) MarshalJSON() ([]byte, error) {
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
		toSerialize["Commands"] = o.Commands
	}
	return json.Marshal(toSerialize)
}

type NullableAssignedModeCommand struct {
	value *AssignedModeCommand
	isSet bool
}

func (v NullableAssignedModeCommand) Get() *AssignedModeCommand {
	return v.value
}

func (v *NullableAssignedModeCommand) Set(val *AssignedModeCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignedModeCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignedModeCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignedModeCommand(val *AssignedModeCommand) *NullableAssignedModeCommand {
	return &NullableAssignedModeCommand{value: val, isSet: true}
}

func (v NullableAssignedModeCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignedModeCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


