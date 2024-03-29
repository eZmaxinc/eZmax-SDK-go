/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignbulksendGetListV1ResponseMPayloadAllOf struct for EzsignbulksendGetListV1ResponseMPayloadAllOf
type EzsignbulksendGetListV1ResponseMPayloadAllOf struct {
	AObjEzsignbulksend []EzsignbulksendListElement `json:"a_objEzsignbulksend"`
}

// NewEzsignbulksendGetListV1ResponseMPayloadAllOf instantiates a new EzsignbulksendGetListV1ResponseMPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendGetListV1ResponseMPayloadAllOf(aObjEzsignbulksend []EzsignbulksendListElement) *EzsignbulksendGetListV1ResponseMPayloadAllOf {
	this := EzsignbulksendGetListV1ResponseMPayloadAllOf{}
	this.AObjEzsignbulksend = aObjEzsignbulksend
	return &this
}

// NewEzsignbulksendGetListV1ResponseMPayloadAllOfWithDefaults instantiates a new EzsignbulksendGetListV1ResponseMPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendGetListV1ResponseMPayloadAllOfWithDefaults() *EzsignbulksendGetListV1ResponseMPayloadAllOf {
	this := EzsignbulksendGetListV1ResponseMPayloadAllOf{}
	return &this
}

// GetAObjEzsignbulksend returns the AObjEzsignbulksend field value
func (o *EzsignbulksendGetListV1ResponseMPayloadAllOf) GetAObjEzsignbulksend() []EzsignbulksendListElement {
	if o == nil {
		var ret []EzsignbulksendListElement
		return ret
	}

	return o.AObjEzsignbulksend
}

// GetAObjEzsignbulksendOk returns a tuple with the AObjEzsignbulksend field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetListV1ResponseMPayloadAllOf) GetAObjEzsignbulksendOk() (*[]EzsignbulksendListElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjEzsignbulksend, true
}

// SetAObjEzsignbulksend sets field value
func (o *EzsignbulksendGetListV1ResponseMPayloadAllOf) SetAObjEzsignbulksend(v []EzsignbulksendListElement) {
	o.AObjEzsignbulksend = v
}

func (o EzsignbulksendGetListV1ResponseMPayloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objEzsignbulksend"] = o.AObjEzsignbulksend
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignbulksendGetListV1ResponseMPayloadAllOf struct {
	value *EzsignbulksendGetListV1ResponseMPayloadAllOf
	isSet bool
}

func (v NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) Get() *EzsignbulksendGetListV1ResponseMPayloadAllOf {
	return v.value
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) Set(val *EzsignbulksendGetListV1ResponseMPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendGetListV1ResponseMPayloadAllOf(val *EzsignbulksendGetListV1ResponseMPayloadAllOf) *NullableEzsignbulksendGetListV1ResponseMPayloadAllOf {
	return &NullableEzsignbulksendGetListV1ResponseMPayloadAllOf{value: val, isSet: true}
}

func (v NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendGetListV1ResponseMPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


