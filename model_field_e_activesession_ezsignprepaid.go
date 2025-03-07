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

// FieldEActivesessionEzsignprepaid eZsign subscription level
type FieldEActivesessionEzsignprepaid string

// List of Field-eActivesessionEzsignprepaid
const (
	NO FieldEActivesessionEzsignprepaid = "No"
	BASIC FieldEActivesessionEzsignprepaid = "Basic"
	STANDARD FieldEActivesessionEzsignprepaid = "Standard"
	PRO FieldEActivesessionEzsignprepaid = "Pro"
)

// All allowed values of FieldEActivesessionEzsignprepaid enum
var AllowedFieldEActivesessionEzsignprepaidEnumValues = []FieldEActivesessionEzsignprepaid{
	"No",
	"Basic",
	"Standard",
	"Pro",
}

func (v *FieldEActivesessionEzsignprepaid) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEActivesessionEzsignprepaid(value)
	for _, existing := range AllowedFieldEActivesessionEzsignprepaidEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEActivesessionEzsignprepaid", value)
}

// NewFieldEActivesessionEzsignprepaidFromValue returns a pointer to a valid FieldEActivesessionEzsignprepaid
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEActivesessionEzsignprepaidFromValue(v string) (*FieldEActivesessionEzsignprepaid, error) {
	ev := FieldEActivesessionEzsignprepaid(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEActivesessionEzsignprepaid: valid values are %v", v, AllowedFieldEActivesessionEzsignprepaidEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEActivesessionEzsignprepaid) IsValid() bool {
	for _, existing := range AllowedFieldEActivesessionEzsignprepaidEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eActivesessionEzsignprepaid value
func (v FieldEActivesessionEzsignprepaid) Ptr() *FieldEActivesessionEzsignprepaid {
	return &v
}

type NullableFieldEActivesessionEzsignprepaid struct {
	value *FieldEActivesessionEzsignprepaid
	isSet bool
}

func (v NullableFieldEActivesessionEzsignprepaid) Get() *FieldEActivesessionEzsignprepaid {
	return v.value
}

func (v *NullableFieldEActivesessionEzsignprepaid) Set(val *FieldEActivesessionEzsignprepaid) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEActivesessionEzsignprepaid) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEActivesessionEzsignprepaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEActivesessionEzsignprepaid(val *FieldEActivesessionEzsignprepaid) *NullableFieldEActivesessionEzsignprepaid {
	return &NullableFieldEActivesessionEzsignprepaid{value: val, isSet: true}
}

func (v NullableFieldEActivesessionEzsignprepaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEActivesessionEzsignprepaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

