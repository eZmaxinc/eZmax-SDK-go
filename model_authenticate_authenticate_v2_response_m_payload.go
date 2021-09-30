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

// AuthenticateAuthenticateV2ResponseMPayload Payload for the /2/module/authenticate/authenticate API Request
type AuthenticateAuthenticateV2ResponseMPayload struct {
	// The Authorization key
	SAuthorization string `json:"sAuthorization"`
	// The secret key
	SSecret string `json:"sSecret"`
}

// NewAuthenticateAuthenticateV2ResponseMPayload instantiates a new AuthenticateAuthenticateV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateAuthenticateV2ResponseMPayload(sAuthorization string, sSecret string) *AuthenticateAuthenticateV2ResponseMPayload {
	this := AuthenticateAuthenticateV2ResponseMPayload{}
	this.SAuthorization = sAuthorization
	this.SSecret = sSecret
	return &this
}

// NewAuthenticateAuthenticateV2ResponseMPayloadWithDefaults instantiates a new AuthenticateAuthenticateV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateAuthenticateV2ResponseMPayloadWithDefaults() *AuthenticateAuthenticateV2ResponseMPayload {
	this := AuthenticateAuthenticateV2ResponseMPayload{}
	return &this
}

// GetSAuthorization returns the SAuthorization field value
func (o *AuthenticateAuthenticateV2ResponseMPayload) GetSAuthorization() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAuthorization
}

// GetSAuthorizationOk returns a tuple with the SAuthorization field value
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2ResponseMPayload) GetSAuthorizationOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAuthorization, true
}

// SetSAuthorization sets field value
func (o *AuthenticateAuthenticateV2ResponseMPayload) SetSAuthorization(v string) {
	o.SAuthorization = v
}

// GetSSecret returns the SSecret field value
func (o *AuthenticateAuthenticateV2ResponseMPayload) GetSSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSecret
}

// GetSSecretOk returns a tuple with the SSecret field value
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2ResponseMPayload) GetSSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SSecret, true
}

// SetSSecret sets field value
func (o *AuthenticateAuthenticateV2ResponseMPayload) SetSSecret(v string) {
	o.SSecret = v
}

func (o AuthenticateAuthenticateV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sAuthorization"] = o.SAuthorization
	}
	if true {
		toSerialize["sSecret"] = o.SSecret
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateAuthenticateV2ResponseMPayload struct {
	value *AuthenticateAuthenticateV2ResponseMPayload
	isSet bool
}

func (v NullableAuthenticateAuthenticateV2ResponseMPayload) Get() *AuthenticateAuthenticateV2ResponseMPayload {
	return v.value
}

func (v *NullableAuthenticateAuthenticateV2ResponseMPayload) Set(val *AuthenticateAuthenticateV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateAuthenticateV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateAuthenticateV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateAuthenticateV2ResponseMPayload(val *AuthenticateAuthenticateV2ResponseMPayload) *NullableAuthenticateAuthenticateV2ResponseMPayload {
	return &NullableAuthenticateAuthenticateV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableAuthenticateAuthenticateV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateAuthenticateV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


