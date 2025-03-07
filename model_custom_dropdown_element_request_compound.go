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
	"bytes"
	"fmt"
)

// checks if the CustomDropdownElementRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDropdownElementRequestCompound{}

// CustomDropdownElementRequestCompound A Generic DropdownElement Object and children to create a complete structure
type CustomDropdownElementRequestCompound struct {
	// The Description of the element
	SLabel string `json:"sLabel"`
	// The Value of the element
	SValue string `json:"sValue"`
}

type _CustomDropdownElementRequestCompound CustomDropdownElementRequestCompound

// NewCustomDropdownElementRequestCompound instantiates a new CustomDropdownElementRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDropdownElementRequestCompound(sLabel string, sValue string) *CustomDropdownElementRequestCompound {
	this := CustomDropdownElementRequestCompound{}
	this.SLabel = sLabel
	this.SValue = sValue
	return &this
}

// NewCustomDropdownElementRequestCompoundWithDefaults instantiates a new CustomDropdownElementRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDropdownElementRequestCompoundWithDefaults() *CustomDropdownElementRequestCompound {
	this := CustomDropdownElementRequestCompound{}
	return &this
}

// GetSLabel returns the SLabel field value
func (o *CustomDropdownElementRequestCompound) GetSLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLabel
}

// GetSLabelOk returns a tuple with the SLabel field value
// and a boolean to check if the value has been set.
func (o *CustomDropdownElementRequestCompound) GetSLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLabel, true
}

// SetSLabel sets field value
func (o *CustomDropdownElementRequestCompound) SetSLabel(v string) {
	o.SLabel = v
}

// GetSValue returns the SValue field value
func (o *CustomDropdownElementRequestCompound) GetSValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SValue
}

// GetSValueOk returns a tuple with the SValue field value
// and a boolean to check if the value has been set.
func (o *CustomDropdownElementRequestCompound) GetSValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SValue, true
}

// SetSValue sets field value
func (o *CustomDropdownElementRequestCompound) SetSValue(v string) {
	o.SValue = v
}

func (o CustomDropdownElementRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDropdownElementRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sLabel"] = o.SLabel
	toSerialize["sValue"] = o.SValue
	return toSerialize, nil
}

func (o *CustomDropdownElementRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sLabel",
		"sValue",
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

	varCustomDropdownElementRequestCompound := _CustomDropdownElementRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomDropdownElementRequestCompound)

	if err != nil {
		return err
	}

	*o = CustomDropdownElementRequestCompound(varCustomDropdownElementRequestCompound)

	return err
}

type NullableCustomDropdownElementRequestCompound struct {
	value *CustomDropdownElementRequestCompound
	isSet bool
}

func (v NullableCustomDropdownElementRequestCompound) Get() *CustomDropdownElementRequestCompound {
	return v.value
}

func (v *NullableCustomDropdownElementRequestCompound) Set(val *CustomDropdownElementRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDropdownElementRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDropdownElementRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDropdownElementRequestCompound(val *CustomDropdownElementRequestCompound) *NullableCustomDropdownElementRequestCompound {
	return &NullableCustomDropdownElementRequestCompound{value: val, isSet: true}
}

func (v NullableCustomDropdownElementRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDropdownElementRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


