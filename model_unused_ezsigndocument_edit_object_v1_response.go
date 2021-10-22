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

// UNUSEDEzsigndocumentEditObjectV1Response Response for the /1/object/ezsigndocument/editObject API Request
type UNUSEDEzsigndocumentEditObjectV1Response struct {
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewUNUSEDEzsigndocumentEditObjectV1Response instantiates a new UNUSEDEzsigndocumentEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUNUSEDEzsigndocumentEditObjectV1Response() *UNUSEDEzsigndocumentEditObjectV1Response {
	this := UNUSEDEzsigndocumentEditObjectV1Response{}
	return &this
}

// NewUNUSEDEzsigndocumentEditObjectV1ResponseWithDefaults instantiates a new UNUSEDEzsigndocumentEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUNUSEDEzsigndocumentEditObjectV1ResponseWithDefaults() *UNUSEDEzsigndocumentEditObjectV1Response {
	this := UNUSEDEzsigndocumentEditObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *UNUSEDEzsigndocumentEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o UNUSEDEzsigndocumentEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableUNUSEDEzsigndocumentEditObjectV1Response struct {
	value *UNUSEDEzsigndocumentEditObjectV1Response
	isSet bool
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Response) Get() *UNUSEDEzsigndocumentEditObjectV1Response {
	return v.value
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Response) Set(val *UNUSEDEzsigndocumentEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUNUSEDEzsigndocumentEditObjectV1Response(val *UNUSEDEzsigndocumentEditObjectV1Response) *NullableUNUSEDEzsigndocumentEditObjectV1Response {
	return &NullableUNUSEDEzsigndocumentEditObjectV1Response{value: val, isSet: true}
}

func (v NullableUNUSEDEzsigndocumentEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUNUSEDEzsigndocumentEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


