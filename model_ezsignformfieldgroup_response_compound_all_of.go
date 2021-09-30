/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignformfieldgroupResponseCompoundAllOf struct for EzsignformfieldgroupResponseCompoundAllOf
type EzsignformfieldgroupResponseCompoundAllOf struct {
	AObjEzsignformfield []EzsignformfieldResponse `json:"a_objEzsignformfield"`
}

// NewEzsignformfieldgroupResponseCompoundAllOf instantiates a new EzsignformfieldgroupResponseCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupResponseCompoundAllOf(aObjEzsignformfield []EzsignformfieldResponse) *EzsignformfieldgroupResponseCompoundAllOf {
	this := EzsignformfieldgroupResponseCompoundAllOf{}
	this.AObjEzsignformfield = aObjEzsignformfield
	return &this
}

// NewEzsignformfieldgroupResponseCompoundAllOfWithDefaults instantiates a new EzsignformfieldgroupResponseCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupResponseCompoundAllOfWithDefaults() *EzsignformfieldgroupResponseCompoundAllOf {
	this := EzsignformfieldgroupResponseCompoundAllOf{}
	return &this
}

// GetAObjEzsignformfield returns the AObjEzsignformfield field value
func (o *EzsignformfieldgroupResponseCompoundAllOf) GetAObjEzsignformfield() []EzsignformfieldResponse {
	if o == nil {
		var ret []EzsignformfieldResponse
		return ret
	}

	return o.AObjEzsignformfield
}

// GetAObjEzsignformfieldOk returns a tuple with the AObjEzsignformfield field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupResponseCompoundAllOf) GetAObjEzsignformfieldOk() (*[]EzsignformfieldResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjEzsignformfield, true
}

// SetAObjEzsignformfield sets field value
func (o *EzsignformfieldgroupResponseCompoundAllOf) SetAObjEzsignformfield(v []EzsignformfieldResponse) {
	o.AObjEzsignformfield = v
}

func (o EzsignformfieldgroupResponseCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objEzsignformfield"] = o.AObjEzsignformfield
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignformfieldgroupResponseCompoundAllOf struct {
	value *EzsignformfieldgroupResponseCompoundAllOf
	isSet bool
}

func (v NullableEzsignformfieldgroupResponseCompoundAllOf) Get() *EzsignformfieldgroupResponseCompoundAllOf {
	return v.value
}

func (v *NullableEzsignformfieldgroupResponseCompoundAllOf) Set(val *EzsignformfieldgroupResponseCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupResponseCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupResponseCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupResponseCompoundAllOf(val *EzsignformfieldgroupResponseCompoundAllOf) *NullableEzsignformfieldgroupResponseCompoundAllOf {
	return &NullableEzsignformfieldgroupResponseCompoundAllOf{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupResponseCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupResponseCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


