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

// FieldECommunicationImportance The importance of the Communication
type FieldECommunicationImportance string

// List of Field-eCommunicationImportance
const (
	HIGH FieldECommunicationImportance = "High"
	NORMAL FieldECommunicationImportance = "Normal"
	LOW FieldECommunicationImportance = "Low"
)

// All allowed values of FieldECommunicationImportance enum
var AllowedFieldECommunicationImportanceEnumValues = []FieldECommunicationImportance{
	"High",
	"Normal",
	"Low",
}

func (v *FieldECommunicationImportance) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldECommunicationImportance(value)
	for _, existing := range AllowedFieldECommunicationImportanceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldECommunicationImportance", value)
}

// NewFieldECommunicationImportanceFromValue returns a pointer to a valid FieldECommunicationImportance
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldECommunicationImportanceFromValue(v string) (*FieldECommunicationImportance, error) {
	ev := FieldECommunicationImportance(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldECommunicationImportance: valid values are %v", v, AllowedFieldECommunicationImportanceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldECommunicationImportance) IsValid() bool {
	for _, existing := range AllowedFieldECommunicationImportanceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eCommunicationImportance value
func (v FieldECommunicationImportance) Ptr() *FieldECommunicationImportance {
	return &v
}

type NullableFieldECommunicationImportance struct {
	value *FieldECommunicationImportance
	isSet bool
}

func (v NullableFieldECommunicationImportance) Get() *FieldECommunicationImportance {
	return v.value
}

func (v *NullableFieldECommunicationImportance) Set(val *FieldECommunicationImportance) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldECommunicationImportance) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldECommunicationImportance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldECommunicationImportance(val *FieldECommunicationImportance) *NullableFieldECommunicationImportance {
	return &NullableFieldECommunicationImportance{value: val, isSet: true}
}

func (v NullableFieldECommunicationImportance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldECommunicationImportance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

