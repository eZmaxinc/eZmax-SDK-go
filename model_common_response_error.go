/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the CommonResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseError{}

// CommonResponseError Generic Error Message
type CommonResponseError struct {
	// The message giving details about the error
	SErrorMessage string `json:"sErrorMessage"`
	EErrorCode FieldEErrorCode `json:"eErrorCode"`
}

// NewCommonResponseError instantiates a new CommonResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseError(sErrorMessage string, eErrorCode FieldEErrorCode) *CommonResponseError {
	this := CommonResponseError{}
	this.SErrorMessage = sErrorMessage
	this.EErrorCode = eErrorCode
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
	if o == nil {
		var ret string
		return ret
	}

	return o.SErrorMessage
}

// GetSErrorMessageOk returns a tuple with the SErrorMessage field value
// and a boolean to check if the value has been set.
func (o *CommonResponseError) GetSErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SErrorMessage, true
}

// SetSErrorMessage sets field value
func (o *CommonResponseError) SetSErrorMessage(v string) {
	o.SErrorMessage = v
}

// GetEErrorCode returns the EErrorCode field value
func (o *CommonResponseError) GetEErrorCode() FieldEErrorCode {
	if o == nil {
		var ret FieldEErrorCode
		return ret
	}

	return o.EErrorCode
}

// GetEErrorCodeOk returns a tuple with the EErrorCode field value
// and a boolean to check if the value has been set.
func (o *CommonResponseError) GetEErrorCodeOk() (*FieldEErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EErrorCode, true
}

// SetEErrorCode sets field value
func (o *CommonResponseError) SetEErrorCode(v FieldEErrorCode) {
	o.EErrorCode = v
}

func (o CommonResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sErrorMessage"] = o.SErrorMessage
	toSerialize["eErrorCode"] = o.EErrorCode
	return toSerialize, nil
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


