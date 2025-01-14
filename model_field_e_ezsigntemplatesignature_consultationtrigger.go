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

// FieldEEzsigntemplatesignatureConsultationtrigger Indicates when the “consultation” type signature must be signed.  1. **Automatic** When the document is displayed . 2. **Manual** The user must indicate that he has viewed the document.
type FieldEEzsigntemplatesignatureConsultationtrigger string

// List of Field-eEzsigntemplatesignatureConsultationtrigger
const (
	AUTOMATIC FieldEEzsigntemplatesignatureConsultationtrigger = "Automatic"
	MANUAL FieldEEzsigntemplatesignatureConsultationtrigger = "Manual"
)

// All allowed values of FieldEEzsigntemplatesignatureConsultationtrigger enum
var AllowedFieldEEzsigntemplatesignatureConsultationtriggerEnumValues = []FieldEEzsigntemplatesignatureConsultationtrigger{
	"Automatic",
	"Manual",
}

func (v *FieldEEzsigntemplatesignatureConsultationtrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplatesignatureConsultationtrigger(value)
	for _, existing := range AllowedFieldEEzsigntemplatesignatureConsultationtriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplatesignatureConsultationtrigger", value)
}

// NewFieldEEzsigntemplatesignatureConsultationtriggerFromValue returns a pointer to a valid FieldEEzsigntemplatesignatureConsultationtrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplatesignatureConsultationtriggerFromValue(v string) (*FieldEEzsigntemplatesignatureConsultationtrigger, error) {
	ev := FieldEEzsigntemplatesignatureConsultationtrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplatesignatureConsultationtrigger: valid values are %v", v, AllowedFieldEEzsigntemplatesignatureConsultationtriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplatesignatureConsultationtrigger) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplatesignatureConsultationtriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplatesignatureConsultationtrigger value
func (v FieldEEzsigntemplatesignatureConsultationtrigger) Ptr() *FieldEEzsigntemplatesignatureConsultationtrigger {
	return &v
}

type NullableFieldEEzsigntemplatesignatureConsultationtrigger struct {
	value *FieldEEzsigntemplatesignatureConsultationtrigger
	isSet bool
}

func (v NullableFieldEEzsigntemplatesignatureConsultationtrigger) Get() *FieldEEzsigntemplatesignatureConsultationtrigger {
	return v.value
}

func (v *NullableFieldEEzsigntemplatesignatureConsultationtrigger) Set(val *FieldEEzsigntemplatesignatureConsultationtrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplatesignatureConsultationtrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplatesignatureConsultationtrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplatesignatureConsultationtrigger(val *FieldEEzsigntemplatesignatureConsultationtrigger) *NullableFieldEEzsigntemplatesignatureConsultationtrigger {
	return &NullableFieldEEzsigntemplatesignatureConsultationtrigger{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplatesignatureConsultationtrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplatesignatureConsultationtrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

