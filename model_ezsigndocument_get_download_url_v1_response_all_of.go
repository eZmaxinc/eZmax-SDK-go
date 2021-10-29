/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentGetDownloadUrlV1ResponseAllOf struct for EzsigndocumentGetDownloadUrlV1ResponseAllOf
type EzsigndocumentGetDownloadUrlV1ResponseAllOf struct {
	MPayload EzsigndocumentGetDownloadUrlV1ResponseMPayload `json:"mPayload"`
}

// NewEzsigndocumentGetDownloadUrlV1ResponseAllOf instantiates a new EzsigndocumentGetDownloadUrlV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetDownloadUrlV1ResponseAllOf(mPayload EzsigndocumentGetDownloadUrlV1ResponseMPayload) *EzsigndocumentGetDownloadUrlV1ResponseAllOf {
	this := EzsigndocumentGetDownloadUrlV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetDownloadUrlV1ResponseAllOfWithDefaults instantiates a new EzsigndocumentGetDownloadUrlV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetDownloadUrlV1ResponseAllOfWithDefaults() *EzsigndocumentGetDownloadUrlV1ResponseAllOf {
	this := EzsigndocumentGetDownloadUrlV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetDownloadUrlV1ResponseAllOf) GetMPayload() EzsigndocumentGetDownloadUrlV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetDownloadUrlV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetDownloadUrlV1ResponseAllOf) GetMPayloadOk() (*EzsigndocumentGetDownloadUrlV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetDownloadUrlV1ResponseAllOf) SetMPayload(v EzsigndocumentGetDownloadUrlV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentGetDownloadUrlV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf struct {
	value *EzsigndocumentGetDownloadUrlV1ResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) Get() *EzsigndocumentGetDownloadUrlV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) Set(val *EzsigndocumentGetDownloadUrlV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetDownloadUrlV1ResponseAllOf(val *EzsigndocumentGetDownloadUrlV1ResponseAllOf) *NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf {
	return &NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


