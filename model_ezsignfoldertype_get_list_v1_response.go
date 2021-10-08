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

// EzsignfoldertypeGetListV1Response Response for the /1/object/ezsignfoldertype/getList API Request
type EzsignfoldertypeGetListV1Response struct {
	MPayload EzsignfoldertypeGetListV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayloadGetList `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfoldertypeGetListV1Response instantiates a new EzsignfoldertypeGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeGetListV1Response(mPayload EzsignfoldertypeGetListV1ResponseMPayload) *EzsignfoldertypeGetListV1Response {
	this := EzsignfoldertypeGetListV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldertypeGetListV1ResponseWithDefaults instantiates a new EzsignfoldertypeGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeGetListV1ResponseWithDefaults() *EzsignfoldertypeGetListV1Response {
	this := EzsignfoldertypeGetListV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldertypeGetListV1Response) GetMPayload() EzsignfoldertypeGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsignfoldertypeGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetListV1Response) GetMPayloadOk() (*EzsignfoldertypeGetListV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldertypeGetListV1Response) SetMPayload(v EzsignfoldertypeGetListV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsignfoldertypeGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsignfoldertypeGetListV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayloadGetList and assigns it to the ObjDebugPayload field.
func (o *EzsignfoldertypeGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfoldertypeGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfoldertypeGetListV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfoldertypeGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfoldertypeGetListV1Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	if o.ObjDebugPayload != nil {
		toSerialize["objDebugPayload"] = o.ObjDebugPayload
	}
	if o.ObjDebug != nil {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldertypeGetListV1Response struct {
	value *EzsignfoldertypeGetListV1Response
	isSet bool
}

func (v NullableEzsignfoldertypeGetListV1Response) Get() *EzsignfoldertypeGetListV1Response {
	return v.value
}

func (v *NullableEzsignfoldertypeGetListV1Response) Set(val *EzsignfoldertypeGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeGetListV1Response(val *EzsignfoldertypeGetListV1Response) *NullableEzsignfoldertypeGetListV1Response {
	return &NullableEzsignfoldertypeGetListV1Response{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


