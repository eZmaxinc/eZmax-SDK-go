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

// FieldEEzsigntemplateglobalModule The Module for the Ezsigntemplateglobal
type FieldEEzsigntemplateglobalModule string

// List of Field-eEzsigntemplateglobalModule
const (
	ALL FieldEEzsigntemplateglobalModule = "All"
	INSCRIPTION FieldEEzsigntemplateglobalModule = "Inscription"
)

// All allowed values of FieldEEzsigntemplateglobalModule enum
var AllowedFieldEEzsigntemplateglobalModuleEnumValues = []FieldEEzsigntemplateglobalModule{
	"All",
	"Inscription",
}

func (v *FieldEEzsigntemplateglobalModule) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateglobalModule(value)
	for _, existing := range AllowedFieldEEzsigntemplateglobalModuleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateglobalModule", value)
}

// NewFieldEEzsigntemplateglobalModuleFromValue returns a pointer to a valid FieldEEzsigntemplateglobalModule
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateglobalModuleFromValue(v string) (*FieldEEzsigntemplateglobalModule, error) {
	ev := FieldEEzsigntemplateglobalModule(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateglobalModule: valid values are %v", v, AllowedFieldEEzsigntemplateglobalModuleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateglobalModule) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateglobalModuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateglobalModule value
func (v FieldEEzsigntemplateglobalModule) Ptr() *FieldEEzsigntemplateglobalModule {
	return &v
}

type NullableFieldEEzsigntemplateglobalModule struct {
	value *FieldEEzsigntemplateglobalModule
	isSet bool
}

func (v NullableFieldEEzsigntemplateglobalModule) Get() *FieldEEzsigntemplateglobalModule {
	return v.value
}

func (v *NullableFieldEEzsigntemplateglobalModule) Set(val *FieldEEzsigntemplateglobalModule) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateglobalModule) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateglobalModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateglobalModule(val *FieldEEzsigntemplateglobalModule) *NullableFieldEEzsigntemplateglobalModule {
	return &NullableFieldEEzsigntemplateglobalModule{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateglobalModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateglobalModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

