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

// FieldECustomerType The type of the Customer
type FieldECustomerType string

// List of Field-eCustomerType
const (
	NORMAL FieldECustomerType = "Normal"
	VETRX_SERVER FieldECustomerType = "Vetrx-Server"
	REWARD_ADMINISTRATION FieldECustomerType = "Reward-Administration"
	REWARD_REPRESENTATIVE FieldECustomerType = "Reward-Representative"
	REWARD_SERVER FieldECustomerType = "Reward-Server"
)

// All allowed values of FieldECustomerType enum
var AllowedFieldECustomerTypeEnumValues = []FieldECustomerType{
	"Normal",
	"Vetrx-Server",
	"Reward-Administration",
	"Reward-Representative",
	"Reward-Server",
}

func (v *FieldECustomerType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldECustomerType(value)
	for _, existing := range AllowedFieldECustomerTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldECustomerType", value)
}

// NewFieldECustomerTypeFromValue returns a pointer to a valid FieldECustomerType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldECustomerTypeFromValue(v string) (*FieldECustomerType, error) {
	ev := FieldECustomerType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldECustomerType: valid values are %v", v, AllowedFieldECustomerTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldECustomerType) IsValid() bool {
	for _, existing := range AllowedFieldECustomerTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eCustomerType value
func (v FieldECustomerType) Ptr() *FieldECustomerType {
	return &v
}

type NullableFieldECustomerType struct {
	value *FieldECustomerType
	isSet bool
}

func (v NullableFieldECustomerType) Get() *FieldECustomerType {
	return v.value
}

func (v *NullableFieldECustomerType) Set(val *FieldECustomerType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldECustomerType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldECustomerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldECustomerType(val *FieldECustomerType) *NullableFieldECustomerType {
	return &NullableFieldECustomerType{value: val, isSet: true}
}

func (v NullableFieldECustomerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldECustomerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

