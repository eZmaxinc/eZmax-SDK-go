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

// checks if the ApikeyCreateObjectV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyCreateObjectV2Request{}

// ApikeyCreateObjectV2Request Request for POST /2/object/apikey
type ApikeyCreateObjectV2Request struct {
	AObjApikey []ApikeyRequestCompound `json:"a_objApikey"`
}

type _ApikeyCreateObjectV2Request ApikeyCreateObjectV2Request

// NewApikeyCreateObjectV2Request instantiates a new ApikeyCreateObjectV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyCreateObjectV2Request(aObjApikey []ApikeyRequestCompound) *ApikeyCreateObjectV2Request {
	this := ApikeyCreateObjectV2Request{}
	this.AObjApikey = aObjApikey
	return &this
}

// NewApikeyCreateObjectV2RequestWithDefaults instantiates a new ApikeyCreateObjectV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyCreateObjectV2RequestWithDefaults() *ApikeyCreateObjectV2Request {
	this := ApikeyCreateObjectV2Request{}
	return &this
}

// GetAObjApikey returns the AObjApikey field value
func (o *ApikeyCreateObjectV2Request) GetAObjApikey() []ApikeyRequestCompound {
	if o == nil {
		var ret []ApikeyRequestCompound
		return ret
	}

	return o.AObjApikey
}

// GetAObjApikeyOk returns a tuple with the AObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV2Request) GetAObjApikeyOk() ([]ApikeyRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjApikey, true
}

// SetAObjApikey sets field value
func (o *ApikeyCreateObjectV2Request) SetAObjApikey(v []ApikeyRequestCompound) {
	o.AObjApikey = v
}

func (o ApikeyCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyCreateObjectV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objApikey"] = o.AObjApikey
	return toSerialize, nil
}

func (o *ApikeyCreateObjectV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objApikey",
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

	varApikeyCreateObjectV2Request := _ApikeyCreateObjectV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyCreateObjectV2Request)

	if err != nil {
		return err
	}

	*o = ApikeyCreateObjectV2Request(varApikeyCreateObjectV2Request)

	return err
}

type NullableApikeyCreateObjectV2Request struct {
	value *ApikeyCreateObjectV2Request
	isSet bool
}

func (v NullableApikeyCreateObjectV2Request) Get() *ApikeyCreateObjectV2Request {
	return v.value
}

func (v *NullableApikeyCreateObjectV2Request) Set(val *ApikeyCreateObjectV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyCreateObjectV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyCreateObjectV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyCreateObjectV2Request(val *ApikeyCreateObjectV2Request) *NullableApikeyCreateObjectV2Request {
	return &NullableApikeyCreateObjectV2Request{value: val, isSet: true}
}

func (v NullableApikeyCreateObjectV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyCreateObjectV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


