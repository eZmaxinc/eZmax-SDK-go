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

// checks if the EzsigndocumentEditEzsignsignaturesV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEditEzsignsignaturesV1Request{}

// EzsigndocumentEditEzsignsignaturesV1Request Request for PUT /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignsignatures
type EzsigndocumentEditEzsignsignaturesV1Request struct {
	AObjEzsignsignature []EzsignsignatureRequestCompound `json:"a_objEzsignsignature"`
}

// NewEzsigndocumentEditEzsignsignaturesV1Request instantiates a new EzsigndocumentEditEzsignsignaturesV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEditEzsignsignaturesV1Request(aObjEzsignsignature []EzsignsignatureRequestCompound) *EzsigndocumentEditEzsignsignaturesV1Request {
	this := EzsigndocumentEditEzsignsignaturesV1Request{}
	this.AObjEzsignsignature = aObjEzsignsignature
	return &this
}

// NewEzsigndocumentEditEzsignsignaturesV1RequestWithDefaults instantiates a new EzsigndocumentEditEzsignsignaturesV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEditEzsignsignaturesV1RequestWithDefaults() *EzsigndocumentEditEzsignsignaturesV1Request {
	this := EzsigndocumentEditEzsignsignaturesV1Request{}
	return &this
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsigndocumentEditEzsignsignaturesV1Request) GetAObjEzsignsignature() []EzsignsignatureRequestCompound {
	if o == nil {
		var ret []EzsignsignatureRequestCompound
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEditEzsignsignaturesV1Request) GetAObjEzsignsignatureOk() ([]EzsignsignatureRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsigndocumentEditEzsignsignaturesV1Request) SetAObjEzsignsignature(v []EzsignsignatureRequestCompound) {
	o.AObjEzsignsignature = v
}

func (o EzsigndocumentEditEzsignsignaturesV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEditEzsignsignaturesV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	return toSerialize, nil
}

type NullableEzsigndocumentEditEzsignsignaturesV1Request struct {
	value *EzsigndocumentEditEzsignsignaturesV1Request
	isSet bool
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1Request) Get() *EzsigndocumentEditEzsignsignaturesV1Request {
	return v.value
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1Request) Set(val *EzsigndocumentEditEzsignsignaturesV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEditEzsignsignaturesV1Request(val *EzsigndocumentEditEzsignsignaturesV1Request) *NullableEzsigndocumentEditEzsignsignaturesV1Request {
	return &NullableEzsigndocumentEditEzsignsignaturesV1Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


