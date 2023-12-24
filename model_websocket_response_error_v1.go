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
	"bytes"
	"fmt"
)

// checks if the WebsocketResponseErrorV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketResponseErrorV1{}

// WebsocketResponseErrorV1 Response for Websocket Error V1
type WebsocketResponseErrorV1 struct {
	// The Type of message
	EWebsocketMessagetype string `json:"eWebsocketMessagetype"`
	// The Channel on which to route the websocket message
	SWebsocketChannel string `json:"sWebsocketChannel"`
	MPayload WebsocketResponseErrorV1MPayload `json:"mPayload"`
}

type _WebsocketResponseErrorV1 WebsocketResponseErrorV1

// NewWebsocketResponseErrorV1 instantiates a new WebsocketResponseErrorV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketResponseErrorV1(eWebsocketMessagetype string, sWebsocketChannel string, mPayload WebsocketResponseErrorV1MPayload) *WebsocketResponseErrorV1 {
	this := WebsocketResponseErrorV1{}
	this.EWebsocketMessagetype = eWebsocketMessagetype
	this.SWebsocketChannel = sWebsocketChannel
	this.MPayload = mPayload
	return &this
}

// NewWebsocketResponseErrorV1WithDefaults instantiates a new WebsocketResponseErrorV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketResponseErrorV1WithDefaults() *WebsocketResponseErrorV1 {
	this := WebsocketResponseErrorV1{}
	return &this
}

// GetEWebsocketMessagetype returns the EWebsocketMessagetype field value
func (o *WebsocketResponseErrorV1) GetEWebsocketMessagetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EWebsocketMessagetype
}

// GetEWebsocketMessagetypeOk returns a tuple with the EWebsocketMessagetype field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseErrorV1) GetEWebsocketMessagetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebsocketMessagetype, true
}

// SetEWebsocketMessagetype sets field value
func (o *WebsocketResponseErrorV1) SetEWebsocketMessagetype(v string) {
	o.EWebsocketMessagetype = v
}

// GetSWebsocketChannel returns the SWebsocketChannel field value
func (o *WebsocketResponseErrorV1) GetSWebsocketChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWebsocketChannel
}

// GetSWebsocketChannelOk returns a tuple with the SWebsocketChannel field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseErrorV1) GetSWebsocketChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWebsocketChannel, true
}

// SetSWebsocketChannel sets field value
func (o *WebsocketResponseErrorV1) SetSWebsocketChannel(v string) {
	o.SWebsocketChannel = v
}

// GetMPayload returns the MPayload field value
func (o *WebsocketResponseErrorV1) GetMPayload() WebsocketResponseErrorV1MPayload {
	if o == nil {
		var ret WebsocketResponseErrorV1MPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *WebsocketResponseErrorV1) GetMPayloadOk() (*WebsocketResponseErrorV1MPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *WebsocketResponseErrorV1) SetMPayload(v WebsocketResponseErrorV1MPayload) {
	o.MPayload = v
}

func (o WebsocketResponseErrorV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketResponseErrorV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eWebsocketMessagetype"] = o.EWebsocketMessagetype
	toSerialize["sWebsocketChannel"] = o.SWebsocketChannel
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *WebsocketResponseErrorV1) UnmarshalJSON(data []byte) (err error) {
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

	varWebsocketResponseErrorV1 := _WebsocketResponseErrorV1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebsocketResponseErrorV1)

	if err != nil {
		return err
	}

	*o = WebsocketResponseErrorV1(varWebsocketResponseErrorV1)

	return err
}

type NullableWebsocketResponseErrorV1 struct {
	value *WebsocketResponseErrorV1
	isSet bool
}

func (v NullableWebsocketResponseErrorV1) Get() *WebsocketResponseErrorV1 {
	return v.value
}

func (v *NullableWebsocketResponseErrorV1) Set(val *WebsocketResponseErrorV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketResponseErrorV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketResponseErrorV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketResponseErrorV1(val *WebsocketResponseErrorV1) *NullableWebsocketResponseErrorV1 {
	return &NullableWebsocketResponseErrorV1{value: val, isSet: true}
}

func (v NullableWebsocketResponseErrorV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketResponseErrorV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


