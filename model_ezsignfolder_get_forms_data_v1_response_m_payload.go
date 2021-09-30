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

// EzsignfolderGetFormsDataV1ResponseMPayload Payload for the /1/object/ezsignfolder/{pkiEzsigndocument}/getFormsData API Request
type EzsignfolderGetFormsDataV1ResponseMPayload struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The description of the Ezsign Folder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	AObjFormDataDocument []CustomFormDataDocumentResponse `json:"a_objFormDataDocument"`
}

// NewEzsignfolderGetFormsDataV1ResponseMPayload instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetFormsDataV1ResponseMPayload(pkiEzsignfolderID int32, sEzsignfolderDescription string, aObjFormDataDocument []CustomFormDataDocumentResponse) *EzsignfolderGetFormsDataV1ResponseMPayload {
	this := EzsignfolderGetFormsDataV1ResponseMPayload{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.AObjFormDataDocument = aObjFormDataDocument
	return &this
}

// NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults() *EzsignfolderGetFormsDataV1ResponseMPayload {
	this := EzsignfolderGetFormsDataV1ResponseMPayload{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetAObjFormDataDocument returns the AObjFormDataDocument field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetAObjFormDataDocument() []CustomFormDataDocumentResponse {
	if o == nil {
		var ret []CustomFormDataDocumentResponse
		return ret
	}

	return o.AObjFormDataDocument
}

// GetAObjFormDataDocumentOk returns a tuple with the AObjFormDataDocument field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetAObjFormDataDocumentOk() (*[]CustomFormDataDocumentResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjFormDataDocument, true
}

// SetAObjFormDataDocument sets field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetAObjFormDataDocument(v []CustomFormDataDocumentResponse) {
	o.AObjFormDataDocument = v
}

func (o EzsignfolderGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsignfolderID"] = o.PkiEzsignfolderID
	}
	if true {
		toSerialize["sEzsignfolderDescription"] = o.SEzsignfolderDescription
	}
	if true {
		toSerialize["a_objFormDataDocument"] = o.AObjFormDataDocument
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderGetFormsDataV1ResponseMPayload struct {
	value *EzsignfolderGetFormsDataV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) Get() *EzsignfolderGetFormsDataV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) Set(val *EzsignfolderGetFormsDataV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetFormsDataV1ResponseMPayload(val *EzsignfolderGetFormsDataV1ResponseMPayload) *NullableEzsignfolderGetFormsDataV1ResponseMPayload {
	return &NullableEzsignfolderGetFormsDataV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


