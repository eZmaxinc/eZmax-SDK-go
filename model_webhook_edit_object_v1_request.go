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

// checks if the WebhookEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookEditObjectV1Request{}

// WebhookEditObjectV1Request Request for PUT /1/object/webhook/{pkiWebhookID}
type WebhookEditObjectV1Request struct {
	ObjWebhook WebhookRequestCompound `json:"objWebhook"`
}

type _WebhookEditObjectV1Request WebhookEditObjectV1Request

// NewWebhookEditObjectV1Request instantiates a new WebhookEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookEditObjectV1Request(objWebhook WebhookRequestCompound) *WebhookEditObjectV1Request {
	this := WebhookEditObjectV1Request{}
	this.ObjWebhook = objWebhook
	return &this
}

// NewWebhookEditObjectV1RequestWithDefaults instantiates a new WebhookEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookEditObjectV1RequestWithDefaults() *WebhookEditObjectV1Request {
	this := WebhookEditObjectV1Request{}
	return &this
}

// GetObjWebhook returns the ObjWebhook field value
func (o *WebhookEditObjectV1Request) GetObjWebhook() WebhookRequestCompound {
	if o == nil {
		var ret WebhookRequestCompound
		return ret
	}

	return o.ObjWebhook
}

// GetObjWebhookOk returns a tuple with the ObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookEditObjectV1Request) GetObjWebhookOk() (*WebhookRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjWebhook, true
}

// SetObjWebhook sets field value
func (o *WebhookEditObjectV1Request) SetObjWebhook(v WebhookRequestCompound) {
	o.ObjWebhook = v
}

func (o WebhookEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objWebhook"] = o.ObjWebhook
	return toSerialize, nil
}

func (o *WebhookEditObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objWebhook",
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

	varWebhookEditObjectV1Request := _WebhookEditObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookEditObjectV1Request)

	if err != nil {
		return err
	}

	*o = WebhookEditObjectV1Request(varWebhookEditObjectV1Request)

	return err
}

type NullableWebhookEditObjectV1Request struct {
	value *WebhookEditObjectV1Request
	isSet bool
}

func (v NullableWebhookEditObjectV1Request) Get() *WebhookEditObjectV1Request {
	return v.value
}

func (v *NullableWebhookEditObjectV1Request) Set(val *WebhookEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookEditObjectV1Request(val *WebhookEditObjectV1Request) *NullableWebhookEditObjectV1Request {
	return &NullableWebhookEditObjectV1Request{value: val, isSet: true}
}

func (v NullableWebhookEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


