/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.0.48
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ContactinformationsRequestCompoundAllOf struct for ContactinformationsRequestCompoundAllOf
type ContactinformationsRequestCompoundAllOf struct {
	AObjAddress []AddressRequest `json:"a_objAddress"`
	AObjPhone []PhoneRequest `json:"a_objPhone"`
	AObjEmail []EmailRequest `json:"a_objEmail"`
	AObjWebsite []WebsiteRequest `json:"a_objWebsite"`
}

// NewContactinformationsRequestCompoundAllOf instantiates a new ContactinformationsRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactinformationsRequestCompoundAllOf(aObjAddress []AddressRequest, aObjPhone []PhoneRequest, aObjEmail []EmailRequest, aObjWebsite []WebsiteRequest) *ContactinformationsRequestCompoundAllOf {
	this := ContactinformationsRequestCompoundAllOf{}
	this.AObjAddress = aObjAddress
	this.AObjPhone = aObjPhone
	this.AObjEmail = aObjEmail
	this.AObjWebsite = aObjWebsite
	return &this
}

// NewContactinformationsRequestCompoundAllOfWithDefaults instantiates a new ContactinformationsRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactinformationsRequestCompoundAllOfWithDefaults() *ContactinformationsRequestCompoundAllOf {
	this := ContactinformationsRequestCompoundAllOf{}
	return &this
}

// GetAObjAddress returns the AObjAddress field value
func (o *ContactinformationsRequestCompoundAllOf) GetAObjAddress() []AddressRequest {
	if o == nil {
		var ret []AddressRequest
		return ret
	}

	return o.AObjAddress
}

// GetAObjAddressOk returns a tuple with the AObjAddress field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompoundAllOf) GetAObjAddressOk() (*[]AddressRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjAddress, true
}

// SetAObjAddress sets field value
func (o *ContactinformationsRequestCompoundAllOf) SetAObjAddress(v []AddressRequest) {
	o.AObjAddress = v
}

// GetAObjPhone returns the AObjPhone field value
func (o *ContactinformationsRequestCompoundAllOf) GetAObjPhone() []PhoneRequest {
	if o == nil {
		var ret []PhoneRequest
		return ret
	}

	return o.AObjPhone
}

// GetAObjPhoneOk returns a tuple with the AObjPhone field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompoundAllOf) GetAObjPhoneOk() (*[]PhoneRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjPhone, true
}

// SetAObjPhone sets field value
func (o *ContactinformationsRequestCompoundAllOf) SetAObjPhone(v []PhoneRequest) {
	o.AObjPhone = v
}

// GetAObjEmail returns the AObjEmail field value
func (o *ContactinformationsRequestCompoundAllOf) GetAObjEmail() []EmailRequest {
	if o == nil {
		var ret []EmailRequest
		return ret
	}

	return o.AObjEmail
}

// GetAObjEmailOk returns a tuple with the AObjEmail field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompoundAllOf) GetAObjEmailOk() (*[]EmailRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjEmail, true
}

// SetAObjEmail sets field value
func (o *ContactinformationsRequestCompoundAllOf) SetAObjEmail(v []EmailRequest) {
	o.AObjEmail = v
}

// GetAObjWebsite returns the AObjWebsite field value
func (o *ContactinformationsRequestCompoundAllOf) GetAObjWebsite() []WebsiteRequest {
	if o == nil {
		var ret []WebsiteRequest
		return ret
	}

	return o.AObjWebsite
}

// GetAObjWebsiteOk returns a tuple with the AObjWebsite field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompoundAllOf) GetAObjWebsiteOk() (*[]WebsiteRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjWebsite, true
}

// SetAObjWebsite sets field value
func (o *ContactinformationsRequestCompoundAllOf) SetAObjWebsite(v []WebsiteRequest) {
	o.AObjWebsite = v
}

func (o ContactinformationsRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objAddress"] = o.AObjAddress
	}
	if true {
		toSerialize["a_objPhone"] = o.AObjPhone
	}
	if true {
		toSerialize["a_objEmail"] = o.AObjEmail
	}
	if true {
		toSerialize["a_objWebsite"] = o.AObjWebsite
	}
	return json.Marshal(toSerialize)
}

type NullableContactinformationsRequestCompoundAllOf struct {
	value *ContactinformationsRequestCompoundAllOf
	isSet bool
}

func (v NullableContactinformationsRequestCompoundAllOf) Get() *ContactinformationsRequestCompoundAllOf {
	return v.value
}

func (v *NullableContactinformationsRequestCompoundAllOf) Set(val *ContactinformationsRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContactinformationsRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContactinformationsRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactinformationsRequestCompoundAllOf(val *ContactinformationsRequestCompoundAllOf) *NullableContactinformationsRequestCompoundAllOf {
	return &NullableContactinformationsRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableContactinformationsRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactinformationsRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

