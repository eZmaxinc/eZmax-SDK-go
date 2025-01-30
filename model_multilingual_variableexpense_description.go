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

// checks if the MultilingualVariableexpenseDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualVariableexpenseDescription{}

// MultilingualVariableexpenseDescription The description of the Variableexpense
type MultilingualVariableexpenseDescription struct {
	// The description of the Variableexpense in French
	SVariableexpenseDescription1 *string `json:"sVariableexpenseDescription1,omitempty" validate:"regexp=^.{0,40}$"`
	// The description of the Variableexpense in English
	SVariableexpenseDescription2 *string `json:"sVariableexpenseDescription2,omitempty" validate:"regexp=^.{0,40}$"`
}

// NewMultilingualVariableexpenseDescription instantiates a new MultilingualVariableexpenseDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualVariableexpenseDescription() *MultilingualVariableexpenseDescription {
	this := MultilingualVariableexpenseDescription{}
	return &this
}

// NewMultilingualVariableexpenseDescriptionWithDefaults instantiates a new MultilingualVariableexpenseDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualVariableexpenseDescriptionWithDefaults() *MultilingualVariableexpenseDescription {
	this := MultilingualVariableexpenseDescription{}
	return &this
}

// GetSVariableexpenseDescription1 returns the SVariableexpenseDescription1 field value if set, zero value otherwise.
func (o *MultilingualVariableexpenseDescription) GetSVariableexpenseDescription1() string {
	if o == nil || IsNil(o.SVariableexpenseDescription1) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseDescription1
}

// GetSVariableexpenseDescription1Ok returns a tuple with the SVariableexpenseDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualVariableexpenseDescription) GetSVariableexpenseDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseDescription1) {
		return nil, false
	}
	return o.SVariableexpenseDescription1, true
}

// HasSVariableexpenseDescription1 returns a boolean if a field has been set.
func (o *MultilingualVariableexpenseDescription) HasSVariableexpenseDescription1() bool {
	if o != nil && !IsNil(o.SVariableexpenseDescription1) {
		return true
	}

	return false
}

// SetSVariableexpenseDescription1 gets a reference to the given string and assigns it to the SVariableexpenseDescription1 field.
func (o *MultilingualVariableexpenseDescription) SetSVariableexpenseDescription1(v string) {
	o.SVariableexpenseDescription1 = &v
}

// GetSVariableexpenseDescription2 returns the SVariableexpenseDescription2 field value if set, zero value otherwise.
func (o *MultilingualVariableexpenseDescription) GetSVariableexpenseDescription2() string {
	if o == nil || IsNil(o.SVariableexpenseDescription2) {
		var ret string
		return ret
	}
	return *o.SVariableexpenseDescription2
}

// GetSVariableexpenseDescription2Ok returns a tuple with the SVariableexpenseDescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualVariableexpenseDescription) GetSVariableexpenseDescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.SVariableexpenseDescription2) {
		return nil, false
	}
	return o.SVariableexpenseDescription2, true
}

// HasSVariableexpenseDescription2 returns a boolean if a field has been set.
func (o *MultilingualVariableexpenseDescription) HasSVariableexpenseDescription2() bool {
	if o != nil && !IsNil(o.SVariableexpenseDescription2) {
		return true
	}

	return false
}

// SetSVariableexpenseDescription2 gets a reference to the given string and assigns it to the SVariableexpenseDescription2 field.
func (o *MultilingualVariableexpenseDescription) SetSVariableexpenseDescription2(v string) {
	o.SVariableexpenseDescription2 = &v
}

func (o MultilingualVariableexpenseDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualVariableexpenseDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SVariableexpenseDescription1) {
		toSerialize["sVariableexpenseDescription1"] = o.SVariableexpenseDescription1
	}
	if !IsNil(o.SVariableexpenseDescription2) {
		toSerialize["sVariableexpenseDescription2"] = o.SVariableexpenseDescription2
	}
	return toSerialize, nil
}

type NullableMultilingualVariableexpenseDescription struct {
	value *MultilingualVariableexpenseDescription
	isSet bool
}

func (v NullableMultilingualVariableexpenseDescription) Get() *MultilingualVariableexpenseDescription {
	return v.value
}

func (v *NullableMultilingualVariableexpenseDescription) Set(val *MultilingualVariableexpenseDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualVariableexpenseDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualVariableexpenseDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualVariableexpenseDescription(val *MultilingualVariableexpenseDescription) *NullableMultilingualVariableexpenseDescription {
	return &NullableMultilingualVariableexpenseDescription{value: val, isSet: true}
}

func (v NullableMultilingualVariableexpenseDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualVariableexpenseDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


