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

// FieldEAttachmentlogType The Type for the Attachmentlog
type FieldEAttachmentlogType string

// List of Field-eAttachmentlogType
const (
	AUTO_VALIDATION FieldEAttachmentlogType = "AutoValidation"
	COPY_FROM FieldEAttachmentlogType = "CopyFrom"
	COPY_TO FieldEAttachmentlogType = "CopyTo"
	COPY_TO_EZSIGN FieldEAttachmentlogType = "CopyToEzsign"
	CREATE_BY_EZSIGN FieldEAttachmentlogType = "CreateByEzsign"
	DOWNLOAD FieldEAttachmentlogType = "Download"
	DELETED FieldEAttachmentlogType = "Deleted"
	DESTROYED FieldEAttachmentlogType = "Destroyed"
	EMAIL FieldEAttachmentlogType = "Email"
	EMAIL_CC FieldEAttachmentlogType = "EmailCC"
	EMAIL_CCI FieldEAttachmentlogType = "EmailCCI"
	FAX FieldEAttachmentlogType = "Fax"
	IMPORTED_FROM_EXTERNAL_SYSTEM FieldEAttachmentlogType = "ImportedFromExternalSystem"
	IMPORTED_FROM_EZA FieldEAttachmentlogType = "ImportedFromEZA"
	IMPORTED_FROM_FALTOUR FieldEAttachmentlogType = "ImportedFromFaltour"
	IMPORTED_FROM_LONEWOLF FieldEAttachmentlogType = "ImportedFromLonewolf"
	IMPORTED_FROM_PROSPECTS FieldEAttachmentlogType = "ImportedFromProspects"
	MOVE FieldEAttachmentlogType = "Move"
	OPEN_FROM_EMAIL FieldEAttachmentlogType = "OpenFromEmail"
	PURGED FieldEAttachmentlogType = "Purged"
	REJECT FieldEAttachmentlogType = "Reject"
	RENAME FieldEAttachmentlogType = "Rename"
	RESTORE FieldEAttachmentlogType = "Restore"
	SCANNED FieldEAttachmentlogType = "Scanned"
	SEND_TO_GED FieldEAttachmentlogType = "SendToGED"
	UNVALIDATED_BY FieldEAttachmentlogType = "UnvalidatedBy"
	UPLOAD FieldEAttachmentlogType = "Upload"
	VALIDATED_BY FieldEAttachmentlogType = "ValidatedBy"
	VETINFO_UPLOAD FieldEAttachmentlogType = "VetinfoUpload"
)

// All allowed values of FieldEAttachmentlogType enum
var AllowedFieldEAttachmentlogTypeEnumValues = []FieldEAttachmentlogType{
	"AutoValidation",
	"CopyFrom",
	"CopyTo",
	"CopyToEzsign",
	"CreateByEzsign",
	"Download",
	"Deleted",
	"Destroyed",
	"Email",
	"EmailCC",
	"EmailCCI",
	"Fax",
	"ImportedFromExternalSystem",
	"ImportedFromEZA",
	"ImportedFromFaltour",
	"ImportedFromLonewolf",
	"ImportedFromProspects",
	"Move",
	"OpenFromEmail",
	"Purged",
	"Reject",
	"Rename",
	"Restore",
	"Scanned",
	"SendToGED",
	"UnvalidatedBy",
	"Upload",
	"ValidatedBy",
	"VetinfoUpload",
}

func (v *FieldEAttachmentlogType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEAttachmentlogType(value)
	for _, existing := range AllowedFieldEAttachmentlogTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEAttachmentlogType", value)
}

// NewFieldEAttachmentlogTypeFromValue returns a pointer to a valid FieldEAttachmentlogType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEAttachmentlogTypeFromValue(v string) (*FieldEAttachmentlogType, error) {
	ev := FieldEAttachmentlogType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEAttachmentlogType: valid values are %v", v, AllowedFieldEAttachmentlogTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEAttachmentlogType) IsValid() bool {
	for _, existing := range AllowedFieldEAttachmentlogTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eAttachmentlogType value
func (v FieldEAttachmentlogType) Ptr() *FieldEAttachmentlogType {
	return &v
}

type NullableFieldEAttachmentlogType struct {
	value *FieldEAttachmentlogType
	isSet bool
}

func (v NullableFieldEAttachmentlogType) Get() *FieldEAttachmentlogType {
	return v.value
}

func (v *NullableFieldEAttachmentlogType) Set(val *FieldEAttachmentlogType) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEAttachmentlogType) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEAttachmentlogType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEAttachmentlogType(val *FieldEAttachmentlogType) *NullableFieldEAttachmentlogType {
	return &NullableFieldEAttachmentlogType{value: val, isSet: true}
}

func (v NullableFieldEAttachmentlogType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEAttachmentlogType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

