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

// ContactinformationsRequestCompound A Contactinformations Object and children to create a complete structure
type ContactinformationsRequestCompound struct {
	AObjAddress []AddressRequest `json:"a_objAddress"`
	AObjPhone []PhoneRequest `json:"a_objPhone"`
	AObjEmail []EmailRequest `json:"a_objEmail"`
	AObjWebsite []WebsiteRequest `json:"a_objWebsite"`
	// The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty.
	IAddressDefault int32 `json:"iAddressDefault"`
	// The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty.
	IPhoneDefault int32 `json:"iPhoneDefault"`
	// The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty.
	IEmailDefault int32 `json:"iEmailDefault"`
	// The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty.
	IWebsiteDefault int32 `json:"iWebsiteDefault"`
}

// NewContactinformationsRequestCompound instantiates a new ContactinformationsRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactinformationsRequestCompound(aObjAddress []AddressRequest, aObjPhone []PhoneRequest, aObjEmail []EmailRequest, aObjWebsite []WebsiteRequest, iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32) *ContactinformationsRequestCompound {
	this := ContactinformationsRequestCompound{}
	this.AObjAddress = aObjAddress
	this.AObjPhone = aObjPhone
	this.AObjEmail = aObjEmail
	this.AObjWebsite = aObjWebsite
	this.IAddressDefault = iAddressDefault
	this.IPhoneDefault = iPhoneDefault
	this.IEmailDefault = iEmailDefault
	this.IWebsiteDefault = iWebsiteDefault
	return &this
}

// NewContactinformationsRequestCompoundWithDefaults instantiates a new ContactinformationsRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactinformationsRequestCompoundWithDefaults() *ContactinformationsRequestCompound {
	this := ContactinformationsRequestCompound{}
	return &this
}

// GetAObjAddress returns the AObjAddress field value
func (o *ContactinformationsRequestCompound) GetAObjAddress() []AddressRequest {
	if o == nil {
		var ret []AddressRequest
		return ret
	}

	return o.AObjAddress
}

// GetAObjAddressOk returns a tuple with the AObjAddress field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetAObjAddressOk() ([]AddressRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAddress, true
}

// SetAObjAddress sets field value
func (o *ContactinformationsRequestCompound) SetAObjAddress(v []AddressRequest) {
	o.AObjAddress = v
}

// GetAObjPhone returns the AObjPhone field value
func (o *ContactinformationsRequestCompound) GetAObjPhone() []PhoneRequest {
	if o == nil {
		var ret []PhoneRequest
		return ret
	}

	return o.AObjPhone
}

// GetAObjPhoneOk returns a tuple with the AObjPhone field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetAObjPhoneOk() ([]PhoneRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPhone, true
}

// SetAObjPhone sets field value
func (o *ContactinformationsRequestCompound) SetAObjPhone(v []PhoneRequest) {
	o.AObjPhone = v
}

// GetAObjEmail returns the AObjEmail field value
func (o *ContactinformationsRequestCompound) GetAObjEmail() []EmailRequest {
	if o == nil {
		var ret []EmailRequest
		return ret
	}

	return o.AObjEmail
}

// GetAObjEmailOk returns a tuple with the AObjEmail field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetAObjEmailOk() ([]EmailRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEmail, true
}

// SetAObjEmail sets field value
func (o *ContactinformationsRequestCompound) SetAObjEmail(v []EmailRequest) {
	o.AObjEmail = v
}

// GetAObjWebsite returns the AObjWebsite field value
func (o *ContactinformationsRequestCompound) GetAObjWebsite() []WebsiteRequest {
	if o == nil {
		var ret []WebsiteRequest
		return ret
	}

	return o.AObjWebsite
}

// GetAObjWebsiteOk returns a tuple with the AObjWebsite field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetAObjWebsiteOk() ([]WebsiteRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjWebsite, true
}

// SetAObjWebsite sets field value
func (o *ContactinformationsRequestCompound) SetAObjWebsite(v []WebsiteRequest) {
	o.AObjWebsite = v
}

// GetIAddressDefault returns the IAddressDefault field value
func (o *ContactinformationsRequestCompound) GetIAddressDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IAddressDefault
}

// GetIAddressDefaultOk returns a tuple with the IAddressDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetIAddressDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IAddressDefault, true
}

// SetIAddressDefault sets field value
func (o *ContactinformationsRequestCompound) SetIAddressDefault(v int32) {
	o.IAddressDefault = v
}

// GetIPhoneDefault returns the IPhoneDefault field value
func (o *ContactinformationsRequestCompound) GetIPhoneDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IPhoneDefault
}

// GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetIPhoneDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPhoneDefault, true
}

// SetIPhoneDefault sets field value
func (o *ContactinformationsRequestCompound) SetIPhoneDefault(v int32) {
	o.IPhoneDefault = v
}

// GetIEmailDefault returns the IEmailDefault field value
func (o *ContactinformationsRequestCompound) GetIEmailDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEmailDefault
}

// GetIEmailDefaultOk returns a tuple with the IEmailDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetIEmailDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IEmailDefault, true
}

// SetIEmailDefault sets field value
func (o *ContactinformationsRequestCompound) SetIEmailDefault(v int32) {
	o.IEmailDefault = v
}

// GetIWebsiteDefault returns the IWebsiteDefault field value
func (o *ContactinformationsRequestCompound) GetIWebsiteDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IWebsiteDefault
}

// GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequestCompound) GetIWebsiteDefaultOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IWebsiteDefault, true
}

// SetIWebsiteDefault sets field value
func (o *ContactinformationsRequestCompound) SetIWebsiteDefault(v int32) {
	o.IWebsiteDefault = v
}

func (o ContactinformationsRequestCompound) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["iAddressDefault"] = o.IAddressDefault
	}
	if true {
		toSerialize["iPhoneDefault"] = o.IPhoneDefault
	}
	if true {
		toSerialize["iEmailDefault"] = o.IEmailDefault
	}
	if true {
		toSerialize["iWebsiteDefault"] = o.IWebsiteDefault
	}
	return json.Marshal(toSerialize)
}

type NullableContactinformationsRequestCompound struct {
	value *ContactinformationsRequestCompound
	isSet bool
}

func (v NullableContactinformationsRequestCompound) Get() *ContactinformationsRequestCompound {
	return v.value
}

func (v *NullableContactinformationsRequestCompound) Set(val *ContactinformationsRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableContactinformationsRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableContactinformationsRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactinformationsRequestCompound(val *ContactinformationsRequestCompound) *NullableContactinformationsRequestCompound {
	return &NullableContactinformationsRequestCompound{value: val, isSet: true}
}

func (v NullableContactinformationsRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactinformationsRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


