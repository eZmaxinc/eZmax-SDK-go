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

// EzsigndocumentGetFormDataV1ResponseMPayload Payload for the /1/object/ezsigndocument/{pkiEzsigndocument}/getFormData API Request
type EzsigndocumentGetFormDataV1ResponseMPayload struct {
	// The unique ID of the Ezsigndocument
	PkiEzsigndocumentID int32 `json:"pkiEzsigndocumentID"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// The name of the document that will be presented to Ezsignfoldersignerassociations
	SEzsigndocumentName string `json:"sEzsigndocumentName"`
	// The date and time at which the object was last modified
	DtModifiedDate string `json:"dtModifiedDate"`
	AObjFormDataSigner []CustomFormDataSignerResponse `json:"a_objFormDataSigner"`
}

// NewEzsigndocumentGetFormDataV1ResponseMPayload instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetFormDataV1ResponseMPayload(pkiEzsigndocumentID int32, fkiEzsignfolderID int32, sEzsigndocumentName string, dtModifiedDate string, aObjFormDataSigner []CustomFormDataSignerResponse) *EzsigndocumentGetFormDataV1ResponseMPayload {
	this := EzsigndocumentGetFormDataV1ResponseMPayload{}
	this.PkiEzsigndocumentID = pkiEzsigndocumentID
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.SEzsigndocumentName = sEzsigndocumentName
	this.DtModifiedDate = dtModifiedDate
	this.AObjFormDataSigner = aObjFormDataSigner
	return &this
}

// NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults() *EzsigndocumentGetFormDataV1ResponseMPayload {
	this := EzsigndocumentGetFormDataV1ResponseMPayload{}
	return &this
}

// GetPkiEzsigndocumentID returns the PkiEzsigndocumentID field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetPkiEzsigndocumentID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigndocumentID
}

// GetPkiEzsigndocumentIDOk returns a tuple with the PkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetPkiEzsigndocumentIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsigndocumentID, true
}

// SetPkiEzsigndocumentID sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetPkiEzsigndocumentID(v int32) {
	o.PkiEzsigndocumentID = v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetSEzsigndocumentName returns the SEzsigndocumentName field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetSEzsigndocumentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigndocumentName
}

// GetSEzsigndocumentNameOk returns a tuple with the SEzsigndocumentName field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetSEzsigndocumentNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsigndocumentName, true
}

// SetSEzsigndocumentName sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetSEzsigndocumentName(v string) {
	o.SEzsigndocumentName = v
}

// GetDtModifiedDate returns the DtModifiedDate field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetDtModifiedDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtModifiedDate
}

// GetDtModifiedDateOk returns a tuple with the DtModifiedDate field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetDtModifiedDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtModifiedDate, true
}

// SetDtModifiedDate sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetDtModifiedDate(v string) {
	o.DtModifiedDate = v
}

// GetAObjFormDataSigner returns the AObjFormDataSigner field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetAObjFormDataSigner() []CustomFormDataSignerResponse {
	if o == nil {
		var ret []CustomFormDataSignerResponse
		return ret
	}

	return o.AObjFormDataSigner
}

// GetAObjFormDataSignerOk returns a tuple with the AObjFormDataSigner field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetAObjFormDataSignerOk() (*[]CustomFormDataSignerResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjFormDataSigner, true
}

// SetAObjFormDataSigner sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetAObjFormDataSigner(v []CustomFormDataSignerResponse) {
	o.AObjFormDataSigner = v
}

func (o EzsigndocumentGetFormDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsigndocumentID"] = o.PkiEzsigndocumentID
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	if true {
		toSerialize["sEzsigndocumentName"] = o.SEzsigndocumentName
	}
	if true {
		toSerialize["dtModifiedDate"] = o.DtModifiedDate
	}
	if true {
		toSerialize["a_objFormDataSigner"] = o.AObjFormDataSigner
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetFormDataV1ResponseMPayload struct {
	value *EzsigndocumentGetFormDataV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) Get() *EzsigndocumentGetFormDataV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) Set(val *EzsigndocumentGetFormDataV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetFormDataV1ResponseMPayload(val *EzsigndocumentGetFormDataV1ResponseMPayload) *NullableEzsigndocumentGetFormDataV1ResponseMPayload {
	return &NullableEzsigndocumentGetFormDataV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


