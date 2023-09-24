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

// checks if the EzsigndocumentEditEzsignformfieldgroupsV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEditEzsignformfieldgroupsV1Request{}

// EzsigndocumentEditEzsignformfieldgroupsV1Request Request for PUT /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignformfieldgroups
type EzsigndocumentEditEzsignformfieldgroupsV1Request struct {
	AObjEzsignformfieldgroup []EzsignformfieldgroupRequestCompound `json:"a_objEzsignformfieldgroup"`
}

// NewEzsigndocumentEditEzsignformfieldgroupsV1Request instantiates a new EzsigndocumentEditEzsignformfieldgroupsV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEditEzsignformfieldgroupsV1Request(aObjEzsignformfieldgroup []EzsignformfieldgroupRequestCompound) *EzsigndocumentEditEzsignformfieldgroupsV1Request {
	this := EzsigndocumentEditEzsignformfieldgroupsV1Request{}
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsigndocumentEditEzsignformfieldgroupsV1RequestWithDefaults instantiates a new EzsigndocumentEditEzsignformfieldgroupsV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEditEzsignformfieldgroupsV1RequestWithDefaults() *EzsigndocumentEditEzsignformfieldgroupsV1Request {
	this := EzsigndocumentEditEzsignformfieldgroupsV1Request{}
	return &this
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentEditEzsignformfieldgroupsV1Request) GetAObjEzsignformfieldgroup() []EzsignformfieldgroupRequestCompound {
	if o == nil {
		var ret []EzsignformfieldgroupRequestCompound
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEditEzsignformfieldgroupsV1Request) GetAObjEzsignformfieldgroupOk() ([]EzsignformfieldgroupRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentEditEzsignformfieldgroupsV1Request) SetAObjEzsignformfieldgroup(v []EzsignformfieldgroupRequestCompound) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsigndocumentEditEzsignformfieldgroupsV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEditEzsignformfieldgroupsV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

type NullableEzsigndocumentEditEzsignformfieldgroupsV1Request struct {
	value *EzsigndocumentEditEzsignformfieldgroupsV1Request
	isSet bool
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) Get() *EzsigndocumentEditEzsignformfieldgroupsV1Request {
	return v.value
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) Set(val *EzsigndocumentEditEzsignformfieldgroupsV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEditEzsignformfieldgroupsV1Request(val *EzsigndocumentEditEzsignformfieldgroupsV1Request) *NullableEzsigndocumentEditEzsignformfieldgroupsV1Request {
	return &NullableEzsigndocumentEditEzsignformfieldgroupsV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


