/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfolderEditObjectV1Request Request for the /1/object/ezsignfolder/editObject API Request
type EzsignfolderEditObjectV1Request struct {
	ObjEzsignfolder EzsignfolderRequestCompound `json:"objEzsignfolder"`
}

// NewEzsignfolderEditObjectV1Request instantiates a new EzsignfolderEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderEditObjectV1Request(objEzsignfolder EzsignfolderRequestCompound) *EzsignfolderEditObjectV1Request {
	this := EzsignfolderEditObjectV1Request{}
	this.ObjEzsignfolder = objEzsignfolder
	return &this
}

// NewEzsignfolderEditObjectV1RequestWithDefaults instantiates a new EzsignfolderEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderEditObjectV1RequestWithDefaults() *EzsignfolderEditObjectV1Request {
	this := EzsignfolderEditObjectV1Request{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value
func (o *EzsignfolderEditObjectV1Request) GetObjEzsignfolder() EzsignfolderRequestCompound {
	if o == nil {
		var ret EzsignfolderRequestCompound
		return ret
	}

	return o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderEditObjectV1Request) GetObjEzsignfolderOk() (*EzsignfolderRequestCompound, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjEzsignfolder, true
}

// SetObjEzsignfolder sets field value
func (o *EzsignfolderEditObjectV1Request) SetObjEzsignfolder(v EzsignfolderRequestCompound) {
	o.ObjEzsignfolder = v
}

func (o EzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderEditObjectV1Request struct {
	value *EzsignfolderEditObjectV1Request
	isSet bool
}

func (v NullableEzsignfolderEditObjectV1Request) Get() *EzsignfolderEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignfolderEditObjectV1Request) Set(val *EzsignfolderEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderEditObjectV1Request(val *EzsignfolderEditObjectV1Request) *NullableEzsignfolderEditObjectV1Request {
	return &NullableEzsignfolderEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

