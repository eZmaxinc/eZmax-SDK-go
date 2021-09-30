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

// UserCreateEzsignuserV1ResponseMPayload Payload for the /1/module/user/createEzsignuser API Request
type UserCreateEzsignuserV1ResponseMPayload struct {
	// An array of email addresses that succeeded.
	ASEmailAddressSuccess []string `json:"a_sEmailAddressSuccess"`
	// An array of email addresses that failed.
	ASEmailAddressFailure []string `json:"a_sEmailAddressFailure"`
}

// NewUserCreateEzsignuserV1ResponseMPayload instantiates a new UserCreateEzsignuserV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateEzsignuserV1ResponseMPayload(aSEmailAddressSuccess []string, aSEmailAddressFailure []string) *UserCreateEzsignuserV1ResponseMPayload {
	this := UserCreateEzsignuserV1ResponseMPayload{}
	this.ASEmailAddressSuccess = aSEmailAddressSuccess
	this.ASEmailAddressFailure = aSEmailAddressFailure
	return &this
}

// NewUserCreateEzsignuserV1ResponseMPayloadWithDefaults instantiates a new UserCreateEzsignuserV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateEzsignuserV1ResponseMPayloadWithDefaults() *UserCreateEzsignuserV1ResponseMPayload {
	this := UserCreateEzsignuserV1ResponseMPayload{}
	return &this
}

// GetASEmailAddressSuccess returns the ASEmailAddressSuccess field value
func (o *UserCreateEzsignuserV1ResponseMPayload) GetASEmailAddressSuccess() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEmailAddressSuccess
}

// GetASEmailAddressSuccessOk returns a tuple with the ASEmailAddressSuccess field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1ResponseMPayload) GetASEmailAddressSuccessOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASEmailAddressSuccess, true
}

// SetASEmailAddressSuccess sets field value
func (o *UserCreateEzsignuserV1ResponseMPayload) SetASEmailAddressSuccess(v []string) {
	o.ASEmailAddressSuccess = v
}

// GetASEmailAddressFailure returns the ASEmailAddressFailure field value
func (o *UserCreateEzsignuserV1ResponseMPayload) GetASEmailAddressFailure() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEmailAddressFailure
}

// GetASEmailAddressFailureOk returns a tuple with the ASEmailAddressFailure field value
// and a boolean to check if the value has been set.
func (o *UserCreateEzsignuserV1ResponseMPayload) GetASEmailAddressFailureOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ASEmailAddressFailure, true
}

// SetASEmailAddressFailure sets field value
func (o *UserCreateEzsignuserV1ResponseMPayload) SetASEmailAddressFailure(v []string) {
	o.ASEmailAddressFailure = v
}

func (o UserCreateEzsignuserV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_sEmailAddressSuccess"] = o.ASEmailAddressSuccess
	}
	if true {
		toSerialize["a_sEmailAddressFailure"] = o.ASEmailAddressFailure
	}
	return json.Marshal(toSerialize)
}

type NullableUserCreateEzsignuserV1ResponseMPayload struct {
	value *UserCreateEzsignuserV1ResponseMPayload
	isSet bool
}

func (v NullableUserCreateEzsignuserV1ResponseMPayload) Get() *UserCreateEzsignuserV1ResponseMPayload {
	return v.value
}

func (v *NullableUserCreateEzsignuserV1ResponseMPayload) Set(val *UserCreateEzsignuserV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateEzsignuserV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateEzsignuserV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateEzsignuserV1ResponseMPayload(val *UserCreateEzsignuserV1ResponseMPayload) *NullableUserCreateEzsignuserV1ResponseMPayload {
	return &NullableUserCreateEzsignuserV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserCreateEzsignuserV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateEzsignuserV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


