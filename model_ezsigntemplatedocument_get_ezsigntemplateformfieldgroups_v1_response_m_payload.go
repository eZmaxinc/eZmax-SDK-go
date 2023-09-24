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

// checks if the EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload{}

// EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload Payload for GET /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocument}/getEzsigntemplateformfieldgroups
type EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload struct {
	AObjEzsigntemplateformfieldgroup []EzsigntemplateformfieldgroupResponseCompound `json:"a_objEzsigntemplateformfieldgroup"`
}

// NewEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload instantiates a new EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload(aObjEzsigntemplateformfieldgroup []EzsigntemplateformfieldgroupResponseCompound) *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload {
	this := EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload{}
	this.AObjEzsigntemplateformfieldgroup = aObjEzsigntemplateformfieldgroup
	return &this
}

// NewEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayloadWithDefaults() *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload {
	this := EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsigntemplateformfieldgroup returns the AObjEzsigntemplateformfieldgroup field value
func (o *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) GetAObjEzsigntemplateformfieldgroup() []EzsigntemplateformfieldgroupResponseCompound {
	if o == nil {
		var ret []EzsigntemplateformfieldgroupResponseCompound
		return ret
	}

	return o.AObjEzsigntemplateformfieldgroup
}

// GetAObjEzsigntemplateformfieldgroupOk returns a tuple with the AObjEzsigntemplateformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) GetAObjEzsigntemplateformfieldgroupOk() ([]EzsigntemplateformfieldgroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplateformfieldgroup, true
}

// SetAObjEzsigntemplateformfieldgroup sets field value
func (o *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) SetAObjEzsigntemplateformfieldgroup(v []EzsigntemplateformfieldgroupResponseCompound) {
	o.AObjEzsigntemplateformfieldgroup = v
}

func (o EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplateformfieldgroup"] = o.AObjEzsigntemplateformfieldgroup
	return toSerialize, nil
}

type NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload struct {
	value *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) Get() *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) Set(val *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload(val *EzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) *NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload {
	return &NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplateformfieldgroupsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


