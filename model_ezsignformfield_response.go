/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignformfieldResponse An Ezsignformfield Object
type EzsignformfieldResponse struct {
	// The Label for the Ezsignformfield
	SEzsignformfieldLabel string `json:"sEzsignformfieldLabel"`
	// The Value for the Ezsignformfield
	SEzsignformfieldValue string `json:"sEzsignformfieldValue"`
}

// NewEzsignformfieldResponse instantiates a new EzsignformfieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldResponse(sEzsignformfieldLabel string, sEzsignformfieldValue string) *EzsignformfieldResponse {
	this := EzsignformfieldResponse{}
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.SEzsignformfieldValue = sEzsignformfieldValue
	return &this
}

// NewEzsignformfieldResponseWithDefaults instantiates a new EzsignformfieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldResponseWithDefaults() *EzsignformfieldResponse {
	this := EzsignformfieldResponse{}
	return &this
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *EzsignformfieldResponse) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *EzsignformfieldResponse) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetSEzsignformfieldValue returns the SEzsignformfieldValue field value
func (o *EzsignformfieldResponse) GetSEzsignformfieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldValue
}

// GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldResponse) GetSEzsignformfieldValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignformfieldValue, true
}

// SetSEzsignformfieldValue sets field value
func (o *EzsignformfieldResponse) SetSEzsignformfieldValue(v string) {
	o.SEzsignformfieldValue = v
}

func (o EzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sEzsignformfieldLabel"] = o.SEzsignformfieldLabel
	}
	if true {
		toSerialize["sEzsignformfieldValue"] = o.SEzsignformfieldValue
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignformfieldResponse struct {
	value *EzsignformfieldResponse
	isSet bool
}

func (v NullableEzsignformfieldResponse) Get() *EzsignformfieldResponse {
	return v.value
}

func (v *NullableEzsignformfieldResponse) Set(val *EzsignformfieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldResponse(val *EzsignformfieldResponse) *NullableEzsignformfieldResponse {
	return &NullableEzsignformfieldResponse{value: val, isSet: true}
}

func (v NullableEzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


