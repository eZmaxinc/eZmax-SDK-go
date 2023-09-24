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

// checks if the UsergroupdelegationCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupdelegationCreateObjectV1Request{}

// UsergroupdelegationCreateObjectV1Request Request for POST /1/object/usergroupdelegation
type UsergroupdelegationCreateObjectV1Request struct {
	AObjUsergroupdelegation []UsergroupdelegationRequestCompound `json:"a_objUsergroupdelegation"`
}

// NewUsergroupdelegationCreateObjectV1Request instantiates a new UsergroupdelegationCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupdelegationCreateObjectV1Request(aObjUsergroupdelegation []UsergroupdelegationRequestCompound) *UsergroupdelegationCreateObjectV1Request {
	this := UsergroupdelegationCreateObjectV1Request{}
	this.AObjUsergroupdelegation = aObjUsergroupdelegation
	return &this
}

// NewUsergroupdelegationCreateObjectV1RequestWithDefaults instantiates a new UsergroupdelegationCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupdelegationCreateObjectV1RequestWithDefaults() *UsergroupdelegationCreateObjectV1Request {
	this := UsergroupdelegationCreateObjectV1Request{}
	return &this
}

// GetAObjUsergroupdelegation returns the AObjUsergroupdelegation field value
func (o *UsergroupdelegationCreateObjectV1Request) GetAObjUsergroupdelegation() []UsergroupdelegationRequestCompound {
	if o == nil {
		var ret []UsergroupdelegationRequestCompound
		return ret
	}

	return o.AObjUsergroupdelegation
}

// GetAObjUsergroupdelegationOk returns a tuple with the AObjUsergroupdelegation field value
// and a boolean to check if the value has been set.
func (o *UsergroupdelegationCreateObjectV1Request) GetAObjUsergroupdelegationOk() ([]UsergroupdelegationRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUsergroupdelegation, true
}

// SetAObjUsergroupdelegation sets field value
func (o *UsergroupdelegationCreateObjectV1Request) SetAObjUsergroupdelegation(v []UsergroupdelegationRequestCompound) {
	o.AObjUsergroupdelegation = v
}

func (o UsergroupdelegationCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupdelegationCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUsergroupdelegation"] = o.AObjUsergroupdelegation
	return toSerialize, nil
}

type NullableUsergroupdelegationCreateObjectV1Request struct {
	value *UsergroupdelegationCreateObjectV1Request
	isSet bool
}

func (v NullableUsergroupdelegationCreateObjectV1Request) Get() *UsergroupdelegationCreateObjectV1Request {
	return v.value
}

func (v *NullableUsergroupdelegationCreateObjectV1Request) Set(val *UsergroupdelegationCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupdelegationCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupdelegationCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupdelegationCreateObjectV1Request(val *UsergroupdelegationCreateObjectV1Request) *NullableUsergroupdelegationCreateObjectV1Request {
	return &NullableUsergroupdelegationCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableUsergroupdelegationCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupdelegationCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

