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

// FieldEEzsigntemplateformfieldDependencyrequirement The Dependency requirement of the Ezsigntemplateformfield
type FieldEEzsigntemplateformfieldDependencyrequirement string

// List of Field-eEzsigntemplateformfieldDependencyrequirement
const (
	ALL_OF FieldEEzsigntemplateformfieldDependencyrequirement = "AllOf"
	ANY_OF FieldEEzsigntemplateformfieldDependencyrequirement = "AnyOf"
)

// All allowed values of FieldEEzsigntemplateformfieldDependencyrequirement enum
var AllowedFieldEEzsigntemplateformfieldDependencyrequirementEnumValues = []FieldEEzsigntemplateformfieldDependencyrequirement{
	"AllOf",
	"AnyOf",
}

func (v *FieldEEzsigntemplateformfieldDependencyrequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateformfieldDependencyrequirement(value)
	for _, existing := range AllowedFieldEEzsigntemplateformfieldDependencyrequirementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateformfieldDependencyrequirement", value)
}

// NewFieldEEzsigntemplateformfieldDependencyrequirementFromValue returns a pointer to a valid FieldEEzsigntemplateformfieldDependencyrequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateformfieldDependencyrequirementFromValue(v string) (*FieldEEzsigntemplateformfieldDependencyrequirement, error) {
	ev := FieldEEzsigntemplateformfieldDependencyrequirement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateformfieldDependencyrequirement: valid values are %v", v, AllowedFieldEEzsigntemplateformfieldDependencyrequirementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateformfieldDependencyrequirement) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateformfieldDependencyrequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateformfieldDependencyrequirement value
func (v FieldEEzsigntemplateformfieldDependencyrequirement) Ptr() *FieldEEzsigntemplateformfieldDependencyrequirement {
	return &v
}

type NullableFieldEEzsigntemplateformfieldDependencyrequirement struct {
	value *FieldEEzsigntemplateformfieldDependencyrequirement
	isSet bool
}

func (v NullableFieldEEzsigntemplateformfieldDependencyrequirement) Get() *FieldEEzsigntemplateformfieldDependencyrequirement {
	return v.value
}

func (v *NullableFieldEEzsigntemplateformfieldDependencyrequirement) Set(val *FieldEEzsigntemplateformfieldDependencyrequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateformfieldDependencyrequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateformfieldDependencyrequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateformfieldDependencyrequirement(val *FieldEEzsigntemplateformfieldDependencyrequirement) *NullableFieldEEzsigntemplateformfieldDependencyrequirement {
	return &NullableFieldEEzsigntemplateformfieldDependencyrequirement{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateformfieldDependencyrequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateformfieldDependencyrequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

