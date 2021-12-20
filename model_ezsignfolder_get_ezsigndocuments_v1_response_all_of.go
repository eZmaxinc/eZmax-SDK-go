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

// EzsignfolderGetEzsigndocumentsV1ResponseAllOf struct for EzsignfolderGetEzsigndocumentsV1ResponseAllOf
type EzsignfolderGetEzsigndocumentsV1ResponseAllOf struct {
	MPayload EzsignfolderGetEzsigndocumentsV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderGetEzsigndocumentsV1ResponseAllOf instantiates a new EzsignfolderGetEzsigndocumentsV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetEzsigndocumentsV1ResponseAllOf(mPayload EzsignfolderGetEzsigndocumentsV1ResponseMPayload) *EzsignfolderGetEzsigndocumentsV1ResponseAllOf {
	this := EzsignfolderGetEzsigndocumentsV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetEzsigndocumentsV1ResponseAllOfWithDefaults instantiates a new EzsignfolderGetEzsigndocumentsV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetEzsigndocumentsV1ResponseAllOfWithDefaults() *EzsignfolderGetEzsigndocumentsV1ResponseAllOf {
	this := EzsignfolderGetEzsigndocumentsV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetEzsigndocumentsV1ResponseAllOf) GetMPayload() EzsignfolderGetEzsigndocumentsV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetEzsigndocumentsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsigndocumentsV1ResponseAllOf) GetMPayloadOk() (*EzsignfolderGetEzsigndocumentsV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetEzsigndocumentsV1ResponseAllOf) SetMPayload(v EzsignfolderGetEzsigndocumentsV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderGetEzsigndocumentsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf struct {
	value *EzsignfolderGetEzsigndocumentsV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) Get() *EzsignfolderGetEzsigndocumentsV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) Set(val *EzsignfolderGetEzsigndocumentsV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf(val *EzsignfolderGetEzsigndocumentsV1ResponseAllOf) *NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf {
	return &NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


