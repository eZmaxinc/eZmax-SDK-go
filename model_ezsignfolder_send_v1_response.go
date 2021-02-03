/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.27
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfolderSendV1Response Response for the /1/object/ezsignfolder/{pkiEzsignfolderID}/send API Request
type EzsignfolderSendV1Response struct {
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfolderSendV1Response instantiates a new EzsignfolderSendV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderSendV1Response() *EzsignfolderSendV1Response {
	this := EzsignfolderSendV1Response{}
	return &this
}

// NewEzsignfolderSendV1ResponseWithDefaults instantiates a new EzsignfolderSendV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderSendV1ResponseWithDefaults() *EzsignfolderSendV1Response {
	this := EzsignfolderSendV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsignfolderSendV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsignfolderSendV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsignfolderSendV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderSendV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderSendV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderSendV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfolderSendV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderSendV1Response struct {
	value *EzsignfolderSendV1Response
	isSet bool
}

func (v NullableEzsignfolderSendV1Response) Get() *EzsignfolderSendV1Response {
	return v.value
}

func (v *NullableEzsignfolderSendV1Response) Set(val *EzsignfolderSendV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderSendV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderSendV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderSendV1Response(val *EzsignfolderSendV1Response) *NullableEzsignfolderSendV1Response {
	return &NullableEzsignfolderSendV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderSendV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderSendV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


