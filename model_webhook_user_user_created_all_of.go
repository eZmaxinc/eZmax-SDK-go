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

// WebhookUserUserCreatedAllOf struct for WebhookUserUserCreatedAllOf
type WebhookUserUserCreatedAllOf struct {
	ObjUser UserResponse `json:"objUser"`
}

// NewWebhookUserUserCreatedAllOf instantiates a new WebhookUserUserCreatedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUserUserCreatedAllOf(objUser UserResponse) *WebhookUserUserCreatedAllOf {
	this := WebhookUserUserCreatedAllOf{}
	this.ObjUser = objUser
	return &this
}

// NewWebhookUserUserCreatedAllOfWithDefaults instantiates a new WebhookUserUserCreatedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUserUserCreatedAllOfWithDefaults() *WebhookUserUserCreatedAllOf {
	this := WebhookUserUserCreatedAllOf{}
	return &this
}

// GetObjUser returns the ObjUser field value
func (o *WebhookUserUserCreatedAllOf) GetObjUser() UserResponse {
	if o == nil {
		var ret UserResponse
		return ret
	}

	return o.ObjUser
}

// GetObjUserOk returns a tuple with the ObjUser field value
// and a boolean to check if the value has been set.
func (o *WebhookUserUserCreatedAllOf) GetObjUserOk() (*UserResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjUser, true
}

// SetObjUser sets field value
func (o *WebhookUserUserCreatedAllOf) SetObjUser(v UserResponse) {
	o.ObjUser = v
}

func (o WebhookUserUserCreatedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objUser"] = o.ObjUser
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookUserUserCreatedAllOf struct {
	value *WebhookUserUserCreatedAllOf
	isSet bool
}

func (v NullableWebhookUserUserCreatedAllOf) Get() *WebhookUserUserCreatedAllOf {
	return v.value
}

func (v *NullableWebhookUserUserCreatedAllOf) Set(val *WebhookUserUserCreatedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUserUserCreatedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUserUserCreatedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUserUserCreatedAllOf(val *WebhookUserUserCreatedAllOf) *NullableWebhookUserUserCreatedAllOf {
	return &NullableWebhookUserUserCreatedAllOf{value: val, isSet: true}
}

func (v NullableWebhookUserUserCreatedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUserUserCreatedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


