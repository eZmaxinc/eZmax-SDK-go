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

// checks if the WebsocketResponseInformationV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketResponseInformationV1{}

// WebsocketResponseInformationV1 Response for Websocket Information V1
type WebsocketResponseInformationV1 struct {
	// The Type of message
	EWebsocketMessagetype string `json:"eWebsocketMessagetype"`
	// The Channel on which to route the websocket message
	SWebsocketChannel string `json:"sWebsocketChannel" validate:"regexp=^[a-zA-Z0-9_@.]{32}$"`
	MPayload WebsocketResponseInformationV1MPayload `json:"mPayload"`
}

type _WebsocketResponseInformationV1 WebsocketResponseInformationV1

// NewWebsocketResponseInformationV1 instantiates a new WebsocketResponseInformationV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketResponseInformationV1(eWebsocketMessagetype string, sWebsocketChannel string, mPayload WebsocketResponseInformationV1MPayload) *WebsocketResponseInformationV1 {
	this := WebsocketResponseInformationV1{}
	this.EWebsocketMessagetype = eWebsocketMessagetype
	this.SWebsocketChannel = sWebsocketChannel
	this.MPayload = mPayload
	return &this
}

// NewWebsocketResponseInformationV1WithDefaults instantiates a new WebsocketResponseInformationV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketResponseInformationV1WithDefaults() *WebsocketResponseInformationV1 {
	this := WebsocketResponseInformationV1{}
	return &this
}

// GetEWebsocketMessagetype returns the EWebsocketMessagetype field value
func (o *WebsocketResponseInformationV1) GetEWebsocketMessagetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EWebsocketMessagetype
}

// GetEWebsocketMessagetypeOk returns a tuple with the EWebsocketMessagetype field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseInformationV1) GetEWebsocketMessagetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebsocketMessagetype, true
}

// SetEWebsocketMessagetype sets field value
func (o *WebsocketResponseInformationV1) SetEWebsocketMessagetype(v string) {
	o.EWebsocketMessagetype = v
}

// GetSWebsocketChannel returns the SWebsocketChannel field value
func (o *WebsocketResponseInformationV1) GetSWebsocketChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebsocketChannel
}

// GetSWebsocketChannelOk returns a tuple with the SWebsocketChannel field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseInformationV1) GetSWebsocketChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebsocketChannel, true
}

// SetSWebsocketChannel sets field value
func (o *WebsocketResponseInformationV1) SetSWebsocketChannel(v string) {
	o.SWebsocketChannel = v
}

// GetMPayload returns the MPayload field value
func (o *WebsocketResponseInformationV1) GetMPayload() WebsocketResponseInformationV1MPayload {
	if o == nil {
		var ret WebsocketResponseInformationV1MPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseInformationV1) GetMPayloadOk() (*WebsocketResponseInformationV1MPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *WebsocketResponseInformationV1) SetMPayload(v WebsocketResponseInformationV1MPayload) {
	o.MPayload = v
}

func (o WebsocketResponseInformationV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketResponseInformationV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eWebsocketMessagetype"] = o.EWebsocketMessagetype
	toSerialize["sWebsocketChannel"] = o.SWebsocketChannel
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *WebsocketResponseInformationV1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eWebsocketMessagetype",
		"sWebsocketChannel",
		"mPayload",
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

	varWebsocketResponseInformationV1 := _WebsocketResponseInformationV1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebsocketResponseInformationV1)

	if err != nil {
		return err
	}

	*o = WebsocketResponseInformationV1(varWebsocketResponseInformationV1)

	return err
}

type NullableWebsocketResponseInformationV1 struct {
	value *WebsocketResponseInformationV1
	isSet bool
}

func (v NullableWebsocketResponseInformationV1) Get() *WebsocketResponseInformationV1 {
	return v.value
}

func (v *NullableWebsocketResponseInformationV1) Set(val *WebsocketResponseInformationV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketResponseInformationV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketResponseInformationV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketResponseInformationV1(val *WebsocketResponseInformationV1) *NullableWebsocketResponseInformationV1 {
	return &NullableWebsocketResponseInformationV1{value: val, isSet: true}
}

func (v NullableWebsocketResponseInformationV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketResponseInformationV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


