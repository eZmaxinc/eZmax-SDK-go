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

// checks if the EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload{}

// EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload Payload for GET /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocument}/getEzsigntemplatesignatures
type EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload struct {
	AObjEzsigntemplatesignature []EzsigntemplatesignatureResponseCompound `json:"a_objEzsigntemplatesignature"`
}

// NewEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload instantiates a new EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload(aObjEzsigntemplatesignature []EzsigntemplatesignatureResponseCompound) *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload {
	this := EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload{}
	this.AObjEzsigntemplatesignature = aObjEzsigntemplatesignature
	return &this
}

// NewEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayloadWithDefaults() *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload {
	this := EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload{}
	return &this
}

// GetAObjEzsigntemplatesignature returns the AObjEzsigntemplatesignature field value
func (o *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) GetAObjEzsigntemplatesignature() []EzsigntemplatesignatureResponseCompound {
	if o == nil {
		var ret []EzsigntemplatesignatureResponseCompound
		return ret
	}

	return o.AObjEzsigntemplatesignature
}

// GetAObjEzsigntemplatesignatureOk returns a tuple with the AObjEzsigntemplatesignature field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) GetAObjEzsigntemplatesignatureOk() ([]EzsigntemplatesignatureResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplatesignature, true
}

// SetAObjEzsigntemplatesignature sets field value
func (o *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) SetAObjEzsigntemplatesignature(v []EzsigntemplatesignatureResponseCompound) {
	o.AObjEzsigntemplatesignature = v
}

func (o EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplatesignature"] = o.AObjEzsigntemplatesignature
	return toSerialize, nil
}

type NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload struct {
	value *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) Get() *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) Set(val *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload(val *EzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) *NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload {
	return &NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentGetEzsigntemplatesignaturesV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


