/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.31
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// WebhookEzsignFolderCompleted This is the base Webhook object
type WebhookEzsignFolderCompleted struct {
	ObjEzsignfolder EzsignfolderResponse `json:"objEzsignfolder"`
	ObjWebhook WebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponse `json:"a_objAttempt"`
}

// NewWebhookEzsignFolderCompleted instantiates a new WebhookEzsignFolderCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignFolderCompleted(objEzsignfolder EzsignfolderResponse, objWebhook WebhookResponse, aObjAttempt []AttemptResponse) *WebhookEzsignFolderCompleted {
	this := WebhookEzsignFolderCompleted{}
	this.ObjEzsignfolder = objEzsignfolder
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	return &this
}

// NewWebhookEzsignFolderCompletedWithDefaults instantiates a new WebhookEzsignFolderCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignFolderCompletedWithDefaults() *WebhookEzsignFolderCompleted {
	this := WebhookEzsignFolderCompleted{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value
func (o *WebhookEzsignFolderCompleted) GetObjEzsignfolder() EzsignfolderResponse {
	if o == nil {
		var ret EzsignfolderResponse
		return ret
	}

	return o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderCompleted) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjEzsignfolder, true
}

// SetObjEzsignfolder sets field value
func (o *WebhookEzsignFolderCompleted) SetObjEzsignfolder(v EzsignfolderResponse) {
	o.ObjEzsignfolder = v
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignFolderCompleted) GetObjWebhook() WebhookResponse {
	if o == nil {
		var ret WebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderCompleted) GetObjWebhookOk() (*WebhookResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignFolderCompleted) SetObjWebhook(v WebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignFolderCompleted) GetAObjAttempt() []AttemptResponse {
	if o == nil {
		var ret []AttemptResponse
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderCompleted) GetAObjAttemptOk() (*[]AttemptResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignFolderCompleted) SetAObjAttempt(v []AttemptResponse) {
	o.AObjAttempt = v
}

func (o WebhookEzsignFolderCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	if true {
		toSerialize["objWebhook"] = o.ObjWebhook
	}
	if true {
		toSerialize["a_objAttempt"] = o.AObjAttempt
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookEzsignFolderCompleted struct {
	value *WebhookEzsignFolderCompleted
	isSet bool
}

func (v NullableWebhookEzsignFolderCompleted) Get() *WebhookEzsignFolderCompleted {
	return v.value
}

func (v *NullableWebhookEzsignFolderCompleted) Set(val *WebhookEzsignFolderCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignFolderCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignFolderCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignFolderCompleted(val *WebhookEzsignFolderCompleted) *NullableWebhookEzsignFolderCompleted {
	return &NullableWebhookEzsignFolderCompleted{value: val, isSet: true}
}

func (v NullableWebhookEzsignFolderCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignFolderCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


