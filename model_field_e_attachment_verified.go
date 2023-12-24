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

// FieldEAttachmentVerified The verified of the Attachment
type FieldEAttachmentVerified string

// List of Field-eAttachmentVerified
const (
	NO FieldEAttachmentVerified = "No"
	YES FieldEAttachmentVerified = "Yes"
	REJECTED FieldEAttachmentVerified = "Rejected"
)

// All allowed values of FieldEAttachmentVerified enum
var AllowedFieldEAttachmentVerifiedEnumValues = []FieldEAttachmentVerified{
	"No",
	"Yes",
	"Rejected",
}

func (v *FieldEAttachmentVerified) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEAttachmentVerified(value)
	for _, existing := range AllowedFieldEAttachmentVerifiedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEAttachmentVerified", value)
}

// NewFieldEAttachmentVerifiedFromValue returns a pointer to a valid FieldEAttachmentVerified
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEAttachmentVerifiedFromValue(v string) (*FieldEAttachmentVerified, error) {
	ev := FieldEAttachmentVerified(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEAttachmentVerified: valid values are %v", v, AllowedFieldEAttachmentVerifiedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEAttachmentVerified) IsValid() bool {
	for _, existing := range AllowedFieldEAttachmentVerifiedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eAttachmentVerified value
func (v FieldEAttachmentVerified) Ptr() *FieldEAttachmentVerified {
	return &v
}

type NullableFieldEAttachmentVerified struct {
	value *FieldEAttachmentVerified
	isSet bool
}

func (v NullableFieldEAttachmentVerified) Get() *FieldEAttachmentVerified {
	return v.value
}

func (v *NullableFieldEAttachmentVerified) Set(val *FieldEAttachmentVerified) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEAttachmentVerified) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEAttachmentVerified) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEAttachmentVerified(val *FieldEAttachmentVerified) *NullableFieldEAttachmentVerified {
	return &NullableFieldEAttachmentVerified{value: val, isSet: true}
}

func (v NullableFieldEAttachmentVerified) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEAttachmentVerified) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

