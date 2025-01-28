/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldESystemconfigurationLanguage2 The type of the english for the client
type FieldESystemconfigurationLanguage2 string

// List of Field-eSystemconfigurationLanguage2
const (
	EN_CA FieldESystemconfigurationLanguage2 = "en_CA"
	EN_QC FieldESystemconfigurationLanguage2 = "en_QC"
	EN_US FieldESystemconfigurationLanguage2 = "en_US"
)

// All allowed values of FieldESystemconfigurationLanguage2 enum
var AllowedFieldESystemconfigurationLanguage2EnumValues = []FieldESystemconfigurationLanguage2{
	"en_CA",
	"en_QC",
	"en_US",
}

func (v *FieldESystemconfigurationLanguage2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldESystemconfigurationLanguage2(value)
	for _, existing := range AllowedFieldESystemconfigurationLanguage2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldESystemconfigurationLanguage2", value)
}

// NewFieldESystemconfigurationLanguage2FromValue returns a pointer to a valid FieldESystemconfigurationLanguage2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldESystemconfigurationLanguage2FromValue(v string) (*FieldESystemconfigurationLanguage2, error) {
	ev := FieldESystemconfigurationLanguage2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldESystemconfigurationLanguage2: valid values are %v", v, AllowedFieldESystemconfigurationLanguage2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldESystemconfigurationLanguage2) IsValid() bool {
	for _, existing := range AllowedFieldESystemconfigurationLanguage2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eSystemconfigurationLanguage2 value
func (v FieldESystemconfigurationLanguage2) Ptr() *FieldESystemconfigurationLanguage2 {
	return &v
}

type NullableFieldESystemconfigurationLanguage2 struct {
	value *FieldESystemconfigurationLanguage2
	isSet bool
}

func (v NullableFieldESystemconfigurationLanguage2) Get() *FieldESystemconfigurationLanguage2 {
	return v.value
}

func (v *NullableFieldESystemconfigurationLanguage2) Set(val *FieldESystemconfigurationLanguage2) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldESystemconfigurationLanguage2) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldESystemconfigurationLanguage2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldESystemconfigurationLanguage2(val *FieldESystemconfigurationLanguage2) *NullableFieldESystemconfigurationLanguage2 {
	return &NullableFieldESystemconfigurationLanguage2{value: val, isSet: true}
}

func (v NullableFieldESystemconfigurationLanguage2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldESystemconfigurationLanguage2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

