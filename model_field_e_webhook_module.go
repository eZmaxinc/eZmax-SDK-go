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

// FieldEWebhookModule The module for the Webhook
type FieldEWebhookModule string

// List of Field-eWebhookModule
const (
	EZSIGN FieldEWebhookModule = "Ezsign"
	MANAGEMENT FieldEWebhookModule = "Management"
)

// All allowed values of FieldEWebhookModule enum
var AllowedFieldEWebhookModuleEnumValues = []FieldEWebhookModule{
	"Ezsign",
	"Management",
}

func (v *FieldEWebhookModule) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEWebhookModule(value)
	for _, existing := range AllowedFieldEWebhookModuleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEWebhookModule", value)
}

// NewFieldEWebhookModuleFromValue returns a pointer to a valid FieldEWebhookModule
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEWebhookModuleFromValue(v string) (*FieldEWebhookModule, error) {
	ev := FieldEWebhookModule(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEWebhookModule: valid values are %v", v, AllowedFieldEWebhookModuleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEWebhookModule) IsValid() bool {
	for _, existing := range AllowedFieldEWebhookModuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eWebhookModule value
func (v FieldEWebhookModule) Ptr() *FieldEWebhookModule {
	return &v
}

type NullableFieldEWebhookModule struct {
	value *FieldEWebhookModule
	isSet bool
}

func (v NullableFieldEWebhookModule) Get() *FieldEWebhookModule {
	return v.value
}

func (v *NullableFieldEWebhookModule) Set(val *FieldEWebhookModule) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEWebhookModule) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEWebhookModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEWebhookModule(val *FieldEWebhookModule) *NullableFieldEWebhookModule {
	return &NullableFieldEWebhookModule{value: val, isSet: true}
}

func (v NullableFieldEWebhookModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEWebhookModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

