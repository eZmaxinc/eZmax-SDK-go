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

// FieldEBrandingLogointerface The logo for the Branding. Select the value 'Default' if you want to use the default logo and delete the custom one if you used one
type FieldEBrandingLogointerface string

// List of Field-eBrandingLogointerface
const (
	DEFAULT FieldEBrandingLogointerface = "Default"
	JPEG FieldEBrandingLogointerface = "JPEG"
	PNG FieldEBrandingLogointerface = "PNG"
)

// All allowed values of FieldEBrandingLogointerface enum
var AllowedFieldEBrandingLogointerfaceEnumValues = []FieldEBrandingLogointerface{
	"Default",
	"JPEG",
	"PNG",
}

func (v *FieldEBrandingLogointerface) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEBrandingLogointerface(value)
	for _, existing := range AllowedFieldEBrandingLogointerfaceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEBrandingLogointerface", value)
}

// NewFieldEBrandingLogointerfaceFromValue returns a pointer to a valid FieldEBrandingLogointerface
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEBrandingLogointerfaceFromValue(v string) (*FieldEBrandingLogointerface, error) {
	ev := FieldEBrandingLogointerface(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEBrandingLogointerface: valid values are %v", v, AllowedFieldEBrandingLogointerfaceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEBrandingLogointerface) IsValid() bool {
	for _, existing := range AllowedFieldEBrandingLogointerfaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eBrandingLogointerface value
func (v FieldEBrandingLogointerface) Ptr() *FieldEBrandingLogointerface {
	return &v
}

type NullableFieldEBrandingLogointerface struct {
	value *FieldEBrandingLogointerface
	isSet bool
}

func (v NullableFieldEBrandingLogointerface) Get() *FieldEBrandingLogointerface {
	return v.value
}

func (v *NullableFieldEBrandingLogointerface) Set(val *FieldEBrandingLogointerface) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEBrandingLogointerface) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEBrandingLogointerface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEBrandingLogointerface(val *FieldEBrandingLogointerface) *NullableFieldEBrandingLogointerface {
	return &NullableFieldEBrandingLogointerface{value: val, isSet: true}
}

func (v NullableFieldEBrandingLogointerface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEBrandingLogointerface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

