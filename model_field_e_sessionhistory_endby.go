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

// FieldESessionhistoryEndby The Type of the Sessionhistory
type FieldESessionhistoryEndby string

// List of Field-eSessionhistoryEndby
const (
	DECRYPTION FieldESessionhistoryEndby = "Decryption"
	HACK FieldESessionhistoryEndby = "Hack"
	EXPIRED FieldESessionhistoryEndby = "Expired"
	HIJACK FieldESessionhistoryEndby = "Hijack"
	DOUBLE_LOGON FieldESessionhistoryEndby = "DoubleLogon"
	GARBAGE FieldESessionhistoryEndby = "Garbage"
	LOGOFF FieldESessionhistoryEndby = "Logoff"
	BAD_AUTH FieldESessionhistoryEndby = "BadAuth"
	LOCKED FieldESessionhistoryEndby = "Locked"
	INACTIVE FieldESessionhistoryEndby = "Inactive"
	INVALID_USER FieldESessionhistoryEndby = "InvalidUser"
	BAD_USER_TYPE FieldESessionhistoryEndby = "BadUserType"
	BAD_IP FieldESessionhistoryEndby = "BadIP"
	FORCED_LOGOFF FieldESessionhistoryEndby = "ForcedLogoff"
)

// All allowed values of FieldESessionhistoryEndby enum
var AllowedFieldESessionhistoryEndbyEnumValues = []FieldESessionhistoryEndby{
	"Decryption",
	"Hack",
	"Expired",
	"Hijack",
	"DoubleLogon",
	"Garbage",
	"Logoff",
	"BadAuth",
	"Locked",
	"Inactive",
	"InvalidUser",
	"BadUserType",
	"BadIP",
	"ForcedLogoff",
}

func (v *FieldESessionhistoryEndby) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FieldESessionhistoryEndby(value)
	for _, existing := range AllowedFieldESessionhistoryEndbyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FieldESessionhistoryEndby", value)
}

// NewFieldESessionhistoryEndbyFromValue returns a pointer to a valid FieldESessionhistoryEndby
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFieldESessionhistoryEndbyFromValue(v string) (*FieldESessionhistoryEndby, error) {
	ev := FieldESessionhistoryEndby(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FieldESessionhistoryEndby: valid values are %v", v, AllowedFieldESessionhistoryEndbyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FieldESessionhistoryEndby) IsValid() bool {
	for _, existing := range AllowedFieldESessionhistoryEndbyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Field-eSessionhistoryEndby value
func (v FieldESessionhistoryEndby) Ptr() *FieldESessionhistoryEndby {
	return &v
}

type NullableFieldESessionhistoryEndby struct {
	value *FieldESessionhistoryEndby
	isSet bool
}

func (v NullableFieldESessionhistoryEndby) Get() *FieldESessionhistoryEndby {
	return v.value
}

func (v *NullableFieldESessionhistoryEndby) Set(val *FieldESessionhistoryEndby) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldESessionhistoryEndby) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldESessionhistoryEndby) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldESessionhistoryEndby(val *FieldESessionhistoryEndby) *NullableFieldESessionhistoryEndby {
	return &NullableFieldESessionhistoryEndby{value: val, isSet: true}
}

func (v NullableFieldESessionhistoryEndby) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldESessionhistoryEndby) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
