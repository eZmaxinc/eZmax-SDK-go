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

// FieldEUserEzsignprepaid Subscription level when a user has a Prepaid subscription.
type FieldEUserEzsignprepaid string

// List of Field-eUserEzsignprepaid
const (
	NO FieldEUserEzsignprepaid = "No"
	BASIC FieldEUserEzsignprepaid = "Basic"
	STANDARD FieldEUserEzsignprepaid = "Standard"
	PRO FieldEUserEzsignprepaid = "Pro"
)

// All allowed values of FieldEUserEzsignprepaid enum
var AllowedFieldEUserEzsignprepaidEnumValues = []FieldEUserEzsignprepaid{
	"No",
	"Basic",
	"Standard",
	"Pro",
}

func (v *FieldEUserEzsignprepaid) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEUserEzsignprepaid(value)
	for _, existing := range AllowedFieldEUserEzsignprepaidEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEUserEzsignprepaid", value)
}

// NewFieldEUserEzsignprepaidFromValue returns a pointer to a valid FieldEUserEzsignprepaid
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEUserEzsignprepaidFromValue(v string) (*FieldEUserEzsignprepaid, error) {
	ev := FieldEUserEzsignprepaid(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEUserEzsignprepaid: valid values are %v", v, AllowedFieldEUserEzsignprepaidEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEUserEzsignprepaid) IsValid() bool {
	for _, existing := range AllowedFieldEUserEzsignprepaidEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eUserEzsignprepaid value
func (v FieldEUserEzsignprepaid) Ptr() *FieldEUserEzsignprepaid {
	return &v
}

type NullableFieldEUserEzsignprepaid struct {
	value *FieldEUserEzsignprepaid
	isSet bool
}

func (v NullableFieldEUserEzsignprepaid) Get() *FieldEUserEzsignprepaid {
	return v.value
}

func (v *NullableFieldEUserEzsignprepaid) Set(val *FieldEUserEzsignprepaid) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEUserEzsignprepaid) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEUserEzsignprepaid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEUserEzsignprepaid(val *FieldEUserEzsignprepaid) *NullableFieldEUserEzsignprepaid {
	return &NullableFieldEUserEzsignprepaid{value: val, isSet: true}
}

func (v NullableFieldEUserEzsignprepaid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEUserEzsignprepaid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

