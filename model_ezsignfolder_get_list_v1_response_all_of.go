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

// EzsignfolderGetListV1ResponseAllOf struct for EzsignfolderGetListV1ResponseAllOf
type EzsignfolderGetListV1ResponseAllOf struct {
	MPayload EzsignfolderGetListV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderGetListV1ResponseAllOf instantiates a new EzsignfolderGetListV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetListV1ResponseAllOf(mPayload EzsignfolderGetListV1ResponseMPayload) *EzsignfolderGetListV1ResponseAllOf {
	this := EzsignfolderGetListV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetListV1ResponseAllOfWithDefaults instantiates a new EzsignfolderGetListV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetListV1ResponseAllOfWithDefaults() *EzsignfolderGetListV1ResponseAllOf {
	this := EzsignfolderGetListV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetListV1ResponseAllOf) GetMPayload() EzsignfolderGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetListV1ResponseAllOf) GetMPayloadOk() (*EzsignfolderGetListV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetListV1ResponseAllOf) SetMPayload(v EzsignfolderGetListV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderGetListV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetListV1ResponseAllOf struct {
	value *EzsignfolderGetListV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfolderGetListV1ResponseAllOf) Get() *EzsignfolderGetListV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfolderGetListV1ResponseAllOf) Set(val *EzsignfolderGetListV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetListV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetListV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetListV1ResponseAllOf(val *EzsignfolderGetListV1ResponseAllOf) *NullableEzsignfolderGetListV1ResponseAllOf {
	return &NullableEzsignfolderGetListV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderGetListV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetListV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

