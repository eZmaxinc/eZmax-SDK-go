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

// checks if the CustomFormDataEzsignformfieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFormDataEzsignformfieldResponse{}

// CustomFormDataEzsignformfieldResponse An Ezsignformfield Object
type CustomFormDataEzsignformfieldResponse struct {
	// The Label for the Ezsignformfield
	SEzsignformfieldLabel string `json:"sEzsignformfieldLabel"`
	// The value for the Ezsignformfield  This can only be set if eEzsignformfieldgroupType is Checkbox or Radio
	SEzsignformfieldValue string `json:"sEzsignformfieldValue"`
}

type _CustomFormDataEzsignformfieldResponse CustomFormDataEzsignformfieldResponse

// NewCustomFormDataEzsignformfieldResponse instantiates a new CustomFormDataEzsignformfieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormDataEzsignformfieldResponse(sEzsignformfieldLabel string, sEzsignformfieldValue string) *CustomFormDataEzsignformfieldResponse {
	this := CustomFormDataEzsignformfieldResponse{}
	this.SEzsignformfieldLabel = sEzsignformfieldLabel
	this.SEzsignformfieldValue = sEzsignformfieldValue
	return &this
}

// NewCustomFormDataEzsignformfieldResponseWithDefaults instantiates a new CustomFormDataEzsignformfieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormDataEzsignformfieldResponseWithDefaults() *CustomFormDataEzsignformfieldResponse {
	this := CustomFormDataEzsignformfieldResponse{}
	return &this
}

// GetSEzsignformfieldLabel returns the SEzsignformfieldLabel field value
func (o *CustomFormDataEzsignformfieldResponse) GetSEzsignformfieldLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldLabel
}

// GetSEzsignformfieldLabelOk returns a tuple with the SEzsignformfieldLabel field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataEzsignformfieldResponse) GetSEzsignformfieldLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldLabel, true
}

// SetSEzsignformfieldLabel sets field value
func (o *CustomFormDataEzsignformfieldResponse) SetSEzsignformfieldLabel(v string) {
	o.SEzsignformfieldLabel = v
}

// GetSEzsignformfieldValue returns the SEzsignformfieldValue field value
func (o *CustomFormDataEzsignformfieldResponse) GetSEzsignformfieldValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignformfieldValue
}

// GetSEzsignformfieldValueOk returns a tuple with the SEzsignformfieldValue field value
// and a boolean to check if the value has been set.
func (o *CustomFormDataEzsignformfieldResponse) GetSEzsignformfieldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignformfieldValue, true
}

// SetSEzsignformfieldValue sets field value
func (o *CustomFormDataEzsignformfieldResponse) SetSEzsignformfieldValue(v string) {
	o.SEzsignformfieldValue = v
}

func (o CustomFormDataEzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFormDataEzsignformfieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEzsignformfieldLabel"] = o.SEzsignformfieldLabel
	toSerialize["sEzsignformfieldValue"] = o.SEzsignformfieldValue
	return toSerialize, nil
}

func (o *CustomFormDataEzsignformfieldResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sEzsignformfieldLabel",
		"sEzsignformfieldValue",
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

	varCustomFormDataEzsignformfieldResponse := _CustomFormDataEzsignformfieldResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomFormDataEzsignformfieldResponse)

	if err != nil {
		return err
	}

	*o = CustomFormDataEzsignformfieldResponse(varCustomFormDataEzsignformfieldResponse)

	return err
}

type NullableCustomFormDataEzsignformfieldResponse struct {
	value *CustomFormDataEzsignformfieldResponse
	isSet bool
}

func (v NullableCustomFormDataEzsignformfieldResponse) Get() *CustomFormDataEzsignformfieldResponse {
	return v.value
}

func (v *NullableCustomFormDataEzsignformfieldResponse) Set(val *CustomFormDataEzsignformfieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormDataEzsignformfieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormDataEzsignformfieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormDataEzsignformfieldResponse(val *CustomFormDataEzsignformfieldResponse) *NullableCustomFormDataEzsignformfieldResponse {
	return &NullableCustomFormDataEzsignformfieldResponse{value: val, isSet: true}
}

func (v NullableCustomFormDataEzsignformfieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormDataEzsignformfieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


