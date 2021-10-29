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

// WebhookEzsignFolderCompletedAllOf struct for WebhookEzsignFolderCompletedAllOf
type WebhookEzsignFolderCompletedAllOf struct {
	ObjEzsignfolder EzsignfolderResponse `json:"objEzsignfolder"`
}

// NewWebhookEzsignFolderCompletedAllOf instantiates a new WebhookEzsignFolderCompletedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignFolderCompletedAllOf(objEzsignfolder EzsignfolderResponse) *WebhookEzsignFolderCompletedAllOf {
	this := WebhookEzsignFolderCompletedAllOf{}
	this.ObjEzsignfolder = objEzsignfolder
	return &this
}

// NewWebhookEzsignFolderCompletedAllOfWithDefaults instantiates a new WebhookEzsignFolderCompletedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignFolderCompletedAllOfWithDefaults() *WebhookEzsignFolderCompletedAllOf {
	this := WebhookEzsignFolderCompletedAllOf{}
	return &this
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value
func (o *WebhookEzsignFolderCompletedAllOf) GetObjEzsignfolder() EzsignfolderResponse {
	if o == nil {
		var ret EzsignfolderResponse
		return ret
	}

	return o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderCompletedAllOf) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjEzsignfolder, true
}

// SetObjEzsignfolder sets field value
func (o *WebhookEzsignFolderCompletedAllOf) SetObjEzsignfolder(v EzsignfolderResponse) {
	o.ObjEzsignfolder = v
}

func (o WebhookEzsignFolderCompletedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	}
	return json.Marshal(toSerialize)
}

type NullableWebhookEzsignFolderCompletedAllOf struct {
	value *WebhookEzsignFolderCompletedAllOf
	isSet bool
}

func (v NullableWebhookEzsignFolderCompletedAllOf) Get() *WebhookEzsignFolderCompletedAllOf {
	return v.value
}

func (v *NullableWebhookEzsignFolderCompletedAllOf) Set(val *WebhookEzsignFolderCompletedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignFolderCompletedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignFolderCompletedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignFolderCompletedAllOf(val *WebhookEzsignFolderCompletedAllOf) *NullableWebhookEzsignFolderCompletedAllOf {
	return &NullableWebhookEzsignFolderCompletedAllOf{value: val, isSet: true}
}

func (v NullableWebhookEzsignFolderCompletedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignFolderCompletedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


