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

// FieldEBrandingLogo The logo for the Branding. Select the value 'Default' if you want to use the default logo and delete the custom one if you used one
type FieldEBrandingLogo string

// List of Field-eBrandingLogo
const (
	DEFAULT FieldEBrandingLogo = "Default"
	JPEG FieldEBrandingLogo = "JPEG"
	PNG FieldEBrandingLogo = "PNG"
)

// All allowed values of FieldEBrandingLogo enum
var AllowedFieldEBrandingLogoEnumValues = []FieldEBrandingLogo{
	"Default",
	"JPEG",
	"PNG",
}

func (v *FieldEBrandingLogo) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEBrandingLogo(value)
	for _, existing := range AllowedFieldEBrandingLogoEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEBrandingLogo", value)
}

// NewFieldEBrandingLogoFromValue returns a pointer to a valid FieldEBrandingLogo
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEBrandingLogoFromValue(v string) (*FieldEBrandingLogo, error) {
	ev := FieldEBrandingLogo(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEBrandingLogo: valid values are %v", v, AllowedFieldEBrandingLogoEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEBrandingLogo) IsValid() bool {
	for _, existing := range AllowedFieldEBrandingLogoEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eBrandingLogo value
func (v FieldEBrandingLogo) Ptr() *FieldEBrandingLogo {
	return &v
}

type NullableFieldEBrandingLogo struct {
	value *FieldEBrandingLogo
	isSet bool
}

func (v NullableFieldEBrandingLogo) Get() *FieldEBrandingLogo {
	return v.value
}

func (v *NullableFieldEBrandingLogo) Set(val *FieldEBrandingLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEBrandingLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEBrandingLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEBrandingLogo(val *FieldEBrandingLogo) *NullableFieldEBrandingLogo {
	return &NullableFieldEBrandingLogo{value: val, isSet: true}
}

func (v NullableFieldEBrandingLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEBrandingLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

