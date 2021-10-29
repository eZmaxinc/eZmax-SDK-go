/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ApikeyRequestCompound An Apikey Object and children to create a complete structure
type ApikeyRequestCompound struct {
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	ObjApikeyDescription MultilingualApikeyDescription `json:"objApikeyDescription"`
}

// NewApikeyRequestCompound instantiates a new ApikeyRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyRequestCompound(fkiUserID int32, objApikeyDescription MultilingualApikeyDescription) *ApikeyRequestCompound {
	this := ApikeyRequestCompound{}
	this.FkiUserID = fkiUserID
	this.ObjApikeyDescription = objApikeyDescription
	return &this
}

// NewApikeyRequestCompoundWithDefaults instantiates a new ApikeyRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyRequestCompoundWithDefaults() *ApikeyRequestCompound {
	this := ApikeyRequestCompound{}
	return &this
}

// GetFkiUserID returns the FkiUserID field value
func (o *ApikeyRequestCompound) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *ApikeyRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *ApikeyRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetObjApikeyDescription returns the ObjApikeyDescription field value
func (o *ApikeyRequestCompound) GetObjApikeyDescription() MultilingualApikeyDescription {
	if o == nil {
		var ret MultilingualApikeyDescription
		return ret
	}

	return o.ObjApikeyDescription
}

// GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field value
// and a boolean to check if the value has been set.
func (o *ApikeyRequestCompound) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjApikeyDescription, true
}

// SetObjApikeyDescription sets field value
func (o *ApikeyRequestCompound) SetObjApikeyDescription(v MultilingualApikeyDescription) {
	o.ObjApikeyDescription = v
}

func (o ApikeyRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if true {
		toSerialize["objApikeyDescription"] = o.ObjApikeyDescription
	}
	return json.Marshal(toSerialize)
}

type NullableApikeyRequestCompound struct {
	value *ApikeyRequestCompound
	isSet bool
}

func (v NullableApikeyRequestCompound) Get() *ApikeyRequestCompound {
	return v.value
}

func (v *NullableApikeyRequestCompound) Set(val *ApikeyRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyRequestCompound(val *ApikeyRequestCompound) *NullableApikeyRequestCompound {
	return &NullableApikeyRequestCompound{value: val, isSet: true}
}

func (v NullableApikeyRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


