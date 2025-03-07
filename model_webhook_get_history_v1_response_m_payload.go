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

// checks if the WebhookGetHistoryV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookGetHistoryV1ResponseMPayload{}

// WebhookGetHistoryV1ResponseMPayload Payload for GET /1/object/webhook/{pkiWebhookID}/getHistory
type WebhookGetHistoryV1ResponseMPayload struct {
	AObjWebhooklog []CustomWebhooklogResponse `json:"a_objWebhooklog"`
}

type _WebhookGetHistoryV1ResponseMPayload WebhookGetHistoryV1ResponseMPayload

// NewWebhookGetHistoryV1ResponseMPayload instantiates a new WebhookGetHistoryV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookGetHistoryV1ResponseMPayload(aObjWebhooklog []CustomWebhooklogResponse) *WebhookGetHistoryV1ResponseMPayload {
	this := WebhookGetHistoryV1ResponseMPayload{}
	this.AObjWebhooklog = aObjWebhooklog
	return &this
}

// NewWebhookGetHistoryV1ResponseMPayloadWithDefaults instantiates a new WebhookGetHistoryV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookGetHistoryV1ResponseMPayloadWithDefaults() *WebhookGetHistoryV1ResponseMPayload {
	this := WebhookGetHistoryV1ResponseMPayload{}
	return &this
}

// GetAObjWebhooklog returns the AObjWebhooklog field value
func (o *WebhookGetHistoryV1ResponseMPayload) GetAObjWebhooklog() []CustomWebhooklogResponse {
	if o == nil {
		var ret []CustomWebhooklogResponse
		return ret
	}

	return o.AObjWebhooklog
}

// GetAObjWebhooklogOk returns a tuple with the AObjWebhooklog field value
// and a boolean to check if the value has been set.
func (o *WebhookGetHistoryV1ResponseMPayload) GetAObjWebhooklogOk() ([]CustomWebhooklogResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjWebhooklog, true
}

// SetAObjWebhooklog sets field value
func (o *WebhookGetHistoryV1ResponseMPayload) SetAObjWebhooklog(v []CustomWebhooklogResponse) {
	o.AObjWebhooklog = v
}

func (o WebhookGetHistoryV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookGetHistoryV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objWebhooklog"] = o.AObjWebhooklog
	return toSerialize, nil
}

func (o *WebhookGetHistoryV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objWebhooklog",
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

	varWebhookGetHistoryV1ResponseMPayload := _WebhookGetHistoryV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookGetHistoryV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = WebhookGetHistoryV1ResponseMPayload(varWebhookGetHistoryV1ResponseMPayload)

	return err
}

type NullableWebhookGetHistoryV1ResponseMPayload struct {
	value *WebhookGetHistoryV1ResponseMPayload
	isSet bool
}

func (v NullableWebhookGetHistoryV1ResponseMPayload) Get() *WebhookGetHistoryV1ResponseMPayload {
	return v.value
}

func (v *NullableWebhookGetHistoryV1ResponseMPayload) Set(val *WebhookGetHistoryV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookGetHistoryV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookGetHistoryV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookGetHistoryV1ResponseMPayload(val *WebhookGetHistoryV1ResponseMPayload) *NullableWebhookGetHistoryV1ResponseMPayload {
	return &NullableWebhookGetHistoryV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableWebhookGetHistoryV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookGetHistoryV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


