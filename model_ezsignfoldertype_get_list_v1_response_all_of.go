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

// EzsignfoldertypeGetListV1ResponseAllOf struct for EzsignfoldertypeGetListV1ResponseAllOf
type EzsignfoldertypeGetListV1ResponseAllOf struct {
	MPayload EzsignfoldertypeGetListV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfoldertypeGetListV1ResponseAllOf instantiates a new EzsignfoldertypeGetListV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeGetListV1ResponseAllOf(mPayload EzsignfoldertypeGetListV1ResponseMPayload) *EzsignfoldertypeGetListV1ResponseAllOf {
	this := EzsignfoldertypeGetListV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldertypeGetListV1ResponseAllOfWithDefaults instantiates a new EzsignfoldertypeGetListV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeGetListV1ResponseAllOfWithDefaults() *EzsignfoldertypeGetListV1ResponseAllOf {
	this := EzsignfoldertypeGetListV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldertypeGetListV1ResponseAllOf) GetMPayload() EzsignfoldertypeGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsignfoldertypeGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetListV1ResponseAllOf) GetMPayloadOk() (*EzsignfoldertypeGetListV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldertypeGetListV1ResponseAllOf) SetMPayload(v EzsignfoldertypeGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfoldertypeGetListV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldertypeGetListV1ResponseAllOf struct {
	value *EzsignfoldertypeGetListV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfoldertypeGetListV1ResponseAllOf) Get() *EzsignfoldertypeGetListV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfoldertypeGetListV1ResponseAllOf) Set(val *EzsignfoldertypeGetListV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeGetListV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeGetListV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeGetListV1ResponseAllOf(val *EzsignfoldertypeGetListV1ResponseAllOf) *NullableEzsignfoldertypeGetListV1ResponseAllOf {
	return &NullableEzsignfoldertypeGetListV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeGetListV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeGetListV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


