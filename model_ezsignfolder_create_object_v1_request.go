/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfolderCreateObjectV1Request Request for the /1/object/ezsignfolder/createObject API Request
type EzsignfolderCreateObjectV1Request struct {
	ObjEzsignfolder *EzsignfolderRequest `json:"objEzsignfolder,omitempty"`
	ObjEzsignfolderCompound *EzsignfolderRequestCompound `json:"objEzsignfolderCompound,omitempty"`
}

// NewEzsignfolderCreateObjectV1Request instantiates a new EzsignfolderCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderCreateObjectV1Request() *EzsignfolderCreateObjectV1Request {
	this := EzsignfolderCreateObjectV1Request{}
	return &this
}

// NewEzsignfolderCreateObjectV1RequestWithDefaults instantiates a new EzsignfolderCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderCreateObjectV1RequestWithDefaults() *EzsignfolderCreateObjectV1Request {
	this := EzsignfolderCreateObjectV1Request{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value if set, zero value otherwise.
func (o *EzsignfolderCreateObjectV1Request) GetObjEzsignfolder() EzsignfolderRequest {
	if o == nil || o.ObjEzsignfolder == nil {
		var ret EzsignfolderRequest
		return ret
	}
	return *o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderCreateObjectV1Request) GetObjEzsignfolderOk() (*EzsignfolderRequest, bool) {
	if o == nil || o.ObjEzsignfolder == nil {
		return nil, false
	}
	return o.ObjEzsignfolder, true
}

// HasObjEzsignfolder returns a boolean if a field has been set.
func (o *EzsignfolderCreateObjectV1Request) HasObjEzsignfolder() bool {
	if o != nil && o.ObjEzsignfolder != nil {
		return true
	}

	return false
}

// SetObjEzsignfolder gets a reference to the given EzsignfolderRequest and assigns it to the ObjEzsignfolder field.
func (o *EzsignfolderCreateObjectV1Request) SetObjEzsignfolder(v EzsignfolderRequest) {
	o.ObjEzsignfolder = &v
}

// GetObjEzsignfolderCompound returns the ObjEzsignfolderCompound field value if set, zero value otherwise.
func (o *EzsignfolderCreateObjectV1Request) GetObjEzsignfolderCompound() EzsignfolderRequestCompound {
	if o == nil || o.ObjEzsignfolderCompound == nil {
		var ret EzsignfolderRequestCompound
		return ret
	}
	return *o.ObjEzsignfolderCompound
}

// GetObjEzsignfolderCompoundOk returns a tuple with the ObjEzsignfolderCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderCreateObjectV1Request) GetObjEzsignfolderCompoundOk() (*EzsignfolderRequestCompound, bool) {
	if o == nil || o.ObjEzsignfolderCompound == nil {
		return nil, false
	}
	return o.ObjEzsignfolderCompound, true
}

// HasObjEzsignfolderCompound returns a boolean if a field has been set.
func (o *EzsignfolderCreateObjectV1Request) HasObjEzsignfolderCompound() bool {
	if o != nil && o.ObjEzsignfolderCompound != nil {
		return true
	}

	return false
}

// SetObjEzsignfolderCompound gets a reference to the given EzsignfolderRequestCompound and assigns it to the ObjEzsignfolderCompound field.
func (o *EzsignfolderCreateObjectV1Request) SetObjEzsignfolderCompound(v EzsignfolderRequestCompound) {
	o.ObjEzsignfolderCompound = &v
}

func (o EzsignfolderCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignfolder != nil {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	if o.ObjEzsignfolderCompound != nil {
		toSerialize["objEzsignfolderCompound"] = o.ObjEzsignfolderCompound
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderCreateObjectV1Request struct {
	value *EzsignfolderCreateObjectV1Request
	isSet bool
}

func (v NullableEzsignfolderCreateObjectV1Request) Get() *EzsignfolderCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsignfolderCreateObjectV1Request) Set(val *EzsignfolderCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderCreateObjectV1Request(val *EzsignfolderCreateObjectV1Request) *NullableEzsignfolderCreateObjectV1Request {
	return &NullableEzsignfolderCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignfolderCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


