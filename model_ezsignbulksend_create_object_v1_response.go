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

// checks if the EzsignbulksendCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendCreateObjectV1Response{}

// EzsignbulksendCreateObjectV1Response Response for POST /1/object/ezsignbulksend
type EzsignbulksendCreateObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignbulksendCreateObjectV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignbulksendCreateObjectV1Response instantiates a new EzsignbulksendCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendCreateObjectV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignbulksendCreateObjectV1ResponseMPayload) *EzsignbulksendCreateObjectV1Response {
	this := EzsignbulksendCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignbulksendCreateObjectV1ResponseWithDefaults instantiates a new EzsignbulksendCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendCreateObjectV1ResponseWithDefaults() *EzsignbulksendCreateObjectV1Response {
	this := EzsignbulksendCreateObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignbulksendCreateObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignbulksendCreateObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignbulksendCreateObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignbulksendCreateObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignbulksendCreateObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignbulksendCreateObjectV1Response) GetMPayload() EzsignbulksendCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret EzsignbulksendCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendCreateObjectV1Response) GetMPayloadOk() (*EzsignbulksendCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignbulksendCreateObjectV1Response) SetMPayload(v EzsignbulksendCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignbulksendCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignbulksendCreateObjectV1Response struct {
	value *EzsignbulksendCreateObjectV1Response
	isSet bool
}

func (v NullableEzsignbulksendCreateObjectV1Response) Get() *EzsignbulksendCreateObjectV1Response {
	return v.value
}

func (v *NullableEzsignbulksendCreateObjectV1Response) Set(val *EzsignbulksendCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendCreateObjectV1Response(val *EzsignbulksendCreateObjectV1Response) *NullableEzsignbulksendCreateObjectV1Response {
	return &NullableEzsignbulksendCreateObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

