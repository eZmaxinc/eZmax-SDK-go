/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the WebhookRegenerateApikeyV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookRegenerateApikeyV1Request{}

// WebhookRegenerateApikeyV1Request Request for POST /1/object/webhook/{pkiWebhookID}/regenerateApikey
type WebhookRegenerateApikeyV1Request struct {
	// Whether the requests will be signed or not
	BWebhookIssigned *bool `json:"bWebhookIssigned,omitempty"`
}

// NewWebhookRegenerateApikeyV1Request instantiates a new WebhookRegenerateApikeyV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRegenerateApikeyV1Request() *WebhookRegenerateApikeyV1Request {
	this := WebhookRegenerateApikeyV1Request{}
	return &this
}

// NewWebhookRegenerateApikeyV1RequestWithDefaults instantiates a new WebhookRegenerateApikeyV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRegenerateApikeyV1RequestWithDefaults() *WebhookRegenerateApikeyV1Request {
	this := WebhookRegenerateApikeyV1Request{}
	return &this
}

// GetBWebhookIssigned returns the BWebhookIssigned field value if set, zero value otherwise.
func (o *WebhookRegenerateApikeyV1Request) GetBWebhookIssigned() bool {
	if o == nil || IsNil(o.BWebhookIssigned) {
		var ret bool
		return ret
	}
	return *o.BWebhookIssigned
}

// GetBWebhookIssignedOk returns a tuple with the BWebhookIssigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRegenerateApikeyV1Request) GetBWebhookIssignedOk() (*bool, bool) {
	if o == nil || IsNil(o.BWebhookIssigned) {
		return nil, false
	}
	return o.BWebhookIssigned, true
}

// HasBWebhookIssigned returns a boolean if a field has been set.
func (o *WebhookRegenerateApikeyV1Request) HasBWebhookIssigned() bool {
	if o != nil && !IsNil(o.BWebhookIssigned) {
		return true
	}

	return false
}

// SetBWebhookIssigned gets a reference to the given bool and assigns it to the BWebhookIssigned field.
func (o *WebhookRegenerateApikeyV1Request) SetBWebhookIssigned(v bool) {
	o.BWebhookIssigned = &v
}

func (o WebhookRegenerateApikeyV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookRegenerateApikeyV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BWebhookIssigned) {
		toSerialize["bWebhookIssigned"] = o.BWebhookIssigned
	}
	return toSerialize, nil
}

type NullableWebhookRegenerateApikeyV1Request struct {
	value *WebhookRegenerateApikeyV1Request
	isSet bool
}

func (v NullableWebhookRegenerateApikeyV1Request) Get() *WebhookRegenerateApikeyV1Request {
	return v.value
}

func (v *NullableWebhookRegenerateApikeyV1Request) Set(val *WebhookRegenerateApikeyV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRegenerateApikeyV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRegenerateApikeyV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRegenerateApikeyV1Request(val *WebhookRegenerateApikeyV1Request) *NullableWebhookRegenerateApikeyV1Request {
	return &NullableWebhookRegenerateApikeyV1Request{value: val, isSet: true}
}

func (v NullableWebhookRegenerateApikeyV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRegenerateApikeyV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


