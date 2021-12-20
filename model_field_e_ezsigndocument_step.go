/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsigndocumentStep The signature step of the Ezsigndocument.
type FieldEEzsigndocumentStep string

// List of Field-eEzsigndocumentStep
const (
	UNSENT FieldEEzsigndocumentStep = "Unsent"
	UNSIGNED FieldEEzsigndocumentStep = "Unsigned"
	PARTIALLY_SIGNED FieldEEzsigndocumentStep = "PartiallySigned"
	DECLINED_TO_SIGN FieldEEzsigndocumentStep = "DeclinedToSign"
	PREMATURELY_ENDED FieldEEzsigndocumentStep = "PrematurelyEnded"
	COMPLETED FieldEEzsigndocumentStep = "Completed"
)

// All allowed values of FieldEEzsigndocumentStep enum
var AllowedFieldEEzsigndocumentStepEnumValues = []FieldEEzsigndocumentStep{
	"Unsent",
	"Unsigned",
	"PartiallySigned",
	"DeclinedToSign",
	"PrematurelyEnded",
	"Completed",
}

func (v *FieldEEzsigndocumentStep) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigndocumentStep(value)
	for _, existing := range AllowedFieldEEzsigndocumentStepEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigndocumentStep", value)
}

// NewFieldEEzsigndocumentStepFromValue returns a pointer to a valid FieldEEzsigndocumentStep
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigndocumentStepFromValue(v string) (*FieldEEzsigndocumentStep, error) {
	ev := FieldEEzsigndocumentStep(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigndocumentStep: valid values are %v", v, AllowedFieldEEzsigndocumentStepEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigndocumentStep) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigndocumentStepEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigndocumentStep value
func (v FieldEEzsigndocumentStep) Ptr() *FieldEEzsigndocumentStep {
	return &v
}

type NullableFieldEEzsigndocumentStep struct {
	value *FieldEEzsigndocumentStep
	isSet bool
}

func (v NullableFieldEEzsigndocumentStep) Get() *FieldEEzsigndocumentStep {
	return v.value
}

func (v *NullableFieldEEzsigndocumentStep) Set(val *FieldEEzsigndocumentStep) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigndocumentStep) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigndocumentStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigndocumentStep(val *FieldEEzsigndocumentStep) *NullableFieldEEzsigndocumentStep {
	return &NullableFieldEEzsigndocumentStep{value: val, isSet: true}
}

func (v NullableFieldEEzsigndocumentStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigndocumentStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

