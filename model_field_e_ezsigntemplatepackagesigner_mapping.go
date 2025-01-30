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

// FieldEEzsigntemplatepackagesignerMapping Mapping configuration to use when package is apply.  1. **Manual** User need to choose mapping manually. 2. **Creator** mapping will be set to creator of folderS. 3. **User** mapping will be set to fkiUserID 4. **Usergroup** mapping will be set to fkiUsergroupID.
type FieldEEzsigntemplatepackagesignerMapping string

// List of Field-eEzsigntemplatepackagesignerMapping
const (
	MANUAL FieldEEzsigntemplatepackagesignerMapping = "Manual"
	CREATOR FieldEEzsigntemplatepackagesignerMapping = "Creator"
	USER FieldEEzsigntemplatepackagesignerMapping = "User"
	USERGROUP FieldEEzsigntemplatepackagesignerMapping = "Usergroup"
)

// All allowed values of FieldEEzsigntemplatepackagesignerMapping enum
var AllowedFieldEEzsigntemplatepackagesignerMappingEnumValues = []FieldEEzsigntemplatepackagesignerMapping{
	"Manual",
	"Creator",
	"User",
	"Usergroup",
}

func (v *FieldEEzsigntemplatepackagesignerMapping) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEEzsigntemplatepackagesignerMapping(value)
	for _, existing := range AllowedFieldEEzsigntemplatepackagesignerMappingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEEzsigntemplatepackagesignerMapping", value)
}

// NewFieldEEzsigntemplatepackagesignerMappingFromValue returns a pointer to a valid FieldEEzsigntemplatepackagesignerMapping
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEEzsigntemplatepackagesignerMappingFromValue(v string) (*FieldEEzsigntemplatepackagesignerMapping, error) {
	ev := FieldEEzsigntemplatepackagesignerMapping(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEEzsigntemplatepackagesignerMapping: valid values are %v", v, AllowedFieldEEzsigntemplatepackagesignerMappingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEEzsigntemplatepackagesignerMapping) IsValid() bool {
	for _, existing := range AllowedFieldEEzsigntemplatepackagesignerMappingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eEzsigntemplatepackagesignerMapping value
func (v FieldEEzsigntemplatepackagesignerMapping) Ptr() *FieldEEzsigntemplatepackagesignerMapping {
	return &v
}

type NullableFieldEEzsigntemplatepackagesignerMapping struct {
	value *FieldEEzsigntemplatepackagesignerMapping
	isSet bool
}

func (v NullableFieldEEzsigntemplatepackagesignerMapping) Get() *FieldEEzsigntemplatepackagesignerMapping {
	return v.value
}

func (v *NullableFieldEEzsigntemplatepackagesignerMapping) Set(val *FieldEEzsigntemplatepackagesignerMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEEzsigntemplatepackagesignerMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEEzsigntemplatepackagesignerMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEEzsigntemplatepackagesignerMapping(val *FieldEEzsigntemplatepackagesignerMapping) *NullableFieldEEzsigntemplatepackagesignerMapping {
	return &NullableFieldEEzsigntemplatepackagesignerMapping{value: val, isSet: true}
}

func (v NullableFieldEEzsigntemplatepackagesignerMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEEzsigntemplatepackagesignerMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

