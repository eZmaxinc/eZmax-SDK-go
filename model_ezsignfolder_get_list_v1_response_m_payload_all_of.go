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

// EzsignfolderGetListV1ResponseMPayloadAllOf struct for EzsignfolderGetListV1ResponseMPayloadAllOf
type EzsignfolderGetListV1ResponseMPayloadAllOf struct {
	AObjEzsignfolder []EzsignfolderListElement `json:"a_objEzsignfolder"`
}

// NewEzsignfolderGetListV1ResponseMPayloadAllOf instantiates a new EzsignfolderGetListV1ResponseMPayloadAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetListV1ResponseMPayloadAllOf(aObjEzsignfolder []EzsignfolderListElement) *EzsignfolderGetListV1ResponseMPayloadAllOf {
	this := EzsignfolderGetListV1ResponseMPayloadAllOf{}
	this.AObjEzsignfolder = aObjEzsignfolder
	return &this
}

// NewEzsignfolderGetListV1ResponseMPayloadAllOfWithDefaults instantiates a new EzsignfolderGetListV1ResponseMPayloadAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetListV1ResponseMPayloadAllOfWithDefaults() *EzsignfolderGetListV1ResponseMPayloadAllOf {
	this := EzsignfolderGetListV1ResponseMPayloadAllOf{}
	return &this
}

// GetAObjEzsignfolder returns the AObjEzsignfolder field value
func (o *EzsignfolderGetListV1ResponseMPayloadAllOf) GetAObjEzsignfolder() []EzsignfolderListElement {
	if o == nil {
		var ret []EzsignfolderListElement
		return ret
	}

	return o.AObjEzsignfolder
}

// GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetListV1ResponseMPayloadAllOf) GetAObjEzsignfolderOk() (*[]EzsignfolderListElement, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjEzsignfolder, true
}

// SetAObjEzsignfolder sets field value
func (o *EzsignfolderGetListV1ResponseMPayloadAllOf) SetAObjEzsignfolder(v []EzsignfolderListElement) {
	o.AObjEzsignfolder = v
}

func (o EzsignfolderGetListV1ResponseMPayloadAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objEzsignfolder"] = o.AObjEzsignfolder
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetListV1ResponseMPayloadAllOf struct {
	value *EzsignfolderGetListV1ResponseMPayloadAllOf
	isSet bool
}

func (v NullableEzsignfolderGetListV1ResponseMPayloadAllOf) Get() *EzsignfolderGetListV1ResponseMPayloadAllOf {
	return v.value
}

func (v *NullableEzsignfolderGetListV1ResponseMPayloadAllOf) Set(val *EzsignfolderGetListV1ResponseMPayloadAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetListV1ResponseMPayloadAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetListV1ResponseMPayloadAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetListV1ResponseMPayloadAllOf(val *EzsignfolderGetListV1ResponseMPayloadAllOf) *NullableEzsignfolderGetListV1ResponseMPayloadAllOf {
	return &NullableEzsignfolderGetListV1ResponseMPayloadAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderGetListV1ResponseMPayloadAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetListV1ResponseMPayloadAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

