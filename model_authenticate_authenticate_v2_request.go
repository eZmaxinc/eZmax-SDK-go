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

// AuthenticateAuthenticateV2Request Request for the /2/module/authenticate/authenticate API Request
type AuthenticateAuthenticateV2Request struct {
	// The customer code assigned to your account
	PksCustomerCode string `json:"pksCustomerCode"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Login name of the User.
	SUserLoginname *string `json:"sUserLoginname,omitempty"`
	// A Password.  Must meet complexity requirements
	SPassword *string `json:"sPassword,omitempty"`
	// A Password encrypted and encoded in Base64  Must meet complexity requirements
	SPasswordEncrypted *string `json:"sPasswordEncrypted,omitempty"`
}

// NewAuthenticateAuthenticateV2Request instantiates a new AuthenticateAuthenticateV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticateAuthenticateV2Request(pksCustomerCode string) *AuthenticateAuthenticateV2Request {
	this := AuthenticateAuthenticateV2Request{}
	this.PksCustomerCode = pksCustomerCode
	return &this
}

// NewAuthenticateAuthenticateV2RequestWithDefaults instantiates a new AuthenticateAuthenticateV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticateAuthenticateV2RequestWithDefaults() *AuthenticateAuthenticateV2Request {
	this := AuthenticateAuthenticateV2Request{}
	return &this
}

// GetPksCustomerCode returns the PksCustomerCode field value
func (o *AuthenticateAuthenticateV2Request) GetPksCustomerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PksCustomerCode
}

// GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field value
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Request) GetPksCustomerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PksCustomerCode, true
}

// SetPksCustomerCode sets field value
func (o *AuthenticateAuthenticateV2Request) SetPksCustomerCode(v string) {
	o.PksCustomerCode = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Request) GetSEmailAddress() string {
	if o == nil || o.SEmailAddress == nil {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Request) GetSEmailAddressOk() (*string, bool) {
	if o == nil || o.SEmailAddress == nil {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Request) HasSEmailAddress() bool {
	if o != nil && o.SEmailAddress != nil {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *AuthenticateAuthenticateV2Request) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUserLoginname returns the SUserLoginname field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Request) GetSUserLoginname() string {
	if o == nil || o.SUserLoginname == nil {
		var ret string
		return ret
	}
	return *o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Request) GetSUserLoginnameOk() (*string, bool) {
	if o == nil || o.SUserLoginname == nil {
		return nil, false
	}
	return o.SUserLoginname, true
}

// HasSUserLoginname returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Request) HasSUserLoginname() bool {
	if o != nil && o.SUserLoginname != nil {
		return true
	}

	return false
}

// SetSUserLoginname gets a reference to the given string and assigns it to the SUserLoginname field.
func (o *AuthenticateAuthenticateV2Request) SetSUserLoginname(v string) {
	o.SUserLoginname = &v
}

// GetSPassword returns the SPassword field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Request) GetSPassword() string {
	if o == nil || o.SPassword == nil {
		var ret string
		return ret
	}
	return *o.SPassword
}

// GetSPasswordOk returns a tuple with the SPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Request) GetSPasswordOk() (*string, bool) {
	if o == nil || o.SPassword == nil {
		return nil, false
	}
	return o.SPassword, true
}

// HasSPassword returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Request) HasSPassword() bool {
	if o != nil && o.SPassword != nil {
		return true
	}

	return false
}

// SetSPassword gets a reference to the given string and assigns it to the SPassword field.
func (o *AuthenticateAuthenticateV2Request) SetSPassword(v string) {
	o.SPassword = &v
}

// GetSPasswordEncrypted returns the SPasswordEncrypted field value if set, zero value otherwise.
func (o *AuthenticateAuthenticateV2Request) GetSPasswordEncrypted() string {
	if o == nil || o.SPasswordEncrypted == nil {
		var ret string
		return ret
	}
	return *o.SPasswordEncrypted
}

// GetSPasswordEncryptedOk returns a tuple with the SPasswordEncrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticateAuthenticateV2Request) GetSPasswordEncryptedOk() (*string, bool) {
	if o == nil || o.SPasswordEncrypted == nil {
		return nil, false
	}
	return o.SPasswordEncrypted, true
}

// HasSPasswordEncrypted returns a boolean if a field has been set.
func (o *AuthenticateAuthenticateV2Request) HasSPasswordEncrypted() bool {
	if o != nil && o.SPasswordEncrypted != nil {
		return true
	}

	return false
}

// SetSPasswordEncrypted gets a reference to the given string and assigns it to the SPasswordEncrypted field.
func (o *AuthenticateAuthenticateV2Request) SetSPasswordEncrypted(v string) {
	o.SPasswordEncrypted = &v
}

func (o AuthenticateAuthenticateV2Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pksCustomerCode"] = o.PksCustomerCode
	}
	if o.SEmailAddress != nil {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	if o.SUserLoginname != nil {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if o.SPassword != nil {
		toSerialize["sPassword"] = o.SPassword
	}
	if o.SPasswordEncrypted != nil {
		toSerialize["sPasswordEncrypted"] = o.SPasswordEncrypted
	}
	return json.Marshal(toSerialize)
}

type NullableAuthenticateAuthenticateV2Request struct {
	value *AuthenticateAuthenticateV2Request
	isSet bool
}

func (v NullableAuthenticateAuthenticateV2Request) Get() *AuthenticateAuthenticateV2Request {
	return v.value
}

func (v *NullableAuthenticateAuthenticateV2Request) Set(val *AuthenticateAuthenticateV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticateAuthenticateV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticateAuthenticateV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticateAuthenticateV2Request(val *AuthenticateAuthenticateV2Request) *NullableAuthenticateAuthenticateV2Request {
	return &NullableAuthenticateAuthenticateV2Request{value: val, isSet: true}
}

func (v NullableAuthenticateAuthenticateV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticateAuthenticateV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


