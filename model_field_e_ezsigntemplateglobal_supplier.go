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

// FieldEEzsigntemplateglobalSupplier The Supplier of the Ezsigntemplateglobal
type FieldEEzsigntemplateglobalSupplier string

// List of Field-eEzsigntemplateglobalSupplier
const (
	CENTRIS FieldEEzsigntemplateglobalSupplier = "Centris"
	WEBFORMS FieldEEzsigntemplateglobalSupplier = "Webforms"
	GHACQ FieldEEzsigntemplateglobalSupplier = "GHACQ"
)

// All allowed values of FieldEEzsigntemplateglobalSupplier enum
var AllowedFieldEEzsigntemplateglobalSupplierEnumValues = []FieldEEzsigntemplateglobalSupplier{
	"Centris",
	"Webforms",
	"GHACQ",
}

func (v *FieldEEzsigntemplateglobalSupplier) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateglobalSupplier(value)
	for _, existing := range AllowedFieldEEzsigntemplateglobalSupplierEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateglobalSupplier", value)
}

// NewFieldEEzsigntemplateglobalSupplierFromValue returns a pointer to a valid FieldEEzsigntemplateglobalSupplier
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateglobalSupplierFromValue(v string) (*FieldEEzsigntemplateglobalSupplier, error) {
	ev := FieldEEzsigntemplateglobalSupplier(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateglobalSupplier: valid values are %v", v, AllowedFieldEEzsigntemplateglobalSupplierEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateglobalSupplier) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateglobalSupplierEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateglobalSupplier value
func (v FieldEEzsigntemplateglobalSupplier) Ptr() *FieldEEzsigntemplateglobalSupplier {
	return &v
}

type NullableFieldEEzsigntemplateglobalSupplier struct {
	value *FieldEEzsigntemplateglobalSupplier
	isSet bool
}

func (v NullableFieldEEzsigntemplateglobalSupplier) Get() *FieldEEzsigntemplateglobalSupplier {
	return v.value
}

func (v *NullableFieldEEzsigntemplateglobalSupplier) Set(val *FieldEEzsigntemplateglobalSupplier) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateglobalSupplier) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateglobalSupplier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateglobalSupplier(val *FieldEEzsigntemplateglobalSupplier) *NullableFieldEEzsigntemplateglobalSupplier {
	return &NullableFieldEEzsigntemplateglobalSupplier{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateglobalSupplier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateglobalSupplier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

