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

// checks if the EzsigndocumentGetEzsignsignaturesV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetEzsignsignaturesV1ResponseMPayload{}

// EzsigndocumentGetEzsignsignaturesV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getEzsignsignatures
type EzsigndocumentGetEzsignsignaturesV1ResponseMPayload struct {
	AObjEzsignsignature []EzsignsignatureResponseCompound `json:"a_objEzsignsignature"`
}

// NewEzsigndocumentGetEzsignsignaturesV1ResponseMPayload instantiates a new EzsigndocumentGetEzsignsignaturesV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignsignaturesV1ResponseMPayload(aObjEzsignsignature []EzsignsignatureResponseCompound) *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload {
	this := EzsigndocumentGetEzsignsignaturesV1ResponseMPayload{}
	this.AObjEzsignsignature = aObjEzsignsignature
	return &this
}

// NewEzsigndocumentGetEzsignsignaturesV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetEzsignsignaturesV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignsignaturesV1ResponseMPayloadWithDefaults() *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload {
	this := EzsigndocumentGetEzsignsignaturesV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignsignature returns the AObjEzsignsignature field value
func (o *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) GetAObjEzsignsignature() []EzsignsignatureResponseCompound {
	if o == nil {
		var ret []EzsignsignatureResponseCompound
		return ret
	}

	return o.AObjEzsignsignature
}

// GetAObjEzsignsignatureOk returns a tuple with the AObjEzsignsignature field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) GetAObjEzsignsignatureOk() ([]EzsignsignatureResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignsignature, true
}

// SetAObjEzsignsignature sets field value
func (o *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) SetAObjEzsignsignature(v []EzsignsignatureResponseCompound) {
	o.AObjEzsignsignature = v
}

func (o EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignsignature"] = o.AObjEzsignsignature
	return toSerialize, nil
}

type NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload struct {
	value *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) Get() *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) Set(val *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload(val *EzsigndocumentGetEzsignsignaturesV1ResponseMPayload) *NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload {
	return &NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignsignaturesV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


