/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsignfolderStep The signature step of the Ezsignfolder.
type FieldEEzsignfolderStep string

// List of Field-eEzsignfolderStep
const (
	UNSENT FieldEEzsignfolderStep = "Unsent"
	SENT FieldEEzsignfolderStep = "Sent"
	PARTIALLY_SIGNED FieldEEzsignfolderStep = "PartiallySigned"
	EXPIRED FieldEEzsignfolderStep = "Expired"
	COMPLETED FieldEEzsignfolderStep = "Completed"
	ARCHIVED FieldEEzsignfolderStep = "Archived"
)

// All allowed values of FieldEEzsignfolderStep enum
var AllowedFieldEEzsignfolderStepEnumValues = []FieldEEzsignfolderStep{
	"Unsent",
	"Sent",
	"PartiallySigned",
	"Expired",
	"Completed",
	"Archived",
}

func (v *FieldEEzsignfolderStep) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfolderStep(value)
	for _, existing := range AllowedFieldEEzsignfolderStepEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfolderStep", value)
}

// NewFieldEEzsignfolderStepFromValue returns a pointer to a valid FieldEEzsignfolderStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignfolderStepFromValue(v string) (*FieldEEzsignfolderStep, error) {
	ev := FieldEEzsignfolderStep(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignfolderStep: valid values are %v", v, AllowedFieldEEzsignfolderStepEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignfolderStep) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignfolderStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignfolderStep value
func (v FieldEEzsignfolderStep) Ptr() *FieldEEzsignfolderStep {
	return &v
}

type NullableFieldEEzsignfolderStep struct {
	value *FieldEEzsignfolderStep
	isSet bool
}

func (v NullableFieldEEzsignfolderStep) Get() *FieldEEzsignfolderStep {
	return v.value
}

func (v *NullableFieldEEzsignfolderStep) Set(val *FieldEEzsignfolderStep) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfolderStep) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfolderStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfolderStep(val *FieldEEzsignfolderStep) *NullableFieldEEzsignfolderStep {
	return &NullableFieldEEzsignfolderStep{value: val, isSet: true}
}

func (v NullableFieldEEzsignfolderStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfolderStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

