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

// FieldEPaymenttermType The type of the Paymentterm
type FieldEPaymenttermType string

// List of Field-ePaymenttermType
const (
	DAYS FieldEPaymenttermType = "Days"
	DAYOFTHEMONTH FieldEPaymenttermType = "Dayofthemonth"
)

// All allowed values of FieldEPaymenttermType enum
var AllowedFieldEPaymenttermTypeEnumValues = []FieldEPaymenttermType{
	"Days",
	"Dayofthemonth",
}

func (v *FieldEPaymenttermType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEPaymenttermType(value)
	for _, existing := range AllowedFieldEPaymenttermTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEPaymenttermType", value)
}

// NewFieldEPaymenttermTypeFromValue returns a pointer to a valid FieldEPaymenttermType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEPaymenttermTypeFromValue(v string) (*FieldEPaymenttermType, error) {
	ev := FieldEPaymenttermType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEPaymenttermType: valid values are %v", v, AllowedFieldEPaymenttermTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEPaymenttermType) IsValid() bool {
	for _, existing := range AllowedFieldEPaymenttermTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-ePaymenttermType value
func (v FieldEPaymenttermType) Ptr() *FieldEPaymenttermType {
	return &v
}

type NullableFieldEPaymenttermType struct {
	value *FieldEPaymenttermType
	isSet bool
}

func (v NullableFieldEPaymenttermType) Get() *FieldEPaymenttermType {
	return v.value
}

func (v *NullableFieldEPaymenttermType) Set(val *FieldEPaymenttermType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEPaymenttermType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEPaymenttermType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEPaymenttermType(val *FieldEPaymenttermType) *NullableFieldEPaymenttermType {
	return &NullableFieldEPaymenttermType{value: val, isSet: true}
}

func (v NullableFieldEPaymenttermType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEPaymenttermType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

