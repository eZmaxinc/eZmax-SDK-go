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

// FieldEEzsignformfieldgroupType The Type of Ezsignformfieldgroup
type FieldEEzsignformfieldgroupType string

// List of Field-eEzsignformfieldgroupType
const (
	TEXT FieldEEzsignformfieldgroupType = "Text"
	TEXTAREA FieldEEzsignformfieldgroupType = "Textarea"
	DROPDOWN FieldEEzsignformfieldgroupType = "Dropdown"
	RADIO FieldEEzsignformfieldgroupType = "Radio"
	CHECKBOX FieldEEzsignformfieldgroupType = "Checkbox"
	NUMBER FieldEEzsignformfieldgroupType = "Number"
	DATE FieldEEzsignformfieldgroupType = "Date"
)

// All allowed values of FieldEEzsignformfieldgroupType enum
var AllowedFieldEEzsignformfieldgroupTypeEnumValues = []FieldEEzsignformfieldgroupType{
	"Text",
	"Textarea",
	"Dropdown",
	"Radio",
	"Checkbox",
	"Number",
	"Date",
}

func (v *FieldEEzsignformfieldgroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignformfieldgroupType(value)
	for _, existing := range AllowedFieldEEzsignformfieldgroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignformfieldgroupType", value)
}

// NewFieldEEzsignformfieldgroupTypeFromValue returns a pointer to a valid FieldEEzsignformfieldgroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignformfieldgroupTypeFromValue(v string) (*FieldEEzsignformfieldgroupType, error) {
	ev := FieldEEzsignformfieldgroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignformfieldgroupType: valid values are %v", v, AllowedFieldEEzsignformfieldgroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignformfieldgroupType) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignformfieldgroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignformfieldgroupType value
func (v FieldEEzsignformfieldgroupType) Ptr() *FieldEEzsignformfieldgroupType {
	return &v
}

type NullableFieldEEzsignformfieldgroupType struct {
	value *FieldEEzsignformfieldgroupType
	isSet bool
}

func (v NullableFieldEEzsignformfieldgroupType) Get() *FieldEEzsignformfieldgroupType {
	return v.value
}

func (v *NullableFieldEEzsignformfieldgroupType) Set(val *FieldEEzsignformfieldgroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignformfieldgroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignformfieldgroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignformfieldgroupType(val *FieldEEzsignformfieldgroupType) *NullableFieldEEzsignformfieldgroupType {
	return &NullableFieldEEzsignformfieldgroupType{value: val, isSet: true}
}

func (v NullableFieldEEzsignformfieldgroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignformfieldgroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

