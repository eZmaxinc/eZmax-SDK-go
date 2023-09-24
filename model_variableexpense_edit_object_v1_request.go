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
)

// checks if the VariableexpenseEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseEditObjectV1Request{}

// VariableexpenseEditObjectV1Request Request for PUT /1/object/variableexpense/{pkiVariableexpenseID}
type VariableexpenseEditObjectV1Request struct {
	ObjVariableexpense VariableexpenseRequestCompound `json:"objVariableexpense"`
}

// NewVariableexpenseEditObjectV1Request instantiates a new VariableexpenseEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseEditObjectV1Request(objVariableexpense VariableexpenseRequestCompound) *VariableexpenseEditObjectV1Request {
	this := VariableexpenseEditObjectV1Request{}
	this.ObjVariableexpense = objVariableexpense
	return &this
}

// NewVariableexpenseEditObjectV1RequestWithDefaults instantiates a new VariableexpenseEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseEditObjectV1RequestWithDefaults() *VariableexpenseEditObjectV1Request {
	this := VariableexpenseEditObjectV1Request{}
	return &this
}

// GetObjVariableexpense returns the ObjVariableexpense field value
func (o *VariableexpenseEditObjectV1Request) GetObjVariableexpense() VariableexpenseRequestCompound {
	if o == nil {
		var ret VariableexpenseRequestCompound
		return ret
	}

	return o.ObjVariableexpense
}

// GetObjVariableexpenseOk returns a tuple with the ObjVariableexpense field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseEditObjectV1Request) GetObjVariableexpenseOk() (*VariableexpenseRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjVariableexpense, true
}

// SetObjVariableexpense sets field value
func (o *VariableexpenseEditObjectV1Request) SetObjVariableexpense(v VariableexpenseRequestCompound) {
	o.ObjVariableexpense = v
}

func (o VariableexpenseEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objVariableexpense"] = o.ObjVariableexpense
	return toSerialize, nil
}

type NullableVariableexpenseEditObjectV1Request struct {
	value *VariableexpenseEditObjectV1Request
	isSet bool
}

func (v NullableVariableexpenseEditObjectV1Request) Get() *VariableexpenseEditObjectV1Request {
	return v.value
}

func (v *NullableVariableexpenseEditObjectV1Request) Set(val *VariableexpenseEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseEditObjectV1Request(val *VariableexpenseEditObjectV1Request) *NullableVariableexpenseEditObjectV1Request {
	return &NullableVariableexpenseEditObjectV1Request{value: val, isSet: true}
}

func (v NullableVariableexpenseEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

