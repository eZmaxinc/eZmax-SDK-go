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

// FieldEEzsignformfieldgroupTooltipposition The location of the tooltip relative to the Ezsignformfieldgroup's location.
type FieldEEzsignformfieldgroupTooltipposition string

// List of Field-eEzsignformfieldgroupTooltipposition
const (
	TOP_LEFT FieldEEzsignformfieldgroupTooltipposition = "TopLeft"
	TOP_CENTER FieldEEzsignformfieldgroupTooltipposition = "TopCenter"
	TOP_RIGHT FieldEEzsignformfieldgroupTooltipposition = "TopRight"
	MIDDLE_LEFT FieldEEzsignformfieldgroupTooltipposition = "MiddleLeft"
	MIDDLE_RIGHT FieldEEzsignformfieldgroupTooltipposition = "MiddleRight"
	BOTTOM_LEFT FieldEEzsignformfieldgroupTooltipposition = "BottomLeft"
	BOTTOM_CENTER FieldEEzsignformfieldgroupTooltipposition = "BottomCenter"
	BOTTOM_RIGHT FieldEEzsignformfieldgroupTooltipposition = "BottomRight"
)

// All allowed values of FieldEEzsignformfieldgroupTooltipposition enum
var AllowedFieldEEzsignformfieldgroupTooltippositionEnumValues = []FieldEEzsignformfieldgroupTooltipposition{
	"TopLeft",
	"TopCenter",
	"TopRight",
	"MiddleLeft",
	"MiddleRight",
	"BottomLeft",
	"BottomCenter",
	"BottomRight",
}

func (v *FieldEEzsignformfieldgroupTooltipposition) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsignformfieldgroupTooltipposition(value)
	for _, existing := range AllowedFieldEEzsignformfieldgroupTooltippositionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsignformfieldgroupTooltipposition", value)
}

// NewFieldEEzsignformfieldgroupTooltippositionFromValue returns a pointer to a valid FieldEEzsignformfieldgroupTooltipposition
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsignformfieldgroupTooltippositionFromValue(v string) (*FieldEEzsignformfieldgroupTooltipposition, error) {
	ev := FieldEEzsignformfieldgroupTooltipposition(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsignformfieldgroupTooltipposition: valid values are %v", v, AllowedFieldEEzsignformfieldgroupTooltippositionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsignformfieldgroupTooltipposition) IsValid() bool {
	for _, existing := range AllowedFieldEEzsignformfieldgroupTooltippositionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsignformfieldgroupTooltipposition value
func (v FieldEEzsignformfieldgroupTooltipposition) Ptr() *FieldEEzsignformfieldgroupTooltipposition {
	return &v
}

type NullableFieldEEzsignformfieldgroupTooltipposition struct {
	value *FieldEEzsignformfieldgroupTooltipposition
	isSet bool
}

func (v NullableFieldEEzsignformfieldgroupTooltipposition) Get() *FieldEEzsignformfieldgroupTooltipposition {
	return v.value
}

func (v *NullableFieldEEzsignformfieldgroupTooltipposition) Set(val *FieldEEzsignformfieldgroupTooltipposition) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsignformfieldgroupTooltipposition) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsignformfieldgroupTooltipposition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsignformfieldgroupTooltipposition(val *FieldEEzsignformfieldgroupTooltipposition) *NullableFieldEEzsignformfieldgroupTooltipposition {
	return &NullableFieldEEzsignformfieldgroupTooltipposition{value: val, isSet: true}
}

func (v NullableFieldEEzsignformfieldgroupTooltipposition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsignformfieldgroupTooltipposition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

