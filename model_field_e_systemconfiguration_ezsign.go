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

// FieldESystemconfigurationEzsign Whether if Ezsign is paid by the company or not
type FieldESystemconfigurationEzsign string

// List of Field-eSystemconfigurationEzsign
const (
	NO FieldESystemconfigurationEzsign = "No"
	YES FieldESystemconfigurationEzsign = "Yes"
)

// All allowed values of FieldESystemconfigurationEzsign enum
var AllowedFieldESystemconfigurationEzsignEnumValues = []FieldESystemconfigurationEzsign{
	"No",
	"Yes",
}

func (v *FieldESystemconfigurationEzsign) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldESystemconfigurationEzsign(value)
	for _, existing := range AllowedFieldESystemconfigurationEzsignEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldESystemconfigurationEzsign", value)
}

// NewFieldESystemconfigurationEzsignFromValue returns a pointer to a valid FieldESystemconfigurationEzsign
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldESystemconfigurationEzsignFromValue(v string) (*FieldESystemconfigurationEzsign, error) {
	ev := FieldESystemconfigurationEzsign(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldESystemconfigurationEzsign: valid values are %v", v, AllowedFieldESystemconfigurationEzsignEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldESystemconfigurationEzsign) IsValid() bool {
	for _, existing := range AllowedFieldESystemconfigurationEzsignEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eSystemconfigurationEzsign value
func (v FieldESystemconfigurationEzsign) Ptr() *FieldESystemconfigurationEzsign {
	return &v
}

type NullableFieldESystemconfigurationEzsign struct {
	value *FieldESystemconfigurationEzsign
	isSet bool
}

func (v NullableFieldESystemconfigurationEzsign) Get() *FieldESystemconfigurationEzsign {
	return v.value
}

func (v *NullableFieldESystemconfigurationEzsign) Set(val *FieldESystemconfigurationEzsign) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldESystemconfigurationEzsign) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldESystemconfigurationEzsign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldESystemconfigurationEzsign(val *FieldESystemconfigurationEzsign) *NullableFieldESystemconfigurationEzsign {
	return &NullableFieldESystemconfigurationEzsign{value: val, isSet: true}
}

func (v NullableFieldESystemconfigurationEzsign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldESystemconfigurationEzsign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

