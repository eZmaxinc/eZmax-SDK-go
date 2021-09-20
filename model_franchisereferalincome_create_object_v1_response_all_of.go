/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// FranchisereferalincomeCreateObjectV1ResponseAllOf struct for FranchisereferalincomeCreateObjectV1ResponseAllOf
type FranchisereferalincomeCreateObjectV1ResponseAllOf struct {
	MPayload FranchisereferalincomeCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewFranchisereferalincomeCreateObjectV1ResponseAllOf instantiates a new FranchisereferalincomeCreateObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV1ResponseAllOf(mPayload FranchisereferalincomeCreateObjectV1ResponseMPayload) *FranchisereferalincomeCreateObjectV1ResponseAllOf {
	this := FranchisereferalincomeCreateObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewFranchisereferalincomeCreateObjectV1ResponseAllOfWithDefaults instantiates a new FranchisereferalincomeCreateObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV1ResponseAllOfWithDefaults() *FranchisereferalincomeCreateObjectV1ResponseAllOf {
	this := FranchisereferalincomeCreateObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *FranchisereferalincomeCreateObjectV1ResponseAllOf) GetMPayload() FranchisereferalincomeCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret FranchisereferalincomeCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1ResponseAllOf) GetMPayloadOk() (*FranchisereferalincomeCreateObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *FranchisereferalincomeCreateObjectV1ResponseAllOf) SetMPayload(v FranchisereferalincomeCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o FranchisereferalincomeCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableFranchisereferalincomeCreateObjectV1ResponseAllOf struct {
	value *FranchisereferalincomeCreateObjectV1ResponseAllOf
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) Get() *FranchisereferalincomeCreateObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) Set(val *FranchisereferalincomeCreateObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV1ResponseAllOf(val *FranchisereferalincomeCreateObjectV1ResponseAllOf) *NullableFranchisereferalincomeCreateObjectV1ResponseAllOf {
	return &NullableFranchisereferalincomeCreateObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


