/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// CommonGetListV1ResponseMPayload Generic List Response
type CommonGetListV1ResponseMPayload struct {
	// The number of rows returned
	IRowReturned int32 `json:"iRowReturned"`
	// The number of rows matching your filters (if any) or the total number of rows
	IRowFiltered int32 `json:"iRowFiltered"`
}

// NewCommonGetListV1ResponseMPayload instantiates a new CommonGetListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonGetListV1ResponseMPayload(iRowReturned int32, iRowFiltered int32) *CommonGetListV1ResponseMPayload {
	this := CommonGetListV1ResponseMPayload{}
	this.IRowReturned = iRowReturned
	this.IRowFiltered = iRowFiltered
	return &this
}

// NewCommonGetListV1ResponseMPayloadWithDefaults instantiates a new CommonGetListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonGetListV1ResponseMPayloadWithDefaults() *CommonGetListV1ResponseMPayload {
	this := CommonGetListV1ResponseMPayload{}
	return &this
}

// GetIRowReturned returns the IRowReturned field value
func (o *CommonGetListV1ResponseMPayload) GetIRowReturned() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowReturned
}

// GetIRowReturnedOk returns a tuple with the IRowReturned field value
// and a boolean to check if the value has been set.
func (o *CommonGetListV1ResponseMPayload) GetIRowReturnedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IRowReturned, true
}

// SetIRowReturned sets field value
func (o *CommonGetListV1ResponseMPayload) SetIRowReturned(v int32) {
	o.IRowReturned = v
}

// GetIRowFiltered returns the IRowFiltered field value
func (o *CommonGetListV1ResponseMPayload) GetIRowFiltered() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IRowFiltered
}

// GetIRowFilteredOk returns a tuple with the IRowFiltered field value
// and a boolean to check if the value has been set.
func (o *CommonGetListV1ResponseMPayload) GetIRowFilteredOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IRowFiltered, true
}

// SetIRowFiltered sets field value
func (o *CommonGetListV1ResponseMPayload) SetIRowFiltered(v int32) {
	o.IRowFiltered = v
}

func (o CommonGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["iRowReturned"] = o.IRowReturned
	}
	if true {
		toSerialize["iRowFiltered"] = o.IRowFiltered
	}
	return json.Marshal(toSerialize)
}

type NullableCommonGetListV1ResponseMPayload struct {
	value *CommonGetListV1ResponseMPayload
	isSet bool
}

func (v NullableCommonGetListV1ResponseMPayload) Get() *CommonGetListV1ResponseMPayload {
	return v.value
}

func (v *NullableCommonGetListV1ResponseMPayload) Set(val *CommonGetListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonGetListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonGetListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonGetListV1ResponseMPayload(val *CommonGetListV1ResponseMPayload) *NullableCommonGetListV1ResponseMPayload {
	return &NullableCommonGetListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableCommonGetListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonGetListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


