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

// ComputedEEzsignfolderAccess Indicates the user's access level to this folder.
type ComputedEEzsignfolderAccess string

// List of Computed-eEzsignfolderAccess
const (
	SIGNER ComputedEEzsignfolderAccess = "Signer"
	READ ComputedEEzsignfolderAccess = "Read"
	MODIFY ComputedEEzsignfolderAccess = "Modify"
	FULL ComputedEEzsignfolderAccess = "Full"
)

// All allowed values of ComputedEEzsignfolderAccess enum
var AllowedComputedEEzsignfolderAccessEnumValues = []ComputedEEzsignfolderAccess{
	"Signer",
	"Read",
	"Modify",
	"Full",
}

func (v *ComputedEEzsignfolderAccess) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ComputedEEzsignfolderAccess(value)
	for _, existing := range AllowedComputedEEzsignfolderAccessEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ComputedEEzsignfolderAccess", value)
}

// NewComputedEEzsignfolderAccessFromValue returns a pointer to a valid ComputedEEzsignfolderAccess
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewComputedEEzsignfolderAccessFromValue(v string) (*ComputedEEzsignfolderAccess, error) {
	ev := ComputedEEzsignfolderAccess(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ComputedEEzsignfolderAccess: valid values are %v", v, AllowedComputedEEzsignfolderAccessEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ComputedEEzsignfolderAccess) IsValid() bool {
	for _, existing := range AllowedComputedEEzsignfolderAccessEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Computed-eEzsignfolderAccess value
func (v ComputedEEzsignfolderAccess) Ptr() *ComputedEEzsignfolderAccess {
	return &v
}

type NullableComputedEEzsignfolderAccess struct {
	value *ComputedEEzsignfolderAccess
	isSet bool
}

func (v NullableComputedEEzsignfolderAccess) Get() *ComputedEEzsignfolderAccess {
	return v.value
}

func (v *NullableComputedEEzsignfolderAccess) Set(val *ComputedEEzsignfolderAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableComputedEEzsignfolderAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableComputedEEzsignfolderAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputedEEzsignfolderAccess(val *ComputedEEzsignfolderAccess) *NullableComputedEEzsignfolderAccess {
	return &NullableComputedEEzsignfolderAccess{value: val, isSet: true}
}

func (v NullableComputedEEzsignfolderAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputedEEzsignfolderAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

