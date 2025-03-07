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

// FieldEEzsigntemplateformfieldgroupType The Type of Ezsigntemplateformfieldgroup
type FieldEEzsigntemplateformfieldgroupType string

// List of Field-eEzsigntemplateformfieldgroupType
const (
	TEXT FieldEEzsigntemplateformfieldgroupType = "Text"
	TEXTAREA FieldEEzsigntemplateformfieldgroupType = "Textarea"
	DROPDOWN FieldEEzsigntemplateformfieldgroupType = "Dropdown"
	RADIO FieldEEzsigntemplateformfieldgroupType = "Radio"
	CHECKBOX FieldEEzsigntemplateformfieldgroupType = "Checkbox"
	NUMBER FieldEEzsigntemplateformfieldgroupType = "Number"
	DATE FieldEEzsigntemplateformfieldgroupType = "Date"
)

// All allowed values of FieldEEzsigntemplateformfieldgroupType enum
var AllowedFieldEEzsigntemplateformfieldgroupTypeEnumValues = []FieldEEzsigntemplateformfieldgroupType{
	"Text",
	"Textarea",
	"Dropdown",
	"Radio",
	"Checkbox",
	"Number",
	"Date",
}

func (v *FieldEEzsigntemplateformfieldgroupType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateformfieldgroupType(value)
	for _, existing := range AllowedFieldEEzsigntemplateformfieldgroupTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateformfieldgroupType", value)
}

// NewFieldEEzsigntemplateformfieldgroupTypeFromValue returns a pointer to a valid FieldEEzsigntemplateformfieldgroupType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateformfieldgroupTypeFromValue(v string) (*FieldEEzsigntemplateformfieldgroupType, error) {
	ev := FieldEEzsigntemplateformfieldgroupType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateformfieldgroupType: valid values are %v", v, AllowedFieldEEzsigntemplateformfieldgroupTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateformfieldgroupType) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateformfieldgroupTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateformfieldgroupType value
func (v FieldEEzsigntemplateformfieldgroupType) Ptr() *FieldEEzsigntemplateformfieldgroupType {
	return &v
}

type NullableFieldEEzsigntemplateformfieldgroupType struct {
	value *FieldEEzsigntemplateformfieldgroupType
	isSet bool
}

func (v NullableFieldEEzsigntemplateformfieldgroupType) Get() *FieldEEzsigntemplateformfieldgroupType {
	return v.value
}

func (v *NullableFieldEEzsigntemplateformfieldgroupType) Set(val *FieldEEzsigntemplateformfieldgroupType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateformfieldgroupType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateformfieldgroupType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateformfieldgroupType(val *FieldEEzsigntemplateformfieldgroupType) *NullableFieldEEzsigntemplateformfieldgroupType {
	return &NullableFieldEEzsigntemplateformfieldgroupType{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateformfieldgroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateformfieldgroupType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

