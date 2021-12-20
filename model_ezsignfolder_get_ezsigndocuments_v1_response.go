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

// EzsignfolderGetEzsigndocumentsV1Response Response for the /1/object/ezsignfolder/{pkiEzsignfolder}/getEzsigndocuments API Request
type EzsignfolderGetEzsigndocumentsV1Response struct {
	MPayload EzsignfolderGetEzsigndocumentsV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfolderGetEzsigndocumentsV1Response instantiates a new EzsignfolderGetEzsigndocumentsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetEzsigndocumentsV1Response(mPayload EzsignfolderGetEzsigndocumentsV1ResponseMPayload) *EzsignfolderGetEzsigndocumentsV1Response {
	this := EzsignfolderGetEzsigndocumentsV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderGetEzsigndocumentsV1ResponseWithDefaults instantiates a new EzsignfolderGetEzsigndocumentsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetEzsigndocumentsV1ResponseWithDefaults() *EzsignfolderGetEzsigndocumentsV1Response {
	this := EzsignfolderGetEzsigndocumentsV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetMPayload() EzsignfolderGetEzsigndocumentsV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderGetEzsigndocumentsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetMPayloadOk() (*EzsignfolderGetEzsigndocumentsV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderGetEzsigndocumentsV1Response) SetMPayload(v EzsignfolderGetEzsigndocumentsV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsignfolderGetEzsigndocumentsV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsignfolderGetEzsigndocumentsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsigndocumentsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderGetEzsigndocumentsV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderGetEzsigndocumentsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfolderGetEzsigndocumentsV1Response) MarshalJSON() ([]byte, error) {
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

type NullableEzsignfolderGetEzsigndocumentsV1Response struct {
	value *EzsignfolderGetEzsigndocumentsV1Response
	isSet bool
}

func (v NullableEzsignfolderGetEzsigndocumentsV1Response) Get() *EzsignfolderGetEzsigndocumentsV1Response {
	return v.value
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1Response) Set(val *EzsignfolderGetEzsigndocumentsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetEzsigndocumentsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetEzsigndocumentsV1Response(val *EzsignfolderGetEzsigndocumentsV1Response) *NullableEzsignfolderGetEzsigndocumentsV1Response {
	return &NullableEzsignfolderGetEzsigndocumentsV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderGetEzsigndocumentsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetEzsigndocumentsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


