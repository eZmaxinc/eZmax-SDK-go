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

// EzsigndocumentGetEzsignpagesV1ResponseMPayload Payload for the /1/object/ezsigndocument/{pkiEzsigndocument}/getEzsignpages API Request
type EzsigndocumentGetEzsignpagesV1ResponseMPayload struct {
	AObjEzsignpage []EzsignpageResponse `json:"a_objEzsignpage"`
}

// NewEzsigndocumentGetEzsignpagesV1ResponseMPayload instantiates a new EzsigndocumentGetEzsignpagesV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignpagesV1ResponseMPayload(aObjEzsignpage []EzsignpageResponse) *EzsigndocumentGetEzsignpagesV1ResponseMPayload {
	this := EzsigndocumentGetEzsignpagesV1ResponseMPayload{}
	this.AObjEzsignpage = aObjEzsignpage
	return &this
}

// NewEzsigndocumentGetEzsignpagesV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetEzsignpagesV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignpagesV1ResponseMPayloadWithDefaults() *EzsigndocumentGetEzsignpagesV1ResponseMPayload {
	this := EzsigndocumentGetEzsignpagesV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignpage returns the AObjEzsignpage field value
func (o *EzsigndocumentGetEzsignpagesV1ResponseMPayload) GetAObjEzsignpage() []EzsignpageResponse {
	if o == nil {
		var ret []EzsignpageResponse
		return ret
	}

	return o.AObjEzsignpage
}

// GetAObjEzsignpageOk returns a tuple with the AObjEzsignpage field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignpagesV1ResponseMPayload) GetAObjEzsignpageOk() (*[]EzsignpageResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjEzsignpage, true
}

// SetAObjEzsignpage sets field value
func (o *EzsigndocumentGetEzsignpagesV1ResponseMPayload) SetAObjEzsignpage(v []EzsignpageResponse) {
	o.AObjEzsignpage = v
}

func (o EzsigndocumentGetEzsignpagesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objEzsignpage"] = o.AObjEzsignpage
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload struct {
	value *EzsigndocumentGetEzsignpagesV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) Get() *EzsigndocumentGetEzsignpagesV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) Set(val *EzsigndocumentGetEzsignpagesV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignpagesV1ResponseMPayload(val *EzsigndocumentGetEzsignpagesV1ResponseMPayload) *NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload {
	return &NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignpagesV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

