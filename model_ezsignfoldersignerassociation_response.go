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

// EzsignfoldersignerassociationResponse An Ezsignfoldersignerassociation Object
type EzsignfoldersignerassociationResponse struct {
	// The unique ID of the Ezsignfoldersignerassociation
	PkiEzsignfoldersignerassociationID int32 `json:"pkiEzsignfoldersignerassociationID"`
	// The unique ID of the Ezsignfolder
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
	// If this flag is true. The signatory will receive a copy of every signed Ezsigndocument even if it ain't required to sign the document.
	BEzsignfoldersignerassociationReceivecopy bool `json:"bEzsignfoldersignerassociationReceivecopy"`
}

// NewEzsignfoldersignerassociationResponse instantiates a new EzsignfoldersignerassociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationResponse(pkiEzsignfoldersignerassociationID int32, fkiEzsignfolderID int32, bEzsignfoldersignerassociationReceivecopy bool) *EzsignfoldersignerassociationResponse {
	this := EzsignfoldersignerassociationResponse{}
	this.PkiEzsignfoldersignerassociationID = pkiEzsignfoldersignerassociationID
	this.FkiEzsignfolderID = fkiEzsignfolderID
	this.BEzsignfoldersignerassociationReceivecopy = bEzsignfoldersignerassociationReceivecopy
	return &this
}

// NewEzsignfoldersignerassociationResponseWithDefaults instantiates a new EzsignfoldersignerassociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationResponseWithDefaults() *EzsignfoldersignerassociationResponse {
	this := EzsignfoldersignerassociationResponse{}
	return &this
}

// GetPkiEzsignfoldersignerassociationID returns the PkiEzsignfoldersignerassociationID field value
func (o *EzsignfoldersignerassociationResponse) GetPkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfoldersignerassociationID
}

// GetPkiEzsignfoldersignerassociationIDOk returns a tuple with the PkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponse) GetPkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsignfoldersignerassociationID, true
}

// SetPkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfoldersignerassociationResponse) SetPkiEzsignfoldersignerassociationID(v int32) {
	o.PkiEzsignfoldersignerassociationID = v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsignfoldersignerassociationResponse) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponse) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsignfoldersignerassociationResponse) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

// GetBEzsignfoldersignerassociationReceivecopy returns the BEzsignfoldersignerassociationReceivecopy field value
func (o *EzsignfoldersignerassociationResponse) GetBEzsignfoldersignerassociationReceivecopy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignfoldersignerassociationReceivecopy
}

// GetBEzsignfoldersignerassociationReceivecopyOk returns a tuple with the BEzsignfoldersignerassociationReceivecopy field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponse) GetBEzsignfoldersignerassociationReceivecopyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BEzsignfoldersignerassociationReceivecopy, true
}

// SetBEzsignfoldersignerassociationReceivecopy sets field value
func (o *EzsignfoldersignerassociationResponse) SetBEzsignfoldersignerassociationReceivecopy(v bool) {
	o.BEzsignfoldersignerassociationReceivecopy = v
}

func (o EzsignfoldersignerassociationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsignfoldersignerassociationID"] = o.PkiEzsignfoldersignerassociationID
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	if true {
		toSerialize["bEzsignfoldersignerassociationReceivecopy"] = o.BEzsignfoldersignerassociationReceivecopy
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationResponse struct {
	value *EzsignfoldersignerassociationResponse
	isSet bool
}

func (v NullableEzsignfoldersignerassociationResponse) Get() *EzsignfoldersignerassociationResponse {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationResponse) Set(val *EzsignfoldersignerassociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationResponse(val *EzsignfoldersignerassociationResponse) *NullableEzsignfoldersignerassociationResponse {
	return &NullableEzsignfoldersignerassociationResponse{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

