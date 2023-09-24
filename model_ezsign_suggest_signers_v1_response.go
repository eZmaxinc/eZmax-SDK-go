/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the EzsignSuggestSignersV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignSuggestSignersV1Response{}

// EzsignSuggestSignersV1Response Response for GET /1/module/ezsign/suggestSigners
type EzsignSuggestSignersV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignSuggestSignersV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignSuggestSignersV1Response instantiates a new EzsignSuggestSignersV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignSuggestSignersV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignSuggestSignersV1ResponseMPayload) *EzsignSuggestSignersV1Response {
	this := EzsignSuggestSignersV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignSuggestSignersV1ResponseWithDefaults instantiates a new EzsignSuggestSignersV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignSuggestSignersV1ResponseWithDefaults() *EzsignSuggestSignersV1Response {
	this := EzsignSuggestSignersV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignSuggestSignersV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignSuggestSignersV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignSuggestSignersV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignSuggestSignersV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignSuggestSignersV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignSuggestSignersV1Response) GetMPayload() EzsignSuggestSignersV1ResponseMPayload {
	if o == nil {
		var ret EzsignSuggestSignersV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1Response) GetMPayloadOk() (*EzsignSuggestSignersV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignSuggestSignersV1Response) SetMPayload(v EzsignSuggestSignersV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignSuggestSignersV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignSuggestSignersV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignSuggestSignersV1Response struct {
	value *EzsignSuggestSignersV1Response
	isSet bool
}

func (v NullableEzsignSuggestSignersV1Response) Get() *EzsignSuggestSignersV1Response {
	return v.value
}

func (v *NullableEzsignSuggestSignersV1Response) Set(val *EzsignSuggestSignersV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignSuggestSignersV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignSuggestSignersV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignSuggestSignersV1Response(val *EzsignSuggestSignersV1Response) *NullableEzsignSuggestSignersV1Response {
	return &NullableEzsignSuggestSignersV1Response{value: val, isSet: true}
}

func (v NullableEzsignSuggestSignersV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignSuggestSignersV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


