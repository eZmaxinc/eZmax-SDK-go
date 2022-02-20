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

// EzsignfolderCreateObjectV1ResponseAllOf struct for EzsignfolderCreateObjectV1ResponseAllOf
type EzsignfolderCreateObjectV1ResponseAllOf struct {
	MPayload EzsignfolderCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderCreateObjectV1ResponseAllOf instantiates a new EzsignfolderCreateObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderCreateObjectV1ResponseAllOf(mPayload EzsignfolderCreateObjectV1ResponseMPayload) *EzsignfolderCreateObjectV1ResponseAllOf {
	this := EzsignfolderCreateObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderCreateObjectV1ResponseAllOfWithDefaults instantiates a new EzsignfolderCreateObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderCreateObjectV1ResponseAllOfWithDefaults() *EzsignfolderCreateObjectV1ResponseAllOf {
	this := EzsignfolderCreateObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderCreateObjectV1ResponseAllOf) GetMPayload() EzsignfolderCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderCreateObjectV1ResponseAllOf) GetMPayloadOk() (*EzsignfolderCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderCreateObjectV1ResponseAllOf) SetMPayload(v EzsignfolderCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderCreateObjectV1ResponseAllOf struct {
	value *EzsignfolderCreateObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfolderCreateObjectV1ResponseAllOf) Get() *EzsignfolderCreateObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfolderCreateObjectV1ResponseAllOf) Set(val *EzsignfolderCreateObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderCreateObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderCreateObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderCreateObjectV1ResponseAllOf(val *EzsignfolderCreateObjectV1ResponseAllOf) *NullableEzsignfolderCreateObjectV1ResponseAllOf {
	return &NullableEzsignfolderCreateObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderCreateObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


