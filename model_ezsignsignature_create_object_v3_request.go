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

// checks if the EzsignsignatureCreateObjectV3Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureCreateObjectV3Request{}

// EzsignsignatureCreateObjectV3Request Request for POST /3/object/ezsignsignature
type EzsignsignatureCreateObjectV3Request struct {
	AObjEzsignsignature []EzsignsignatureRequestCompoundV2 `json:"a_objEzsignsignature"`
}

type _EzsignsignatureCreateObjectV3Request EzsignsignatureCreateObjectV3Request

// NewEzsignsignatureCreateObjectV3Request instantiates a new EzsignsignatureCreateObjectV3Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV3Request(aObjEzsignsignature []EzsignsignatureRequestCompoundV2) *EzsignsignatureCreateObjectV3Request {
	this := EzsignsignatureCreateObjectV3Request{}
	this.AObjEzsignsignature = aObjEzsignsignature
	return &this
}

// NewEzsignsignatureCreateObjectV3RequestWithDefaults instantiates a new EzsignsignatureCreateObjectV3Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV3RequestWithDefaults() *EzsignsignatureCreateObjectV3Request {
	this := EzsignsignatureCreateObjectV3Request{}
	return &this
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsignsignatureCreateObjectV3Request) GetAObjEzsignsignature() []EzsignsignatureRequestCompoundV2 {
	if o == nil {
		var ret []EzsignsignatureRequestCompoundV2
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV3Request) GetAObjEzsignsignatureOk() ([]EzsignsignatureRequestCompoundV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsignsignatureCreateObjectV3Request) SetAObjEzsignsignature(v []EzsignsignatureRequestCompoundV2) {
	o.AObjEzsignsignature = v
}

func (o EzsignsignatureCreateObjectV3Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureCreateObjectV3Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	return toSerialize, nil
}

func (o *EzsignsignatureCreateObjectV3Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignsignature",
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

	varEzsignsignatureCreateObjectV3Request := _EzsignsignatureCreateObjectV3Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsignatureCreateObjectV3Request)

	if err != nil {
		return err
	}

	*o = EzsignsignatureCreateObjectV3Request(varEzsignsignatureCreateObjectV3Request)

	return err
}

type NullableEzsignsignatureCreateObjectV3Request struct {
	value *EzsignsignatureCreateObjectV3Request
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV3Request) Get() *EzsignsignatureCreateObjectV3Request {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV3Request) Set(val *EzsignsignatureCreateObjectV3Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV3Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV3Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV3Request(val *EzsignsignatureCreateObjectV3Request) *NullableEzsignsignatureCreateObjectV3Request {
	return &NullableEzsignsignatureCreateObjectV3Request{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV3Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV3Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


