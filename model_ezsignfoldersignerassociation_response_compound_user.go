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

// EzsignfoldersignerassociationResponseCompoundUser A Ezsignfoldersignerassociation->User Object and children to create a complete structure
type EzsignfoldersignerassociationResponseCompoundUser struct {
	// The unique ID of the User
	PkiUserID int32 `json:"pkiUserID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	// The First name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The Last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
}

// NewEzsignfoldersignerassociationResponseCompoundUser instantiates a new EzsignfoldersignerassociationResponseCompoundUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationResponseCompoundUser(pkiUserID int32, fkiLanguageID int32, sUserFirstname string, sUserLastname string, sEmailAddress string) *EzsignfoldersignerassociationResponseCompoundUser {
	this := EzsignfoldersignerassociationResponseCompoundUser{}
	this.PkiUserID = pkiUserID
	this.FkiLanguageID = fkiLanguageID
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SEmailAddress = sEmailAddress
	return &this
}

// NewEzsignfoldersignerassociationResponseCompoundUserWithDefaults instantiates a new EzsignfoldersignerassociationResponseCompoundUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationResponseCompoundUserWithDefaults() *EzsignfoldersignerassociationResponseCompoundUser {
	this := EzsignfoldersignerassociationResponseCompoundUser{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetPkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetPkiUserIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiUserID, true
}

// SetPkiUserID sets field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) SetPkiUserID(v int32) {
	o.PkiUserID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSUserFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSUserLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationResponseCompoundUser) GetSEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *EzsignfoldersignerassociationResponseCompoundUser) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

func (o EzsignfoldersignerassociationResponseCompoundUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiUserID"] = o.PkiUserID
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["sUserFirstname"] = o.SUserFirstname
	}
	if true {
		toSerialize["sUserLastname"] = o.SUserLastname
	}
	if true {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationResponseCompoundUser struct {
	value *EzsignfoldersignerassociationResponseCompoundUser
	isSet bool
}

func (v NullableEzsignfoldersignerassociationResponseCompoundUser) Get() *EzsignfoldersignerassociationResponseCompoundUser {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationResponseCompoundUser) Set(val *EzsignfoldersignerassociationResponseCompoundUser) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationResponseCompoundUser) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationResponseCompoundUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationResponseCompoundUser(val *EzsignfoldersignerassociationResponseCompoundUser) *NullableEzsignfoldersignerassociationResponseCompoundUser {
	return &NullableEzsignfoldersignerassociationResponseCompoundUser{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationResponseCompoundUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationResponseCompoundUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


