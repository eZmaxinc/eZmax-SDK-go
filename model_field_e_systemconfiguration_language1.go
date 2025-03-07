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
	"fmt"
)

// FieldESystemconfigurationLanguage1 The type of the french for the client
type FieldESystemconfigurationLanguage1 string

// List of Field-eSystemconfigurationLanguage1
const (
	FR_QC FieldESystemconfigurationLanguage1 = "fr_QC"
)

// All allowed values of FieldESystemconfigurationLanguage1 enum
var AllowedFieldESystemconfigurationLanguage1EnumValues = []FieldESystemconfigurationLanguage1{
	"fr_QC",
}

func (v *FieldESystemconfigurationLanguage1) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldESystemconfigurationLanguage1(value)
	for _, existing := range AllowedFieldESystemconfigurationLanguage1EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldESystemconfigurationLanguage1", value)
}

// NewFieldESystemconfigurationLanguage1FromValue returns a pointer to a valid FieldESystemconfigurationLanguage1
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldESystemconfigurationLanguage1FromValue(v string) (*FieldESystemconfigurationLanguage1, error) {
	ev := FieldESystemconfigurationLanguage1(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldESystemconfigurationLanguage1: valid values are %v", v, AllowedFieldESystemconfigurationLanguage1EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldESystemconfigurationLanguage1) IsValid() bool {
	for _, existing := range AllowedFieldESystemconfigurationLanguage1EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eSystemconfigurationLanguage1 value
func (v FieldESystemconfigurationLanguage1) Ptr() *FieldESystemconfigurationLanguage1 {
	return &v
}

type NullableFieldESystemconfigurationLanguage1 struct {
	value *FieldESystemconfigurationLanguage1
	isSet bool
}

func (v NullableFieldESystemconfigurationLanguage1) Get() *FieldESystemconfigurationLanguage1 {
	return v.value
}

func (v *NullableFieldESystemconfigurationLanguage1) Set(val *FieldESystemconfigurationLanguage1) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldESystemconfigurationLanguage1) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldESystemconfigurationLanguage1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldESystemconfigurationLanguage1(val *FieldESystemconfigurationLanguage1) *NullableFieldESystemconfigurationLanguage1 {
	return &NullableFieldESystemconfigurationLanguage1{value: val, isSet: true}
}

func (v NullableFieldESystemconfigurationLanguage1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldESystemconfigurationLanguage1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

