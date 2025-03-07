/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WebhookEzsignSignatureSigned type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEzsignSignatureSigned{}

// WebhookEzsignSignatureSigned This is the base Webhook object
type WebhookEzsignSignatureSigned struct {
	ObjWebhook CustomWebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponseCompound `json:"a_objAttempt"`
	ObjEzsignsignature EzsignsignatureResponse `json:"objEzsignsignature"`
}

type _WebhookEzsignSignatureSigned WebhookEzsignSignatureSigned

// NewWebhookEzsignSignatureSigned instantiates a new WebhookEzsignSignatureSigned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEzsignSignatureSigned(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound, objEzsignsignature EzsignsignatureResponse) *WebhookEzsignSignatureSigned {
	this := WebhookEzsignSignatureSigned{}
	this.ObjWebhook = objWebhook
	this.AObjAttempt = aObjAttempt
	this.ObjEzsignsignature = objEzsignsignature
	return &this
}

// NewWebhookEzsignSignatureSignedWithDefaults instantiates a new WebhookEzsignSignatureSigned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEzsignSignatureSignedWithDefaults() *WebhookEzsignSignatureSigned {
	this := WebhookEzsignSignatureSigned{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEzsignSignatureSigned) GetObjWebhook() CustomWebhookResponse {
	if o == nil {
		var ret CustomWebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignSignatureSigned) GetObjWebhookOk() (*CustomWebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEzsignSignatureSigned) SetObjWebhook(v CustomWebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *WebhookEzsignSignatureSigned) GetAObjAttempt() []AttemptResponseCompound {
	if o == nil {
		var ret []AttemptResponseCompound
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignSignatureSigned) GetAObjAttemptOk() ([]AttemptResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *WebhookEzsignSignatureSigned) SetAObjAttempt(v []AttemptResponseCompound) {
	o.AObjAttempt = v
}

// GetObjEzsignsignature returns the ObjEzsignsignature field value
func (o *WebhookEzsignSignatureSigned) GetObjEzsignsignature() EzsignsignatureResponse {
	if o == nil {
		var ret EzsignsignatureResponse
		return ret
	}

	return o.ObjEzsignsignature
}

// GetObjEzsignsignatureOk returns a tuple with the ObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *WebhookEzsignSignatureSigned) GetObjEzsignsignatureOk() (*EzsignsignatureResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignsignature, true
}

// SetObjEzsignsignature sets field value
func (o *WebhookEzsignSignatureSigned) SetObjEzsignsignature(v EzsignsignatureResponse) {
	o.ObjEzsignsignature = v
}

func (o WebhookEzsignSignatureSigned) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEzsignSignatureSigned) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	toSerialize["a_objAttempt"] = o.AObjAttempt
	toSerialize["objEzsignsignature"] = o.ObjEzsignsignature
	return toSerialize, nil
}

func (o *WebhookEzsignSignatureSigned) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objWebhook",
		"a_objAttempt",
		"objEzsignsignature",
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

	varWebhookEzsignSignatureSigned := _WebhookEzsignSignatureSigned{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEzsignSignatureSigned)

	if err != nil {
		return err
	}

	*o = WebhookEzsignSignatureSigned(varWebhookEzsignSignatureSigned)

	return err
}

type NullableWebhookEzsignSignatureSigned struct {
	value *WebhookEzsignSignatureSigned
	isSet bool
}

func (v NullableWebhookEzsignSignatureSigned) Get() *WebhookEzsignSignatureSigned {
	return v.value
}

func (v *NullableWebhookEzsignSignatureSigned) Set(val *WebhookEzsignSignatureSigned) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEzsignSignatureSigned) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEzsignSignatureSigned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEzsignSignatureSigned(val *WebhookEzsignSignatureSigned) *NullableWebhookEzsignSignatureSigned {
	return &NullableWebhookEzsignSignatureSigned{value: val, isSet: true}
}

func (v NullableWebhookEzsignSignatureSigned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEzsignSignatureSigned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


