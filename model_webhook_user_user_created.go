/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.
 *
 * API version: 1.0.30
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// WebhookUserUserCreated This is the base Webhook object
type WebhookUserUserCreated struct {
	ObjUser UserResponse `json:"objUser"`
	ObjWebhook WebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponse `json:"a_objAttempt"`
}

// NewWebhookUserUserCreated instantiates a new WebhookUserUserCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookUserUserCreated(objUser UserResponse, objWebhook WebhookResponse, aObjAttempt []AttemptResponse, ) *WebhookUserUserCreated {
	this := WebhookUserUserCreated{}
	this.ObjUser = objUser
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	return &this
}

// NewWebhookUserUserCreatedWithDefaults instantiates a new WebhookUserUserCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookUserUserCreatedWithDefaults() *WebhookUserUserCreated {
	this := WebhookUserUserCreated{}
	return &this
}

// GetObjUser returns the ObjUser field value
func (o *WebhookUserUserCreated) GetObjUser() UserResponse {
	if o == nil  {
		var ret UserResponse
		return ret
	}

	return o.ObjUser
}

// GetObjUserOk returns a tuple with the ObjUser field value
// and a boolean to check if the value has been set.
func (o *WebhookUserUserCreated) GetObjUserOk() (*UserResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjUser, true
}

// SetObjUser sets field value
func (o *WebhookUserUserCreated) SetObjUser(v UserResponse) {
	o.ObjUser = v
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookUserUserCreated) GetObjWebhook() WebhookResponse {
	if o == nil  {
		var ret WebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookUserUserCreated) GetObjWebhookOk() (*WebhookResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookUserUserCreated) SetObjWebhook(v WebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookUserUserCreated) GetAObjAttempt() []AttemptResponse {
	if o == nil  {
		var ret []AttemptResponse
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookUserUserCreated) GetAObjAttemptOk() (*[]AttemptResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookUserUserCreated) SetAObjAttempt(v []AttemptResponse) {
	o.AObjAttempt = v
}

func (o WebhookUserUserCreated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objUser"] = o.ObjUser
	}
	if true {
		toSerialize["objWebhook"] = o.ObjWebhook
	}
	if true {
		toSerialize["a_objAttempt"] = o.AObjAttempt
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookUserUserCreated struct {
	value *WebhookUserUserCreated
	isSet bool
}

func (v NullableWebhookUserUserCreated) Get() *WebhookUserUserCreated {
	return v.value
}

func (v *NullableWebhookUserUserCreated) Set(val *WebhookUserUserCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookUserUserCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookUserUserCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookUserUserCreated(val *WebhookUserUserCreated) *NullableWebhookUserUserCreated {
	return &NullableWebhookUserUserCreated{value: val, isSet: true}
}

func (v NullableWebhookUserUserCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookUserUserCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


