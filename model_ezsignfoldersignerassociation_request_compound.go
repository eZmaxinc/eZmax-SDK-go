/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.28
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfoldersignerassociationRequestCompound An Ezsignfoldersignerassociation Object and children to create a complete structure
type EzsignfoldersignerassociationRequestCompound struct {
	ObjEzsignsigner *EzsignsignerRequestCompound `json:"objEzsignsigner,omitempty"`
	// A reference to a valid User.  This is only used if the signatory will be a user from the system.
	FkiUserID *int32 `json:"fkiUserID,omitempty"`
	// A reference to a valid Ezsignfolder.  That value is returned after a successful Ezsignfolder Creation.
	FkiEzsignfolderID int32 `json:"fkiEzsignfolderID"`
}

// NewEzsignfoldersignerassociationRequestCompound instantiates a new EzsignfoldersignerassociationRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationRequestCompound(fkiEzsignfolderID int32, ) *EzsignfoldersignerassociationRequestCompound {
	this := EzsignfoldersignerassociationRequestCompound{}
	this.FkiEzsignfolderID = fkiEzsignfolderID
	return &this
}

// NewEzsignfoldersignerassociationRequestCompoundWithDefaults instantiates a new EzsignfoldersignerassociationRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationRequestCompoundWithDefaults() *EzsignfoldersignerassociationRequestCompound {
	this := EzsignfoldersignerassociationRequestCompound{}
	return &this
}

// GetObjEzsignsigner returns the ObjEzsignsigner field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsigner() EzsignsignerRequestCompound {
	if o == nil || o.ObjEzsignsigner == nil {
		var ret EzsignsignerRequestCompound
		return ret
	}
	return *o.ObjEzsignsigner
}

// GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetObjEzsignsignerOk() (*EzsignsignerRequestCompound, bool) {
	if o == nil || o.ObjEzsignsigner == nil {
		return nil, false
	}
	return o.ObjEzsignsigner, true
}

// HasObjEzsignsigner returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasObjEzsignsigner() bool {
	if o != nil && o.ObjEzsignsigner != nil {
		return true
	}

	return false
}

// SetObjEzsignsigner gets a reference to the given EzsignsignerRequestCompound and assigns it to the ObjEzsignsigner field.
func (o *EzsignfoldersignerassociationRequestCompound) SetObjEzsignsigner(v EzsignsignerRequestCompound) {
	o.ObjEzsignsigner = &v
}

// GetFkiUserID returns the FkiUserID field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserID() int32 {
	if o == nil || o.FkiUserID == nil {
		var ret int32
		return ret
	}
	return *o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiUserIDOk() (*int32, bool) {
	if o == nil || o.FkiUserID == nil {
		return nil, false
	}
	return o.FkiUserID, true
}

// HasFkiUserID returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompound) HasFkiUserID() bool {
	if o != nil && o.FkiUserID != nil {
		return true
	}

	return false
}

// SetFkiUserID gets a reference to the given int32 and assigns it to the FkiUserID field.
func (o *EzsignfoldersignerassociationRequestCompound) SetFkiUserID(v int32) {
	o.FkiUserID = &v
}

// GetFkiEzsignfolderID returns the FkiEzsignfolderID field value
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiEzsignfolderID
}

// GetFkiEzsignfolderIDOk returns a tuple with the FkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompound) GetFkiEzsignfolderIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEzsignfolderID, true
}

// SetFkiEzsignfolderID sets field value
func (o *EzsignfoldersignerassociationRequestCompound) SetFkiEzsignfolderID(v int32) {
	o.FkiEzsignfolderID = v
}

func (o EzsignfoldersignerassociationRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignsigner != nil {
		toSerialize["objEzsignsigner"] = o.ObjEzsignsigner
	}
	if o.FkiUserID != nil {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if true {
		toSerialize["fkiEzsignfolderID"] = o.FkiEzsignfolderID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationRequestCompound struct {
	value *EzsignfoldersignerassociationRequestCompound
	isSet bool
}

func (v NullableEzsignfoldersignerassociationRequestCompound) Get() *EzsignfoldersignerassociationRequestCompound {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) Set(val *EzsignfoldersignerassociationRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationRequestCompound(val *EzsignfoldersignerassociationRequestCompound) *NullableEzsignfoldersignerassociationRequestCompound {
	return &NullableEzsignfoldersignerassociationRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

