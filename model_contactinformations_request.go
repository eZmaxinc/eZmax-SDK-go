/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.31
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// ContactinformationsRequest A Contactinformations Object
type ContactinformationsRequest struct {
	// The index in the a_objAddress array (zero based index) representing the Address object that should become the default one.  You can leave the value to 0 if the array is empty.
	IAddressDefault int32 `json:"iAddressDefault"`
	// The index in the a_objPhone array (zero based index) representing the Phone object that should become the default one.  You can leave the value to 0 if the array is empty.
	IPhoneDefault int32 `json:"iPhoneDefault"`
	// The index in the a_objEmail array (zero based index) representing the Email object that should become the default one.  You can leave the value to 0 if the array is empty.
	IEmailDefault int32 `json:"iEmailDefault"`
	// The index in the a_objWebsite array (zero based index) representing the Website object that should become the default one.  You can leave the value to 0 if the array is empty.
	IWebsiteDefault int32 `json:"iWebsiteDefault"`
}

// NewContactinformationsRequest instantiates a new ContactinformationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactinformationsRequest(iAddressDefault int32, iPhoneDefault int32, iEmailDefault int32, iWebsiteDefault int32) *ContactinformationsRequest {
	this := ContactinformationsRequest{}
	this.IAddressDefault = iAddressDefault
	this.IPhoneDefault = iPhoneDefault
	this.IEmailDefault = iEmailDefault
	this.IWebsiteDefault = iWebsiteDefault
	return &this
}

// NewContactinformationsRequestWithDefaults instantiates a new ContactinformationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactinformationsRequestWithDefaults() *ContactinformationsRequest {
	this := ContactinformationsRequest{}
	return &this
}

// GetIAddressDefault returns the IAddressDefault field value
func (o *ContactinformationsRequest) GetIAddressDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IAddressDefault
}

// GetIAddressDefaultOk returns a tuple with the IAddressDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequest) GetIAddressDefaultOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IAddressDefault, true
}

// SetIAddressDefault sets field value
func (o *ContactinformationsRequest) SetIAddressDefault(v int32) {
	o.IAddressDefault = v
}

// GetIPhoneDefault returns the IPhoneDefault field value
func (o *ContactinformationsRequest) GetIPhoneDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IPhoneDefault
}

// GetIPhoneDefaultOk returns a tuple with the IPhoneDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequest) GetIPhoneDefaultOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IPhoneDefault, true
}

// SetIPhoneDefault sets field value
func (o *ContactinformationsRequest) SetIPhoneDefault(v int32) {
	o.IPhoneDefault = v
}

// GetIEmailDefault returns the IEmailDefault field value
func (o *ContactinformationsRequest) GetIEmailDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IEmailDefault
}

// GetIEmailDefaultOk returns a tuple with the IEmailDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequest) GetIEmailDefaultOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IEmailDefault, true
}

// SetIEmailDefault sets field value
func (o *ContactinformationsRequest) SetIEmailDefault(v int32) {
	o.IEmailDefault = v
}

// GetIWebsiteDefault returns the IWebsiteDefault field value
func (o *ContactinformationsRequest) GetIWebsiteDefault() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IWebsiteDefault
}

// GetIWebsiteDefaultOk returns a tuple with the IWebsiteDefault field value
// and a boolean to check if the value has been set.
func (o *ContactinformationsRequest) GetIWebsiteDefaultOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IWebsiteDefault, true
}

// SetIWebsiteDefault sets field value
func (o *ContactinformationsRequest) SetIWebsiteDefault(v int32) {
	o.IWebsiteDefault = v
}

func (o ContactinformationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableContactinformationsRequest struct {
	value *ContactinformationsRequest
	isSet bool
}

func (v NullableContactinformationsRequest) Get() *ContactinformationsRequest {
	return v.value
}

func (v *NullableContactinformationsRequest) Set(val *ContactinformationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactinformationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactinformationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactinformationsRequest(val *ContactinformationsRequest) *NullableContactinformationsRequest {
	return &NullableContactinformationsRequest{value: val, isSet: true}
}

func (v NullableContactinformationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactinformationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


