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

// checks if the WebhookGetListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookGetListV1ResponseMPayload{}

// WebhookGetListV1ResponseMPayload Payload for GET /1/object/webhook/getList
type WebhookGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
	AObjWebhook []WebhookListElement `json:"a_objWebhook"`
}

// NewWebhookGetListV1ResponseMPayload instantiates a new WebhookGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32, aObjWebhook []WebhookListElement) *WebhookGetListV1ResponseMPayload {
	this := WebhookGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	this.AObjWebhook = aObjWebhook
	return &this
}

// NewWebhookGetListV1ResponseMPayloadWithDefaults instantiates a new WebhookGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookGetListV1ResponseMPayloadWithDefaults() *WebhookGetListV1ResponseMPayload {
	this := WebhookGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *WebhookGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *WebhookGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *WebhookGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *WebhookGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

// GetAObjWebhook returns the AObjWebhook field value
func (o *WebhookGetListV1ResponseMPayload) GetAObjWebhook() []WebhookListElement {
	if o == nil {
		var ret []WebhookListElement
		return ret
	}

	return o.AObjWebhook
}

// GetAObjWebhookOk returns a tuple with the AObjWebhook field value
// and a boolean to check if the value has been set.
func (o *WebhookGetListV1ResponseMPayload) GetAObjWebhookOk() ([]WebhookListElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjWebhook, true
}

// SetAObjWebhook sets field value
func (o *WebhookGetListV1ResponseMPayload) SetAObjWebhook(v []WebhookListElement) {
	o.AObjWebhook = v
}

func (o WebhookGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookGetListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iRowReturned"] = o.IRowReturned
	toSerialize["iRowFiltered"] = o.IRowFiltered
	toSerialize["a_objWebhook"] = o.AObjWebhook
	return toSerialize, nil
}

type NullableWebhookGetListV1ResponseMPayload struct {
	value *WebhookGetListV1ResponseMPayload
	isSet bool
}

func (v NullableWebhookGetListV1ResponseMPayload) Get() *WebhookGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableWebhookGetListV1ResponseMPayload) Set(val *WebhookGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookGetListV1ResponseMPayload(val *WebhookGetListV1ResponseMPayload) *NullableWebhookGetListV1ResponseMPayload {
	return &NullableWebhookGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableWebhookGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

