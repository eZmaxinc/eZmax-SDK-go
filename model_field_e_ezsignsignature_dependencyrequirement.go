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

// FieldEEzsignsignatureDependencyrequirement The Dependency requirement of the Ezsignsignature
type FieldEEzsignsignatureDependencyrequirement string

// List of Field-eEzsignsignatureDependencyrequirement
const (
	ALL_OF FieldEEzsignsignatureDependencyrequirement = "AllOf"
	ANY_OF FieldEEzsignsignatureDependencyrequirement = "AnyOf"
)

// All allowed values of FieldEEzsignsignatureDependencyrequirement enum
var AllowedFieldEEzsignsignatureDependencyrequirementEnumValues = []FieldEEzsignsignatureDependencyrequirement{
	"AllOf",
	"AnyOf",
}

func (v *FieldEEzsignsignatureDependencyrequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignsignatureDependencyrequirement(value)
	for _, existing := range AllowedFieldEEzsignsignatureDependencyrequirementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignsignatureDependencyrequirement", value)
}

// NewFieldEEzsignsignatureDependencyrequirementFromValue returns a pointer to a valid FieldEEzsignsignatureDependencyrequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignsignatureDependencyrequirementFromValue(v string) (*FieldEEzsignsignatureDependencyrequirement, error) {
	ev := FieldEEzsignsignatureDependencyrequirement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignsignatureDependencyrequirement: valid values are %v", v, AllowedFieldEEzsignsignatureDependencyrequirementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignsignatureDependencyrequirement) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignsignatureDependencyrequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignsignatureDependencyrequirement value
func (v FieldEEzsignsignatureDependencyrequirement) Ptr() *FieldEEzsignsignatureDependencyrequirement {
	return &v
}

type NullableFieldEEzsignsignatureDependencyrequirement struct {
	value *FieldEEzsignsignatureDependencyrequirement
	isSet bool
}

func (v NullableFieldEEzsignsignatureDependencyrequirement) Get() *FieldEEzsignsignatureDependencyrequirement {
	return v.value
}

func (v *NullableFieldEEzsignsignatureDependencyrequirement) Set(val *FieldEEzsignsignatureDependencyrequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignsignatureDependencyrequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignsignatureDependencyrequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignsignatureDependencyrequirement(val *FieldEEzsignsignatureDependencyrequirement) *NullableFieldEEzsignsignatureDependencyrequirement {
	return &NullableFieldEEzsignsignatureDependencyrequirement{value: val, isSet: true}
}

func (v NullableFieldEEzsignsignatureDependencyrequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignsignatureDependencyrequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
