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

// EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload Payload for the /1/object/ezsignfoldersignerassociation/getInPersonLoginUrl API Request
type EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload struct {
	// The Url to login to the signing application.    Url will expire after 30 minutes.  
	SLoginUrl string `json:"sLoginUrl"`
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload(sLoginUrl string) *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload{}
	this.SLoginUrl = sLoginUrl
	return &this
}

// NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	this := EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload{}
	return &this
}

// GetSLoginUrl returns the SLoginUrl field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) GetSLoginUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SLoginUrl
}

// GetSLoginUrlOk returns a tuple with the SLoginUrl field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) GetSLoginUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SLoginUrl, true
}

// SetSLoginUrl sets field value
func (o *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) SetSLoginUrl(v string) {
	o.SLoginUrl = v
}

func (o EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sLoginUrl"] = o.SLoginUrl
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload struct {
	value *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) Get() *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) Set(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload(val *EzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload {
	return &NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationGetInPersonLoginUrlV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


