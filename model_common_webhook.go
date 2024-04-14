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

// checks if the CommonWebhook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonWebhook{}

// CommonWebhook This is the base Webhook object
type CommonWebhook struct {
	ObjWebhook CustomWebhookResponse `json:"objWebhook"`
	// An array containing details of previous attempts that were made to deliver the message. The array is empty if it's the first attempt.
	AObjAttempt []AttemptResponseCompound `json:"a_objAttempt"`
}

type _CommonWebhook CommonWebhook

// NewCommonWebhook instantiates a new CommonWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonWebhook(objWebhook CustomWebhookResponse, aObjAttempt []AttemptResponseCompound) *CommonWebhook {
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
func (o *CommonWebhook) GetObjWebhook() CustomWebhookResponse {
	if o == nil {
		var ret CustomWebhookResponse
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *CommonWebhook) GetObjWebhookOk() (*CustomWebhookResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *CommonWebhook) SetObjWebhook(v CustomWebhookResponse) {
	o.ObjWebhook = v
}

// GetAObjAttempt returns the AObjAttempt field value
func (o *CommonWebhook) GetAObjAttempt() []AttemptResponseCompound {
	if o == nil {
		var ret []AttemptResponseCompound
		return ret
	}

	return o.AObjAttempt
}

// GetAObjAttemptOk returns a tuple with the AObjAttempt field value
// and a boolean to check if the value has been set.
func (o *CommonWebhook) GetAObjAttemptOk() ([]AttemptResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttempt, true
}

// SetAObjAttempt sets field value
func (o *CommonWebhook) SetAObjAttempt(v []AttemptResponseCompound) {
	o.AObjAttempt = v
}

func (o CommonWebhook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonWebhook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	toSerialize["a_objAttempt"] = o.AObjAttempt
	return toSerialize, nil
}

func (o *CommonWebhook) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objWebhook",
		"a_objAttempt",
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

	varCommonWebhook := _CommonWebhook{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonWebhook)

	if err != nil {
		return err
	}

	*o = CommonWebhook(varCommonWebhook)

	return err
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


