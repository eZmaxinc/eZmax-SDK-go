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

// FieldEEzsignelementdependencyValidation The validation type of the Ezsignelementdependency
type FieldEEzsignelementdependencyValidation string

// List of Field-eEzsignelementdependencyValidation
const (
	VALUE FieldEEzsignelementdependencyValidation = "Value"
	SELECTED FieldEEzsignelementdependencyValidation = "Selected"
	FILLED FieldEEzsignelementdependencyValidation = "Filled"
)

// All allowed values of FieldEEzsignelementdependencyValidation enum
var AllowedFieldEEzsignelementdependencyValidationEnumValues = []FieldEEzsignelementdependencyValidation{
	"Value",
	"Selected",
	"Filled",
}

func (v *FieldEEzsignelementdependencyValidation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignelementdependencyValidation(value)
	for _, existing := range AllowedFieldEEzsignelementdependencyValidationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignelementdependencyValidation", value)
}

// NewFieldEEzsignelementdependencyValidationFromValue returns a pointer to a valid FieldEEzsignelementdependencyValidation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignelementdependencyValidationFromValue(v string) (*FieldEEzsignelementdependencyValidation, error) {
	ev := FieldEEzsignelementdependencyValidation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignelementdependencyValidation: valid values are %v", v, AllowedFieldEEzsignelementdependencyValidationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignelementdependencyValidation) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignelementdependencyValidationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignelementdependencyValidation value
func (v FieldEEzsignelementdependencyValidation) Ptr() *FieldEEzsignelementdependencyValidation {
	return &v
}

type NullableFieldEEzsignelementdependencyValidation struct {
	value *FieldEEzsignelementdependencyValidation
	isSet bool
}

func (v NullableFieldEEzsignelementdependencyValidation) Get() *FieldEEzsignelementdependencyValidation {
	return v.value
}

func (v *NullableFieldEEzsignelementdependencyValidation) Set(val *FieldEEzsignelementdependencyValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignelementdependencyValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignelementdependencyValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignelementdependencyValidation(val *FieldEEzsignelementdependencyValidation) *NullableFieldEEzsignelementdependencyValidation {
	return &NullableFieldEEzsignelementdependencyValidation{value: val, isSet: true}
}

func (v NullableFieldEEzsignelementdependencyValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignelementdependencyValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

