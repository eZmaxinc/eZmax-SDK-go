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

// EzsignsignatureCreateObjectV1ResponseAllOf struct for EzsignsignatureCreateObjectV1ResponseAllOf
type EzsignsignatureCreateObjectV1ResponseAllOf struct {
	MPayload EzsignsignatureCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignsignatureCreateObjectV1ResponseAllOf instantiates a new EzsignsignatureCreateObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV1ResponseAllOf(mPayload EzsignsignatureCreateObjectV1ResponseMPayload) *EzsignsignatureCreateObjectV1ResponseAllOf {
	this := EzsignsignatureCreateObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignsignatureCreateObjectV1ResponseAllOfWithDefaults instantiates a new EzsignsignatureCreateObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV1ResponseAllOfWithDefaults() *EzsignsignatureCreateObjectV1ResponseAllOf {
	this := EzsignsignatureCreateObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignsignatureCreateObjectV1ResponseAllOf) GetMPayload() EzsignsignatureCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignsignatureCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV1ResponseAllOf) GetMPayloadOk() (*EzsignsignatureCreateObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignsignatureCreateObjectV1ResponseAllOf) SetMPayload(v EzsignsignatureCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignsignatureCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignatureCreateObjectV1ResponseAllOf struct {
	value *EzsignsignatureCreateObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV1ResponseAllOf) Get() *EzsignsignatureCreateObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseAllOf) Set(val *EzsignsignatureCreateObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV1ResponseAllOf(val *EzsignsignatureCreateObjectV1ResponseAllOf) *NullableEzsignsignatureCreateObjectV1ResponseAllOf {
	return &NullableEzsignsignatureCreateObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


