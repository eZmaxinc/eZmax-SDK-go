/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsignfolderSendreminderfrequency Frequency at which reminders will be sent to signers that haven't signed the documents
type FieldEEzsignfolderSendreminderfrequency string

// List of Field-eEzsignfolderSendreminderfrequency
const (
	NONE FieldEEzsignfolderSendreminderfrequency = "None"
	DAILY FieldEEzsignfolderSendreminderfrequency = "Daily"
	WEEKLY FieldEEzsignfolderSendreminderfrequency = "Weekly"
)

var allowedFieldEEzsignfolderSendreminderfrequencyEnumValues = []FieldEEzsignfolderSendreminderfrequency{
	"None",
	"Daily",
	"Weekly",
}

func (v *FieldEEzsignfolderSendreminderfrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfolderSendreminderfrequency(value)
	for _, existing := range allowedFieldEEzsignfolderSendreminderfrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfolderSendreminderfrequency", value)
}

// NewFieldEEzsignfolderSendreminderfrequencyFromValue returns a pointer to a valid FieldEEzsignfolderSendreminderfrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignfolderSendreminderfrequencyFromValue(v string) (*FieldEEzsignfolderSendreminderfrequency, error) {
	ev := FieldEEzsignfolderSendreminderfrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignfolderSendreminderfrequency: valid values are %v", v, allowedFieldEEzsignfolderSendreminderfrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignfolderSendreminderfrequency) IsValid() bool {
	for _, existing := range allowedFieldEEzsignfolderSendreminderfrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignfolderSendreminderfrequency value
func (v FieldEEzsignfolderSendreminderfrequency) Ptr() *FieldEEzsignfolderSendreminderfrequency {
	return &v
}

type NullableFieldEEzsignfolderSendreminderfrequency struct {
	value *FieldEEzsignfolderSendreminderfrequency
	isSet bool
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) Get() *FieldEEzsignfolderSendreminderfrequency {
	return v.value
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) Set(val *FieldEEzsignfolderSendreminderfrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfolderSendreminderfrequency(val *FieldEEzsignfolderSendreminderfrequency) *NullableFieldEEzsignfolderSendreminderfrequency {
	return &NullableFieldEEzsignfolderSendreminderfrequency{value: val, isSet: true}
}

func (v NullableFieldEEzsignfolderSendreminderfrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfolderSendreminderfrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

