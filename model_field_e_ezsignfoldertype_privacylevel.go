/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// FieldEEzsignfoldertypePrivacylevel The Privacy level of the Ezsignfolder type.  * **User** is for personal folders use and cannot be shared * **Usergroup** is for shared folders and complex permission can be configured to control access
type FieldEEzsignfoldertypePrivacylevel string

// List of Field-eEzsignfoldertypePrivacylevel
const (
	USER FieldEEzsignfoldertypePrivacylevel = "User"
	USERGROUP FieldEEzsignfoldertypePrivacylevel = "Usergroup"
)

// All allowed values of FieldEEzsignfoldertypePrivacylevel enum
var AllowedFieldEEzsignfoldertypePrivacylevelEnumValues = []FieldEEzsignfoldertypePrivacylevel{
	"User",
	"Usergroup",
}

func (v *FieldEEzsignfoldertypePrivacylevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignfoldertypePrivacylevel(value)
	for _, existing := range AllowedFieldEEzsignfoldertypePrivacylevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignfoldertypePrivacylevel", value)
}

// NewFieldEEzsignfoldertypePrivacylevelFromValue returns a pointer to a valid FieldEEzsignfoldertypePrivacylevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignfoldertypePrivacylevelFromValue(v string) (*FieldEEzsignfoldertypePrivacylevel, error) {
	ev := FieldEEzsignfoldertypePrivacylevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignfoldertypePrivacylevel: valid values are %v", v, AllowedFieldEEzsignfoldertypePrivacylevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignfoldertypePrivacylevel) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignfoldertypePrivacylevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignfoldertypePrivacylevel value
func (v FieldEEzsignfoldertypePrivacylevel) Ptr() *FieldEEzsignfoldertypePrivacylevel {
	return &v
}

type NullableFieldEEzsignfoldertypePrivacylevel struct {
	value *FieldEEzsignfoldertypePrivacylevel
	isSet bool
}

func (v NullableFieldEEzsignfoldertypePrivacylevel) Get() *FieldEEzsignfoldertypePrivacylevel {
	return v.value
}

func (v *NullableFieldEEzsignfoldertypePrivacylevel) Set(val *FieldEEzsignfoldertypePrivacylevel) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignfoldertypePrivacylevel) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignfoldertypePrivacylevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignfoldertypePrivacylevel(val *FieldEEzsignfoldertypePrivacylevel) *NullableFieldEEzsignfoldertypePrivacylevel {
	return &NullableFieldEEzsignfoldertypePrivacylevel{value: val, isSet: true}
}

func (v NullableFieldEEzsignfoldertypePrivacylevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignfoldertypePrivacylevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

