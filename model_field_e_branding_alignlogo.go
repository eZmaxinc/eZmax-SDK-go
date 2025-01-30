/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEBrandingAlignlogo Alignment of logo for the Branding.
type FieldEBrandingAlignlogo string

// List of Field-eBrandingAlignlogo
const (
	CENTER FieldEBrandingAlignlogo = "Center"
	LEFT FieldEBrandingAlignlogo = "Left"
	RIGHT FieldEBrandingAlignlogo = "Right"
)

// All allowed values of FieldEBrandingAlignlogo enum
var AllowedFieldEBrandingAlignlogoEnumValues = []FieldEBrandingAlignlogo{
	"Center",
	"Left",
	"Right",
}

func (v *FieldEBrandingAlignlogo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEBrandingAlignlogo(value)
	for _, existing := range AllowedFieldEBrandingAlignlogoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEBrandingAlignlogo", value)
}

// NewFieldEBrandingAlignlogoFromValue returns a pointer to a valid FieldEBrandingAlignlogo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEBrandingAlignlogoFromValue(v string) (*FieldEBrandingAlignlogo, error) {
	ev := FieldEBrandingAlignlogo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEBrandingAlignlogo: valid values are %v", v, AllowedFieldEBrandingAlignlogoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEBrandingAlignlogo) IsValid() bool {
	for _, existing := range AllowedFieldEBrandingAlignlogoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eBrandingAlignlogo value
func (v FieldEBrandingAlignlogo) Ptr() *FieldEBrandingAlignlogo {
	return &v
}

type NullableFieldEBrandingAlignlogo struct {
	value *FieldEBrandingAlignlogo
	isSet bool
}

func (v NullableFieldEBrandingAlignlogo) Get() *FieldEBrandingAlignlogo {
	return v.value
}

func (v *NullableFieldEBrandingAlignlogo) Set(val *FieldEBrandingAlignlogo) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEBrandingAlignlogo) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEBrandingAlignlogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEBrandingAlignlogo(val *FieldEBrandingAlignlogo) *NullableFieldEBrandingAlignlogo {
	return &NullableFieldEBrandingAlignlogo{value: val, isSet: true}
}

func (v NullableFieldEBrandingAlignlogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEBrandingAlignlogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

