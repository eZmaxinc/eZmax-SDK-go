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

// FieldEUserEzsignsendreminderfrequency Frequency at which reminders will be sent to signers that haven't signed the documents
type FieldEUserEzsignsendreminderfrequency string

// List of Field-eUserEzsignsendreminderfrequency
const (
	NONE FieldEUserEzsignsendreminderfrequency = "None"
	DAILY FieldEUserEzsignsendreminderfrequency = "Daily"
	WEEKLY FieldEUserEzsignsendreminderfrequency = "Weekly"
)

// All allowed values of FieldEUserEzsignsendreminderfrequency enum
var AllowedFieldEUserEzsignsendreminderfrequencyEnumValues = []FieldEUserEzsignsendreminderfrequency{
	"None",
	"Daily",
	"Weekly",
}

func (v *FieldEUserEzsignsendreminderfrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEUserEzsignsendreminderfrequency(value)
	for _, existing := range AllowedFieldEUserEzsignsendreminderfrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEUserEzsignsendreminderfrequency", value)
}

// NewFieldEUserEzsignsendreminderfrequencyFromValue returns a pointer to a valid FieldEUserEzsignsendreminderfrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEUserEzsignsendreminderfrequencyFromValue(v string) (*FieldEUserEzsignsendreminderfrequency, error) {
	ev := FieldEUserEzsignsendreminderfrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEUserEzsignsendreminderfrequency: valid values are %v", v, AllowedFieldEUserEzsignsendreminderfrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEUserEzsignsendreminderfrequency) IsValid() bool {
	for _, existing := range AllowedFieldEUserEzsignsendreminderfrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eUserEzsignsendreminderfrequency value
func (v FieldEUserEzsignsendreminderfrequency) Ptr() *FieldEUserEzsignsendreminderfrequency {
	return &v
}

type NullableFieldEUserEzsignsendreminderfrequency struct {
	value *FieldEUserEzsignsendreminderfrequency
	isSet bool
}

func (v NullableFieldEUserEzsignsendreminderfrequency) Get() *FieldEUserEzsignsendreminderfrequency {
	return v.value
}

func (v *NullableFieldEUserEzsignsendreminderfrequency) Set(val *FieldEUserEzsignsendreminderfrequency) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEUserEzsignsendreminderfrequency) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEUserEzsignsendreminderfrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEUserEzsignsendreminderfrequency(val *FieldEUserEzsignsendreminderfrequency) *NullableFieldEUserEzsignsendreminderfrequency {
	return &NullableFieldEUserEzsignsendreminderfrequency{value: val, isSet: true}
}

func (v NullableFieldEUserEzsignsendreminderfrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEUserEzsignsendreminderfrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

