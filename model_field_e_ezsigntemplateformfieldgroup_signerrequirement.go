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

// FieldEEzsigntemplateformfieldgroupSignerrequirement The Signer requirement of the Ezsigntemplateformfieldgroup. **All** means anyone can fill it, **One** means a specific person must fill it.
type FieldEEzsigntemplateformfieldgroupSignerrequirement string

// List of Field-eEzsigntemplateformfieldgroupSignerrequirement
const (
	ALL FieldEEzsigntemplateformfieldgroupSignerrequirement = "All"
	ONE FieldEEzsigntemplateformfieldgroupSignerrequirement = "One"
)

// All allowed values of FieldEEzsigntemplateformfieldgroupSignerrequirement enum
var AllowedFieldEEzsigntemplateformfieldgroupSignerrequirementEnumValues = []FieldEEzsigntemplateformfieldgroupSignerrequirement{
	"All",
	"One",
}

func (v *FieldEEzsigntemplateformfieldgroupSignerrequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateformfieldgroupSignerrequirement(value)
	for _, existing := range AllowedFieldEEzsigntemplateformfieldgroupSignerrequirementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateformfieldgroupSignerrequirement", value)
}

// NewFieldEEzsigntemplateformfieldgroupSignerrequirementFromValue returns a pointer to a valid FieldEEzsigntemplateformfieldgroupSignerrequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateformfieldgroupSignerrequirementFromValue(v string) (*FieldEEzsigntemplateformfieldgroupSignerrequirement, error) {
	ev := FieldEEzsigntemplateformfieldgroupSignerrequirement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateformfieldgroupSignerrequirement: valid values are %v", v, AllowedFieldEEzsigntemplateformfieldgroupSignerrequirementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateformfieldgroupSignerrequirement) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateformfieldgroupSignerrequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateformfieldgroupSignerrequirement value
func (v FieldEEzsigntemplateformfieldgroupSignerrequirement) Ptr() *FieldEEzsigntemplateformfieldgroupSignerrequirement {
	return &v
}

type NullableFieldEEzsigntemplateformfieldgroupSignerrequirement struct {
	value *FieldEEzsigntemplateformfieldgroupSignerrequirement
	isSet bool
}

func (v NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) Get() *FieldEEzsigntemplateformfieldgroupSignerrequirement {
	return v.value
}

func (v *NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) Set(val *FieldEEzsigntemplateformfieldgroupSignerrequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateformfieldgroupSignerrequirement(val *FieldEEzsigntemplateformfieldgroupSignerrequirement) *NullableFieldEEzsigntemplateformfieldgroupSignerrequirement {
	return &NullableFieldEEzsigntemplateformfieldgroupSignerrequirement{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateformfieldgroupSignerrequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

