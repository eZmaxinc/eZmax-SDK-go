/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookEzsignFolderSent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEzsignFolderSent{}

// WebhookEzsignFolderSent This is the base Webhook object
type WebhookEzsignFolderSent struct {
	ObjWebhook CustomWebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponseCompound `json:"a_objAttempt"`
	ObjEzsignfolder EzsignfolderResponse `json:"objEzsignfolder"`
}

type _WebhookEzsignFolderSent WebhookEzsignFolderSent

// NewWebhookEzsignFolderSent instantiates a new WebhookEzsignFolderSent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignFolderSent(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignfolder EzsignfolderResponse) *WebhookEzsignFolderSent {
	this := WebhookEzsignFolderSent{}
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	this.ObjEzsignfolder = objEzsignfolder
	return &this
}

// NewWebhookEzsignFolderSentWithDefaults instantiates a new WebhookEzsignFolderSent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignFolderSentWithDefaults() *WebhookEzsignFolderSent {
	this := WebhookEzsignFolderSent{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignFolderSent) GetObjWebhook() CustomWebhookResponse {
	if o == nil {
		var ret CustomWebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderSent) GetObjWebhookOk() (*CustomWebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignFolderSent) SetObjWebhook(v CustomWebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignFolderSent) GetAObjAttempt() []AttemptResponseCompound {
	if o == nil {
		var ret []AttemptResponseCompound
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderSent) GetAObjAttemptOk() ([]AttemptResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignFolderSent) SetAObjAttempt(v []AttemptResponseCompound) {
	o.AObjAttempt = v
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value
func (o *WebhookEzsignFolderSent) GetObjEzsignfolder() EzsignfolderResponse {
	if o == nil {
		var ret EzsignfolderResponse
		return ret
	}

	return o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderSent) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfolder, true
}

// SetObjEzsignfolder sets field value
func (o *WebhookEzsignFolderSent) SetObjEzsignfolder(v EzsignfolderResponse) {
	o.ObjEzsignfolder = v
}

func (o WebhookEzsignFolderSent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEzsignFolderSent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	toSerialize["a_objAttempt"] = o.AObjAttempt
	toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	return toSerialize, nil
}

func (o *WebhookEzsignFolderSent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objWebhook",
		"a_objAttempt",
		"objEzsignfolder",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookEzsignFolderSent := _WebhookEzsignFolderSent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEzsignFolderSent)

	if err != nil {
		return err
	}

	*o = WebhookEzsignFolderSent(varWebhookEzsignFolderSent)

	return err
}

type NullableWebhookEzsignFolderSent struct {
	value *WebhookEzsignFolderSent
	isSet bool
}

func (v NullableWebhookEzsignFolderSent) Get() *WebhookEzsignFolderSent {
	return v.value
}

func (v *NullableWebhookEzsignFolderSent) Set(val *WebhookEzsignFolderSent) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignFolderSent) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignFolderSent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignFolderSent(val *WebhookEzsignFolderSent) *NullableWebhookEzsignFolderSent {
	return &NullableWebhookEzsignFolderSent{value: val, isSet: true}
}

func (v NullableWebhookEzsignFolderSent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignFolderSent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


