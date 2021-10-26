/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"fmt"
)

// HeaderAcceptLanguage The language of the returned content.  1. **\\*** (or header not defined) Default language 2. **en** English 2. **fr** French  
type HeaderAcceptLanguage string

// List of Header-Accept-Language
const (
	STAR HeaderAcceptLanguage = "*"
	EN HeaderAcceptLanguage = "en"
	FR HeaderAcceptLanguage = "fr"
)

// All allowed values of HeaderAcceptLanguage enum
var AllowedHeaderAcceptLanguageEnumValues = []HeaderAcceptLanguage{
	"*",
	"en",
	"fr",
}

func (v *HeaderAcceptLanguage) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HeaderAcceptLanguage(value)
	for _, existing := range AllowedHeaderAcceptLanguageEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HeaderAcceptLanguage", value)
}

// NewHeaderAcceptLanguageFromValue returns a pointer to a valid HeaderAcceptLanguage
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHeaderAcceptLanguageFromValue(v string) (*HeaderAcceptLanguage, error) {
	ev := HeaderAcceptLanguage(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HeaderAcceptLanguage: valid values are %v", v, AllowedHeaderAcceptLanguageEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HeaderAcceptLanguage) IsValid() bool {
	for _, existing := range AllowedHeaderAcceptLanguageEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Header-Accept-Language value
func (v HeaderAcceptLanguage) Ptr() *HeaderAcceptLanguage {
	return &v
}

type NullableHeaderAcceptLanguage struct {
	value *HeaderAcceptLanguage
	isSet bool
}

func (v NullableHeaderAcceptLanguage) Get() *HeaderAcceptLanguage {
	return v.value
}

func (v *NullableHeaderAcceptLanguage) Set(val *HeaderAcceptLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableHeaderAcceptLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableHeaderAcceptLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeaderAcceptLanguage(val *HeaderAcceptLanguage) *NullableHeaderAcceptLanguage {
	return &NullableHeaderAcceptLanguage{value: val, isSet: true}
}

func (v NullableHeaderAcceptLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeaderAcceptLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

