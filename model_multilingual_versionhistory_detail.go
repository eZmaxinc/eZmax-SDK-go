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

// checks if the MultilingualVersionhistoryDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultilingualVersionhistoryDetail{}

// MultilingualVersionhistoryDetail Detail of the Versionhistory
type MultilingualVersionhistoryDetail struct {
	// Detail of the Versionhistory in French
	TVersionhistoryDetail1 *string `json:"tVersionhistoryDetail1,omitempty"`
	// Detail of the Versionhistory in English
	TVersionhistoryDetail2 *string `json:"tVersionhistoryDetail2,omitempty"`
}

// NewMultilingualVersionhistoryDetail instantiates a new MultilingualVersionhistoryDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultilingualVersionhistoryDetail() *MultilingualVersionhistoryDetail {
	this := MultilingualVersionhistoryDetail{}
	return &this
}

// NewMultilingualVersionhistoryDetailWithDefaults instantiates a new MultilingualVersionhistoryDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultilingualVersionhistoryDetailWithDefaults() *MultilingualVersionhistoryDetail {
	this := MultilingualVersionhistoryDetail{}
	return &this
}

// GetTVersionhistoryDetail1 returns the TVersionhistoryDetail1 field value if set, zero value otherwise.
func (o *MultilingualVersionhistoryDetail) GetTVersionhistoryDetail1() string {
	if o == nil || IsNil(o.TVersionhistoryDetail1) {
		var ret string
		return ret
	}
	return *o.TVersionhistoryDetail1
}

// GetTVersionhistoryDetail1Ok returns a tuple with the TVersionhistoryDetail1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualVersionhistoryDetail) GetTVersionhistoryDetail1Ok() (*string, bool) {
	if o == nil || IsNil(o.TVersionhistoryDetail1) {
		return nil, false
	}
	return o.TVersionhistoryDetail1, true
}

// HasTVersionhistoryDetail1 returns a boolean if a field has been set.
func (o *MultilingualVersionhistoryDetail) HasTVersionhistoryDetail1() bool {
	if o != nil && !IsNil(o.TVersionhistoryDetail1) {
		return true
	}

	return false
}

// SetTVersionhistoryDetail1 gets a reference to the given string and assigns it to the TVersionhistoryDetail1 field.
func (o *MultilingualVersionhistoryDetail) SetTVersionhistoryDetail1(v string) {
	o.TVersionhistoryDetail1 = &v
}

// GetTVersionhistoryDetail2 returns the TVersionhistoryDetail2 field value if set, zero value otherwise.
func (o *MultilingualVersionhistoryDetail) GetTVersionhistoryDetail2() string {
	if o == nil || IsNil(o.TVersionhistoryDetail2) {
		var ret string
		return ret
	}
	return *o.TVersionhistoryDetail2
}

// GetTVersionhistoryDetail2Ok returns a tuple with the TVersionhistoryDetail2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultilingualVersionhistoryDetail) GetTVersionhistoryDetail2Ok() (*string, bool) {
	if o == nil || IsNil(o.TVersionhistoryDetail2) {
		return nil, false
	}
	return o.TVersionhistoryDetail2, true
}

// HasTVersionhistoryDetail2 returns a boolean if a field has been set.
func (o *MultilingualVersionhistoryDetail) HasTVersionhistoryDetail2() bool {
	if o != nil && !IsNil(o.TVersionhistoryDetail2) {
		return true
	}

	return false
}

// SetTVersionhistoryDetail2 gets a reference to the given string and assigns it to the TVersionhistoryDetail2 field.
func (o *MultilingualVersionhistoryDetail) SetTVersionhistoryDetail2(v string) {
	o.TVersionhistoryDetail2 = &v
}

func (o MultilingualVersionhistoryDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultilingualVersionhistoryDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TVersionhistoryDetail1) {
		toSerialize["tVersionhistoryDetail1"] = o.TVersionhistoryDetail1
	}
	if !IsNil(o.TVersionhistoryDetail2) {
		toSerialize["tVersionhistoryDetail2"] = o.TVersionhistoryDetail2
	}
	return toSerialize, nil
}

type NullableMultilingualVersionhistoryDetail struct {
	value *MultilingualVersionhistoryDetail
	isSet bool
}

func (v NullableMultilingualVersionhistoryDetail) Get() *MultilingualVersionhistoryDetail {
	return v.value
}

func (v *NullableMultilingualVersionhistoryDetail) Set(val *MultilingualVersionhistoryDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableMultilingualVersionhistoryDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableMultilingualVersionhistoryDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultilingualVersionhistoryDetail(val *MultilingualVersionhistoryDetail) *NullableMultilingualVersionhistoryDetail {
	return &NullableMultilingualVersionhistoryDetail{value: val, isSet: true}
}

func (v NullableMultilingualVersionhistoryDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultilingualVersionhistoryDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


