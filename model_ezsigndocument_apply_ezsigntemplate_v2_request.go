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

// EzsigndocumentApplyEzsigntemplateV2Request Request for the /2/object/ezsigndocument/{pkiEzsigndocumentID}/applyezsigntemplate API Request
type EzsigndocumentApplyEzsigntemplateV2Request struct {
	// The unique ID of the Ezsigndocument
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// 
	ASEzsigntemplatesigner []string `json:"a_sEzsigntemplatesigner"`
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

// NewEzsigndocumentApplyEzsigntemplateV2Request instantiates a new EzsigndocumentApplyEzsigntemplateV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentApplyEzsigntemplateV2Request(fkiEzsigntemplateID int32, aSEzsigntemplatesigner []string, aPkiEzsignfoldersignerassociationID []int32) *EzsigndocumentApplyEzsigntemplateV2Request {
	this := EzsigndocumentApplyEzsigntemplateV2Request{}
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.ASEzsigntemplatesigner = aSEzsigntemplatesigner
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsigndocumentApplyEzsigntemplateV2RequestWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentApplyEzsigntemplateV2RequestWithDefaults() *EzsigndocumentApplyEzsigntemplateV2Request {
	this := EzsigndocumentApplyEzsigntemplateV2Request{}
	return &this
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetASEzsigntemplatesigner returns the ASEzsigntemplatesigner field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetASEzsigntemplatesigner() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ASEzsigntemplatesigner
}

// GetASEzsigntemplatesignerOk returns a tuple with the ASEzsigntemplatesigner field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetASEzsigntemplatesignerOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ASEzsigntemplatesigner, true
}

// SetASEzsigntemplatesigner sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetASEzsigntemplatesigner(v []string) {
	o.ASEzsigntemplatesigner = v
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Request) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsigndocumentApplyEzsigntemplateV2Request) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsigndocumentApplyEzsigntemplateV2Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	}
	if true {
		toSerialize["a_sEzsigntemplatesigner"] = o.ASEzsigntemplatesigner
	}
	if true {
		toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentApplyEzsigntemplateV2Request struct {
	value *EzsigndocumentApplyEzsigntemplateV2Request
	isSet bool
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) Get() *EzsigndocumentApplyEzsigntemplateV2Request {
	return v.value
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) Set(val *EzsigndocumentApplyEzsigntemplateV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentApplyEzsigntemplateV2Request(val *EzsigndocumentApplyEzsigntemplateV2Request) *NullableEzsigndocumentApplyEzsigntemplateV2Request {
	return &NullableEzsigndocumentApplyEzsigntemplateV2Request{value: val, isSet: true}
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


