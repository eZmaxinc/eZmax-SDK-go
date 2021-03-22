/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.38
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// MultilingualApikeyDescription Description of the API Key  
type MultilingualApikeyDescription struct {
	// Value in French
	SApikeyDescription1 *string `json:"sApikeyDescription1,omitempty"`
	// Value in English
	SApikeyDescription2 *string `json:"sApikeyDescription2,omitempty"`
}

// NewMultilingualApikeyDescription instantiates a new MultilingualApikeyDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualApikeyDescription() *MultilingualApikeyDescription {
	this := MultilingualApikeyDescription{}
	return &this
}

// NewMultilingualApikeyDescriptionWithDefaults instantiates a new MultilingualApikeyDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualApikeyDescriptionWithDefaults() *MultilingualApikeyDescription {
	this := MultilingualApikeyDescription{}
	return &this
}

// GetSApikeyDescription1 returns the SApikeyDescription1 field value if set, zero value otherwise.
func (o *MultilingualApikeyDescription) GetSApikeyDescription1() string {
	if o == nil || o.SApikeyDescription1 == nil {
		var ret string
		return ret
	}
	return *o.SApikeyDescription1
}

// GetSApikeyDescription1Ok returns a tuple with the SApikeyDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualApikeyDescription) GetSApikeyDescription1Ok() (*string, bool) {
	if o == nil || o.SApikeyDescription1 == nil {
		return nil, false
	}
	return o.SApikeyDescription1, true
}

// HasSApikeyDescription1 returns a boolean if a field has been set.
func (o *MultilingualApikeyDescription) HasSApikeyDescription1() bool {
	if o != nil && o.SApikeyDescription1 != nil {
		return true
	}

	return false
}

// SetSApikeyDescription1 gets a reference to the given string and assigns it to the SApikeyDescription1 field.
func (o *MultilingualApikeyDescription) SetSApikeyDescription1(v string) {
	o.SApikeyDescription1 = &v
}

// GetSApikeyDescription2 returns the SApikeyDescription2 field value if set, zero value otherwise.
func (o *MultilingualApikeyDescription) GetSApikeyDescription2() string {
	if o == nil || o.SApikeyDescription2 == nil {
		var ret string
		return ret
	}
	return *o.SApikeyDescription2
}

// GetSApikeyDescription2Ok returns a tuple with the SApikeyDescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualApikeyDescription) GetSApikeyDescription2Ok() (*string, bool) {
	if o == nil || o.SApikeyDescription2 == nil {
		return nil, false
	}
	return o.SApikeyDescription2, true
}

// HasSApikeyDescription2 returns a boolean if a field has been set.
func (o *MultilingualApikeyDescription) HasSApikeyDescription2() bool {
	if o != nil && o.SApikeyDescription2 != nil {
		return true
	}

	return false
}

// SetSApikeyDescription2 gets a reference to the given string and assigns it to the SApikeyDescription2 field.
func (o *MultilingualApikeyDescription) SetSApikeyDescription2(v string) {
	o.SApikeyDescription2 = &v
}

func (o MultilingualApikeyDescription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SApikeyDescription1 != nil {
		toSerialize["sApikeyDescription1"] = o.SApikeyDescription1
	}
	if o.SApikeyDescription2 != nil {
		toSerialize["sApikeyDescription2"] = o.SApikeyDescription2
	}
	return json.Marshal(toSerialize)
}

type NullableMultilingualApikeyDescription struct {
	value *MultilingualApikeyDescription
	isSet bool
}

func (v NullableMultilingualApikeyDescription) Get() *MultilingualApikeyDescription {
	return v.value
}

func (v *NullableMultilingualApikeyDescription) Set(val *MultilingualApikeyDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualApikeyDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualApikeyDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualApikeyDescription(val *MultilingualApikeyDescription) *NullableMultilingualApikeyDescription {
	return &NullableMultilingualApikeyDescription{value: val, isSet: true}
}

func (v NullableMultilingualApikeyDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualApikeyDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


