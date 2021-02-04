/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.28
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// CommonResponseError Generic Error Message
type CommonResponseError struct {
	// More detail about the error
	SErrorMessage string `json:"sErrorMessage"`
}

// NewCommonResponseError instantiates a new CommonResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseError(sErrorMessage string, ) *CommonResponseError {
	this := CommonResponseError{}
	this.SErrorMessage = sErrorMessage
	return &this
}

// NewCommonResponseErrorWithDefaults instantiates a new CommonResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseErrorWithDefaults() *CommonResponseError {
	this := CommonResponseError{}
	return &this
}

// GetSErrorMessage returns the SErrorMessage field value
func (o *CommonResponseError) GetSErrorMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SErrorMessage
}

// GetSErrorMessageOk returns a tuple with the SErrorMessage field value
// and a boolean to check if the value has been set.
func (o *CommonResponseError) GetSErrorMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SErrorMessage, true
}

// SetSErrorMessage sets field value
func (o *CommonResponseError) SetSErrorMessage(v string) {
	o.SErrorMessage = v
}

func (o CommonResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sErrorMessage"] = o.SErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseError struct {
	value *CommonResponseError
	isSet bool
}

func (v NullableCommonResponseError) Get() *CommonResponseError {
	return v.value
}

func (v *NullableCommonResponseError) Set(val *CommonResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseError(val *CommonResponseError) *NullableCommonResponseError {
	return &NullableCommonResponseError{value: val, isSet: true}
}

func (v NullableCommonResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

