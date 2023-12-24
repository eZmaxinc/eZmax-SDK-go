/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEAttachmentType The type of the Attachment
type FieldEAttachmentType string

// List of Field-eAttachmentType
const (
	OTHER FieldEAttachmentType = "Other"
	PDF FieldEAttachmentType = "Pdf"
	PDF_GENERATED FieldEAttachmentType = "PdfGenerated"
	PDF_SCANNED FieldEAttachmentType = "PdfScanned"
	EZSIGN FieldEAttachmentType = "Ezsign"
)

// All allowed values of FieldEAttachmentType enum
var AllowedFieldEAttachmentTypeEnumValues = []FieldEAttachmentType{
	"Other",
	"Pdf",
	"PdfGenerated",
	"PdfScanned",
	"Ezsign",
}

func (v *FieldEAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEAttachmentType(value)
	for _, existing := range AllowedFieldEAttachmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEAttachmentType", value)
}

// NewFieldEAttachmentTypeFromValue returns a pointer to a valid FieldEAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEAttachmentTypeFromValue(v string) (*FieldEAttachmentType, error) {
	ev := FieldEAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEAttachmentType: valid values are %v", v, AllowedFieldEAttachmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEAttachmentType) IsValid() bool {
	for _, existing := range AllowedFieldEAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eAttachmentType value
func (v FieldEAttachmentType) Ptr() *FieldEAttachmentType {
	return &v
}

type NullableFieldEAttachmentType struct {
	value *FieldEAttachmentType
	isSet bool
}

func (v NullableFieldEAttachmentType) Get() *FieldEAttachmentType {
	return v.value
}

func (v *NullableFieldEAttachmentType) Set(val *FieldEAttachmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEAttachmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEAttachmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEAttachmentType(val *FieldEAttachmentType) *NullableFieldEAttachmentType {
	return &NullableFieldEAttachmentType{value: val, isSet: true}
}

func (v NullableFieldEAttachmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEAttachmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

