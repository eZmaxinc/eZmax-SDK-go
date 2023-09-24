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

// checks if the UsergroupmembershipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipResponse{}

// UsergroupmembershipResponse A Usergroupmembership Object
type UsergroupmembershipResponse struct {
	// The unique ID of the Usergroupmembership
	PkiUsergroupmembershipID int32 `json:"pkiUsergroupmembershipID"`
	// The unique ID of the Usergroup
	FkiUsergroupID int32 `json:"fkiUsergroupID"`
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	// The first name of the user
	SUserFirstname string `json:"sUserFirstname"`
	// The last name of the user
	SUserLastname string `json:"sUserLastname"`
	// The login name of the User.
	SUserLoginname string `json:"sUserLoginname"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Name of the Usergroup in the language of the requester
	SUsergroupNameX string `json:"sUsergroupNameX"`
}

// NewUsergroupmembershipResponse instantiates a new UsergroupmembershipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipResponse(pkiUsergroupmembershipID int32, fkiUsergroupID int32, fkiUserID int32, sUserFirstname string, sUserLastname string, sUserLoginname string, sUsergroupNameX string) *UsergroupmembershipResponse {
	this := UsergroupmembershipResponse{}
	this.PkiUsergroupmembershipID = pkiUsergroupmembershipID
	this.FkiUsergroupID = fkiUsergroupID
	this.FkiUserID = fkiUserID
	this.SUserFirstname = sUserFirstname
	this.SUserLastname = sUserLastname
	this.SUserLoginname = sUserLoginname
	this.SUsergroupNameX = sUsergroupNameX
	return &this
}

// NewUsergroupmembershipResponseWithDefaults instantiates a new UsergroupmembershipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipResponseWithDefaults() *UsergroupmembershipResponse {
	this := UsergroupmembershipResponse{}
	return &this
}

// GetPkiUsergroupmembershipID returns the PkiUsergroupmembershipID field value
func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupmembershipID
}

// GetPkiUsergroupmembershipIDOk returns a tuple with the PkiUsergroupmembershipID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetPkiUsergroupmembershipIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupmembershipID, true
}

// SetPkiUsergroupmembershipID sets field value
func (o *UsergroupmembershipResponse) SetPkiUsergroupmembershipID(v int32) {
	o.PkiUsergroupmembershipID = v
}

// GetFkiUsergroupID returns the FkiUsergroupID field value
func (o *UsergroupmembershipResponse) GetFkiUsergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUsergroupID
}

// GetFkiUsergroupIDOk returns a tuple with the FkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetFkiUsergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUsergroupID, true
}

// SetFkiUsergroupID sets field value
func (o *UsergroupmembershipResponse) SetFkiUsergroupID(v int32) {
	o.FkiUsergroupID = v
}

// GetFkiUserID returns the FkiUserID field value
func (o *UsergroupmembershipResponse) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetFkiUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *UsergroupmembershipResponse) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetSUserFirstname returns the SUserFirstname field value
func (o *UsergroupmembershipResponse) GetSUserFirstname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserFirstname
}

// GetSUserFirstnameOk returns a tuple with the SUserFirstname field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserFirstnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserFirstname, true
}

// SetSUserFirstname sets field value
func (o *UsergroupmembershipResponse) SetSUserFirstname(v string) {
	o.SUserFirstname = v
}

// GetSUserLastname returns the SUserLastname field value
func (o *UsergroupmembershipResponse) GetSUserLastname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLastname
}

// GetSUserLastnameOk returns a tuple with the SUserLastname field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserLastnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLastname, true
}

// SetSUserLastname sets field value
func (o *UsergroupmembershipResponse) SetSUserLastname(v string) {
	o.SUserLastname = v
}

// GetSUserLoginname returns the SUserLoginname field value
func (o *UsergroupmembershipResponse) GetSUserLoginname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUserLoginnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUserLoginname, true
}

// SetSUserLoginname sets field value
func (o *UsergroupmembershipResponse) SetSUserLoginname(v string) {
	o.SUserLoginname = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *UsergroupmembershipResponse) GetSEmailAddress() string {
	if o == nil || IsNil(o.SEmailAddress) {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.SEmailAddress) {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *UsergroupmembershipResponse) HasSEmailAddress() bool {
	if o != nil && !IsNil(o.SEmailAddress) {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *UsergroupmembershipResponse) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUsergroupNameX returns the SUsergroupNameX field value
func (o *UsergroupmembershipResponse) GetSUsergroupNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupNameX
}

// GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field value
// and a boolean to check if the value has been set.
func (o *UsergroupmembershipResponse) GetSUsergroupNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupNameX, true
}

// SetSUsergroupNameX sets field value
func (o *UsergroupmembershipResponse) SetSUsergroupNameX(v string) {
	o.SUsergroupNameX = v
}

func (o UsergroupmembershipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupmembershipID"] = o.PkiUsergroupmembershipID
	toSerialize["fkiUsergroupID"] = o.FkiUsergroupID
	toSerialize["fkiUserID"] = o.FkiUserID
	toSerialize["sUserFirstname"] = o.SUserFirstname
	toSerialize["sUserLastname"] = o.SUserLastname
	toSerialize["sUserLoginname"] = o.SUserLoginname
	if !IsNil(o.SEmailAddress) {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	toSerialize["sUsergroupNameX"] = o.SUsergroupNameX
	return toSerialize, nil
}

type NullableUsergroupmembershipResponse struct {
	value *UsergroupmembershipResponse
	isSet bool
}

func (v NullableUsergroupmembershipResponse) Get() *UsergroupmembershipResponse {
	return v.value
}

func (v *NullableUsergroupmembershipResponse) Set(val *UsergroupmembershipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipResponse(val *UsergroupmembershipResponse) *NullableUsergroupmembershipResponse {
	return &NullableUsergroupmembershipResponse{value: val, isSet: true}
}

func (v NullableUsergroupmembershipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

