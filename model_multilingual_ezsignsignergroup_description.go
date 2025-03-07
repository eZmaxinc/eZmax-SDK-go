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

// checks if the MultilingualEzsignsignergroupDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualEzsignsignergroupDescription{}

// MultilingualEzsignsignergroupDescription Description of the Ezsignsignergroup
type MultilingualEzsignsignergroupDescription struct {
	// The description of the Ezsignsignergroup in French
	SEzsignsignergroupDescription1 *string `json:"sEzsignsignergroupDescription1,omitempty"`
	// The description of the Ezsignsignergroup in English
	SEzsignsignergroupDescription2 *string `json:"sEzsignsignergroupDescription2,omitempty"`
}

// NewMultilingualEzsignsignergroupDescription instantiates a new MultilingualEzsignsignergroupDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualEzsignsignergroupDescription() *MultilingualEzsignsignergroupDescription {
	this := MultilingualEzsignsignergroupDescription{}
	return &this
}

// NewMultilingualEzsignsignergroupDescriptionWithDefaults instantiates a new MultilingualEzsignsignergroupDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualEzsignsignergroupDescriptionWithDefaults() *MultilingualEzsignsignergroupDescription {
	this := MultilingualEzsignsignergroupDescription{}
	return &this
}

// GetSEzsignsignergroupDescription1 returns the SEzsignsignergroupDescription1 field value if set, zero value otherwise.
func (o *MultilingualEzsignsignergroupDescription) GetSEzsignsignergroupDescription1() string {
	if o == nil || IsNil(o.SEzsignsignergroupDescription1) {
		var ret string
		return ret
	}
	return *o.SEzsignsignergroupDescription1
}

// GetSEzsignsignergroupDescription1Ok returns a tuple with the SEzsignsignergroupDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualEzsignsignergroupDescription) GetSEzsignsignergroupDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.SEzsignsignergroupDescription1) {
		return nil, false
	}
	return o.SEzsignsignergroupDescription1, true
}

// HasSEzsignsignergroupDescription1 returns a boolean if a field has been set.
func (o *MultilingualEzsignsignergroupDescription) HasSEzsignsignergroupDescription1() bool {
	if o != nil && !IsNil(o.SEzsignsignergroupDescription1) {
		return true
	}

	return false
}

// SetSEzsignsignergroupDescription1 gets a reference to the given string and assigns it to the SEzsignsignergroupDescription1 field.
func (o *MultilingualEzsignsignergroupDescription) SetSEzsignsignergroupDescription1(v string) {
	o.SEzsignsignergroupDescription1 = &v
}

// GetSEzsignsignergroupDescription2 returns the SEzsignsignergroupDescription2 field value if set, zero value otherwise.
func (o *MultilingualEzsignsignergroupDescription) GetSEzsignsignergroupDescription2() string {
	if o == nil || IsNil(o.SEzsignsignergroupDescription2) {
		var ret string
		return ret
	}
	return *o.SEzsignsignergroupDescription2
}

// GetSEzsignsignergroupDescription2Ok returns a tuple with the SEzsignsignergroupDescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualEzsignsignergroupDescription) GetSEzsignsignergroupDescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.SEzsignsignergroupDescription2) {
		return nil, false
	}
	return o.SEzsignsignergroupDescription2, true
}

// HasSEzsignsignergroupDescription2 returns a boolean if a field has been set.
func (o *MultilingualEzsignsignergroupDescription) HasSEzsignsignergroupDescription2() bool {
	if o != nil && !IsNil(o.SEzsignsignergroupDescription2) {
		return true
	}

	return false
}

// SetSEzsignsignergroupDescription2 gets a reference to the given string and assigns it to the SEzsignsignergroupDescription2 field.
func (o *MultilingualEzsignsignergroupDescription) SetSEzsignsignergroupDescription2(v string) {
	o.SEzsignsignergroupDescription2 = &v
}

func (o MultilingualEzsignsignergroupDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualEzsignsignergroupDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SEzsignsignergroupDescription1) {
		toSerialize["sEzsignsignergroupDescription1"] = o.SEzsignsignergroupDescription1
	}
	if !IsNil(o.SEzsignsignergroupDescription2) {
		toSerialize["sEzsignsignergroupDescription2"] = o.SEzsignsignergroupDescription2
	}
	return toSerialize, nil
}

type NullableMultilingualEzsignsignergroupDescription struct {
	value *MultilingualEzsignsignergroupDescription
	isSet bool
}

func (v NullableMultilingualEzsignsignergroupDescription) Get() *MultilingualEzsignsignergroupDescription {
	return v.value
}

func (v *NullableMultilingualEzsignsignergroupDescription) Set(val *MultilingualEzsignsignergroupDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualEzsignsignergroupDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualEzsignsignergroupDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualEzsignsignergroupDescription(val *MultilingualEzsignsignergroupDescription) *NullableMultilingualEzsignsignergroupDescription {
	return &NullableMultilingualEzsignsignergroupDescription{value: val, isSet: true}
}

func (v NullableMultilingualEzsignsignergroupDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualEzsignsignergroupDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


