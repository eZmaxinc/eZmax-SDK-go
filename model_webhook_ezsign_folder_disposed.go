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

// checks if the WebhookEzsignFolderDisposed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEzsignFolderDisposed{}

// WebhookEzsignFolderDisposed This is the base Webhook object
type WebhookEzsignFolderDisposed struct {
	ObjWebhook CustomWebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponseCompound `json:"a_objAttempt"`
	ObjEzsignfolder EzsignfolderResponse `json:"objEzsignfolder"`
}

type _WebhookEzsignFolderDisposed WebhookEzsignFolderDisposed

// NewWebhookEzsignFolderDisposed instantiates a new WebhookEzsignFolderDisposed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignFolderDisposed(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignfolder EzsignfolderResponse) *WebhookEzsignFolderDisposed {
	this := WebhookEzsignFolderDisposed{}
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	this.ObjEzsignfolder = objEzsignfolder
	return &this
}

// NewWebhookEzsignFolderDisposedWithDefaults instantiates a new WebhookEzsignFolderDisposed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignFolderDisposedWithDefaults() *WebhookEzsignFolderDisposed {
	this := WebhookEzsignFolderDisposed{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignFolderDisposed) GetObjWebhook() CustomWebhookResponse {
	if o == nil {
		var ret CustomWebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderDisposed) GetObjWebhookOk() (*CustomWebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignFolderDisposed) SetObjWebhook(v CustomWebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignFolderDisposed) GetAObjAttempt() []AttemptResponseCompound {
	if o == nil {
		var ret []AttemptResponseCompound
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderDisposed) GetAObjAttemptOk() ([]AttemptResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignFolderDisposed) SetAObjAttempt(v []AttemptResponseCompound) {
	o.AObjAttempt = v
}

// GetObjEzsignfolder returns the ObjEzsignfolder field value
func (o *WebhookEzsignFolderDisposed) GetObjEzsignfolder() EzsignfolderResponse {
	if o == nil {
		var ret EzsignfolderResponse
		return ret
	}

	return o.ObjEzsignfolder
}

// GetObjEzsignfolderOk returns a tuple with the ObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignFolderDisposed) GetObjEzsignfolderOk() (*EzsignfolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignfolder, true
}

// SetObjEzsignfolder sets field value
func (o *WebhookEzsignFolderDisposed) SetObjEzsignfolder(v EzsignfolderResponse) {
	o.ObjEzsignfolder = v
}

func (o WebhookEzsignFolderDisposed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEzsignFolderDisposed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	toSerialize["a_objAttempt"] = o.AObjAttempt
	toSerialize["objEzsignfolder"] = o.ObjEzsignfolder
	return toSerialize, nil
}

func (o *WebhookEzsignFolderDisposed) UnmarshalJSON(data []byte) (err error) {
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

	varWebhookEzsignFolderDisposed := _WebhookEzsignFolderDisposed{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEzsignFolderDisposed)

	if err != nil {
		return err
	}

	*o = WebhookEzsignFolderDisposed(varWebhookEzsignFolderDisposed)

	return err
}

type NullableWebhookEzsignFolderDisposed struct {
	value *WebhookEzsignFolderDisposed
	isSet bool
}

func (v NullableWebhookEzsignFolderDisposed) Get() *WebhookEzsignFolderDisposed {
	return v.value
}

func (v *NullableWebhookEzsignFolderDisposed) Set(val *WebhookEzsignFolderDisposed) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignFolderDisposed) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignFolderDisposed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignFolderDisposed(val *WebhookEzsignFolderDisposed) *NullableWebhookEzsignFolderDisposed {
	return &NullableWebhookEzsignFolderDisposed{value: val, isSet: true}
}

func (v NullableWebhookEzsignFolderDisposed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignFolderDisposed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


