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

// CustomFormsDataFolderResponse A forms Data Folder Object
type CustomFormsDataFolderResponse struct {
	// The unique ID of the Ezsignfolder
	PkiEzsignfolderID int32 `json:"pkiEzsignfolderID"`
	// The description of the Ezsignfolder
	SEzsignfolderDescription string `json:"sEzsignfolderDescription"`
	AObjFormDataDocument []CustomFormDataDocumentResponse `json:"a_objFormDataDocument"`
}

// NewCustomFormsDataFolderResponse instantiates a new CustomFormsDataFolderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormsDataFolderResponse(pkiEzsignfolderID int32, sEzsignfolderDescription string, aObjFormDataDocument []CustomFormDataDocumentResponse) *CustomFormsDataFolderResponse {
	this := CustomFormsDataFolderResponse{}
	this.PkiEzsignfolderID = pkiEzsignfolderID
	this.SEzsignfolderDescription = sEzsignfolderDescription
	this.AObjFormDataDocument = aObjFormDataDocument
	return &this
}

// NewCustomFormsDataFolderResponseWithDefaults instantiates a new CustomFormsDataFolderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormsDataFolderResponseWithDefaults() *CustomFormsDataFolderResponse {
	this := CustomFormsDataFolderResponse{}
	return &this
}

// GetPkiEzsignfolderID returns the PkiEzsignfolderID field value
func (o *CustomFormsDataFolderResponse) GetPkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfolderID
}

// GetPkiEzsignfolderIDOk returns a tuple with the PkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *CustomFormsDataFolderResponse) GetPkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignfolderID, true
}

// SetPkiEzsignfolderID sets field value
func (o *CustomFormsDataFolderResponse) SetPkiEzsignfolderID(v int32) {
	o.PkiEzsignfolderID = v
}

// GetSEzsignfolderDescription returns the SEzsignfolderDescription field value
func (o *CustomFormsDataFolderResponse) GetSEzsignfolderDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfolderDescription
}

// GetSEzsignfolderDescriptionOk returns a tuple with the SEzsignfolderDescription field value
// and a boolean to check if the value has been set.
func (o *CustomFormsDataFolderResponse) GetSEzsignfolderDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsignfolderDescription, true
}

// SetSEzsignfolderDescription sets field value
func (o *CustomFormsDataFolderResponse) SetSEzsignfolderDescription(v string) {
	o.SEzsignfolderDescription = v
}

// GetAObjFormDataDocument returns the AObjFormDataDocument field value
func (o *CustomFormsDataFolderResponse) GetAObjFormDataDocument() []CustomFormDataDocumentResponse {
	if o == nil {
		var ret []CustomFormDataDocumentResponse
		return ret
	}

	return o.AObjFormDataDocument
}

// GetAObjFormDataDocumentOk returns a tuple with the AObjFormDataDocument field value
// and a boolean to check if the value has been set.
func (o *CustomFormsDataFolderResponse) GetAObjFormDataDocumentOk() ([]CustomFormDataDocumentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFormDataDocument, true
}

// SetAObjFormDataDocument sets field value
func (o *CustomFormsDataFolderResponse) SetAObjFormDataDocument(v []CustomFormDataDocumentResponse) {
	o.AObjFormDataDocument = v
}

func (o CustomFormsDataFolderResponse) MarshalJSON() ([]byte, error) {
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

type NullableCustomFormsDataFolderResponse struct {
	value *CustomFormsDataFolderResponse
	isSet bool
}

func (v NullableCustomFormsDataFolderResponse) Get() *CustomFormsDataFolderResponse {
	return v.value
}

func (v *NullableCustomFormsDataFolderResponse) Set(val *CustomFormsDataFolderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormsDataFolderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormsDataFolderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormsDataFolderResponse(val *CustomFormsDataFolderResponse) *NullableCustomFormsDataFolderResponse {
	return &NullableCustomFormsDataFolderResponse{value: val, isSet: true}
}

func (v NullableCustomFormsDataFolderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormsDataFolderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


