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

// EzsigndocumentGetFormDataV1Response Response for the /1/object/ezsigndocument/{pkiEzsigndocument}/getFormData API Request
type EzsigndocumentGetFormDataV1Response struct {
	MPayload EzsigndocumentGetFormDataV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsigndocumentGetFormDataV1Response instantiates a new EzsigndocumentGetFormDataV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetFormDataV1Response(mPayload EzsigndocumentGetFormDataV1ResponseMPayload) *EzsigndocumentGetFormDataV1Response {
	this := EzsigndocumentGetFormDataV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetFormDataV1ResponseWithDefaults instantiates a new EzsigndocumentGetFormDataV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetFormDataV1ResponseWithDefaults() *EzsigndocumentGetFormDataV1Response {
	this := EzsigndocumentGetFormDataV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetFormDataV1Response) GetMPayload() EzsigndocumentGetFormDataV1ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetFormDataV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1Response) GetMPayloadOk() (*EzsigndocumentGetFormDataV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetFormDataV1Response) SetMPayload(v EzsigndocumentGetFormDataV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *EzsigndocumentGetFormDataV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *EzsigndocumentGetFormDataV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *EzsigndocumentGetFormDataV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentGetFormDataV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentGetFormDataV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentGetFormDataV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsigndocumentGetFormDataV1Response) MarshalJSON() ([]byte, error) {
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

type NullableEzsigndocumentGetFormDataV1Response struct {
	value *EzsigndocumentGetFormDataV1Response
	isSet bool
}

func (v NullableEzsigndocumentGetFormDataV1Response) Get() *EzsigndocumentGetFormDataV1Response {
	return v.value
}

func (v *NullableEzsigndocumentGetFormDataV1Response) Set(val *EzsigndocumentGetFormDataV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetFormDataV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetFormDataV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetFormDataV1Response(val *EzsigndocumentGetFormDataV1Response) *NullableEzsigndocumentGetFormDataV1Response {
	return &NullableEzsigndocumentGetFormDataV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetFormDataV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetFormDataV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


