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

// checks if the MultilingualNotificationtestName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualNotificationtestName{}

// MultilingualNotificationtestName Name of the Notificationtest
type MultilingualNotificationtestName struct {
	// The name of the Notificationtest in French
	SNotificationtestName1 *string `json:"sNotificationtestName1,omitempty"`
	// The name of the Notificationtest in English
	SNotificationtestName2 *string `json:"sNotificationtestName2,omitempty"`
}

// NewMultilingualNotificationtestName instantiates a new MultilingualNotificationtestName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualNotificationtestName() *MultilingualNotificationtestName {
	this := MultilingualNotificationtestName{}
	return &this
}

// NewMultilingualNotificationtestNameWithDefaults instantiates a new MultilingualNotificationtestName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualNotificationtestNameWithDefaults() *MultilingualNotificationtestName {
	this := MultilingualNotificationtestName{}
	return &this
}

// GetSNotificationtestName1 returns the SNotificationtestName1 field value if set, zero value otherwise.
func (o *MultilingualNotificationtestName) GetSNotificationtestName1() string {
	if o == nil || IsNil(o.SNotificationtestName1) {
		var ret string
		return ret
	}
	return *o.SNotificationtestName1
}

// GetSNotificationtestName1Ok returns a tuple with the SNotificationtestName1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualNotificationtestName) GetSNotificationtestName1Ok() (*string, bool) {
	if o == nil || IsNil(o.SNotificationtestName1) {
		return nil, false
	}
	return o.SNotificationtestName1, true
}

// HasSNotificationtestName1 returns a boolean if a field has been set.
func (o *MultilingualNotificationtestName) HasSNotificationtestName1() bool {
	if o != nil && !IsNil(o.SNotificationtestName1) {
		return true
	}

	return false
}

// SetSNotificationtestName1 gets a reference to the given string and assigns it to the SNotificationtestName1 field.
func (o *MultilingualNotificationtestName) SetSNotificationtestName1(v string) {
	o.SNotificationtestName1 = &v
}

// GetSNotificationtestName2 returns the SNotificationtestName2 field value if set, zero value otherwise.
func (o *MultilingualNotificationtestName) GetSNotificationtestName2() string {
	if o == nil || IsNil(o.SNotificationtestName2) {
		var ret string
		return ret
	}
	return *o.SNotificationtestName2
}

// GetSNotificationtestName2Ok returns a tuple with the SNotificationtestName2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualNotificationtestName) GetSNotificationtestName2Ok() (*string, bool) {
	if o == nil || IsNil(o.SNotificationtestName2) {
		return nil, false
	}
	return o.SNotificationtestName2, true
}

// HasSNotificationtestName2 returns a boolean if a field has been set.
func (o *MultilingualNotificationtestName) HasSNotificationtestName2() bool {
	if o != nil && !IsNil(o.SNotificationtestName2) {
		return true
	}

	return false
}

// SetSNotificationtestName2 gets a reference to the given string and assigns it to the SNotificationtestName2 field.
func (o *MultilingualNotificationtestName) SetSNotificationtestName2(v string) {
	o.SNotificationtestName2 = &v
}

func (o MultilingualNotificationtestName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualNotificationtestName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SNotificationtestName1) {
		toSerialize["sNotificationtestName1"] = o.SNotificationtestName1
	}
	if !IsNil(o.SNotificationtestName2) {
		toSerialize["sNotificationtestName2"] = o.SNotificationtestName2
	}
	return toSerialize, nil
}

type NullableMultilingualNotificationtestName struct {
	value *MultilingualNotificationtestName
	isSet bool
}

func (v NullableMultilingualNotificationtestName) Get() *MultilingualNotificationtestName {
	return v.value
}

func (v *NullableMultilingualNotificationtestName) Set(val *MultilingualNotificationtestName) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualNotificationtestName) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualNotificationtestName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualNotificationtestName(val *MultilingualNotificationtestName) *NullableMultilingualNotificationtestName {
	return &NullableMultilingualNotificationtestName{value: val, isSet: true}
}

func (v NullableMultilingualNotificationtestName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualNotificationtestName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


