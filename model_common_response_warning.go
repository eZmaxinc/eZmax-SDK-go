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
	"bytes"
	"fmt"
)

// checks if the CommonResponseWarning type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseWarning{}

// CommonResponseWarning Generic Warning Message
type CommonResponseWarning struct {
	// More detail about the warning
	SWarningMessage string `json:"sWarningMessage"`
	// The warning code. See documentation for valid values
	EWarningCode string `json:"eWarningCode"`
}

type _CommonResponseWarning CommonResponseWarning

// NewCommonResponseWarning instantiates a new CommonResponseWarning object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseWarning(sWarningMessage string, eWarningCode string) *CommonResponseWarning {
	this := CommonResponseWarning{}
	this.SWarningMessage = sWarningMessage
	this.EWarningCode = eWarningCode
	return &this
}

// NewCommonResponseWarningWithDefaults instantiates a new CommonResponseWarning object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseWarningWithDefaults() *CommonResponseWarning {
	this := CommonResponseWarning{}
	return &this
}

// GetSWarningMessage returns the SWarningMessage field value
func (o *CommonResponseWarning) GetSWarningMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SWarningMessage
}

// GetSWarningMessageOk returns a tuple with the SWarningMessage field value
// and a boolean to check if the value has been set.
func (o *CommonResponseWarning) GetSWarningMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SWarningMessage, true
}

// SetSWarningMessage sets field value
func (o *CommonResponseWarning) SetSWarningMessage(v string) {
	o.SWarningMessage = v
}

// GetEWarningCode returns the EWarningCode field value
func (o *CommonResponseWarning) GetEWarningCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EWarningCode
}

// GetEWarningCodeOk returns a tuple with the EWarningCode field value
// and a boolean to check if the value has been set.
func (o *CommonResponseWarning) GetEWarningCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EWarningCode, true
}

// SetEWarningCode sets field value
func (o *CommonResponseWarning) SetEWarningCode(v string) {
	o.EWarningCode = v
}

func (o CommonResponseWarning) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseWarning) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sWarningMessage"] = o.SWarningMessage
	toSerialize["eWarningCode"] = o.EWarningCode
	return toSerialize, nil
}

func (o *CommonResponseWarning) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sWarningMessage",
		"eWarningCode",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCommonResponseWarning := _CommonResponseWarning{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseWarning)

	if err != nil {
		return err
	}

	*o = CommonResponseWarning(varCommonResponseWarning)

	return err
}

type NullableCommonResponseWarning struct {
	value *CommonResponseWarning
	isSet bool
}

func (v NullableCommonResponseWarning) Get() *CommonResponseWarning {
	return v.value
}

func (v *NullableCommonResponseWarning) Set(val *CommonResponseWarning) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseWarning) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseWarning) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseWarning(val *CommonResponseWarning) *NullableCommonResponseWarning {
	return &NullableCommonResponseWarning{value: val, isSet: true}
}

func (v NullableCommonResponseWarning) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseWarning) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


