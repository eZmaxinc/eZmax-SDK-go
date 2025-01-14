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

// CustomEWebhookEzsignevent This Ezsign Event. This property will be set only if the Module is \"Ezsign\"
type CustomEWebhookEzsignevent string

// List of Custom-eWebhookEzsignevent
const (
	DOCUMENT_COMPLETED CustomEWebhookEzsignevent = "DocumentCompleted"
	DOCUMENT_FORM_COMPLETED CustomEWebhookEzsignevent = "DocumentFormCompleted"
	DOCUMENT_UNSENT CustomEWebhookEzsignevent = "DocumentUnsent"
	EZSIGNSIGNER_ACCEPTCLAUSE CustomEWebhookEzsignevent = "EzsignsignerAcceptclause"
	EZSIGNSIGNER_CONNECT CustomEWebhookEzsignevent = "EzsignsignerConnect"
	FOLDER_COMPLETED CustomEWebhookEzsignevent = "FolderCompleted"
	FOLDER_DISPOSED CustomEWebhookEzsignevent = "FolderDisposed"
	FOLDER_SENT CustomEWebhookEzsignevent = "FolderSent"
	FOLDER_UNSENT CustomEWebhookEzsignevent = "FolderUnsent"
)

// All allowed values of CustomEWebhookEzsignevent enum
var AllowedCustomEWebhookEzsigneventEnumValues = []CustomEWebhookEzsignevent{
	"DocumentCompleted",
	"DocumentFormCompleted",
	"DocumentUnsent",
	"EzsignsignerAcceptclause",
	"EzsignsignerConnect",
	"FolderCompleted",
	"FolderDisposed",
	"FolderSent",
	"FolderUnsent",
}

func (v *CustomEWebhookEzsignevent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomEWebhookEzsignevent(value)
	for _, existing := range AllowedCustomEWebhookEzsigneventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomEWebhookEzsignevent", value)
}

// NewCustomEWebhookEzsigneventFromValue returns a pointer to a valid CustomEWebhookEzsignevent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomEWebhookEzsigneventFromValue(v string) (*CustomEWebhookEzsignevent, error) {
	ev := CustomEWebhookEzsignevent(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomEWebhookEzsignevent: valid values are %v", v, AllowedCustomEWebhookEzsigneventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomEWebhookEzsignevent) IsValid() bool {
	for _, existing := range AllowedCustomEWebhookEzsigneventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Custom-eWebhookEzsignevent value
func (v CustomEWebhookEzsignevent) Ptr() *CustomEWebhookEzsignevent {
	return &v
}

type NullableCustomEWebhookEzsignevent struct {
	value *CustomEWebhookEzsignevent
	isSet bool
}

func (v NullableCustomEWebhookEzsignevent) Get() *CustomEWebhookEzsignevent {
	return v.value
}

func (v *NullableCustomEWebhookEzsignevent) Set(val *CustomEWebhookEzsignevent) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomEWebhookEzsignevent) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomEWebhookEzsignevent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomEWebhookEzsignevent(val *CustomEWebhookEzsignevent) *NullableCustomEWebhookEzsignevent {
	return &NullableCustomEWebhookEzsignevent{value: val, isSet: true}
}

func (v NullableCustomEWebhookEzsignevent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomEWebhookEzsignevent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

