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

// checks if the EzsigndocumentGetCompletedElementsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetCompletedElementsV1ResponseMPayload{}

// EzsigndocumentGetCompletedElementsV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocumentID}/getCompletedElements
type EzsigndocumentGetCompletedElementsV1ResponseMPayload struct {
	AObjEzsignsignature []EzsignsignatureResponseCompound `json:"a_objEzsignsignature"`
	AObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound `json:"a_objEzsignformfieldgroup"`
}

// NewEzsigndocumentGetCompletedElementsV1ResponseMPayload instantiates a new EzsigndocumentGetCompletedElementsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetCompletedElementsV1ResponseMPayload(aObjEzsignsignature []EzsignsignatureResponseCompound, aObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound) *EzsigndocumentGetCompletedElementsV1ResponseMPayload {
	this := EzsigndocumentGetCompletedElementsV1ResponseMPayload{}
	this.AObjEzsignsignature = aObjEzsignsignature
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsigndocumentGetCompletedElementsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetCompletedElementsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetCompletedElementsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetCompletedElementsV1ResponseMPayload {
	this := EzsigndocumentGetCompletedElementsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) GetAObjEzsignsignature() []EzsignsignatureResponseCompound {
	if o == nil {
		var ret []EzsignsignatureResponseCompound
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) GetAObjEzsignsignatureOk() ([]EzsignsignatureResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) SetAObjEzsignsignature(v []EzsignsignatureResponseCompound) {
	o.AObjEzsignsignature = v
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) GetAObjEzsignformfieldgroup() []EzsignformfieldgroupResponseCompound {
	if o == nil {
		var ret []EzsignformfieldgroupResponseCompound
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) GetAObjEzsignformfieldgroupOk() ([]EzsignformfieldgroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentGetCompletedElementsV1ResponseMPayload) SetAObjEzsignformfieldgroup(v []EzsignformfieldgroupResponseCompound) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsigndocumentGetCompletedElementsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetCompletedElementsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

type NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload struct {
	value *EzsigndocumentGetCompletedElementsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) Get() *EzsigndocumentGetCompletedElementsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) Set(val *EzsigndocumentGetCompletedElementsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetCompletedElementsV1ResponseMPayload(val *EzsigndocumentGetCompletedElementsV1ResponseMPayload) *NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload {
	return &NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetCompletedElementsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


