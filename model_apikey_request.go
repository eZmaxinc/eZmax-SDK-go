/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ApikeyRequest An Apikey Object
type ApikeyRequest struct {
	// The unique ID of the User
	FkiUserID int32 `json:"fkiUserID"`
	ObjApikeyDescription MultilingualApikeyDescription `json:"objApikeyDescription"`
}

// NewApikeyRequest instantiates a new ApikeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyRequest(fkiUserID int32, objApikeyDescription MultilingualApikeyDescription) *ApikeyRequest {
	this := ApikeyRequest{}
	this.FkiUserID = fkiUserID
	this.ObjApikeyDescription = objApikeyDescription
	return &this
}

// NewApikeyRequestWithDefaults instantiates a new ApikeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyRequestWithDefaults() *ApikeyRequest {
	this := ApikeyRequest{}
	return &this
}

// GetFkiUserID returns the FkiUserID field value
func (o *ApikeyRequest) GetFkiUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiUserID
}

// GetFkiUserIDOk returns a tuple with the FkiUserID field value
// and a boolean to check if the value has been set.
func (o *ApikeyRequest) GetFkiUserIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FkiUserID, true
}

// SetFkiUserID sets field value
func (o *ApikeyRequest) SetFkiUserID(v int32) {
	o.FkiUserID = v
}

// GetObjApikeyDescription returns the ObjApikeyDescription field value
func (o *ApikeyRequest) GetObjApikeyDescription() MultilingualApikeyDescription {
	if o == nil {
		var ret MultilingualApikeyDescription
		return ret
	}

	return o.ObjApikeyDescription
}

// GetObjApikeyDescriptionOk returns a tuple with the ObjApikeyDescription field value
// and a boolean to check if the value has been set.
func (o *ApikeyRequest) GetObjApikeyDescriptionOk() (*MultilingualApikeyDescription, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjApikeyDescription, true
}

// SetObjApikeyDescription sets field value
func (o *ApikeyRequest) SetObjApikeyDescription(v MultilingualApikeyDescription) {
	o.ObjApikeyDescription = v
}

func (o ApikeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fkiUserID"] = o.FkiUserID
	}
	if true {
		toSerialize["objApikeyDescription"] = o.ObjApikeyDescription
	}
	return json.Marshal(toSerialize)
}

type NullableApikeyRequest struct {
	value *ApikeyRequest
	isSet bool
}

func (v NullableApikeyRequest) Get() *ApikeyRequest {
	return v.value
}

func (v *NullableApikeyRequest) Set(val *ApikeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyRequest(val *ApikeyRequest) *NullableApikeyRequest {
	return &NullableApikeyRequest{value: val, isSet: true}
}

func (v NullableApikeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


