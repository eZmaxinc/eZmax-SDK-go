/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.35
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// CommonWebhook This is the base Webhook object
type CommonWebhook struct {
	ObjWebhook WebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponse `json:"a_objAttempt"`
}

// NewCommonWebhook instantiates a new CommonWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonWebhook(objWebhook WebhookResponse, aObjAttempt []AttemptResponse) *CommonWebhook {
	this := CommonWebhook{}
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	return &this
}

// NewCommonWebhookWithDefaults instantiates a new CommonWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonWebhookWithDefaults() *CommonWebhook {
	this := CommonWebhook{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *CommonWebhook) GetObjWebhook() WebhookResponse {
	if o == nil {
		var ret WebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *CommonWebhook) GetObjWebhookOk() (*WebhookResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *CommonWebhook) SetObjWebhook(v WebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *CommonWebhook) GetAObjAttempt() []AttemptResponse {
	if o == nil {
		var ret []AttemptResponse
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *CommonWebhook) GetAObjAttemptOk() (*[]AttemptResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *CommonWebhook) SetAObjAttempt(v []AttemptResponse) {
	o.AObjAttempt = v
}

func (o CommonWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objWebhook"] = o.ObjWebhook
	}
	if true {
		toSerialize["a_objAttempt"] = o.AObjAttempt
	}
	return json.Marshal(toSerialize)
}

type NullableCommonWebhook struct {
	value *CommonWebhook
	isSet bool
}

func (v NullableCommonWebhook) Get() *CommonWebhook {
	return v.value
}

func (v *NullableCommonWebhook) Set(val *CommonWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonWebhook(val *CommonWebhook) *NullableCommonWebhook {
	return &NullableCommonWebhook{value: val, isSet: true}
}

func (v NullableCommonWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


