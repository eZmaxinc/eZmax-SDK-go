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

// ComputedEEzsigndocumentSteptype Indicates whether the current step is a form or signatures.
type ComputedEEzsigndocumentSteptype string

// List of Computed-eEzsigndocumentSteptype
const (
	FORM ComputedEEzsigndocumentSteptype = "Form"
	SIGN ComputedEEzsigndocumentSteptype = "Sign"
	NONE ComputedEEzsigndocumentSteptype = "None"
)

// All allowed values of ComputedEEzsigndocumentSteptype enum
var AllowedComputedEEzsigndocumentSteptypeEnumValues = []ComputedEEzsigndocumentSteptype{
	"Form",
	"Sign",
	"None",
}

func (v *ComputedEEzsigndocumentSteptype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComputedEEzsigndocumentSteptype(value)
	for _, existing := range AllowedComputedEEzsigndocumentSteptypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComputedEEzsigndocumentSteptype", value)
}

// NewComputedEEzsigndocumentSteptypeFromValue returns a pointer to a valid ComputedEEzsigndocumentSteptype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputedEEzsigndocumentSteptypeFromValue(v string) (*ComputedEEzsigndocumentSteptype, error) {
	ev := ComputedEEzsigndocumentSteptype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComputedEEzsigndocumentSteptype: valid values are %v", v, AllowedComputedEEzsigndocumentSteptypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputedEEzsigndocumentSteptype) IsValid() bool {
	for _, existing := range AllowedComputedEEzsigndocumentSteptypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Computed-eEzsigndocumentSteptype value
func (v ComputedEEzsigndocumentSteptype) Ptr() *ComputedEEzsigndocumentSteptype {
	return &v
}

type NullableComputedEEzsigndocumentSteptype struct {
	value *ComputedEEzsigndocumentSteptype
	isSet bool
}

func (v NullableComputedEEzsigndocumentSteptype) Get() *ComputedEEzsigndocumentSteptype {
	return v.value
}

func (v *NullableComputedEEzsigndocumentSteptype) Set(val *ComputedEEzsigndocumentSteptype) {
	v.value = val
	v.isSet = true
}

func (v NullableComputedEEzsigndocumentSteptype) IsSet() bool {
	return v.isSet
}

func (v *NullableComputedEEzsigndocumentSteptype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputedEEzsigndocumentSteptype(val *ComputedEEzsigndocumentSteptype) *NullableComputedEEzsigndocumentSteptype {
	return &NullableComputedEEzsigndocumentSteptype{value: val, isSet: true}
}

func (v NullableComputedEEzsigndocumentSteptype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputedEEzsigndocumentSteptype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

