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

// EzsignfolderGetObjectV1ResponseAllOf struct for EzsignfolderGetObjectV1ResponseAllOf
type EzsignfolderGetObjectV1ResponseAllOf struct {
	MPayload EzsignfolderGetObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderGetObjectV1ResponseAllOf instantiates a new EzsignfolderGetObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetObjectV1ResponseAllOf(mPayload EzsignfolderGetObjectV1ResponseMPayload) *EzsignfolderGetObjectV1ResponseAllOf {
	this := EzsignfolderGetObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetObjectV1ResponseAllOfWithDefaults instantiates a new EzsignfolderGetObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetObjectV1ResponseAllOfWithDefaults() *EzsignfolderGetObjectV1ResponseAllOf {
	this := EzsignfolderGetObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetObjectV1ResponseAllOf) GetMPayload() EzsignfolderGetObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetObjectV1ResponseAllOf) GetMPayloadOk() (*EzsignfolderGetObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetObjectV1ResponseAllOf) SetMPayload(v EzsignfolderGetObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetObjectV1ResponseAllOf struct {
	value *EzsignfolderGetObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfolderGetObjectV1ResponseAllOf) Get() *EzsignfolderGetObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfolderGetObjectV1ResponseAllOf) Set(val *EzsignfolderGetObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetObjectV1ResponseAllOf(val *EzsignfolderGetObjectV1ResponseAllOf) *NullableEzsignfolderGetObjectV1ResponseAllOf {
	return &NullableEzsignfolderGetObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


