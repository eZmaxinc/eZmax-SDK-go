/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentGetObjectV1ResponseAllOf struct for EzsigndocumentGetObjectV1ResponseAllOf
type EzsigndocumentGetObjectV1ResponseAllOf struct {
	MPayload EzsigndocumentGetObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsigndocumentGetObjectV1ResponseAllOf instantiates a new EzsigndocumentGetObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetObjectV1ResponseAllOf(mPayload EzsigndocumentGetObjectV1ResponseMPayload) *EzsigndocumentGetObjectV1ResponseAllOf {
	this := EzsigndocumentGetObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetObjectV1ResponseAllOfWithDefaults instantiates a new EzsigndocumentGetObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetObjectV1ResponseAllOfWithDefaults() *EzsigndocumentGetObjectV1ResponseAllOf {
	this := EzsigndocumentGetObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetObjectV1ResponseAllOf) GetMPayload() EzsigndocumentGetObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetObjectV1ResponseAllOf) GetMPayloadOk() (*EzsigndocumentGetObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetObjectV1ResponseAllOf) SetMPayload(v EzsigndocumentGetObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetObjectV1ResponseAllOf struct {
	value *EzsigndocumentGetObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentGetObjectV1ResponseAllOf) Get() *EzsigndocumentGetObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentGetObjectV1ResponseAllOf) Set(val *EzsigndocumentGetObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetObjectV1ResponseAllOf(val *EzsigndocumentGetObjectV1ResponseAllOf) *NullableEzsigndocumentGetObjectV1ResponseAllOf {
	return &NullableEzsigndocumentGetObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


