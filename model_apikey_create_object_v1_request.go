/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.48
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ApikeyCreateObjectV1Request Request for the /1/object/apikey/createObject API Request
type ApikeyCreateObjectV1Request struct {
	ObjApikey *ApikeyRequest `json:"objApikey,omitempty"`
	ObjApikeyCompound *ApikeyRequestCompound `json:"objApikeyCompound,omitempty"`
}

// NewApikeyCreateObjectV1Request instantiates a new ApikeyCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyCreateObjectV1Request() *ApikeyCreateObjectV1Request {
	this := ApikeyCreateObjectV1Request{}
	return &this
}

// NewApikeyCreateObjectV1RequestWithDefaults instantiates a new ApikeyCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyCreateObjectV1RequestWithDefaults() *ApikeyCreateObjectV1Request {
	this := ApikeyCreateObjectV1Request{}
	return &this
}

// GetObjApikey returns the ObjApikey field value if set, zero value otherwise.
func (o *ApikeyCreateObjectV1Request) GetObjApikey() ApikeyRequest {
	if o == nil || o.ObjApikey == nil {
		var ret ApikeyRequest
		return ret
	}
	return *o.ObjApikey
}

// GetObjApikeyOk returns a tuple with the ObjApikey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV1Request) GetObjApikeyOk() (*ApikeyRequest, bool) {
	if o == nil || o.ObjApikey == nil {
		return nil, false
	}
	return o.ObjApikey, true
}

// HasObjApikey returns a boolean if a field has been set.
func (o *ApikeyCreateObjectV1Request) HasObjApikey() bool {
	if o != nil && o.ObjApikey != nil {
		return true
	}

	return false
}

// SetObjApikey gets a reference to the given ApikeyRequest and assigns it to the ObjApikey field.
func (o *ApikeyCreateObjectV1Request) SetObjApikey(v ApikeyRequest) {
	o.ObjApikey = &v
}

// GetObjApikeyCompound returns the ObjApikeyCompound field value if set, zero value otherwise.
func (o *ApikeyCreateObjectV1Request) GetObjApikeyCompound() ApikeyRequestCompound {
	if o == nil || o.ObjApikeyCompound == nil {
		var ret ApikeyRequestCompound
		return ret
	}
	return *o.ObjApikeyCompound
}

// GetObjApikeyCompoundOk returns a tuple with the ObjApikeyCompound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV1Request) GetObjApikeyCompoundOk() (*ApikeyRequestCompound, bool) {
	if o == nil || o.ObjApikeyCompound == nil {
		return nil, false
	}
	return o.ObjApikeyCompound, true
}

// HasObjApikeyCompound returns a boolean if a field has been set.
func (o *ApikeyCreateObjectV1Request) HasObjApikeyCompound() bool {
	if o != nil && o.ObjApikeyCompound != nil {
		return true
	}

	return false
}

// SetObjApikeyCompound gets a reference to the given ApikeyRequestCompound and assigns it to the ObjApikeyCompound field.
func (o *ApikeyCreateObjectV1Request) SetObjApikeyCompound(v ApikeyRequestCompound) {
	o.ObjApikeyCompound = &v
}

func (o ApikeyCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjApikey != nil {
		toSerialize["objApikey"] = o.ObjApikey
	}
	if o.ObjApikeyCompound != nil {
		toSerialize["objApikeyCompound"] = o.ObjApikeyCompound
	}
	return json.Marshal(toSerialize)
}

type NullableApikeyCreateObjectV1Request struct {
	value *ApikeyCreateObjectV1Request
	isSet bool
}

func (v NullableApikeyCreateObjectV1Request) Get() *ApikeyCreateObjectV1Request {
	return v.value
}

func (v *NullableApikeyCreateObjectV1Request) Set(val *ApikeyCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyCreateObjectV1Request(val *ApikeyCreateObjectV1Request) *NullableApikeyCreateObjectV1Request {
	return &NullableApikeyCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableApikeyCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


