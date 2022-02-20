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

// EzsigndocumentCreateObjectV1ResponseAllOf struct for EzsigndocumentCreateObjectV1ResponseAllOf
type EzsigndocumentCreateObjectV1ResponseAllOf struct {
	MPayload EzsigndocumentCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsigndocumentCreateObjectV1ResponseAllOf instantiates a new EzsigndocumentCreateObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateObjectV1ResponseAllOf(mPayload EzsigndocumentCreateObjectV1ResponseMPayload) *EzsigndocumentCreateObjectV1ResponseAllOf {
	this := EzsigndocumentCreateObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentCreateObjectV1ResponseAllOfWithDefaults instantiates a new EzsigndocumentCreateObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateObjectV1ResponseAllOfWithDefaults() *EzsigndocumentCreateObjectV1ResponseAllOf {
	this := EzsigndocumentCreateObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentCreateObjectV1ResponseAllOf) GetMPayload() EzsigndocumentCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1ResponseAllOf) GetMPayloadOk() (*EzsigndocumentCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentCreateObjectV1ResponseAllOf) SetMPayload(v EzsigndocumentCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentCreateObjectV1ResponseAllOf struct {
	value *EzsigndocumentCreateObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentCreateObjectV1ResponseAllOf) Get() *EzsigndocumentCreateObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseAllOf) Set(val *EzsigndocumentCreateObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateObjectV1ResponseAllOf(val *EzsigndocumentCreateObjectV1ResponseAllOf) *NullableEzsigndocumentCreateObjectV1ResponseAllOf {
	return &NullableEzsigndocumentCreateObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


