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

// FieldEWebhookEzsignevent This Ezsign Event. This property will be set only if the Module is \"Ezsign\"
type FieldEWebhookEzsignevent string

// List of Field-eWebhookEzsignevent
const (
	DOCUMENT_COMPLETED FieldEWebhookEzsignevent = "DocumentCompleted"
	EZSIGNSIGNER_ACCEPTCLAUSE FieldEWebhookEzsignevent = "EzsignsignerAcceptclause"
	EZSIGNSIGNER_CONNECT FieldEWebhookEzsignevent = "EzsignsignerConnect"
	FOLDER_COMPLETED FieldEWebhookEzsignevent = "FolderCompleted"
)

// All allowed values of FieldEWebhookEzsignevent enum
var AllowedFieldEWebhookEzsigneventEnumValues = []FieldEWebhookEzsignevent{
	"DocumentCompleted",
	"EzsignsignerAcceptclause",
	"EzsignsignerConnect",
	"FolderCompleted",
}

func (v *FieldEWebhookEzsignevent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEWebhookEzsignevent(value)
	for _, existing := range AllowedFieldEWebhookEzsigneventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEWebhookEzsignevent", value)
}

// NewFieldEWebhookEzsigneventFromValue returns a pointer to a valid FieldEWebhookEzsignevent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEWebhookEzsigneventFromValue(v string) (*FieldEWebhookEzsignevent, error) {
	ev := FieldEWebhookEzsignevent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEWebhookEzsignevent: valid values are %v", v, AllowedFieldEWebhookEzsigneventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEWebhookEzsignevent) IsValid() bool {
	for _, existing := range AllowedFieldEWebhookEzsigneventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eWebhookEzsignevent value
func (v FieldEWebhookEzsignevent) Ptr() *FieldEWebhookEzsignevent {
	return &v
}

type NullableFieldEWebhookEzsignevent struct {
	value *FieldEWebhookEzsignevent
	isSet bool
}

func (v NullableFieldEWebhookEzsignevent) Get() *FieldEWebhookEzsignevent {
	return v.value
}

func (v *NullableFieldEWebhookEzsignevent) Set(val *FieldEWebhookEzsignevent) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEWebhookEzsignevent) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEWebhookEzsignevent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEWebhookEzsignevent(val *FieldEWebhookEzsignevent) *NullableFieldEWebhookEzsignevent {
	return &NullableFieldEWebhookEzsignevent{value: val, isSet: true}
}

func (v NullableFieldEWebhookEzsignevent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEWebhookEzsignevent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

