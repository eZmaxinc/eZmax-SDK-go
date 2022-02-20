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

// EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf struct for EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf
type EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf struct {
	MPayload EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf(mPayload EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOfWithDefaults instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOfWithDefaults() *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) GetMPayload() EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	if o == nil {
		var ret EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) GetMPayloadOk() (*EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) SetMPayload(v EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf struct {
	value *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf
	isSet bool
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) Get() *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) Set(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf {
	return &NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


