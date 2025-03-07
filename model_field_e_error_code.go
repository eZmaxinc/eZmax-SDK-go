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

// FieldEErrorCode The error code. See documentation for valid values
type FieldEErrorCode string

// List of Field-eErrorCode
const (
	BADREQUEST FieldEErrorCode = "BADREQUEST"
	BADREQUEST_CLOCKSKEW FieldEErrorCode = "BADREQUEST_CLOCKSKEW"
	UNAUTHORIZED_BADAUTH FieldEErrorCode = "UNAUTHORIZED_BADAUTH"
	UNAUTHORIZED_BADMFA FieldEErrorCode = "UNAUTHORIZED_BADMFA"
	UNAUTHORIZED_EXPIRED FieldEErrorCode = "UNAUTHORIZED_EXPIRED"
	UNAUTHORIZED_REQUEST FieldEErrorCode = "UNAUTHORIZED_REQUEST"
	FORBIDDEN FieldEErrorCode = "FORBIDDEN"
	FORBIDDEN_CONFIGURATION FieldEErrorCode = "FORBIDDEN_CONFIGURATION"
	FORBIDDEN_MODULE FieldEErrorCode = "FORBIDDEN_MODULE"
	FORBIDDEN_NOACCESS FieldEErrorCode = "FORBIDDEN_NOACCESS"
	FORBIDDEN_PERMISSION FieldEErrorCode = "FORBIDDEN_PERMISSION"
	FORBIDDEN_SUBSCRIPTION FieldEErrorCode = "FORBIDDEN_SUBSCRIPTION"
	FORBIDDEN_USERTYPE FieldEErrorCode = "FORBIDDEN_USERTYPE"
	FORBIDDEN_USER_ORIGIN_EXTERNAL FieldEErrorCode = "FORBIDDEN_USER_ORIGIN_EXTERNAL"
	NOTFOUND FieldEErrorCode = "NOTFOUND"
	NOTFOUND_OBJECT FieldEErrorCode = "NOTFOUND_OBJECT"
	NOTFOUND_ROUTE FieldEErrorCode = "NOTFOUND_ROUTE"
	METHODNOTALLOWED FieldEErrorCode = "METHODNOTALLOWED"
	NOTACCEPTABLE_CONTENT FieldEErrorCode = "NOTACCEPTABLE_CONTENT"
	NOTACCEPTABLE_LANGUAGE FieldEErrorCode = "NOTACCEPTABLE_LANGUAGE"
	UNPROCESSABLEENTITY_ACTIVESESSION_ALREADY_CLONING FieldEErrorCode = "UNPROCESSABLEENTITY_ACTIVESESSION_ALREADY_CLONING"
	UNPROCESSABLEENTITY_CANNOTDELETE FieldEErrorCode = "UNPROCESSABLEENTITY_CANNOTDELETE"
	UNPROCESSABLEENTITY_CANNOTMODIFY FieldEErrorCode = "UNPROCESSABLEENTITY_CANNOTMODIFY"
	UNPROCESSABLEENTITY_CHANGEPASSWORD_INVALID_CURRENT FieldEErrorCode = "UNPROCESSABLEENTITY_CHANGEPASSWORD_INVALID_CURRENT"
	UNPROCESSABLEENTITY_CHANGEPASSWORD_SAME FieldEErrorCode = "UNPROCESSABLEENTITY_CHANGEPASSWORD_SAME"
	UNPROCESSABLEENTITY_DATA_MISSING FieldEErrorCode = "UNPROCESSABLEENTITY_DATA_MISSING"
	UNPROCESSABLEENTITY_DATA_UNIQUE FieldEErrorCode = "UNPROCESSABLEENTITY_DATA_UNIQUE"
	UNPROCESSABLEENTITY_DATA_VALIDATION FieldEErrorCode = "UNPROCESSABLEENTITY_DATA_VALIDATION"
	UNPROCESSABLEENTITY_DATA_OUTOFBOUND FieldEErrorCode = "UNPROCESSABLEENTITY_DATA_OUTOFBOUND"
	UNPROCESSABLEENTITY_DOWNLOAD_ERROR FieldEErrorCode = "UNPROCESSABLEENTITY_DOWNLOAD_ERROR"
	UNPROCESSABLEENTITY_EZSIGNFORM_VALIDATION FieldEErrorCode = "UNPROCESSABLEENTITY_EZSIGNFORM_VALIDATION"
	UNPROCESSABLEENTITY_EZSIGNSIGNERCONNECTED FieldEErrorCode = "UNPROCESSABLEENTITY_EZSIGNSIGNERCONNECTED"
	UNPROCESSABLEENTITY_NOTHINGTODO FieldEErrorCode = "UNPROCESSABLEENTITY_NOTHINGTODO"
	UNPROCESSABLEENTITY_NOTREADY FieldEErrorCode = "UNPROCESSABLEENTITY_NOTREADY"
	UNPROCESSABLEENTITY_PDF_FORM FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_FORM"
	UNPROCESSABLEENTITY_PDF_SIGNATURE FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_SIGNATURE"
	UNPROCESSABLEENTITY_PDF_FORM_AND_SIGNATURE FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_FORM_AND_SIGNATURE"
	UNPROCESSABLEENTITY_PDF_INCOMPATIBLE FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_INCOMPATIBLE"
	UNPROCESSABLEENTITY_PDF_PASSWORD FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_PASSWORD"
	UNPROCESSABLEENTITY_PDF_WRONG_PASSWORD FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_WRONG_PASSWORD"
	UNPROCESSABLEENTITY_PDF_REPAIRABLE FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_REPAIRABLE"
	UNPROCESSABLEENTITY_PDF_XFA FieldEErrorCode = "UNPROCESSABLEENTITY_PDF_XFA"
	UNPROCESSABLEENTITY_TEMPLATE_MISMATCH FieldEErrorCode = "UNPROCESSABLEENTITY_TEMPLATE_MISMATCH"
	UNPROCESSABLEENTITY_UNMODIFIABLE_FIELD FieldEErrorCode = "UNPROCESSABLEENTITY_UNMODIFIABLE_FIELD"
	UNPROCESSABLEENTITY_USER_STAGED FieldEErrorCode = "UNPROCESSABLEENTITY_USER_STAGED"
	TOOMANYREQUESTS FieldEErrorCode = "TOOMANYREQUESTS"
	TOOMANYREQUESTS_THIRDPARTY FieldEErrorCode = "TOOMANYREQUESTS_THIRDPARTY"
	ERROR_INTERNAL FieldEErrorCode = "ERROR_INTERNAL"
	ERROR_CONFIGURATION FieldEErrorCode = "ERROR_CONFIGURATION"
	ERROR_NOTIMPLEMENTED FieldEErrorCode = "ERROR_NOTIMPLEMENTED"
)

// All allowed values of FieldEErrorCode enum
var AllowedFieldEErrorCodeEnumValues = []FieldEErrorCode{
	"BADREQUEST",
	"BADREQUEST_CLOCKSKEW",
	"UNAUTHORIZED_BADAUTH",
	"UNAUTHORIZED_BADMFA",
	"UNAUTHORIZED_EXPIRED",
	"UNAUTHORIZED_REQUEST",
	"FORBIDDEN",
	"FORBIDDEN_CONFIGURATION",
	"FORBIDDEN_MODULE",
	"FORBIDDEN_NOACCESS",
	"FORBIDDEN_PERMISSION",
	"FORBIDDEN_SUBSCRIPTION",
	"FORBIDDEN_USERTYPE",
	"FORBIDDEN_USER_ORIGIN_EXTERNAL",
	"NOTFOUND",
	"NOTFOUND_OBJECT",
	"NOTFOUND_ROUTE",
	"METHODNOTALLOWED",
	"NOTACCEPTABLE_CONTENT",
	"NOTACCEPTABLE_LANGUAGE",
	"UNPROCESSABLEENTITY_ACTIVESESSION_ALREADY_CLONING",
	"UNPROCESSABLEENTITY_CANNOTDELETE",
	"UNPROCESSABLEENTITY_CANNOTMODIFY",
	"UNPROCESSABLEENTITY_CHANGEPASSWORD_INVALID_CURRENT",
	"UNPROCESSABLEENTITY_CHANGEPASSWORD_SAME",
	"UNPROCESSABLEENTITY_DATA_MISSING",
	"UNPROCESSABLEENTITY_DATA_UNIQUE",
	"UNPROCESSABLEENTITY_DATA_VALIDATION",
	"UNPROCESSABLEENTITY_DATA_OUTOFBOUND",
	"UNPROCESSABLEENTITY_DOWNLOAD_ERROR",
	"UNPROCESSABLEENTITY_EZSIGNFORM_VALIDATION",
	"UNPROCESSABLEENTITY_EZSIGNSIGNERCONNECTED",
	"UNPROCESSABLEENTITY_NOTHINGTODO",
	"UNPROCESSABLEENTITY_NOTREADY",
	"UNPROCESSABLEENTITY_PDF_FORM",
	"UNPROCESSABLEENTITY_PDF_SIGNATURE",
	"UNPROCESSABLEENTITY_PDF_FORM_AND_SIGNATURE",
	"UNPROCESSABLEENTITY_PDF_INCOMPATIBLE",
	"UNPROCESSABLEENTITY_PDF_PASSWORD",
	"UNPROCESSABLEENTITY_PDF_WRONG_PASSWORD",
	"UNPROCESSABLEENTITY_PDF_REPAIRABLE",
	"UNPROCESSABLEENTITY_PDF_XFA",
	"UNPROCESSABLEENTITY_TEMPLATE_MISMATCH",
	"UNPROCESSABLEENTITY_UNMODIFIABLE_FIELD",
	"UNPROCESSABLEENTITY_USER_STAGED",
	"TOOMANYREQUESTS",
	"TOOMANYREQUESTS_THIRDPARTY",
	"ERROR_INTERNAL",
	"ERROR_CONFIGURATION",
	"ERROR_NOTIMPLEMENTED",
}

func (v *FieldEErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldEErrorCode(value)
	for _, existing := range AllowedFieldEErrorCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldEErrorCode", value)
}

// NewFieldEErrorCodeFromValue returns a pointer to a valid FieldEErrorCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldEErrorCodeFromValue(v string) (*FieldEErrorCode, error) {
	ev := FieldEErrorCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldEErrorCode: valid values are %v", v, AllowedFieldEErrorCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldEErrorCode) IsValid() bool {
	for _, existing := range AllowedFieldEErrorCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eErrorCode value
func (v FieldEErrorCode) Ptr() *FieldEErrorCode {
	return &v
}

type NullableFieldEErrorCode struct {
	value *FieldEErrorCode
	isSet bool
}

func (v NullableFieldEErrorCode) Get() *FieldEErrorCode {
	return v.value
}

func (v *NullableFieldEErrorCode) Set(val *FieldEErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldEErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldEErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldEErrorCode(val *FieldEErrorCode) *NullableFieldEErrorCode {
	return &NullableFieldEErrorCode{value: val, isSet: true}
}

func (v NullableFieldEErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldEErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

