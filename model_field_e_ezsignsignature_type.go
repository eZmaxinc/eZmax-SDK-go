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

// FieldEEzsignsignatureType The type of signature.  1. **Acknowledgement** is for an acknowledgment of receipt. 2. **City** is to request the city where the document is signed. 3. **Handwritten** is for a handwritten kind of signature where users needs to \"draw\" their signature on screen. **DEPRECATED** 4. **Initials** is a simple \"click to add initials\" block. 5. **Name** is a simple \"Click to sign\" block. This is the most common block of signature. **DEPRECATED** 6. **NameReason** is to ask for a signing reason.  7. **Attachments** is to ask for files as attachment that may be validate in another step.  8. **FieldText** is to ask for a short text. 9. **Fieldtextarea** is to ask for a text 10. **Signature** is the type replacing **Name** and **Handwritten** and will support a font or svg
type FieldEEzsignsignatureType string

// List of Field-eEzsignsignatureType
const (
	ACKNOWLEDGEMENT FieldEEzsignsignatureType = "Acknowledgement"
	CITY FieldEEzsignsignatureType = "City"
	HANDWRITTEN FieldEEzsignsignatureType = "Handwritten"
	INITIALS FieldEEzsignsignatureType = "Initials"
	NAME FieldEEzsignsignatureType = "Name"
	NAME_REASON FieldEEzsignsignatureType = "NameReason"
	ATTACHMENTS FieldEEzsignsignatureType = "Attachments"
	ATTACHMENTS_CONFIRMATION FieldEEzsignsignatureType = "AttachmentsConfirmation"
	FIELD_TEXT FieldEEzsignsignatureType = "FieldText"
	FIELD_TEXTAREA FieldEEzsignsignatureType = "FieldTextarea"
	CONSULTATION FieldEEzsignsignatureType = "Consultation"
	SIGNATURE FieldEEzsignsignatureType = "Signature"
)

// All allowed values of FieldEEzsignsignatureType enum
var AllowedFieldEEzsignsignatureTypeEnumValues = []FieldEEzsignsignatureType{
	"Acknowledgement",
	"City",
	"Handwritten",
	"Initials",
	"Name",
	"NameReason",
	"Attachments",
	"AttachmentsConfirmation",
	"FieldText",
	"FieldTextarea",
	"Consultation",
	"Signature",
}

func (v *FieldEEzsignsignatureType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignsignatureType(value)
	for _, existing := range AllowedFieldEEzsignsignatureTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignsignatureType", value)
}

// NewFieldEEzsignsignatureTypeFromValue returns a pointer to a valid FieldEEzsignsignatureType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignsignatureTypeFromValue(v string) (*FieldEEzsignsignatureType, error) {
	ev := FieldEEzsignsignatureType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignsignatureType: valid values are %v", v, AllowedFieldEEzsignsignatureTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignsignatureType) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignsignatureTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignsignatureType value
func (v FieldEEzsignsignatureType) Ptr() *FieldEEzsignsignatureType {
	return &v
}

type NullableFieldEEzsignsignatureType struct {
	value *FieldEEzsignsignatureType
	isSet bool
}

func (v NullableFieldEEzsignsignatureType) Get() *FieldEEzsignsignatureType {
	return v.value
}

func (v *NullableFieldEEzsignsignatureType) Set(val *FieldEEzsignsignatureType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignsignatureType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignsignatureType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignsignatureType(val *FieldEEzsignsignatureType) *NullableFieldEEzsignsignatureType {
	return &NullableFieldEEzsignsignatureType{value: val, isSet: true}
}

func (v NullableFieldEEzsignsignatureType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignsignatureType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

