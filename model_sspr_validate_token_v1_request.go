/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// SsprValidateTokenV1Request Request for the /1/module/sspr/validateToken API Request
type SsprValidateTokenV1Request struct {
	// The customer code assigned to your account
	PksCustomerCode string `json:"pksCustomerCode"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EUserTypeSSPR FieldEUserTypeSSPR `json:"eUserTypeSSPR"`
	// The email address.
	SEmailAddress *string `json:"sEmailAddress,omitempty"`
	// The Login name of the User.
	SUserLoginname *string `json:"sUserLoginname,omitempty"`
	// Hex Encoded Secret SSPR token
	BinUserSSPRtoken string `json:"binUserSSPRtoken"`
}

// NewSsprValidateTokenV1Request instantiates a new SsprValidateTokenV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprValidateTokenV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, binUserSSPRtoken string) *SsprValidateTokenV1Request {
	this := SsprValidateTokenV1Request{}
	this.PksCustomerCode = pksCustomerCode
	this.FkiLanguageID = fkiLanguageID
	this.EUserTypeSSPR = eUserTypeSSPR
	this.BinUserSSPRtoken = binUserSSPRtoken
	return &this
}

// NewSsprValidateTokenV1RequestWithDefaults instantiates a new SsprValidateTokenV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprValidateTokenV1RequestWithDefaults() *SsprValidateTokenV1Request {
	this := SsprValidateTokenV1Request{}
	return &this
}

// GetPksCustomerCode returns the PksCustomerCode field value
func (o *SsprValidateTokenV1Request) GetPksCustomerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PksCustomerCode
}

// GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field value
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetPksCustomerCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PksCustomerCode, true
}

// SetPksCustomerCode sets field value
func (o *SsprValidateTokenV1Request) SetPksCustomerCode(v string) {
	o.PksCustomerCode = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *SsprValidateTokenV1Request) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *SsprValidateTokenV1Request) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEUserTypeSSPR returns the EUserTypeSSPR field value
func (o *SsprValidateTokenV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR {
	if o == nil {
		var ret FieldEUserTypeSSPR
		return ret
	}

	return o.EUserTypeSSPR
}

// GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field value
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EUserTypeSSPR, true
}

// SetEUserTypeSSPR sets field value
func (o *SsprValidateTokenV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR) {
	o.EUserTypeSSPR = v
}

// GetSEmailAddress returns the SEmailAddress field value if set, zero value otherwise.
func (o *SsprValidateTokenV1Request) GetSEmailAddress() string {
	if o == nil || o.SEmailAddress == nil {
		var ret string
		return ret
	}
	return *o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetSEmailAddressOk() (*string, bool) {
	if o == nil || o.SEmailAddress == nil {
		return nil, false
	}
	return o.SEmailAddress, true
}

// HasSEmailAddress returns a boolean if a field has been set.
func (o *SsprValidateTokenV1Request) HasSEmailAddress() bool {
	if o != nil && o.SEmailAddress != nil {
		return true
	}

	return false
}

// SetSEmailAddress gets a reference to the given string and assigns it to the SEmailAddress field.
func (o *SsprValidateTokenV1Request) SetSEmailAddress(v string) {
	o.SEmailAddress = &v
}

// GetSUserLoginname returns the SUserLoginname field value if set, zero value otherwise.
func (o *SsprValidateTokenV1Request) GetSUserLoginname() string {
	if o == nil || o.SUserLoginname == nil {
		var ret string
		return ret
	}
	return *o.SUserLoginname
}

// GetSUserLoginnameOk returns a tuple with the SUserLoginname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetSUserLoginnameOk() (*string, bool) {
	if o == nil || o.SUserLoginname == nil {
		return nil, false
	}
	return o.SUserLoginname, true
}

// HasSUserLoginname returns a boolean if a field has been set.
func (o *SsprValidateTokenV1Request) HasSUserLoginname() bool {
	if o != nil && o.SUserLoginname != nil {
		return true
	}

	return false
}

// SetSUserLoginname gets a reference to the given string and assigns it to the SUserLoginname field.
func (o *SsprValidateTokenV1Request) SetSUserLoginname(v string) {
	o.SUserLoginname = &v
}

// GetBinUserSSPRtoken returns the BinUserSSPRtoken field value
func (o *SsprValidateTokenV1Request) GetBinUserSSPRtoken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinUserSSPRtoken
}

// GetBinUserSSPRtokenOk returns a tuple with the BinUserSSPRtoken field value
// and a boolean to check if the value has been set.
func (o *SsprValidateTokenV1Request) GetBinUserSSPRtokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BinUserSSPRtoken, true
}

// SetBinUserSSPRtoken sets field value
func (o *SsprValidateTokenV1Request) SetBinUserSSPRtoken(v string) {
	o.BinUserSSPRtoken = v
}

func (o SsprValidateTokenV1Request) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pksCustomerCode"] = o.PksCustomerCode
	}
	if true {
		toSerialize["fkiLanguageID"] = o.FkiLanguageID
	}
	if true {
		toSerialize["eUserTypeSSPR"] = o.EUserTypeSSPR
	}
	if o.SEmailAddress != nil {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	if o.SUserLoginname != nil {
		toSerialize["sUserLoginname"] = o.SUserLoginname
	}
	if true {
		toSerialize["binUserSSPRtoken"] = o.BinUserSSPRtoken
	}
	return json.Marshal(toSerialize)
}

type NullableSsprValidateTokenV1Request struct {
	value *SsprValidateTokenV1Request
	isSet bool
}

func (v NullableSsprValidateTokenV1Request) Get() *SsprValidateTokenV1Request {
	return v.value
}

func (v *NullableSsprValidateTokenV1Request) Set(val *SsprValidateTokenV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprValidateTokenV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprValidateTokenV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprValidateTokenV1Request(val *SsprValidateTokenV1Request) *NullableSsprValidateTokenV1Request {
	return &NullableSsprValidateTokenV1Request{value: val, isSet: true}
}

func (v NullableSsprValidateTokenV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprValidateTokenV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


