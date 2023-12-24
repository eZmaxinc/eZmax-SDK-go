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

// checks if the SignatureCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureCreateObjectV1Request{}

// SignatureCreateObjectV1Request Request for POST /1/object/signature
type SignatureCreateObjectV1Request struct {
	AObjSignature []SignatureRequestCompound `json:"a_objSignature"`
}

type _SignatureCreateObjectV1Request SignatureCreateObjectV1Request

// NewSignatureCreateObjectV1Request instantiates a new SignatureCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureCreateObjectV1Request(aObjSignature []SignatureRequestCompound) *SignatureCreateObjectV1Request {
	this := SignatureCreateObjectV1Request{}
	this.AObjSignature = aObjSignature
	return &this
}

// NewSignatureCreateObjectV1RequestWithDefaults instantiates a new SignatureCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureCreateObjectV1RequestWithDefaults() *SignatureCreateObjectV1Request {
	this := SignatureCreateObjectV1Request{}
	return &this
}

// GetAObjSignature returns the AObjSignature field value
func (o *SignatureCreateObjectV1Request) GetAObjSignature() []SignatureRequestCompound {
	if o == nil {
		var ret []SignatureRequestCompound
		return ret
	}

	return o.AObjSignature
}

// GetAObjSignatureOk returns a tuple with the AObjSignature field value
// and a boolean to check if the value has been set.
func (o *SignatureCreateObjectV1Request) GetAObjSignatureOk() ([]SignatureRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjSignature, true
}

// SetAObjSignature sets field value
func (o *SignatureCreateObjectV1Request) SetAObjSignature(v []SignatureRequestCompound) {
	o.AObjSignature = v
}

func (o SignatureCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objSignature"] = o.AObjSignature
	return toSerialize, nil
}

func (o *SignatureCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objSignature",
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

	varSignatureCreateObjectV1Request := _SignatureCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = SignatureCreateObjectV1Request(varSignatureCreateObjectV1Request)

	return err
}

type NullableSignatureCreateObjectV1Request struct {
	value *SignatureCreateObjectV1Request
	isSet bool
}

func (v NullableSignatureCreateObjectV1Request) Get() *SignatureCreateObjectV1Request {
	return v.value
}

func (v *NullableSignatureCreateObjectV1Request) Set(val *SignatureCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureCreateObjectV1Request(val *SignatureCreateObjectV1Request) *NullableSignatureCreateObjectV1Request {
	return &NullableSignatureCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableSignatureCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


