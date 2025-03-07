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

// checks if the EzsigndocumentExtractTextV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentExtractTextV1ResponseMPayload{}

// EzsigndocumentExtractTextV1ResponseMPayload Response for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/ExtractText
type EzsigndocumentExtractTextV1ResponseMPayload struct {
	// The text extract from document
	SText string `json:"sText"`
}

type _EzsigndocumentExtractTextV1ResponseMPayload EzsigndocumentExtractTextV1ResponseMPayload

// NewEzsigndocumentExtractTextV1ResponseMPayload instantiates a new EzsigndocumentExtractTextV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentExtractTextV1ResponseMPayload(sText string) *EzsigndocumentExtractTextV1ResponseMPayload {
	this := EzsigndocumentExtractTextV1ResponseMPayload{}
	this.SText = sText
	return &this
}

// NewEzsigndocumentExtractTextV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentExtractTextV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentExtractTextV1ResponseMPayloadWithDefaults() *EzsigndocumentExtractTextV1ResponseMPayload {
	this := EzsigndocumentExtractTextV1ResponseMPayload{}
	return &this
}

// GetSText returns the SText field value
func (o *EzsigndocumentExtractTextV1ResponseMPayload) GetSText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SText
}

// GetSTextOk returns a tuple with the SText field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentExtractTextV1ResponseMPayload) GetSTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SText, true
}

// SetSText sets field value
func (o *EzsigndocumentExtractTextV1ResponseMPayload) SetSText(v string) {
	o.SText = v
}

func (o EzsigndocumentExtractTextV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentExtractTextV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sText"] = o.SText
	return toSerialize, nil
}

func (o *EzsigndocumentExtractTextV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sText",
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

	varEzsigndocumentExtractTextV1ResponseMPayload := _EzsigndocumentExtractTextV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentExtractTextV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentExtractTextV1ResponseMPayload(varEzsigndocumentExtractTextV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentExtractTextV1ResponseMPayload struct {
	value *EzsigndocumentExtractTextV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentExtractTextV1ResponseMPayload) Get() *EzsigndocumentExtractTextV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentExtractTextV1ResponseMPayload) Set(val *EzsigndocumentExtractTextV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentExtractTextV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentExtractTextV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentExtractTextV1ResponseMPayload(val *EzsigndocumentExtractTextV1ResponseMPayload) *NullableEzsigndocumentExtractTextV1ResponseMPayload {
	return &NullableEzsigndocumentExtractTextV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentExtractTextV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentExtractTextV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


