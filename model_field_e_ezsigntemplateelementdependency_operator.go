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

// FieldEEzsigntemplateelementdependencyOperator The operator of the Ezsigntemplateelementdependency
type FieldEEzsigntemplateelementdependencyOperator string

// List of Field-eEzsigntemplateelementdependencyOperator
const (
	EQ FieldEEzsigntemplateelementdependencyOperator = "eq"
	NEQ FieldEEzsigntemplateelementdependencyOperator = "neq"
	GT FieldEEzsigntemplateelementdependencyOperator = "gt"
	GTE FieldEEzsigntemplateelementdependencyOperator = "gte"
	LT FieldEEzsigntemplateelementdependencyOperator = "lt"
	LTE FieldEEzsigntemplateelementdependencyOperator = "lte"
	IN FieldEEzsigntemplateelementdependencyOperator = "in"
	NIN FieldEEzsigntemplateelementdependencyOperator = "nin"
	RG FieldEEzsigntemplateelementdependencyOperator = "rg"
	LIKE FieldEEzsigntemplateelementdependencyOperator = "like"
	BETWEEN FieldEEzsigntemplateelementdependencyOperator = "between"
)

// All allowed values of FieldEEzsigntemplateelementdependencyOperator enum
var AllowedFieldEEzsigntemplateelementdependencyOperatorEnumValues = []FieldEEzsigntemplateelementdependencyOperator{
	"eq",
	"neq",
	"gt",
	"gte",
	"lt",
	"lte",
	"in",
	"nin",
	"rg",
	"like",
	"between",
}

func (v *FieldEEzsigntemplateelementdependencyOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateelementdependencyOperator(value)
	for _, existing := range AllowedFieldEEzsigntemplateelementdependencyOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateelementdependencyOperator", value)
}

// NewFieldEEzsigntemplateelementdependencyOperatorFromValue returns a pointer to a valid FieldEEzsigntemplateelementdependencyOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateelementdependencyOperatorFromValue(v string) (*FieldEEzsigntemplateelementdependencyOperator, error) {
	ev := FieldEEzsigntemplateelementdependencyOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateelementdependencyOperator: valid values are %v", v, AllowedFieldEEzsigntemplateelementdependencyOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateelementdependencyOperator) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateelementdependencyOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateelementdependencyOperator value
func (v FieldEEzsigntemplateelementdependencyOperator) Ptr() *FieldEEzsigntemplateelementdependencyOperator {
	return &v
}

type NullableFieldEEzsigntemplateelementdependencyOperator struct {
	value *FieldEEzsigntemplateelementdependencyOperator
	isSet bool
}

func (v NullableFieldEEzsigntemplateelementdependencyOperator) Get() *FieldEEzsigntemplateelementdependencyOperator {
	return v.value
}

func (v *NullableFieldEEzsigntemplateelementdependencyOperator) Set(val *FieldEEzsigntemplateelementdependencyOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateelementdependencyOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateelementdependencyOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateelementdependencyOperator(val *FieldEEzsigntemplateelementdependencyOperator) *NullableFieldEEzsigntemplateelementdependencyOperator {
	return &NullableFieldEEzsigntemplateelementdependencyOperator{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateelementdependencyOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateelementdependencyOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

