/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.47
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfoldersignerassociationRequest An Ezsignfoldersignerassociation Object
type EzsignfoldersignerassociationRequest struct {
	// A reference to a valid User.  This is only used if the signatory will be a user from the system.
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation.
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
}

// NewEzsignfoldersignerassociationRequest instantiates a new EzsignfoldersignerassociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationRequest(fkiEzsignfolderID int32) *EzsignfoldersignerassociationRequest {
	this := EzsignfoldersignerassociationRequest{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	return &this
}

// NewEzsignfoldersignerassociationRequestWithDefaults instantiates a new EzsignfoldersignerassociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationRequestWithDefaults() *EzsignfoldersignerassociationRequest {
	this := EzsignfoldersignerassociationRequest{}
	return &this
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequest) GetFkiUserID() int32 {
	if o == nil || o.FkiUserID == nil {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequest) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || o.FkiUserID == nil {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequest) HasFkiUserID() bool {
	if o != nil && o.FkiUserID != nil {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignfoldersignerassociationRequest) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignfolderID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequest) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsignfoldersignerassociationRequest) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

func (o EzsignfoldersignerassociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FkiUserID != nil {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationRequest struct {
	value *EzsignfoldersignerassociationRequest
	isSet bool
}

func (v NullableEzsignfoldersignerassociationRequest) Get() *EzsignfoldersignerassociationRequest {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationRequest) Set(val *EzsignfoldersignerassociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationRequest(val *EzsignfoldersignerassociationRequest) *NullableEzsignfoldersignerassociationRequest {
	return &NullableEzsignfoldersignerassociationRequest{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


