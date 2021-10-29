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

// FranchisereferalincomeCreateObjectV1Response Response for the /1/object/franchisereferalincome/createObject API Request
type FranchisereferalincomeCreateObjectV1Response struct {
	MPayload FranchisereferalincomeCreateObjectV1ResponseMPayload `json:"mPayload"`
	ObjDebugPayload *CommonResponseObjDebugPayload `json:"objDebugPayload,omitempty"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewFranchisereferalincomeCreateObjectV1Response instantiates a new FranchisereferalincomeCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV1Response(mPayload FranchisereferalincomeCreateObjectV1ResponseMPayload) *FranchisereferalincomeCreateObjectV1Response {
	this := FranchisereferalincomeCreateObjectV1Response{}
	this.MPayload = mPayload
	return &this
}

// NewFranchisereferalincomeCreateObjectV1ResponseWithDefaults instantiates a new FranchisereferalincomeCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV1ResponseWithDefaults() *FranchisereferalincomeCreateObjectV1Response {
	this := FranchisereferalincomeCreateObjectV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *FranchisereferalincomeCreateObjectV1Response) GetMPayload() FranchisereferalincomeCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret FranchisereferalincomeCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1Response) GetMPayloadOk() (*FranchisereferalincomeCreateObjectV1ResponseMPayload, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *FranchisereferalincomeCreateObjectV1Response) SetMPayload(v FranchisereferalincomeCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

// GetObjDebugPayload returns the ObjDebugPayload field value if set, zero value otherwise.
func (o *FranchisereferalincomeCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil || o.ObjDebugPayload == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}
	return *o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil || o.ObjDebugPayload == nil {
		return nil, false
	}
	return o.ObjDebugPayload, true
}

// HasObjDebugPayload returns a boolean if a field has been set.
func (o *FranchisereferalincomeCreateObjectV1Response) HasObjDebugPayload() bool {
	if o != nil && o.ObjDebugPayload != nil {
		return true
	}

	return false
}

// SetObjDebugPayload gets a reference to the given CommonResponseObjDebugPayload and assigns it to the ObjDebugPayload field.
func (o *FranchisereferalincomeCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = &v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *FranchisereferalincomeCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || o.ObjDebug == nil {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || o.ObjDebug == nil {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *FranchisereferalincomeCreateObjectV1Response) HasObjDebug() bool {
	if o != nil && o.ObjDebug != nil {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *FranchisereferalincomeCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o FranchisereferalincomeCreateObjectV1Response) MarshalJSON() ([]byte, error) {
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

type NullableFranchisereferalincomeCreateObjectV1Response struct {
	value *FranchisereferalincomeCreateObjectV1Response
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV1Response) Get() *FranchisereferalincomeCreateObjectV1Response {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV1Response) Set(val *FranchisereferalincomeCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV1Response(val *FranchisereferalincomeCreateObjectV1Response) *NullableFranchisereferalincomeCreateObjectV1Response {
	return &NullableFranchisereferalincomeCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


