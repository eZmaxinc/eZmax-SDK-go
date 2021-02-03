/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// UserResponse A User Object
type UserResponse struct {
	// The unique ID of the User
	PkiUserID int32 `json:"pkiUserID"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EUserType FieldEUserType `json:"eUserType"`
	// The First name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The Last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The Login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	ObjAudit CommonAudit `json:"objAudit"`
}

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse(pkiUserID int32, fkiLanguageID int32, eUserType FieldEUserType, sUserFirstname string, sUserLastname string, sUserLoginname string, objAudit CommonAudit, ) *UserResponse {
	this := UserResponse{}
	this.PkiUserID = pkiUserID
	this.FkiLanguageID = fkiLanguageID
	this.EUserType = eUserType
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.ObjAudit = objAudit
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetPkiUserID returns the PkiUserID field value
func (o *UserResponse) GetPkiUserID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.PkiUserID
}

// GetPkiUserIDOk returns a tuple with the PkiUserID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetPkiUserIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiUserID, true
}

// SetPkiUserID sets field value
func (o *UserResponse) SetPkiUserID(v int32) {
	o.PkiUserID = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *UserResponse) GetFkiLanguageID() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *UserResponse) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEUserType returns the EUserType field value
func (o *UserResponse) GetEUserType() FieldEUserType {
	if o == nil  {
		var ret FieldEUserType
		return ret
	}

	return o.EUserType
}

// GetEUserTypeOk returns a tuple with the EUserType field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEUserTypeOk() (*FieldEUserType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EUserType, true
}

// SetEUserType sets field value
func (o *UserResponse) SetEUserType(v FieldEUserType) {
	o.EUserType = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UserResponse) GetSUserFirstname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UserResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UserResponse) GetSUserLastname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UserResponse) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UserResponse) GetSUserLoginname() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetSUserLoginnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UserResponse) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetObjAudit returns the ObjAudit field value
func (o *UserResponse) GetObjAudit() CommonAudit {
	if o == nil  {
		var ret CommonAudit
		return ret
	}

	return o.ObjAudit
}

// GetObjAuditOk returns a tuple with the ObjAudit field value
// and a boolean to check if the value has been set.
func (o *UserResponse) GetObjAuditOk() (*CommonAudit, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjAudit, true
}

// SetObjAudit sets field value
func (o *UserResponse) SetObjAudit(v CommonAudit) {
	o.ObjAudit = v
}

func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiUserID"] = o.PkiUserID
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["eUserType"] = o.EUserType
	}
	if true {
		toSerialize["sUserFirstname"] = o.SUserFirstname
	}
	if true {
		toSerialize["sUserLastname"] = o.SUserLastname
	}
	if true {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if true {
		toSerialize["objAudit"] = o.ObjAudit
	}
	return json.Marshal(toSerialize)
}

type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


