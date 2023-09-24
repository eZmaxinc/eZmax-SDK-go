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

// checks if the EzsigndocumentGetObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetObjectV2Response{}

// EzsigndocumentGetObjectV2Response Response for GET /2/object/ezsigndocument/{pkiEzsigndocumentID}
type EzsigndocumentGetObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigndocumentGetObjectV2ResponseMPayload `json:"mPayload"`
}

// NewEzsigndocumentGetObjectV2Response instantiates a new EzsigndocumentGetObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigndocumentGetObjectV2ResponseMPayload) *EzsigndocumentGetObjectV2Response {
	this := EzsigndocumentGetObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetObjectV2ResponseWithDefaults instantiates a new EzsigndocumentGetObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetObjectV2ResponseWithDefaults() *EzsigndocumentGetObjectV2Response {
	this := EzsigndocumentGetObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigndocumentGetObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigndocumentGetObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigndocumentGetObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigndocumentGetObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigndocumentGetObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetObjectV2Response) GetMPayload() EzsigndocumentGetObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsigndocumentGetObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetObjectV2Response) GetMPayloadOk() (*EzsigndocumentGetObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetObjectV2Response) SetMPayload(v EzsigndocumentGetObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigndocumentGetObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsigndocumentGetObjectV2Response struct {
	value *EzsigndocumentGetObjectV2Response
	isSet bool
}

func (v NullableEzsigndocumentGetObjectV2Response) Get() *EzsigndocumentGetObjectV2Response {
	return v.value
}

func (v *NullableEzsigndocumentGetObjectV2Response) Set(val *EzsigndocumentGetObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetObjectV2Response(val *EzsigndocumentGetObjectV2Response) *NullableEzsigndocumentGetObjectV2Response {
	return &NullableEzsigndocumentGetObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

