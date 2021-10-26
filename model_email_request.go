/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EmailRequest A Contact Object
type EmailRequest struct {
	// The unique ID of the Emailtype.  Valid values:  |Value|Description| |-|-| |1|Office| |2|Home|
	FkiEmailtypeID int32 `json:"fkiEmailtypeID"`
	// The email address.
	SEmailAddress string `json:"sEmailAddress"`
}

// NewEmailRequest instantiates a new EmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailRequest(fkiEmailtypeID int32, sEmailAddress string) *EmailRequest {
	this := EmailRequest{}
	this.FkiEmailtypeID = fkiEmailtypeID
	this.SEmailAddress = sEmailAddress
	return &this
}

// NewEmailRequestWithDefaults instantiates a new EmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailRequestWithDefaults() *EmailRequest {
	this := EmailRequest{}
	return &this
}

// GetFkiEmailtypeID returns the FkiEmailtypeID field value
func (o *EmailRequest) GetFkiEmailtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEmailtypeID
}

// GetFkiEmailtypeIDOk returns a tuple with the FkiEmailtypeID field value
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetFkiEmailtypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiEmailtypeID, true
}

// SetFkiEmailtypeID sets field value
func (o *EmailRequest) SetFkiEmailtypeID(v int32) {
	o.FkiEmailtypeID = v
}

// GetSEmailAddress returns the SEmailAddress field value
func (o *EmailRequest) GetSEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmailAddress
}

// GetSEmailAddressOk returns a tuple with the SEmailAddress field value
// and a boolean to check if the value has been set.
func (o *EmailRequest) GetSEmailAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEmailAddress, true
}

// SetSEmailAddress sets field value
func (o *EmailRequest) SetSEmailAddress(v string) {
	o.SEmailAddress = v
}

func (o EmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiEmailtypeID"] = o.FkiEmailtypeID
	}
	if true {
		toSerialize["sEmailAddress"] = o.SEmailAddress
	}
	return json.Marshal(toSerialize)
}

type NullableEmailRequest struct {
	value *EmailRequest
	isSet bool
}

func (v NullableEmailRequest) Get() *EmailRequest {
	return v.value
}

func (v *NullableEmailRequest) Set(val *EmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailRequest(val *EmailRequest) *NullableEmailRequest {
	return &NullableEmailRequest{value: val, isSet: true}
}

func (v NullableEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


