/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CustomAutocompleteElementResponse Generic Autocomplete Response
type CustomAutocompleteElementResponse struct {
	// The Category for the dropdown or an empty string if not categorized
	SCategory string `json:"sCategory"`
	// The Description of the element
	SLabel string `json:"sLabel"`
	// The Unique ID of the element
	MValue string `json:"mValue"`
}

// NewCustomAutocompleteElementResponse instantiates a new CustomAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomAutocompleteElementResponse(sCategory string, sLabel string, mValue string) *CustomAutocompleteElementResponse {
	this := CustomAutocompleteElementResponse{}
	this.SCategory = sCategory
	this.SLabel = sLabel
	this.MValue = mValue
	return &this
}

// NewCustomAutocompleteElementResponseWithDefaults instantiates a new CustomAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomAutocompleteElementResponseWithDefaults() *CustomAutocompleteElementResponse {
	this := CustomAutocompleteElementResponse{}
	return &this
}

// GetSCategory returns the SCategory field value
func (o *CustomAutocompleteElementResponse) GetSCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCategory
}

// GetSCategoryOk returns a tuple with the SCategory field value
// and a boolean to check if the value has been set.
func (o *CustomAutocompleteElementResponse) GetSCategoryOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SCategory, true
}

// SetSCategory sets field value
func (o *CustomAutocompleteElementResponse) SetSCategory(v string) {
	o.SCategory = v
}

// GetSLabel returns the SLabel field value
func (o *CustomAutocompleteElementResponse) GetSLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLabel
}

// GetSLabelOk returns a tuple with the SLabel field value
// and a boolean to check if the value has been set.
func (o *CustomAutocompleteElementResponse) GetSLabelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SLabel, true
}

// SetSLabel sets field value
func (o *CustomAutocompleteElementResponse) SetSLabel(v string) {
	o.SLabel = v
}

// GetMValue returns the MValue field value
func (o *CustomAutocompleteElementResponse) GetMValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MValue
}

// GetMValueOk returns a tuple with the MValue field value
// and a boolean to check if the value has been set.
func (o *CustomAutocompleteElementResponse) GetMValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MValue, true
}

// SetMValue sets field value
func (o *CustomAutocompleteElementResponse) SetMValue(v string) {
	o.MValue = v
}

func (o CustomAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sCategory"] = o.SCategory
	}
	if true {
		toSerialize["sLabel"] = o.SLabel
	}
	if true {
		toSerialize["mValue"] = o.MValue
	}
	return json.Marshal(toSerialize)
}

type NullableCustomAutocompleteElementResponse struct {
	value *CustomAutocompleteElementResponse
	isSet bool
}

func (v NullableCustomAutocompleteElementResponse) Get() *CustomAutocompleteElementResponse {
	return v.value
}

func (v *NullableCustomAutocompleteElementResponse) Set(val *CustomAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomAutocompleteElementResponse(val *CustomAutocompleteElementResponse) *NullableCustomAutocompleteElementResponse {
	return &NullableCustomAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableCustomAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


