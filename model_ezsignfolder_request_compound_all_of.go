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

// EzsignfolderRequestCompoundAllOf struct for EzsignfolderRequestCompoundAllOf
type EzsignfolderRequestCompoundAllOf struct {
	// An array of signers that will be invited to sign the Ezsigndocuments
	AEzsignfoldersignerassociation []EzsignfoldersignerassociationRequest `json:"a_Ezsignfoldersignerassociation"`
}

// NewEzsignfolderRequestCompoundAllOf instantiates a new EzsignfolderRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderRequestCompoundAllOf(aEzsignfoldersignerassociation []EzsignfoldersignerassociationRequest) *EzsignfolderRequestCompoundAllOf {
	this := EzsignfolderRequestCompoundAllOf{}
	this.AEzsignfoldersignerassociation = aEzsignfoldersignerassociation
	return &this
}

// NewEzsignfolderRequestCompoundAllOfWithDefaults instantiates a new EzsignfolderRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderRequestCompoundAllOfWithDefaults() *EzsignfolderRequestCompoundAllOf {
	this := EzsignfolderRequestCompoundAllOf{}
	return &this
}

// GetAEzsignfoldersignerassociation returns the AEzsignfoldersignerassociation field value
func (o *EzsignfolderRequestCompoundAllOf) GetAEzsignfoldersignerassociation() []EzsignfoldersignerassociationRequest {
	if o == nil {
		var ret []EzsignfoldersignerassociationRequest
		return ret
	}

	return o.AEzsignfoldersignerassociation
}

// GetAEzsignfoldersignerassociationOk returns a tuple with the AEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderRequestCompoundAllOf) GetAEzsignfoldersignerassociationOk() (*[]EzsignfoldersignerassociationRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AEzsignfoldersignerassociation, true
}

// SetAEzsignfoldersignerassociation sets field value
func (o *EzsignfolderRequestCompoundAllOf) SetAEzsignfoldersignerassociation(v []EzsignfoldersignerassociationRequest) {
	o.AEzsignfoldersignerassociation = v
}

func (o EzsignfolderRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_Ezsignfoldersignerassociation"] = o.AEzsignfoldersignerassociation
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderRequestCompoundAllOf struct {
	value *EzsignfolderRequestCompoundAllOf
	isSet bool
}

func (v NullableEzsignfolderRequestCompoundAllOf) Get() *EzsignfolderRequestCompoundAllOf {
	return v.value
}

func (v *NullableEzsignfolderRequestCompoundAllOf) Set(val *EzsignfolderRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderRequestCompoundAllOf(val *EzsignfolderRequestCompoundAllOf) *NullableEzsignfolderRequestCompoundAllOf {
	return &NullableEzsignfolderRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableEzsignfolderRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


