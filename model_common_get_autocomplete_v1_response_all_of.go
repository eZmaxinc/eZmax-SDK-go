/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.48
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonGetAutocompleteV1ResponseAllOf struct for CommonGetAutocompleteV1ResponseAllOf
type CommonGetAutocompleteV1ResponseAllOf struct {
	MPayload []CommonGetAutocompleteV1ResponseMPayload `json:"mPayload"`
}

// NewCommonGetAutocompleteV1ResponseAllOf instantiates a new CommonGetAutocompleteV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonGetAutocompleteV1ResponseAllOf(mPayload []CommonGetAutocompleteV1ResponseMPayload) *CommonGetAutocompleteV1ResponseAllOf {
	this := CommonGetAutocompleteV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewCommonGetAutocompleteV1ResponseAllOfWithDefaults instantiates a new CommonGetAutocompleteV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonGetAutocompleteV1ResponseAllOfWithDefaults() *CommonGetAutocompleteV1ResponseAllOf {
	this := CommonGetAutocompleteV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *CommonGetAutocompleteV1ResponseAllOf) GetMPayload() []CommonGetAutocompleteV1ResponseMPayload {
	if o == nil {
		var ret []CommonGetAutocompleteV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *CommonGetAutocompleteV1ResponseAllOf) GetMPayloadOk() (*[]CommonGetAutocompleteV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *CommonGetAutocompleteV1ResponseAllOf) SetMPayload(v []CommonGetAutocompleteV1ResponseMPayload) {
	o.MPayload = v
}

func (o CommonGetAutocompleteV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableCommonGetAutocompleteV1ResponseAllOf struct {
	value *CommonGetAutocompleteV1ResponseAllOf
	isSet bool
}

func (v NullableCommonGetAutocompleteV1ResponseAllOf) Get() *CommonGetAutocompleteV1ResponseAllOf {
	return v.value
}

func (v *NullableCommonGetAutocompleteV1ResponseAllOf) Set(val *CommonGetAutocompleteV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonGetAutocompleteV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonGetAutocompleteV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonGetAutocompleteV1ResponseAllOf(val *CommonGetAutocompleteV1ResponseAllOf) *NullableCommonGetAutocompleteV1ResponseAllOf {
	return &NullableCommonGetAutocompleteV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableCommonGetAutocompleteV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonGetAutocompleteV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

