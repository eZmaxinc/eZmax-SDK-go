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
)

// checks if the MultilingualBrandingDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualBrandingDescription{}

// MultilingualBrandingDescription Description of the Branding
type MultilingualBrandingDescription struct {
	// The description of the Branding in French
	SBrandingDescription1 *string `json:"sBrandingDescription1,omitempty"`
	// The description of the Branding in English
	SBrandingDescription2 *string `json:"sBrandingDescription2,omitempty"`
}

// NewMultilingualBrandingDescription instantiates a new MultilingualBrandingDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualBrandingDescription() *MultilingualBrandingDescription {
	this := MultilingualBrandingDescription{}
	return &this
}

// NewMultilingualBrandingDescriptionWithDefaults instantiates a new MultilingualBrandingDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualBrandingDescriptionWithDefaults() *MultilingualBrandingDescription {
	this := MultilingualBrandingDescription{}
	return &this
}

// GetSBrandingDescription1 returns the SBrandingDescription1 field value if set, zero value otherwise.
func (o *MultilingualBrandingDescription) GetSBrandingDescription1() string {
	if o == nil || IsNil(o.SBrandingDescription1) {
		var ret string
		return ret
	}
	return *o.SBrandingDescription1
}

// GetSBrandingDescription1Ok returns a tuple with the SBrandingDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualBrandingDescription) GetSBrandingDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.SBrandingDescription1) {
		return nil, false
	}
	return o.SBrandingDescription1, true
}

// HasSBrandingDescription1 returns a boolean if a field has been set.
func (o *MultilingualBrandingDescription) HasSBrandingDescription1() bool {
	if o != nil && !IsNil(o.SBrandingDescription1) {
		return true
	}

	return false
}

// SetSBrandingDescription1 gets a reference to the given string and assigns it to the SBrandingDescription1 field.
func (o *MultilingualBrandingDescription) SetSBrandingDescription1(v string) {
	o.SBrandingDescription1 = &v
}

// GetSBrandingDescription2 returns the SBrandingDescription2 field value if set, zero value otherwise.
func (o *MultilingualBrandingDescription) GetSBrandingDescription2() string {
	if o == nil || IsNil(o.SBrandingDescription2) {
		var ret string
		return ret
	}
	return *o.SBrandingDescription2
}

// GetSBrandingDescription2Ok returns a tuple with the SBrandingDescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualBrandingDescription) GetSBrandingDescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.SBrandingDescription2) {
		return nil, false
	}
	return o.SBrandingDescription2, true
}

// HasSBrandingDescription2 returns a boolean if a field has been set.
func (o *MultilingualBrandingDescription) HasSBrandingDescription2() bool {
	if o != nil && !IsNil(o.SBrandingDescription2) {
		return true
	}

	return false
}

// SetSBrandingDescription2 gets a reference to the given string and assigns it to the SBrandingDescription2 field.
func (o *MultilingualBrandingDescription) SetSBrandingDescription2(v string) {
	o.SBrandingDescription2 = &v
}

func (o MultilingualBrandingDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualBrandingDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SBrandingDescription1) {
		toSerialize["sBrandingDescription1"] = o.SBrandingDescription1
	}
	if !IsNil(o.SBrandingDescription2) {
		toSerialize["sBrandingDescription2"] = o.SBrandingDescription2
	}
	return toSerialize, nil
}

type NullableMultilingualBrandingDescription struct {
	value *MultilingualBrandingDescription
	isSet bool
}

func (v NullableMultilingualBrandingDescription) Get() *MultilingualBrandingDescription {
	return v.value
}

func (v *NullableMultilingualBrandingDescription) Set(val *MultilingualBrandingDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualBrandingDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualBrandingDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualBrandingDescription(val *MultilingualBrandingDescription) *NullableMultilingualBrandingDescription {
	return &NullableMultilingualBrandingDescription{value: val, isSet: true}
}

func (v NullableMultilingualBrandingDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualBrandingDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


