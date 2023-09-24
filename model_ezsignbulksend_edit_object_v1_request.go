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

// checks if the EzsignbulksendEditObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendEditObjectV1Request{}

// EzsignbulksendEditObjectV1Request Request for PUT /1/object/ezsignbulksend/{pkiEzsignbulksendID}
type EzsignbulksendEditObjectV1Request struct {
	ObjEzsignbulksend EzsignbulksendRequestCompound `json:"objEzsignbulksend"`
}

// NewEzsignbulksendEditObjectV1Request instantiates a new EzsignbulksendEditObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendEditObjectV1Request(objEzsignbulksend EzsignbulksendRequestCompound) *EzsignbulksendEditObjectV1Request {
	this := EzsignbulksendEditObjectV1Request{}
	this.ObjEzsignbulksend = objEzsignbulksend
	return &this
}

// NewEzsignbulksendEditObjectV1RequestWithDefaults instantiates a new EzsignbulksendEditObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendEditObjectV1RequestWithDefaults() *EzsignbulksendEditObjectV1Request {
	this := EzsignbulksendEditObjectV1Request{}
	return &this
}

// GetObjEzsignbulksend returns the ObjEzsignbulksend field value
func (o *EzsignbulksendEditObjectV1Request) GetObjEzsignbulksend() EzsignbulksendRequestCompound {
	if o == nil {
		var ret EzsignbulksendRequestCompound
		return ret
	}

	return o.ObjEzsignbulksend
}

// GetObjEzsignbulksendOk returns a tuple with the ObjEzsignbulksend field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendEditObjectV1Request) GetObjEzsignbulksendOk() (*EzsignbulksendRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjEzsignbulksend, true
}

// SetObjEzsignbulksend sets field value
func (o *EzsignbulksendEditObjectV1Request) SetObjEzsignbulksend(v EzsignbulksendRequestCompound) {
	o.ObjEzsignbulksend = v
}

func (o EzsignbulksendEditObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendEditObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objEzsignbulksend"] = o.ObjEzsignbulksend
	return toSerialize, nil
}

type NullableEzsignbulksendEditObjectV1Request struct {
	value *EzsignbulksendEditObjectV1Request
	isSet bool
}

func (v NullableEzsignbulksendEditObjectV1Request) Get() *EzsignbulksendEditObjectV1Request {
	return v.value
}

func (v *NullableEzsignbulksendEditObjectV1Request) Set(val *EzsignbulksendEditObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendEditObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendEditObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendEditObjectV1Request(val *EzsignbulksendEditObjectV1Request) *NullableEzsignbulksendEditObjectV1Request {
	return &NullableEzsignbulksendEditObjectV1Request{value: val, isSet: true}
}

func (v NullableEzsignbulksendEditObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendEditObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

