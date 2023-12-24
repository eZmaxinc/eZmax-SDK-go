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

// FieldEAttachmentPrivacy The privacy of the Attachment
type FieldEAttachmentPrivacy string

// List of Field-eAttachmentPrivacy
const (
	ALL FieldEAttachmentPrivacy = "All"
	INSCRIPTOR FieldEAttachmentPrivacy = "Inscriptor"
	SELLER FieldEAttachmentPrivacy = "Seller"
	ADMINISTRATION FieldEAttachmentPrivacy = "Administration"
	CREATOR FieldEAttachmentPrivacy = "Creator"
	SPECIFICUSER FieldEAttachmentPrivacy = "Specificuser"
)

// All allowed values of FieldEAttachmentPrivacy enum
var AllowedFieldEAttachmentPrivacyEnumValues = []FieldEAttachmentPrivacy{
	"All",
	"Inscriptor",
	"Seller",
	"Administration",
	"Creator",
	"Specificuser",
}

func (v *FieldEAttachmentPrivacy) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEAttachmentPrivacy(value)
	for _, existing := range AllowedFieldEAttachmentPrivacyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEAttachmentPrivacy", value)
}

// NewFieldEAttachmentPrivacyFromValue returns a pointer to a valid FieldEAttachmentPrivacy
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEAttachmentPrivacyFromValue(v string) (*FieldEAttachmentPrivacy, error) {
	ev := FieldEAttachmentPrivacy(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEAttachmentPrivacy: valid values are %v", v, AllowedFieldEAttachmentPrivacyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEAttachmentPrivacy) IsValid() bool {
	for _, existing := range AllowedFieldEAttachmentPrivacyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eAttachmentPrivacy value
func (v FieldEAttachmentPrivacy) Ptr() *FieldEAttachmentPrivacy {
	return &v
}

type NullableFieldEAttachmentPrivacy struct {
	value *FieldEAttachmentPrivacy
	isSet bool
}

func (v NullableFieldEAttachmentPrivacy) Get() *FieldEAttachmentPrivacy {
	return v.value
}

func (v *NullableFieldEAttachmentPrivacy) Set(val *FieldEAttachmentPrivacy) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEAttachmentPrivacy) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEAttachmentPrivacy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEAttachmentPrivacy(val *FieldEAttachmentPrivacy) *NullableFieldEAttachmentPrivacy {
	return &NullableFieldEAttachmentPrivacy{value: val, isSet: true}
}

func (v NullableFieldEAttachmentPrivacy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEAttachmentPrivacy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

