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
	"bytes"
	"fmt"
)

// checks if the CustomDropdownElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDropdownElementResponse{}

// CustomDropdownElementResponse Generic DropdownElement Response
type CustomDropdownElementResponse struct {
	// The Description of the element
	SLabel string `json:"sLabel"`
	// The Value of the element
	SValue string `json:"sValue"`
}

type _CustomDropdownElementResponse CustomDropdownElementResponse

// NewCustomDropdownElementResponse instantiates a new CustomDropdownElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDropdownElementResponse(sLabel string, sValue string) *CustomDropdownElementResponse {
	this := CustomDropdownElementResponse{}
	this.SLabel = sLabel
	this.SValue = sValue
	return &this
}

// NewCustomDropdownElementResponseWithDefaults instantiates a new CustomDropdownElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDropdownElementResponseWithDefaults() *CustomDropdownElementResponse {
	this := CustomDropdownElementResponse{}
	return &this
}

// GetSLabel returns the SLabel field value
func (o *CustomDropdownElementResponse) GetSLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLabel
}

// GetSLabelOk returns a tuple with the SLabel field value
// and a boolean to check if the value has been set.
func (o *CustomDropdownElementResponse) GetSLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLabel, true
}

// SetSLabel sets field value
func (o *CustomDropdownElementResponse) SetSLabel(v string) {
	o.SLabel = v
}

// GetSValue returns the SValue field value
func (o *CustomDropdownElementResponse) GetSValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SValue
}

// GetSValueOk returns a tuple with the SValue field value
// and a boolean to check if the value has been set.
func (o *CustomDropdownElementResponse) GetSValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SValue, true
}

// SetSValue sets field value
func (o *CustomDropdownElementResponse) SetSValue(v string) {
	o.SValue = v
}

func (o CustomDropdownElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDropdownElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sLabel"] = o.SLabel
	toSerialize["sValue"] = o.SValue
	return toSerialize, nil
}

func (o *CustomDropdownElementResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCustomDropdownElementResponse := _CustomDropdownElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomDropdownElementResponse)

	if err != nil {
		return err
	}

	*o = CustomDropdownElementResponse(varCustomDropdownElementResponse)

	return err
}

type NullableCustomDropdownElementResponse struct {
	value *CustomDropdownElementResponse
	isSet bool
}

func (v NullableCustomDropdownElementResponse) Get() *CustomDropdownElementResponse {
	return v.value
}

func (v *NullableCustomDropdownElementResponse) Set(val *CustomDropdownElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDropdownElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDropdownElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDropdownElementResponse(val *CustomDropdownElementResponse) *NullableCustomDropdownElementResponse {
	return &NullableCustomDropdownElementResponse{value: val, isSet: true}
}

func (v NullableCustomDropdownElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDropdownElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


