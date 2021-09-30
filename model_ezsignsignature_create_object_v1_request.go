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

// EzsignsignatureCreateObjectV1Request Request for the /1/object/ezsignsignature/createObject API Request
type EzsignsignatureCreateObjectV1Request struct {
	ObjEzsignsignature *EzsignsignatureRequest `json:"objEzsignsignature,omitempty"`
	ObjEzsignsignatureCompound *EzsignsignatureRequestCompound `json:"objEzsignsignatureCompound,omitempty"`
}

// NewEzsignsignatureCreateObjectV1Request instantiates a new EzsignsignatureCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV1Request() *EzsignsignatureCreateObjectV1Request {
	this := EzsignsignatureCreateObjectV1Request{}
	return &this
}

// NewEzsignsignatureCreateObjectV1RequestWithDefaults instantiates a new EzsignsignatureCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV1RequestWithDefaults() *EzsignsignatureCreateObjectV1Request {
	this := EzsignsignatureCreateObjectV1Request{}
	return &this
}

// GetObjEzsignsignature returns the ObjEzsignsignature field value if set, zero value otherwise.
func (o *EzsignsignatureCreateObjectV1Request) GetObjEzsignsignature() EzsignsignatureRequest {
	if o == nil || o.ObjEzsignsignature == nil {
		var ret EzsignsignatureRequest
		return ret
	}
	return *o.ObjEzsignsignature
}

// GetObjEzsignsignatureOk returns a tuple with the ObjEzsignsignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV1Request) GetObjEzsignsignatureOk() (*EzsignsignatureRequest, bool) {
	if o == nil || o.ObjEzsignsignature == nil {
		return nil, false
	}
	return o.ObjEzsignsignature, true
}

// HasObjEzsignsignature returns a boolean if a field has been set.
func (o *EzsignsignatureCreateObjectV1Request) HasObjEzsignsignature() bool {
	if o != nil && o.ObjEzsignsignature != nil {
		return true
	}

	return false
}

// SetObjEzsignsignature gets a reference to the given EzsignsignatureRequest and assigns it to the ObjEzsignsignature field.
func (o *EzsignsignatureCreateObjectV1Request) SetObjEzsignsignature(v EzsignsignatureRequest) {
	o.ObjEzsignsignature = &v
}

// GetObjEzsignsignatureCompound returns the ObjEzsignsignatureCompound field value if set, zero value otherwise.
func (o *EzsignsignatureCreateObjectV1Request) GetObjEzsignsignatureCompound() EzsignsignatureRequestCompound {
	if o == nil || o.ObjEzsignsignatureCompound == nil {
		var ret EzsignsignatureRequestCompound
		return ret
	}
	return *o.ObjEzsignsignatureCompound
}

// GetObjEzsignsignatureCompoundOk returns a tuple with the ObjEzsignsignatureCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV1Request) GetObjEzsignsignatureCompoundOk() (*EzsignsignatureRequestCompound, bool) {
	if o == nil || o.ObjEzsignsignatureCompound == nil {
		return nil, false
	}
	return o.ObjEzsignsignatureCompound, true
}

// HasObjEzsignsignatureCompound returns a boolean if a field has been set.
func (o *EzsignsignatureCreateObjectV1Request) HasObjEzsignsignatureCompound() bool {
	if o != nil && o.ObjEzsignsignatureCompound != nil {
		return true
	}

	return false
}

// SetObjEzsignsignatureCompound gets a reference to the given EzsignsignatureRequestCompound and assigns it to the ObjEzsignsignatureCompound field.
func (o *EzsignsignatureCreateObjectV1Request) SetObjEzsignsignatureCompound(v EzsignsignatureRequestCompound) {
	o.ObjEzsignsignatureCompound = &v
}

func (o EzsignsignatureCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignsignature != nil {
		toSerialize["objEzsignsignature"] = o.ObjEzsignsignature
	}
	if o.ObjEzsignsignatureCompound != nil {
		toSerialize["objEzsignsignatureCompound"] = o.ObjEzsignsignatureCompound
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignatureCreateObjectV1Request struct {
	value *EzsignsignatureCreateObjectV1Request
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV1Request) Get() *EzsignsignatureCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV1Request) Set(val *EzsignsignatureCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV1Request(val *EzsignsignatureCreateObjectV1Request) *NullableEzsignsignatureCreateObjectV1Request {
	return &NullableEzsignsignatureCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


