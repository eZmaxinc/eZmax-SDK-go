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

// EzsigndocumentEndPrematurelyV1Response Response for the /1/object/ezsigndocument/{pkiEzsigndocument}/endPrematurely API Request
type EzsigndocumentEndPrematurelyV1Response struct {
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsigndocumentEndPrematurelyV1Response instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEndPrematurelyV1Response() *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	return &this
}

// NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults() *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsigndocumentEndPrematurelyV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentEndPrematurelyV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentEndPrematurelyV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentEndPrematurelyV1Response struct {
	value *EzsigndocumentEndPrematurelyV1Response
	isSet bool
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) Get() *EzsigndocumentEndPrematurelyV1Response {
	return v.value
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Set(val *EzsigndocumentEndPrematurelyV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEndPrematurelyV1Response(val *EzsigndocumentEndPrematurelyV1Response) *NullableEzsigndocumentEndPrematurelyV1Response {
	return &NullableEzsigndocumentEndPrematurelyV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

