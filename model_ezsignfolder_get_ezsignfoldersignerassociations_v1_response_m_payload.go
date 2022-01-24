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

// EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload Payload for the /1/object/ezsignfolder/{pkiEzsignfolder}/getEzsignfoldersignerassociations API Request
type EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload struct {
	AObjEzsignfoldersignerassociation []EzsignfoldersignerassociationResponse `json:"a_objEzsignfoldersignerassociation"`
}

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload(aObjEzsignfoldersignerassociation []EzsignfoldersignerassociationResponse) *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}
	this.AObjEzsignfoldersignerassociation = aObjEzsignfoldersignerassociation
	return &this
}

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignfoldersignerassociation returns the AObjEzsignfoldersignerassociation field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) GetAObjEzsignfoldersignerassociation() []EzsignfoldersignerassociationResponse {
	if o == nil {
		var ret []EzsignfoldersignerassociationResponse
		return ret
	}

	return o.AObjEzsignfoldersignerassociation
}

// GetAObjEzsignfoldersignerassociationOk returns a tuple with the AObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) GetAObjEzsignfoldersignerassociationOk() ([]EzsignfoldersignerassociationResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AObjEzsignfoldersignerassociation, true
}

// SetAObjEzsignfoldersignerassociation sets field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) SetAObjEzsignfoldersignerassociation(v []EzsignfoldersignerassociationResponse) {
	o.AObjEzsignfoldersignerassociation = v
}

func (o EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objEzsignfoldersignerassociation"] = o.AObjEzsignfoldersignerassociation
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload struct {
	value *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Get() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Set(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	return &NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


