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

// ActivesessionGetCurrentV1ResponseAllOf struct for ActivesessionGetCurrentV1ResponseAllOf
type ActivesessionGetCurrentV1ResponseAllOf struct {
	MPayload ActivesessionGetCurrentV1ResponseMPayload `json:"mPayload"`
}

// NewActivesessionGetCurrentV1ResponseAllOf instantiates a new ActivesessionGetCurrentV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivesessionGetCurrentV1ResponseAllOf(mPayload ActivesessionGetCurrentV1ResponseMPayload) *ActivesessionGetCurrentV1ResponseAllOf {
	this := ActivesessionGetCurrentV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewActivesessionGetCurrentV1ResponseAllOfWithDefaults instantiates a new ActivesessionGetCurrentV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivesessionGetCurrentV1ResponseAllOfWithDefaults() *ActivesessionGetCurrentV1ResponseAllOf {
	this := ActivesessionGetCurrentV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *ActivesessionGetCurrentV1ResponseAllOf) GetMPayload() ActivesessionGetCurrentV1ResponseMPayload {
	if o == nil {
		var ret ActivesessionGetCurrentV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ActivesessionGetCurrentV1ResponseAllOf) GetMPayloadOk() (*ActivesessionGetCurrentV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ActivesessionGetCurrentV1ResponseAllOf) SetMPayload(v ActivesessionGetCurrentV1ResponseMPayload) {
	o.MPayload = v
}

func (o ActivesessionGetCurrentV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableActivesessionGetCurrentV1ResponseAllOf struct {
	value *ActivesessionGetCurrentV1ResponseAllOf
	isSet bool
}

func (v NullableActivesessionGetCurrentV1ResponseAllOf) Get() *ActivesessionGetCurrentV1ResponseAllOf {
	return v.value
}

func (v *NullableActivesessionGetCurrentV1ResponseAllOf) Set(val *ActivesessionGetCurrentV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableActivesessionGetCurrentV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableActivesessionGetCurrentV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivesessionGetCurrentV1ResponseAllOf(val *ActivesessionGetCurrentV1ResponseAllOf) *NullableActivesessionGetCurrentV1ResponseAllOf {
	return &NullableActivesessionGetCurrentV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableActivesessionGetCurrentV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivesessionGetCurrentV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


