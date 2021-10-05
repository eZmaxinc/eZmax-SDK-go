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

// EzsigndocumentGetEzsignpagesV1ResponseAllOf struct for EzsigndocumentGetEzsignpagesV1ResponseAllOf
type EzsigndocumentGetEzsignpagesV1ResponseAllOf struct {
	MPayload EzsigndocumentGetEzsignpagesV1ResponseMPayload `json:"mPayload"`
}

// NewEzsigndocumentGetEzsignpagesV1ResponseAllOf instantiates a new EzsigndocumentGetEzsignpagesV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignpagesV1ResponseAllOf(mPayload EzsigndocumentGetEzsignpagesV1ResponseMPayload) *EzsigndocumentGetEzsignpagesV1ResponseAllOf {
	this := EzsigndocumentGetEzsignpagesV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetEzsignpagesV1ResponseAllOfWithDefaults instantiates a new EzsigndocumentGetEzsignpagesV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignpagesV1ResponseAllOfWithDefaults() *EzsigndocumentGetEzsignpagesV1ResponseAllOf {
	this := EzsigndocumentGetEzsignpagesV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetEzsignpagesV1ResponseAllOf) GetMPayload() EzsigndocumentGetEzsignpagesV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetEzsignpagesV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignpagesV1ResponseAllOf) GetMPayloadOk() (*EzsigndocumentGetEzsignpagesV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetEzsignpagesV1ResponseAllOf) SetMPayload(v EzsigndocumentGetEzsignpagesV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentGetEzsignpagesV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf struct {
	value *EzsigndocumentGetEzsignpagesV1ResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) Get() *EzsigndocumentGetEzsignpagesV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) Set(val *EzsigndocumentGetEzsignpagesV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignpagesV1ResponseAllOf(val *EzsigndocumentGetEzsignpagesV1ResponseAllOf) *NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf {
	return &NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

