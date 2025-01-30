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

// FieldEEzsignfolderCompletion Indicates if the Ezsigndocument is completed when all signatures of this Ezsigndocument were applied or when all signatures of all Ezsigndocument  were applied
type FieldEEzsignfolderCompletion string

// List of Field-eEzsignfolderCompletion
const (
	PER_EZSIGNDOCUMENT FieldEEzsignfolderCompletion = "PerEzsigndocument"
	PER_EZSIGNFOLDER FieldEEzsignfolderCompletion = "PerEzsignfolder"
)

// All allowed values of FieldEEzsignfolderCompletion enum
var AllowedFieldEEzsignfolderCompletionEnumValues = []FieldEEzsignfolderCompletion{
	"PerEzsigndocument",
	"PerEzsignfolder",
}

func (v *FieldEEzsignfolderCompletion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfolderCompletion(value)
	for _, existing := range AllowedFieldEEzsignfolderCompletionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfolderCompletion", value)
}

// NewFieldEEzsignfolderCompletionFromValue returns a pointer to a valid FieldEEzsignfolderCompletion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignfolderCompletionFromValue(v string) (*FieldEEzsignfolderCompletion, error) {
	ev := FieldEEzsignfolderCompletion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignfolderCompletion: valid values are %v", v, AllowedFieldEEzsignfolderCompletionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignfolderCompletion) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignfolderCompletionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignfolderCompletion value
func (v FieldEEzsignfolderCompletion) Ptr() *FieldEEzsignfolderCompletion {
	return &v
}

type NullableFieldEEzsignfolderCompletion struct {
	value *FieldEEzsignfolderCompletion
	isSet bool
}

func (v NullableFieldEEzsignfolderCompletion) Get() *FieldEEzsignfolderCompletion {
	return v.value
}

func (v *NullableFieldEEzsignfolderCompletion) Set(val *FieldEEzsignfolderCompletion) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfolderCompletion) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfolderCompletion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfolderCompletion(val *FieldEEzsignfolderCompletion) *NullableFieldEEzsignfolderCompletion {
	return &NullableFieldEEzsignfolderCompletion{value: val, isSet: true}
}

func (v NullableFieldEEzsignfolderCompletion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfolderCompletion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

