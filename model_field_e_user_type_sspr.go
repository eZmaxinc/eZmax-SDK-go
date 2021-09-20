/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEUserTypeSSPR The user type of the User for SSPR
type FieldEUserTypeSSPR string

// List of Field-eUserTypeSSPR
const (
	EZSIGN_USER FieldEUserTypeSSPR = "EzsignUser"
	NATIVE FieldEUserTypeSSPR = "Native"
)

var allowedFieldEUserTypeSSPREnumValues = []FieldEUserTypeSSPR{
	"EzsignUser",
	"Native",
}

func (v *FieldEUserTypeSSPR) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEUserTypeSSPR(value)
	for _, existing := range allowedFieldEUserTypeSSPREnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEUserTypeSSPR", value)
}

// NewFieldEUserTypeSSPRFromValue returns a pointer to a valid FieldEUserTypeSSPR
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEUserTypeSSPRFromValue(v string) (*FieldEUserTypeSSPR, error) {
	ev := FieldEUserTypeSSPR(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEUserTypeSSPR: valid values are %v", v, allowedFieldEUserTypeSSPREnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEUserTypeSSPR) IsValid() bool {
	for _, existing := range allowedFieldEUserTypeSSPREnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eUserTypeSSPR value
func (v FieldEUserTypeSSPR) Ptr() *FieldEUserTypeSSPR {
	return &v
}

type NullableFieldEUserTypeSSPR struct {
	value *FieldEUserTypeSSPR
	isSet bool
}

func (v NullableFieldEUserTypeSSPR) Get() *FieldEUserTypeSSPR {
	return v.value
}

func (v *NullableFieldEUserTypeSSPR) Set(val *FieldEUserTypeSSPR) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEUserTypeSSPR) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEUserTypeSSPR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEUserTypeSSPR(val *FieldEUserTypeSSPR) *NullableFieldEUserTypeSSPR {
	return &NullableFieldEUserTypeSSPR{value: val, isSet: true}
}

func (v NullableFieldEUserTypeSSPR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEUserTypeSSPR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

