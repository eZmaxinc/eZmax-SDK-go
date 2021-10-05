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

// EzsigndocumentApplyEzsigntemplateV2Response Response for the /2/object/ezsigndocument/{pkiEzsigndocument}/applyEzsigntemplate API Request
type EzsigndocumentApplyEzsigntemplateV2Response struct {
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsigndocumentApplyEzsigntemplateV2Response instantiates a new EzsigndocumentApplyEzsigntemplateV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentApplyEzsigntemplateV2Response() *EzsigndocumentApplyEzsigntemplateV2Response {
	this := EzsigndocumentApplyEzsigntemplateV2Response{}
	return &this
}

// NewEzsigndocumentApplyEzsigntemplateV2ResponseWithDefaults instantiates a new EzsigndocumentApplyEzsigntemplateV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentApplyEzsigntemplateV2ResponseWithDefaults() *EzsigndocumentApplyEzsigntemplateV2Response {
	this := EzsigndocumentApplyEzsigntemplateV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentApplyEzsigntemplateV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsigndocumentApplyEzsigntemplateV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentApplyEzsigntemplateV2Response struct {
	value *EzsigndocumentApplyEzsigntemplateV2Response
	isSet bool
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Response) Get() *EzsigndocumentApplyEzsigntemplateV2Response {
	return v.value
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Response) Set(val *EzsigndocumentApplyEzsigntemplateV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentApplyEzsigntemplateV2Response(val *EzsigndocumentApplyEzsigntemplateV2Response) *NullableEzsigndocumentApplyEzsigntemplateV2Response {
	return &NullableEzsigndocumentApplyEzsigntemplateV2Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentApplyEzsigntemplateV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentApplyEzsigntemplateV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

