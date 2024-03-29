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

// EzsignsignatureGetObjectV1ResponseAllOf struct for EzsignsignatureGetObjectV1ResponseAllOf
type EzsignsignatureGetObjectV1ResponseAllOf struct {
	// Payload for the /1/object/ezsignsignature/getObject API Request
	MPayload map[string]interface{} `json:"mPayload"`
}

// NewEzsignsignatureGetObjectV1ResponseAllOf instantiates a new EzsignsignatureGetObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureGetObjectV1ResponseAllOf(mPayload map[string]interface{}) *EzsignsignatureGetObjectV1ResponseAllOf {
	this := EzsignsignatureGetObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignsignatureGetObjectV1ResponseAllOfWithDefaults instantiates a new EzsignsignatureGetObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureGetObjectV1ResponseAllOfWithDefaults() *EzsignsignatureGetObjectV1ResponseAllOf {
	this := EzsignsignatureGetObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignsignatureGetObjectV1ResponseAllOf) GetMPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureGetObjectV1ResponseAllOf) GetMPayloadOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignsignatureGetObjectV1ResponseAllOf) SetMPayload(v map[string]interface{}) {
	o.MPayload = v
}

func (o EzsignsignatureGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignatureGetObjectV1ResponseAllOf struct {
	value *EzsignsignatureGetObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignsignatureGetObjectV1ResponseAllOf) Get() *EzsignsignatureGetObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignsignatureGetObjectV1ResponseAllOf) Set(val *EzsignsignatureGetObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureGetObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureGetObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureGetObjectV1ResponseAllOf(val *EzsignsignatureGetObjectV1ResponseAllOf) *NullableEzsignsignatureGetObjectV1ResponseAllOf {
	return &NullableEzsignsignatureGetObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignsignatureGetObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureGetObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


