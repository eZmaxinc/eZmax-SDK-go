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

// checks if the WebsocketRequestServerGetWebsocketIDV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebsocketRequestServerGetWebsocketIDV1{}

// WebsocketRequestServerGetWebsocketIDV1 Request for Websocket GetWebsocketID V1
type WebsocketRequestServerGetWebsocketIDV1 struct {
	// The Type of message
	EWebsocketMessagetype string `json:"eWebsocketMessagetype"`
}

// NewWebsocketRequestServerGetWebsocketIDV1 instantiates a new WebsocketRequestServerGetWebsocketIDV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketRequestServerGetWebsocketIDV1(eWebsocketMessagetype string) *WebsocketRequestServerGetWebsocketIDV1 {
	this := WebsocketRequestServerGetWebsocketIDV1{}
	this.EWebsocketMessagetype = eWebsocketMessagetype
	return &this
}

// NewWebsocketRequestServerGetWebsocketIDV1WithDefaults instantiates a new WebsocketRequestServerGetWebsocketIDV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketRequestServerGetWebsocketIDV1WithDefaults() *WebsocketRequestServerGetWebsocketIDV1 {
	this := WebsocketRequestServerGetWebsocketIDV1{}
	return &this
}

// GetEWebsocketMessagetype returns the EWebsocketMessagetype field value
func (o *WebsocketRequestServerGetWebsocketIDV1) GetEWebsocketMessagetype() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EWebsocketMessagetype
}

// GetEWebsocketMessagetypeOk returns a tuple with the EWebsocketMessagetype field value
// and a boolean to check if the value has been set.
func (o *WebsocketRequestServerGetWebsocketIDV1) GetEWebsocketMessagetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWebsocketMessagetype, true
}

// SetEWebsocketMessagetype sets field value
func (o *WebsocketRequestServerGetWebsocketIDV1) SetEWebsocketMessagetype(v string) {
	o.EWebsocketMessagetype = v
}

func (o WebsocketRequestServerGetWebsocketIDV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebsocketRequestServerGetWebsocketIDV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eWebsocketMessagetype"] = o.EWebsocketMessagetype
	return toSerialize, nil
}

type NullableWebsocketRequestServerGetWebsocketIDV1 struct {
	value *WebsocketRequestServerGetWebsocketIDV1
	isSet bool
}

func (v NullableWebsocketRequestServerGetWebsocketIDV1) Get() *WebsocketRequestServerGetWebsocketIDV1 {
	return v.value
}

func (v *NullableWebsocketRequestServerGetWebsocketIDV1) Set(val *WebsocketRequestServerGetWebsocketIDV1) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketRequestServerGetWebsocketIDV1) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketRequestServerGetWebsocketIDV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketRequestServerGetWebsocketIDV1(val *WebsocketRequestServerGetWebsocketIDV1) *NullableWebsocketRequestServerGetWebsocketIDV1 {
	return &NullableWebsocketRequestServerGetWebsocketIDV1{value: val, isSet: true}
}

func (v NullableWebsocketRequestServerGetWebsocketIDV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketRequestServerGetWebsocketIDV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


