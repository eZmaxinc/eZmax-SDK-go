/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfoldersignerassociationCreateObjectV1ResponseAllOf struct for EzsignfoldersignerassociationCreateObjectV1ResponseAllOf
type EzsignfoldersignerassociationCreateObjectV1ResponseAllOf struct {
	MPayload EzsignfoldersignerassociationCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfoldersignerassociationCreateObjectV1ResponseAllOf instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV1ResponseAllOf(mPayload EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV1ResponseAllOfWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV1ResponseAllOfWithDefaults() *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf {
	this := EzsignfoldersignerassociationCreateObjectV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) GetMPayload() EzsignfoldersignerassociationCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignfoldersignerassociationCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) GetMPayloadOk() (*EzsignfoldersignerassociationCreateObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) SetMPayload(v EzsignfoldersignerassociationCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf struct {
	value *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) Get() *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) Set(val *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf(val *EzsignfoldersignerassociationCreateObjectV1ResponseAllOf) *NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf {
	return &NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


