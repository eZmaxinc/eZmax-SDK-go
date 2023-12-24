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

// checks if the EzsignsigningreasonCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsigningreasonCreateObjectV1Request{}

// EzsignsigningreasonCreateObjectV1Request Request for POST /1/object/ezsignsigningreason
type EzsignsigningreasonCreateObjectV1Request struct {
	AObjEzsignsigningreason []EzsignsigningreasonRequestCompound `json:"a_objEzsignsigningreason"`
}

type _EzsignsigningreasonCreateObjectV1Request EzsignsigningreasonCreateObjectV1Request

// NewEzsignsigningreasonCreateObjectV1Request instantiates a new EzsignsigningreasonCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsigningreasonCreateObjectV1Request(aObjEzsignsigningreason []EzsignsigningreasonRequestCompound) *EzsignsigningreasonCreateObjectV1Request {
	this := EzsignsigningreasonCreateObjectV1Request{}
	this.AObjEzsignsigningreason = aObjEzsignsigningreason
	return &this
}

// NewEzsignsigningreasonCreateObjectV1RequestWithDefaults instantiates a new EzsignsigningreasonCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsigningreasonCreateObjectV1RequestWithDefaults() *EzsignsigningreasonCreateObjectV1Request {
	this := EzsignsigningreasonCreateObjectV1Request{}
	return &this
}

// GetAObjEzsignsigningreason returns the AObjEzsignsigningreason field value
func (o *EzsignsigningreasonCreateObjectV1Request) GetAObjEzsignsigningreason() []EzsignsigningreasonRequestCompound {
	if o == nil {
		var ret []EzsignsigningreasonRequestCompound
		return ret
	}

	return o.AObjEzsignsigningreason
}

// GetAObjEzsignsigningreasonOk returns a tuple with the AObjEzsignsigningreason field value
// and a boolean to check if the value has been set.
func (o *EzsignsigningreasonCreateObjectV1Request) GetAObjEzsignsigningreasonOk() ([]EzsignsigningreasonRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsigningreason, true
}

// SetAObjEzsignsigningreason sets field value
func (o *EzsignsigningreasonCreateObjectV1Request) SetAObjEzsignsigningreason(v []EzsignsigningreasonRequestCompound) {
	o.AObjEzsignsigningreason = v
}

func (o EzsignsigningreasonCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsigningreasonCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsigningreason"] = o.AObjEzsignsigningreason
	return toSerialize, nil
}

func (o *EzsignsigningreasonCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignsigningreason",
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

	varEzsignsigningreasonCreateObjectV1Request := _EzsignsigningreasonCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignsigningreasonCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = EzsignsigningreasonCreateObjectV1Request(varEzsignsigningreasonCreateObjectV1Request)

	return err
}

type NullableEzsignsigningreasonCreateObjectV1Request struct {
	value *EzsignsigningreasonCreateObjectV1Request
	isSet bool
}

func (v NullableEzsignsigningreasonCreateObjectV1Request) Get() *EzsignsigningreasonCreateObjectV1Request {
	return v.value
}

func (v *NullableEzsignsigningreasonCreateObjectV1Request) Set(val *EzsignsigningreasonCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsigningreasonCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsigningreasonCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsigningreasonCreateObjectV1Request(val *EzsignsigningreasonCreateObjectV1Request) *NullableEzsignsigningreasonCreateObjectV1Request {
	return &NullableEzsignsigningreasonCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignsigningreasonCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsigningreasonCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


