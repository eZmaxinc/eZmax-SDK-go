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

// SsprSendUsernamesV1Request Request for the /1/module/sspr/sendUsernames API Request
type SsprSendUsernamesV1Request struct {
	// The customer code assigned to your account
	PksCustomerCode string `json:"pksCustomerCode"`
	// The unique ID of the Language.  Valid values:  |Value|Description| |-|-| |1|French| |2|English|
	FkiLanguageID int32 `json:"fkiLanguageID"`
	EUserTypeSSPR FieldEUserTypeSSPR `json:"eUserTypeSSPR"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
}

// NewSsprSendUsernamesV1Request instantiates a new SsprSendUsernamesV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsprSendUsernamesV1Request(pksCustomerCode string, fkiLanguageID int32, eUserTypeSSPR FieldEUserTypeSSPR, sEmailAddress string) *SsprSendUsernamesV1Request {
	this := SsprSendUsernamesV1Request{}
	this.PksCustomerCode = pksCustomerCode
	this.FkiLanguageID = fkiLanguageID
	this.EUserTypeSSPR = eUserTypeSSPR
	this.SEmailAddress = sEmailAddress
	return &this
}

// NewSsprSendUsernamesV1RequestWithDefaults instantiates a new SsprSendUsernamesV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsprSendUsernamesV1RequestWithDefaults() *SsprSendUsernamesV1Request {
	this := SsprSendUsernamesV1Request{}
	return &this
}

// GetPksCustomerCode returns the PksCustomerCode field value
func (o *SsprSendUsernamesV1Request) GetPksCustomerCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PksCustomerCode
}

// GetPksCustomerCodeOk returns a tuple with the PksCustomerCode field value
// and a boolean to check if the value has been set.
func (o *SsprSendUsernamesV1Request) GetPksCustomerCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PksCustomerCode, true
}

// SetPksCustomerCode sets field value
func (o *SsprSendUsernamesV1Request) SetPksCustomerCode(v string) {
	o.PksCustomerCode = v
}

// GetFkiLanguageID returns the FkiLanguageID field value
func (o *SsprSendUsernamesV1Request) GetFkiLanguageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiLanguageID
}

// GetFkiLanguageIDOk returns a tuple with the FkiLanguageID field value
// and a boolean to check if the value has been set.
func (o *SsprSendUsernamesV1Request) GetFkiLanguageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiLanguageID, true
}

// SetFkiLanguageID sets field value
func (o *SsprSendUsernamesV1Request) SetFkiLanguageID(v int32) {
	o.FkiLanguageID = v
}

// GetEUserTypeSSPR returns the EUserTypeSSPR field value
func (o *SsprSendUsernamesV1Request) GetEUserTypeSSPR() FieldEUserTypeSSPR {
	if o == nil {
		var ret FieldEUserTypeSSPR
		return ret
	}

	return o.EUserTypeSSPR
}

// GetEUserTypeSSPROk returns a tuple with the EUserTypeSSPR field value
// and a boolean to check if the value has been set.
func (o *SsprSendUsernamesV1Request) GetEUserTypeSSPROk() (*FieldEUserTypeSSPR, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EUserTypeSSPR, true
}

// SetEUserTypeSSPR sets field value
func (o *SsprSendUsernamesV1Request) SetEUserTypeSSPR(v FieldEUserTypeSSPR) {
	o.EUserTypeSSPR = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *SsprSendUsernamesV1Request) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *SsprSendUsernamesV1Request) GetSEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *SsprSendUsernamesV1Request) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

func (o SsprSendUsernamesV1Request) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableSsprSendUsernamesV1Request struct {
	value *SsprSendUsernamesV1Request
	isSet bool
}

func (v NullableSsprSendUsernamesV1Request) Get() *SsprSendUsernamesV1Request {
	return v.value
}

func (v *NullableSsprSendUsernamesV1Request) Set(val *SsprSendUsernamesV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableSsprSendUsernamesV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableSsprSendUsernamesV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsprSendUsernamesV1Request(val *SsprSendUsernamesV1Request) *NullableSsprSendUsernamesV1Request {
	return &NullableSsprSendUsernamesV1Request{value: val, isSet: true}
}

func (v NullableSsprSendUsernamesV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsprSendUsernamesV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


