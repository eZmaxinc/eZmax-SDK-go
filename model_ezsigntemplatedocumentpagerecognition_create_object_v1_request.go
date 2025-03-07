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

// checks if the EzsigntemplatedocumentpagerecognitionCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentpagerecognitionCreateObjectV1Request{}

// EzsigntemplatedocumentpagerecognitionCreateObjectV1Request Request for POST /1/object/ezsigntemplatedocumentpagerecognition
type EzsigntemplatedocumentpagerecognitionCreateObjectV1Request struct {
	AObjEzsigntemplatedocumentpagerecognition []EzsigntemplatedocumentpagerecognitionRequestCompound `json:"a_objEzsigntemplatedocumentpagerecognition"`
}

type _EzsigntemplatedocumentpagerecognitionCreateObjectV1Request EzsigntemplatedocumentpagerecognitionCreateObjectV1Request

// NewEzsigntemplatedocumentpagerecognitionCreateObjectV1Request instantiates a new EzsigntemplatedocumentpagerecognitionCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentpagerecognitionCreateObjectV1Request(aObjEzsigntemplatedocumentpagerecognition []EzsigntemplatedocumentpagerecognitionRequestCompound) *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	this := EzsigntemplatedocumentpagerecognitionCreateObjectV1Request{}
	this.AObjEzsigntemplatedocumentpagerecognition = aObjEzsigntemplatedocumentpagerecognition
	return &this
}

// NewEzsigntemplatedocumentpagerecognitionCreateObjectV1RequestWithDefaults instantiates a new EzsigntemplatedocumentpagerecognitionCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentpagerecognitionCreateObjectV1RequestWithDefaults() *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	this := EzsigntemplatedocumentpagerecognitionCreateObjectV1Request{}
	return &this
}

// GetAObjEzsigntemplatedocumentpagerecognition returns the AObjEzsigntemplatedocumentpagerecognition field value
func (o *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) GetAObjEzsigntemplatedocumentpagerecognition() []EzsigntemplatedocumentpagerecognitionRequestCompound {
	if o == nil {
		var ret []EzsigntemplatedocumentpagerecognitionRequestCompound
		return ret
	}

	return o.AObjEzsigntemplatedocumentpagerecognition
}

// GetAObjEzsigntemplatedocumentpagerecognitionOk returns a tuple with the AObjEzsigntemplatedocumentpagerecognition field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) GetAObjEzsigntemplatedocumentpagerecognitionOk() ([]EzsigntemplatedocumentpagerecognitionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatedocumentpagerecognition, true
}

// SetAObjEzsigntemplatedocumentpagerecognition sets field value
func (o *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) SetAObjEzsigntemplatedocumentpagerecognition(v []EzsigntemplatedocumentpagerecognitionRequestCompound) {
	o.AObjEzsigntemplatedocumentpagerecognition = v
}

func (o EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplatedocumentpagerecognition"] = o.AObjEzsigntemplatedocumentpagerecognition
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsigntemplatedocumentpagerecognition",
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

	varEzsigntemplatedocumentpagerecognitionCreateObjectV1Request := _EzsigntemplatedocumentpagerecognitionCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentpagerecognitionCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentpagerecognitionCreateObjectV1Request(varEzsigntemplatedocumentpagerecognitionCreateObjectV1Request)

	return err
}

type NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request struct {
	value *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request
	isSet bool
}

func (v NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) Get() *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) Set(val *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request(val *EzsigntemplatedocumentpagerecognitionCreateObjectV1Request) *NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request {
	return &NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentpagerecognitionCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


