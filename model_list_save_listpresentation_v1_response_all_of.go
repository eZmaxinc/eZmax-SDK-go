/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ListSaveListpresentationV1ResponseAllOf struct for ListSaveListpresentationV1ResponseAllOf
type ListSaveListpresentationV1ResponseAllOf struct {
	MPayload ListSaveListpresentationV1ResponseMPayload `json:"mPayload"`
}

// NewListSaveListpresentationV1ResponseAllOf instantiates a new ListSaveListpresentationV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSaveListpresentationV1ResponseAllOf(mPayload ListSaveListpresentationV1ResponseMPayload) *ListSaveListpresentationV1ResponseAllOf {
	this := ListSaveListpresentationV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewListSaveListpresentationV1ResponseAllOfWithDefaults instantiates a new ListSaveListpresentationV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSaveListpresentationV1ResponseAllOfWithDefaults() *ListSaveListpresentationV1ResponseAllOf {
	this := ListSaveListpresentationV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *ListSaveListpresentationV1ResponseAllOf) GetMPayload() ListSaveListpresentationV1ResponseMPayload {
	if o == nil {
		var ret ListSaveListpresentationV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ListSaveListpresentationV1ResponseAllOf) GetMPayloadOk() (*ListSaveListpresentationV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ListSaveListpresentationV1ResponseAllOf) SetMPayload(v ListSaveListpresentationV1ResponseMPayload) {
	o.MPayload = v
}

func (o ListSaveListpresentationV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableListSaveListpresentationV1ResponseAllOf struct {
	value *ListSaveListpresentationV1ResponseAllOf
	isSet bool
}

func (v NullableListSaveListpresentationV1ResponseAllOf) Get() *ListSaveListpresentationV1ResponseAllOf {
	return v.value
}

func (v *NullableListSaveListpresentationV1ResponseAllOf) Set(val *ListSaveListpresentationV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListSaveListpresentationV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListSaveListpresentationV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSaveListpresentationV1ResponseAllOf(val *ListSaveListpresentationV1ResponseAllOf) *NullableListSaveListpresentationV1ResponseAllOf {
	return &NullableListSaveListpresentationV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableListSaveListpresentationV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSaveListpresentationV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


