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

// checks if the WebsocketResponseGetWebsocketIDV1MPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketResponseGetWebsocketIDV1MPayload{}

// WebsocketResponseGetWebsocketIDV1MPayload Payload for Websocket GetWebsocketID V1
type WebsocketResponseGetWebsocketIDV1MPayload struct {
	// The Unique ID of the Websocket Connection
	SWebsocketID string `json:"sWebsocketID" validate:"regexp=^[a-zA-Z0-9_-]{15}=$"`
}

type _WebsocketResponseGetWebsocketIDV1MPayload WebsocketResponseGetWebsocketIDV1MPayload

// NewWebsocketResponseGetWebsocketIDV1MPayload instantiates a new WebsocketResponseGetWebsocketIDV1MPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketResponseGetWebsocketIDV1MPayload(sWebsocketID string) *WebsocketResponseGetWebsocketIDV1MPayload {
	this := WebsocketResponseGetWebsocketIDV1MPayload{}
	this.SWebsocketID = sWebsocketID
	return &this
}

// NewWebsocketResponseGetWebsocketIDV1MPayloadWithDefaults instantiates a new WebsocketResponseGetWebsocketIDV1MPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketResponseGetWebsocketIDV1MPayloadWithDefaults() *WebsocketResponseGetWebsocketIDV1MPayload {
	this := WebsocketResponseGetWebsocketIDV1MPayload{}
	return &this
}

// GetSWebsocketID returns the SWebsocketID field value
func (o *WebsocketResponseGetWebsocketIDV1MPayload) GetSWebsocketID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebsocketID
}

// GetSWebsocketIDOk returns a tuple with the SWebsocketID field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseGetWebsocketIDV1MPayload) GetSWebsocketIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebsocketID, true
}

// SetSWebsocketID sets field value
func (o *WebsocketResponseGetWebsocketIDV1MPayload) SetSWebsocketID(v string) {
	o.SWebsocketID = v
}

func (o WebsocketResponseGetWebsocketIDV1MPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketResponseGetWebsocketIDV1MPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sWebsocketID"] = o.SWebsocketID
	return toSerialize, nil
}

func (o *WebsocketResponseGetWebsocketIDV1MPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sWebsocketID",
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

	varWebsocketResponseGetWebsocketIDV1MPayload := _WebsocketResponseGetWebsocketIDV1MPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebsocketResponseGetWebsocketIDV1MPayload)

	if err != nil {
		return err
	}

	*o = WebsocketResponseGetWebsocketIDV1MPayload(varWebsocketResponseGetWebsocketIDV1MPayload)

	return err
}

type NullableWebsocketResponseGetWebsocketIDV1MPayload struct {
	value *WebsocketResponseGetWebsocketIDV1MPayload
	isSet bool
}

func (v NullableWebsocketResponseGetWebsocketIDV1MPayload) Get() *WebsocketResponseGetWebsocketIDV1MPayload {
	return v.value
}

func (v *NullableWebsocketResponseGetWebsocketIDV1MPayload) Set(val *WebsocketResponseGetWebsocketIDV1MPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketResponseGetWebsocketIDV1MPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketResponseGetWebsocketIDV1MPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketResponseGetWebsocketIDV1MPayload(val *WebsocketResponseGetWebsocketIDV1MPayload) *NullableWebsocketResponseGetWebsocketIDV1MPayload {
	return &NullableWebsocketResponseGetWebsocketIDV1MPayload{value: val, isSet: true}
}

func (v NullableWebsocketResponseGetWebsocketIDV1MPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketResponseGetWebsocketIDV1MPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


