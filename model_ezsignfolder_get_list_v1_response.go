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

// EzsignfolderGetListV1Response Response for the /1/object/ezsignfolder/getList API Request
type EzsignfolderGetListV1Response struct {
	MPayload EzsignfolderGetListV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayloadGetList `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfolderGetListV1Response instantiates a new EzsignfolderGetListV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetListV1Response(mPayload EzsignfolderGetListV1ResponseMPayload) *EzsignfolderGetListV1Response {
	this := EzsignfolderGetListV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetListV1ResponseWithDefaults instantiates a new EzsignfolderGetListV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetListV1ResponseWithDefaults() *EzsignfolderGetListV1Response {
	this := EzsignfolderGetListV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetListV1Response) GetMPayload() EzsignfolderGetListV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetListV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetListV1Response) GetMPayloadOk() (*EzsignfolderGetListV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetListV1Response) SetMPayload(v EzsignfolderGetListV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsignfolderGetListV1Response) GetObjDebugPayload() CommonResponseObjDebugPayloadGetList {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayloadGetList
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetListV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayloadGetList, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsignfolderGetListV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayloadGetList and assigns it to the ObjDebugPayload field.
func (o *EzsignfolderGetListV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayloadGetList) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderGetListV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetListV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderGetListV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderGetListV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfolderGetListV1Response) MarshalJSON() ([]byte, error) {
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

type NullableEzsignfolderGetListV1Response struct {
	value *EzsignfolderGetListV1Response
	isSet bool
}

func (v NullableEzsignfolderGetListV1Response) Get() *EzsignfolderGetListV1Response {
	return v.value
}

func (v *NullableEzsignfolderGetListV1Response) Set(val *EzsignfolderGetListV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetListV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetListV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetListV1Response(val *EzsignfolderGetListV1Response) *NullableEzsignfolderGetListV1Response {
	return &NullableEzsignfolderGetListV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderGetListV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetListV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


