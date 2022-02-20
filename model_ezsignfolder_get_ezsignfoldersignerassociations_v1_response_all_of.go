/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf struct for EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf
type EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf struct {
	MPayload EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf(mPayload EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOfWithDefaults instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOfWithDefaults() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) GetMPayload() EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) GetMPayloadOk() (*EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) SetMPayload(v EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf struct {
	value *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) Get() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) Set(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf {
	return &NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


