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
	"fmt"
)

// FieldEActivesessionOrigin The Origin of User for the Activesession
type FieldEActivesessionOrigin string

// List of Field-eActivesessionOrigin
const (
	BUILT_IN FieldEActivesessionOrigin = "BuiltIn"
	EXTERNAL FieldEActivesessionOrigin = "External"
)

// All allowed values of FieldEActivesessionOrigin enum
var AllowedFieldEActivesessionOriginEnumValues = []FieldEActivesessionOrigin{
	"BuiltIn",
	"External",
}

func (v *FieldEActivesessionOrigin) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEActivesessionOrigin(value)
	for _, existing := range AllowedFieldEActivesessionOriginEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEActivesessionOrigin", value)
}

// NewFieldEActivesessionOriginFromValue returns a pointer to a valid FieldEActivesessionOrigin
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEActivesessionOriginFromValue(v string) (*FieldEActivesessionOrigin, error) {
	ev := FieldEActivesessionOrigin(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEActivesessionOrigin: valid values are %v", v, AllowedFieldEActivesessionOriginEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEActivesessionOrigin) IsValid() bool {
	for _, existing := range AllowedFieldEActivesessionOriginEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eActivesessionOrigin value
func (v FieldEActivesessionOrigin) Ptr() *FieldEActivesessionOrigin {
	return &v
}

type NullableFieldEActivesessionOrigin struct {
	value *FieldEActivesessionOrigin
	isSet bool
}

func (v NullableFieldEActivesessionOrigin) Get() *FieldEActivesessionOrigin {
	return v.value
}

func (v *NullableFieldEActivesessionOrigin) Set(val *FieldEActivesessionOrigin) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEActivesessionOrigin) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEActivesessionOrigin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEActivesessionOrigin(val *FieldEActivesessionOrigin) *NullableFieldEActivesessionOrigin {
	return &NullableFieldEActivesessionOrigin{value: val, isSet: true}
}

func (v NullableFieldEActivesessionOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEActivesessionOrigin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
