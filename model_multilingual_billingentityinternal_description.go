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

// checks if the MultilingualBillingentityinternalDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualBillingentityinternalDescription{}

// MultilingualBillingentityinternalDescription The description of the Billingentityinternal
type MultilingualBillingentityinternalDescription struct {
	// The description of the Billingentityinternal in French
	SBillingentityinternalDescription1 *string `json:"sBillingentityinternalDescription1,omitempty" validate:"regexp=^.{0,70}$"`
	// The description of the Billingentityinternal in English
	SBillingentityinternalDescription2 *string `json:"sBillingentityinternalDescription2,omitempty" validate:"regexp=^.{0,70}$"`
}

// NewMultilingualBillingentityinternalDescription instantiates a new MultilingualBillingentityinternalDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualBillingentityinternalDescription() *MultilingualBillingentityinternalDescription {
	this := MultilingualBillingentityinternalDescription{}
	return &this
}

// NewMultilingualBillingentityinternalDescriptionWithDefaults instantiates a new MultilingualBillingentityinternalDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualBillingentityinternalDescriptionWithDefaults() *MultilingualBillingentityinternalDescription {
	this := MultilingualBillingentityinternalDescription{}
	return &this
}

// GetSBillingentityinternalDescription1 returns the SBillingentityinternalDescription1 field value if set, zero value otherwise.
func (o *MultilingualBillingentityinternalDescription) GetSBillingentityinternalDescription1() string {
	if o == nil || IsNil(o.SBillingentityinternalDescription1) {
		var ret string
		return ret
	}
	return *o.SBillingentityinternalDescription1
}

// GetSBillingentityinternalDescription1Ok returns a tuple with the SBillingentityinternalDescription1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualBillingentityinternalDescription) GetSBillingentityinternalDescription1Ok() (*string, bool) {
	if o == nil || IsNil(o.SBillingentityinternalDescription1) {
		return nil, false
	}
	return o.SBillingentityinternalDescription1, true
}

// HasSBillingentityinternalDescription1 returns a boolean if a field has been set.
func (o *MultilingualBillingentityinternalDescription) HasSBillingentityinternalDescription1() bool {
	if o != nil && !IsNil(o.SBillingentityinternalDescription1) {
		return true
	}

	return false
}

// SetSBillingentityinternalDescription1 gets a reference to the given string and assigns it to the SBillingentityinternalDescription1 field.
func (o *MultilingualBillingentityinternalDescription) SetSBillingentityinternalDescription1(v string) {
	o.SBillingentityinternalDescription1 = &v
}

// GetSBillingentityinternalDescription2 returns the SBillingentityinternalDescription2 field value if set, zero value otherwise.
func (o *MultilingualBillingentityinternalDescription) GetSBillingentityinternalDescription2() string {
	if o == nil || IsNil(o.SBillingentityinternalDescription2) {
		var ret string
		return ret
	}
	return *o.SBillingentityinternalDescription2
}

// GetSBillingentityinternalDescription2Ok returns a tuple with the SBillingentityinternalDescription2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualBillingentityinternalDescription) GetSBillingentityinternalDescription2Ok() (*string, bool) {
	if o == nil || IsNil(o.SBillingentityinternalDescription2) {
		return nil, false
	}
	return o.SBillingentityinternalDescription2, true
}

// HasSBillingentityinternalDescription2 returns a boolean if a field has been set.
func (o *MultilingualBillingentityinternalDescription) HasSBillingentityinternalDescription2() bool {
	if o != nil && !IsNil(o.SBillingentityinternalDescription2) {
		return true
	}

	return false
}

// SetSBillingentityinternalDescription2 gets a reference to the given string and assigns it to the SBillingentityinternalDescription2 field.
func (o *MultilingualBillingentityinternalDescription) SetSBillingentityinternalDescription2(v string) {
	o.SBillingentityinternalDescription2 = &v
}

func (o MultilingualBillingentityinternalDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualBillingentityinternalDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SBillingentityinternalDescription1) {
		toSerialize["sBillingentityinternalDescription1"] = o.SBillingentityinternalDescription1
	}
	if !IsNil(o.SBillingentityinternalDescription2) {
		toSerialize["sBillingentityinternalDescription2"] = o.SBillingentityinternalDescription2
	}
	return toSerialize, nil
}

type NullableMultilingualBillingentityinternalDescription struct {
	value *MultilingualBillingentityinternalDescription
	isSet bool
}

func (v NullableMultilingualBillingentityinternalDescription) Get() *MultilingualBillingentityinternalDescription {
	return v.value
}

func (v *NullableMultilingualBillingentityinternalDescription) Set(val *MultilingualBillingentityinternalDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualBillingentityinternalDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualBillingentityinternalDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualBillingentityinternalDescription(val *MultilingualBillingentityinternalDescription) *NullableMultilingualBillingentityinternalDescription {
	return &NullableMultilingualBillingentityinternalDescription{value: val, isSet: true}
}

func (v NullableMultilingualBillingentityinternalDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualBillingentityinternalDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


