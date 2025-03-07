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

// checks if the WebsocketResponseInformationV1MPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketResponseInformationV1MPayload{}

// WebsocketResponseInformationV1MPayload Payload for Websocket Information V1
type WebsocketResponseInformationV1MPayload struct {
	// Information message
	SInformationMessage string `json:"sInformationMessage"`
}

type _WebsocketResponseInformationV1MPayload WebsocketResponseInformationV1MPayload

// NewWebsocketResponseInformationV1MPayload instantiates a new WebsocketResponseInformationV1MPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketResponseInformationV1MPayload(sInformationMessage string) *WebsocketResponseInformationV1MPayload {
	this := WebsocketResponseInformationV1MPayload{}
	this.SInformationMessage = sInformationMessage
	return &this
}

// NewWebsocketResponseInformationV1MPayloadWithDefaults instantiates a new WebsocketResponseInformationV1MPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketResponseInformationV1MPayloadWithDefaults() *WebsocketResponseInformationV1MPayload {
	this := WebsocketResponseInformationV1MPayload{}
	return &this
}

// GetSInformationMessage returns the SInformationMessage field value
func (o *WebsocketResponseInformationV1MPayload) GetSInformationMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SInformationMessage
}

// GetSInformationMessageOk returns a tuple with the SInformationMessage field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseInformationV1MPayload) GetSInformationMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SInformationMessage, true
}

// SetSInformationMessage sets field value
func (o *WebsocketResponseInformationV1MPayload) SetSInformationMessage(v string) {
	o.SInformationMessage = v
}

func (o WebsocketResponseInformationV1MPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketResponseInformationV1MPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sInformationMessage"] = o.SInformationMessage
	return toSerialize, nil
}

func (o *WebsocketResponseInformationV1MPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sInformationMessage",
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

	varWebsocketResponseInformationV1MPayload := _WebsocketResponseInformationV1MPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebsocketResponseInformationV1MPayload)

	if err != nil {
		return err
	}

	*o = WebsocketResponseInformationV1MPayload(varWebsocketResponseInformationV1MPayload)

	return err
}

type NullableWebsocketResponseInformationV1MPayload struct {
	value *WebsocketResponseInformationV1MPayload
	isSet bool
}

func (v NullableWebsocketResponseInformationV1MPayload) Get() *WebsocketResponseInformationV1MPayload {
	return v.value
}

func (v *NullableWebsocketResponseInformationV1MPayload) Set(val *WebsocketResponseInformationV1MPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketResponseInformationV1MPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketResponseInformationV1MPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketResponseInformationV1MPayload(val *WebsocketResponseInformationV1MPayload) *NullableWebsocketResponseInformationV1MPayload {
	return &NullableWebsocketResponseInformationV1MPayload{value: val, isSet: true}
}

func (v NullableWebsocketResponseInformationV1MPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketResponseInformationV1MPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


