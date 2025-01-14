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

// FieldEEzsigntemplatesignatureType The type of signature.  1. **Acknowledgement** is for an acknowledgment of receipt. 2. **City** is to request the city where the document is signed. 3. **Handwritten** is for a handwritten kind of signature where users needs to \"draw\" their signature on screen. **Deprecated** 4. **Initials** is a simple \"click to add initials\" block. 5. **Name** is a simple \"Click to sign\" block. This is the most common block of signature. **Deprecated** 6. **NameReason** is to ask for a signing reason.  7. **Attachments** is to ask for files as attachment that may be validate in another step.     8. **Signature** is replacing **Name** and **Handwritten**. Will be normal or handwritten once signed
type FieldEEzsigntemplatesignatureType string

// List of Field-eEzsigntemplatesignatureType
const (
	ACKNOWLEDGEMENT FieldEEzsigntemplatesignatureType = "Acknowledgement"
	CITY FieldEEzsigntemplatesignatureType = "City"
	HANDWRITTEN FieldEEzsigntemplatesignatureType = "Handwritten"
	INITIALS FieldEEzsigntemplatesignatureType = "Initials"
	NAME FieldEEzsigntemplatesignatureType = "Name"
	NAME_REASON FieldEEzsigntemplatesignatureType = "NameReason"
	ATTACHMENTS FieldEEzsigntemplatesignatureType = "Attachments"
	FIELD_TEXT FieldEEzsigntemplatesignatureType = "FieldText"
	FIELD_TEXTAREA FieldEEzsigntemplatesignatureType = "FieldTextarea"
	CONSULTATION FieldEEzsigntemplatesignatureType = "Consultation"
	SIGNATURE FieldEEzsigntemplatesignatureType = "Signature"
)

// All allowed values of FieldEEzsigntemplatesignatureType enum
var AllowedFieldEEzsigntemplatesignatureTypeEnumValues = []FieldEEzsigntemplatesignatureType{
	"Acknowledgement",
	"City",
	"Handwritten",
	"Initials",
	"Name",
	"NameReason",
	"Attachments",
	"FieldText",
	"FieldTextarea",
	"Consultation",
	"Signature",
}

func (v *FieldEEzsigntemplatesignatureType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplatesignatureType(value)
	for _, existing := range AllowedFieldEEzsigntemplatesignatureTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplatesignatureType", value)
}

// NewFieldEEzsigntemplatesignatureTypeFromValue returns a pointer to a valid FieldEEzsigntemplatesignatureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplatesignatureTypeFromValue(v string) (*FieldEEzsigntemplatesignatureType, error) {
	ev := FieldEEzsigntemplatesignatureType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplatesignatureType: valid values are %v", v, AllowedFieldEEzsigntemplatesignatureTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplatesignatureType) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplatesignatureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplatesignatureType value
func (v FieldEEzsigntemplatesignatureType) Ptr() *FieldEEzsigntemplatesignatureType {
	return &v
}

type NullableFieldEEzsigntemplatesignatureType struct {
	value *FieldEEzsigntemplatesignatureType
	isSet bool
}

func (v NullableFieldEEzsigntemplatesignatureType) Get() *FieldEEzsigntemplatesignatureType {
	return v.value
}

func (v *NullableFieldEEzsigntemplatesignatureType) Set(val *FieldEEzsigntemplatesignatureType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplatesignatureType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplatesignatureType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplatesignatureType(val *FieldEEzsigntemplatesignatureType) *NullableFieldEEzsigntemplatesignatureType {
	return &NullableFieldEEzsigntemplatesignatureType{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplatesignatureType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplatesignatureType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

