/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the UserstagedResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserstagedResponseCompound{}

// UserstagedResponseCompound A Userstaged Object
type UserstagedResponseCompound struct {
	// The unique ID of the Userstaged
	PkiUserstagedID int32 `json:"pkiUserstagedID"`
	// The unique ID of the Email
	FkiEmailID int32 `json:"fkiEmailID"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
	// The firstname of the Userstaged
	SUserstagedFirstname string `json:"sUserstagedFirstname"`
	// The lastname of the Userstaged
	SUserstagedLastname string `json:"sUserstagedLastname"`
	// The externalid of the Userstaged
	SUserstagedExternalid string `json:"sUserstagedExternalid"`
}

// NewUserstagedResponseCompound instantiates a new UserstagedResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserstagedResponseCompound(pkiUserstagedID int32, fkiEmailID int32, sEmailAddress string, sUserstagedFirstname string, sUserstagedLastname string, sUserstagedExternalid string) *UserstagedResponseCompound {
	this := UserstagedResponseCompound{}
	this.PkiUserstagedID = pkiUserstagedID
	this.FkiEmailID = fkiEmailID
	this.SEmailAddress = sEmailAddress
	this.SUserstagedFirstname = sUserstagedFirstname
	this.SUserstagedLastname = sUserstagedLastname
	this.SUserstagedExternalid = sUserstagedExternalid
	return &this
}

// NewUserstagedResponseCompoundWithDefaults instantiates a new UserstagedResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserstagedResponseCompoundWithDefaults() *UserstagedResponseCompound {
	this := UserstagedResponseCompound{}
	return &this
}

// GetPkiUserstagedID returns the PkiUserstagedID field value
func (o *UserstagedResponseCompound) GetPkiUserstagedID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUserstagedID
}

// GetPkiUserstagedIDOk returns a tuple with the PkiUserstagedID field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetPkiUserstagedIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUserstagedID, true
}

// SetPkiUserstagedID sets field value
func (o *UserstagedResponseCompound) SetPkiUserstagedID(v int32) {
	o.PkiUserstagedID = v
}

// GetFkiEmailID returns the FkiEmailID field value
func (o *UserstagedResponseCompound) GetFkiEmailID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEmailID
}

// GetFkiEmailIDOk returns a tuple with the FkiEmailID field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetFkiEmailIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEmailID, true
}

// SetFkiEmailID sets field value
func (o *UserstagedResponseCompound) SetFkiEmailID(v int32) {
	o.FkiEmailID = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *UserstagedResponseCompound) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetSEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *UserstagedResponseCompound) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

// GetSUserstagedFirstname returns the SUserstagedFirstname field value
func (o *UserstagedResponseCompound) GetSUserstagedFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedFirstname
}

// GetSUserstagedFirstnameOk returns a tuple with the SUserstagedFirstname field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetSUserstagedFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedFirstname, true
}

// SetSUserstagedFirstname sets field value
func (o *UserstagedResponseCompound) SetSUserstagedFirstname(v string) {
	o.SUserstagedFirstname = v
}

// GetSUserstagedLastname returns the SUserstagedLastname field value
func (o *UserstagedResponseCompound) GetSUserstagedLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedLastname
}

// GetSUserstagedLastnameOk returns a tuple with the SUserstagedLastname field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetSUserstagedLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedLastname, true
}

// SetSUserstagedLastname sets field value
func (o *UserstagedResponseCompound) SetSUserstagedLastname(v string) {
	o.SUserstagedLastname = v
}

// GetSUserstagedExternalid returns the SUserstagedExternalid field value
func (o *UserstagedResponseCompound) GetSUserstagedExternalid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserstagedExternalid
}

// GetSUserstagedExternalidOk returns a tuple with the SUserstagedExternalid field value
// and a boolean to check if the value has been set.
func (o *UserstagedResponseCompound) GetSUserstagedExternalidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserstagedExternalid, true
}

// SetSUserstagedExternalid sets field value
func (o *UserstagedResponseCompound) SetSUserstagedExternalid(v string) {
	o.SUserstagedExternalid = v
}

func (o UserstagedResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserstagedResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUserstagedID"] = o.PkiUserstagedID
	toSerialize["fkiEmailID"] = o.FkiEmailID
	toSerialize["sEmailAddress"] = o.SEmailAddress
	toSerialize["sUserstagedFirstname"] = o.SUserstagedFirstname
	toSerialize["sUserstagedLastname"] = o.SUserstagedLastname
	toSerialize["sUserstagedExternalid"] = o.SUserstagedExternalid
	return toSerialize, nil
}

type NullableUserstagedResponseCompound struct {
	value *UserstagedResponseCompound
	isSet bool
}

func (v NullableUserstagedResponseCompound) Get() *UserstagedResponseCompound {
	return v.value
}

func (v *NullableUserstagedResponseCompound) Set(val *UserstagedResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUserstagedResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUserstagedResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserstagedResponseCompound(val *UserstagedResponseCompound) *NullableUserstagedResponseCompound {
	return &NullableUserstagedResponseCompound{value: val, isSet: true}
}

func (v NullableUserstagedResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserstagedResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


