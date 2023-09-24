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

// FieldEEzsignsignatureFont The font of the signature. This can only be set if eEzsignsignatureType is **Name** or **Initials**
type FieldEEzsignsignatureFont string

// List of Field-eEzsignsignatureFont
const (
	NORMAL FieldEEzsignsignatureFont = "Normal"
	CURSIVE FieldEEzsignsignatureFont = "Cursive"
)

// All allowed values of FieldEEzsignsignatureFont enum
var AllowedFieldEEzsignsignatureFontEnumValues = []FieldEEzsignsignatureFont{
	"Normal",
	"Cursive",
}

func (v *FieldEEzsignsignatureFont) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignsignatureFont(value)
	for _, existing := range AllowedFieldEEzsignsignatureFontEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignsignatureFont", value)
}

// NewFieldEEzsignsignatureFontFromValue returns a pointer to a valid FieldEEzsignsignatureFont
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignsignatureFontFromValue(v string) (*FieldEEzsignsignatureFont, error) {
	ev := FieldEEzsignsignatureFont(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignsignatureFont: valid values are %v", v, AllowedFieldEEzsignsignatureFontEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignsignatureFont) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignsignatureFontEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignsignatureFont value
func (v FieldEEzsignsignatureFont) Ptr() *FieldEEzsignsignatureFont {
	return &v
}

type NullableFieldEEzsignsignatureFont struct {
	value *FieldEEzsignsignatureFont
	isSet bool
}

func (v NullableFieldEEzsignsignatureFont) Get() *FieldEEzsignsignatureFont {
	return v.value
}

func (v *NullableFieldEEzsignsignatureFont) Set(val *FieldEEzsignsignatureFont) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignsignatureFont) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignsignatureFont) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignsignatureFont(val *FieldEEzsignsignatureFont) *NullableFieldEEzsignsignatureFont {
	return &NullableFieldEEzsignsignatureFont{value: val, isSet: true}
}

func (v NullableFieldEEzsignsignatureFont) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignsignatureFont) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
