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

// FieldEEzsigntemplateType The Type of Ezsigntemplate
type FieldEEzsigntemplateType string

// List of Field-eEzsigntemplateType
const (
	USER FieldEEzsigntemplateType = "User"
	USERGROUP FieldEEzsigntemplateType = "Usergroup"
	COMPANY FieldEEzsigntemplateType = "Company"
)

// All allowed values of FieldEEzsigntemplateType enum
var AllowedFieldEEzsigntemplateTypeEnumValues = []FieldEEzsigntemplateType{
	"User",
	"Usergroup",
	"Company",
}

func (v *FieldEEzsigntemplateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateType(value)
	for _, existing := range AllowedFieldEEzsigntemplateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateType", value)
}

// NewFieldEEzsigntemplateTypeFromValue returns a pointer to a valid FieldEEzsigntemplateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateTypeFromValue(v string) (*FieldEEzsigntemplateType, error) {
	ev := FieldEEzsigntemplateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateType: valid values are %v", v, AllowedFieldEEzsigntemplateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateType) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateType value
func (v FieldEEzsigntemplateType) Ptr() *FieldEEzsigntemplateType {
	return &v
}

type NullableFieldEEzsigntemplateType struct {
	value *FieldEEzsigntemplateType
	isSet bool
}

func (v NullableFieldEEzsigntemplateType) Get() *FieldEEzsigntemplateType {
	return v.value
}

func (v *NullableFieldEEzsigntemplateType) Set(val *FieldEEzsigntemplateType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateType(val *FieldEEzsigntemplateType) *NullableFieldEEzsigntemplateType {
	return &NullableFieldEEzsigntemplateType{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

