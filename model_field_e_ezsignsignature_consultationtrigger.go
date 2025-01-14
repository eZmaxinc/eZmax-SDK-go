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

// FieldEEzsignsignatureConsultationtrigger Indicates when the “consultation” type signature must be signed.  1. **Automatic** When the document is displayed . 2. **Manual** The user must indicate that he has viewed the document.
type FieldEEzsignsignatureConsultationtrigger string

// List of Field-eEzsignsignatureConsultationtrigger
const (
	AUTOMATIC FieldEEzsignsignatureConsultationtrigger = "Automatic"
	MANUAL FieldEEzsignsignatureConsultationtrigger = "Manual"
)

// All allowed values of FieldEEzsignsignatureConsultationtrigger enum
var AllowedFieldEEzsignsignatureConsultationtriggerEnumValues = []FieldEEzsignsignatureConsultationtrigger{
	"Automatic",
	"Manual",
}

func (v *FieldEEzsignsignatureConsultationtrigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignsignatureConsultationtrigger(value)
	for _, existing := range AllowedFieldEEzsignsignatureConsultationtriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignsignatureConsultationtrigger", value)
}

// NewFieldEEzsignsignatureConsultationtriggerFromValue returns a pointer to a valid FieldEEzsignsignatureConsultationtrigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignsignatureConsultationtriggerFromValue(v string) (*FieldEEzsignsignatureConsultationtrigger, error) {
	ev := FieldEEzsignsignatureConsultationtrigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignsignatureConsultationtrigger: valid values are %v", v, AllowedFieldEEzsignsignatureConsultationtriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignsignatureConsultationtrigger) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignsignatureConsultationtriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignsignatureConsultationtrigger value
func (v FieldEEzsignsignatureConsultationtrigger) Ptr() *FieldEEzsignsignatureConsultationtrigger {
	return &v
}

type NullableFieldEEzsignsignatureConsultationtrigger struct {
	value *FieldEEzsignsignatureConsultationtrigger
	isSet bool
}

func (v NullableFieldEEzsignsignatureConsultationtrigger) Get() *FieldEEzsignsignatureConsultationtrigger {
	return v.value
}

func (v *NullableFieldEEzsignsignatureConsultationtrigger) Set(val *FieldEEzsignsignatureConsultationtrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignsignatureConsultationtrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignsignatureConsultationtrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignsignatureConsultationtrigger(val *FieldEEzsignsignatureConsultationtrigger) *NullableFieldEEzsignsignatureConsultationtrigger {
	return &NullableFieldEEzsignsignatureConsultationtrigger{value: val, isSet: true}
}

func (v NullableFieldEEzsignsignatureConsultationtrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignsignatureConsultationtrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

