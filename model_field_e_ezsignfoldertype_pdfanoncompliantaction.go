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

// FieldEEzsignfoldertypePdfanoncompliantaction The Action to do if the PDFA is non compliant of the Ezsignfolder type.  * **Reject** is for rejecting the document * **Convert** is for converting to the fkiPdfalevelIDConvert configured
type FieldEEzsignfoldertypePdfanoncompliantaction string

// List of Field-eEzsignfoldertypePdfanoncompliantaction
const (
	REJECT FieldEEzsignfoldertypePdfanoncompliantaction = "Reject"
	CONVERT FieldEEzsignfoldertypePdfanoncompliantaction = "Convert"
)

// All allowed values of FieldEEzsignfoldertypePdfanoncompliantaction enum
var AllowedFieldEEzsignfoldertypePdfanoncompliantactionEnumValues = []FieldEEzsignfoldertypePdfanoncompliantaction{
	"Reject",
	"Convert",
}

func (v *FieldEEzsignfoldertypePdfanoncompliantaction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfoldertypePdfanoncompliantaction(value)
	for _, existing := range AllowedFieldEEzsignfoldertypePdfanoncompliantactionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfoldertypePdfanoncompliantaction", value)
}

// NewFieldEEzsignfoldertypePdfanoncompliantactionFromValue returns a pointer to a valid FieldEEzsignfoldertypePdfanoncompliantaction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignfoldertypePdfanoncompliantactionFromValue(v string) (*FieldEEzsignfoldertypePdfanoncompliantaction, error) {
	ev := FieldEEzsignfoldertypePdfanoncompliantaction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignfoldertypePdfanoncompliantaction: valid values are %v", v, AllowedFieldEEzsignfoldertypePdfanoncompliantactionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignfoldertypePdfanoncompliantaction) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignfoldertypePdfanoncompliantactionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignfoldertypePdfanoncompliantaction value
func (v FieldEEzsignfoldertypePdfanoncompliantaction) Ptr() *FieldEEzsignfoldertypePdfanoncompliantaction {
	return &v
}

type NullableFieldEEzsignfoldertypePdfanoncompliantaction struct {
	value *FieldEEzsignfoldertypePdfanoncompliantaction
	isSet bool
}

func (v NullableFieldEEzsignfoldertypePdfanoncompliantaction) Get() *FieldEEzsignfoldertypePdfanoncompliantaction {
	return v.value
}

func (v *NullableFieldEEzsignfoldertypePdfanoncompliantaction) Set(val *FieldEEzsignfoldertypePdfanoncompliantaction) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfoldertypePdfanoncompliantaction) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfoldertypePdfanoncompliantaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfoldertypePdfanoncompliantaction(val *FieldEEzsignfoldertypePdfanoncompliantaction) *NullableFieldEEzsignfoldertypePdfanoncompliantaction {
	return &NullableFieldEEzsignfoldertypePdfanoncompliantaction{value: val, isSet: true}
}

func (v NullableFieldEEzsignfoldertypePdfanoncompliantaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfoldertypePdfanoncompliantaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

