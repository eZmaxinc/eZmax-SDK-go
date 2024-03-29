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

// AttemptResponse An Attempt object
type AttemptResponse struct {
	// Represent a Date Time. The timezone is the one configured in the User's profile.
	DtAttemptStart string `json:"dtAttemptStart"`
	// The Success or Failure message of the attempt when we tried to call the URL to deliver the webhook event.
	SAttemptResult string `json:"sAttemptResult"`
	// The number of second it took to process the webhook or get an error
	IAttemptDuration int32 `json:"iAttemptDuration"`
}

// NewAttemptResponse instantiates a new AttemptResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttemptResponse(dtAttemptStart string, sAttemptResult string, iAttemptDuration int32) *AttemptResponse {
	this := AttemptResponse{}
	this.DtAttemptStart = dtAttemptStart
	this.SAttemptResult = sAttemptResult
	this.IAttemptDuration = iAttemptDuration
	return &this
}

// NewAttemptResponseWithDefaults instantiates a new AttemptResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttemptResponseWithDefaults() *AttemptResponse {
	this := AttemptResponse{}
	return &this
}

// GetDtAttemptStart returns the DtAttemptStart field value
func (o *AttemptResponse) GetDtAttemptStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DtAttemptStart
}

// GetDtAttemptStartOk returns a tuple with the DtAttemptStart field value
// and a boolean to check if the value has been set.
func (o *AttemptResponse) GetDtAttemptStartOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DtAttemptStart, true
}

// SetDtAttemptStart sets field value
func (o *AttemptResponse) SetDtAttemptStart(v string) {
	o.DtAttemptStart = v
}

// GetSAttemptResult returns the SAttemptResult field value
func (o *AttemptResponse) GetSAttemptResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SAttemptResult
}

// GetSAttemptResultOk returns a tuple with the SAttemptResult field value
// and a boolean to check if the value has been set.
func (o *AttemptResponse) GetSAttemptResultOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SAttemptResult, true
}

// SetSAttemptResult sets field value
func (o *AttemptResponse) SetSAttemptResult(v string) {
	o.SAttemptResult = v
}

// GetIAttemptDuration returns the IAttemptDuration field value
func (o *AttemptResponse) GetIAttemptDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IAttemptDuration
}

// GetIAttemptDurationOk returns a tuple with the IAttemptDuration field value
// and a boolean to check if the value has been set.
func (o *AttemptResponse) GetIAttemptDurationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IAttemptDuration, true
}

// SetIAttemptDuration sets field value
func (o *AttemptResponse) SetIAttemptDuration(v int32) {
	o.IAttemptDuration = v
}

func (o AttemptResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dtAttemptStart"] = o.DtAttemptStart
	}
	if true {
		toSerialize["sAttemptResult"] = o.SAttemptResult
	}
	if true {
		toSerialize["iAttemptDuration"] = o.IAttemptDuration
	}
	return json.Marshal(toSerialize)
}

type NullableAttemptResponse struct {
	value *AttemptResponse
	isSet bool
}

func (v NullableAttemptResponse) Get() *AttemptResponse {
	return v.value
}

func (v *NullableAttemptResponse) Set(val *AttemptResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAttemptResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAttemptResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttemptResponse(val *AttemptResponse) *NullableAttemptResponse {
	return &NullableAttemptResponse{value: val, isSet: true}
}

func (v NullableAttemptResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttemptResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


