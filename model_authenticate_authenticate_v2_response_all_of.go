/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// AuthenticateAuthenticateV2ResponseAllOf struct for AuthenticateAuthenticateV2ResponseAllOf
type AuthenticateAuthenticateV2ResponseAllOf struct {
	MPayload AuthenticateAuthenticateV2ResponseMPayload `json:"mPayload"`
}

// NewAuthenticateAuthenticateV2ResponseAllOf instantiates a new AuthenticateAuthenticateV2ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateAuthenticateV2ResponseAllOf(mPayload AuthenticateAuthenticateV2ResponseMPayload) *AuthenticateAuthenticateV2ResponseAllOf {
	this := AuthenticateAuthenticateV2ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewAuthenticateAuthenticateV2ResponseAllOfWithDefaults instantiates a new AuthenticateAuthenticateV2ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateAuthenticateV2ResponseAllOfWithDefaults() *AuthenticateAuthenticateV2ResponseAllOf {
	this := AuthenticateAuthenticateV2ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *AuthenticateAuthenticateV2ResponseAllOf) GetMPayload() AuthenticateAuthenticateV2ResponseMPayload {
	if o == nil {
		var ret AuthenticateAuthenticateV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2ResponseAllOf) GetMPayloadOk() (*AuthenticateAuthenticateV2ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *AuthenticateAuthenticateV2ResponseAllOf) SetMPayload(v AuthenticateAuthenticateV2ResponseMPayload) {
	o.MPayload = v
}

func (o AuthenticateAuthenticateV2ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateAuthenticateV2ResponseAllOf struct {
	value *AuthenticateAuthenticateV2ResponseAllOf
	isSet bool
}

func (v NullableAuthenticateAuthenticateV2ResponseAllOf) Get() *AuthenticateAuthenticateV2ResponseAllOf {
	return v.value
}

func (v *NullableAuthenticateAuthenticateV2ResponseAllOf) Set(val *AuthenticateAuthenticateV2ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateAuthenticateV2ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateAuthenticateV2ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateAuthenticateV2ResponseAllOf(val *AuthenticateAuthenticateV2ResponseAllOf) *NullableAuthenticateAuthenticateV2ResponseAllOf {
	return &NullableAuthenticateAuthenticateV2ResponseAllOf{value: val, isSet: true}
}

func (v NullableAuthenticateAuthenticateV2ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateAuthenticateV2ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


