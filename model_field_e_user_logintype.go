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

// FieldEUserLogintype The type of authentication for the User
type FieldEUserLogintype string

// List of Field-eUserLogintype
const (
	PASSWORD FieldEUserLogintype = "Password"
	PASSWORD_PHONE FieldEUserLogintype = "PasswordPhone"
	PASSWORD_QUESTION FieldEUserLogintype = "PasswordQuestion"
)

// All allowed values of FieldEUserLogintype enum
var AllowedFieldEUserLogintypeEnumValues = []FieldEUserLogintype{
	"Password",
	"PasswordPhone",
	"PasswordQuestion",
}

func (v *FieldEUserLogintype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEUserLogintype(value)
	for _, existing := range AllowedFieldEUserLogintypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEUserLogintype", value)
}

// NewFieldEUserLogintypeFromValue returns a pointer to a valid FieldEUserLogintype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEUserLogintypeFromValue(v string) (*FieldEUserLogintype, error) {
	ev := FieldEUserLogintype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEUserLogintype: valid values are %v", v, AllowedFieldEUserLogintypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEUserLogintype) IsValid() bool {
	for _, existing := range AllowedFieldEUserLogintypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eUserLogintype value
func (v FieldEUserLogintype) Ptr() *FieldEUserLogintype {
	return &v
}

type NullableFieldEUserLogintype struct {
	value *FieldEUserLogintype
	isSet bool
}

func (v NullableFieldEUserLogintype) Get() *FieldEUserLogintype {
	return v.value
}

func (v *NullableFieldEUserLogintype) Set(val *FieldEUserLogintype) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEUserLogintype) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEUserLogintype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEUserLogintype(val *FieldEUserLogintype) *NullableFieldEUserLogintype {
	return &NullableFieldEUserLogintype{value: val, isSet: true}
}

func (v NullableFieldEUserLogintype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEUserLogintype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

