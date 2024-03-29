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

// CommonResponseGetList All API response will inherit this based Response
type CommonResponseGetList struct {
	ObjDebugPayload *CommonResponseObjDebugPayloadGetList `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewCommonResponseGetList instantiates a new CommonResponseGetList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseGetList() *CommonResponseGetList {
	this := CommonResponseGetList{}
	return &this
}

// NewCommonResponseGetListWithDefaults instantiates a new CommonResponseGetList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseGetListWithDefaults() *CommonResponseGetList {
	this := CommonResponseGetList{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *CommonResponseGetList) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponseGetList) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *CommonResponseGetList) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayloadGetList and assigns it to the ObjDebugPayload field.
func (o *CommonResponseGetList) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *CommonResponseGetList) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonResponseGetList) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *CommonResponseGetList) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *CommonResponseGetList) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o CommonResponseGetList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableCommonResponseGetList struct {
	value *CommonResponseGetList
	isSet bool
}

func (v NullableCommonResponseGetList) Get() *CommonResponseGetList {
	return v.value
}

func (v *NullableCommonResponseGetList) Set(val *CommonResponseGetList) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseGetList) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseGetList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseGetList(val *CommonResponseGetList) *NullableCommonResponseGetList {
	return &NullableCommonResponseGetList{value: val, isSet: true}
}

func (v NullableCommonResponseGetList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseGetList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


