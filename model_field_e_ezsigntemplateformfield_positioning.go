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

// FieldEEzsigntemplateformfieldPositioning How the positioning of the Ezsigntemplateformfield will be done
type FieldEEzsigntemplateformfieldPositioning string

// List of Field-eEzsigntemplateformfieldPositioning
const (
	PER_COORDINATES FieldEEzsigntemplateformfieldPositioning = "PerCoordinates"
	PER_POSITIONING_PATTERN FieldEEzsigntemplateformfieldPositioning = "PerPositioningPattern"
)

// All allowed values of FieldEEzsigntemplateformfieldPositioning enum
var AllowedFieldEEzsigntemplateformfieldPositioningEnumValues = []FieldEEzsigntemplateformfieldPositioning{
	"PerCoordinates",
	"PerPositioningPattern",
}

func (v *FieldEEzsigntemplateformfieldPositioning) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplateformfieldPositioning(value)
	for _, existing := range AllowedFieldEEzsigntemplateformfieldPositioningEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplateformfieldPositioning", value)
}

// NewFieldEEzsigntemplateformfieldPositioningFromValue returns a pointer to a valid FieldEEzsigntemplateformfieldPositioning
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplateformfieldPositioningFromValue(v string) (*FieldEEzsigntemplateformfieldPositioning, error) {
	ev := FieldEEzsigntemplateformfieldPositioning(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplateformfieldPositioning: valid values are %v", v, AllowedFieldEEzsigntemplateformfieldPositioningEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplateformfieldPositioning) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplateformfieldPositioningEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplateformfieldPositioning value
func (v FieldEEzsigntemplateformfieldPositioning) Ptr() *FieldEEzsigntemplateformfieldPositioning {
	return &v
}

type NullableFieldEEzsigntemplateformfieldPositioning struct {
	value *FieldEEzsigntemplateformfieldPositioning
	isSet bool
}

func (v NullableFieldEEzsigntemplateformfieldPositioning) Get() *FieldEEzsigntemplateformfieldPositioning {
	return v.value
}

func (v *NullableFieldEEzsigntemplateformfieldPositioning) Set(val *FieldEEzsigntemplateformfieldPositioning) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplateformfieldPositioning) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplateformfieldPositioning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplateformfieldPositioning(val *FieldEEzsigntemplateformfieldPositioning) *NullableFieldEEzsigntemplateformfieldPositioning {
	return &NullableFieldEEzsigntemplateformfieldPositioning{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplateformfieldPositioning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplateformfieldPositioning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
