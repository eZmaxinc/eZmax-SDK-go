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

// checks if the VariableexpenseEditObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseEditObjectV1Response{}

// VariableexpenseEditObjectV1Response Response for PUT /1/object/variableexpense/{pkiVariableexpenseID}
type VariableexpenseEditObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewVariableexpenseEditObjectV1Response instantiates a new VariableexpenseEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *VariableexpenseEditObjectV1Response {
	this := VariableexpenseEditObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewVariableexpenseEditObjectV1ResponseWithDefaults instantiates a new VariableexpenseEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseEditObjectV1ResponseWithDefaults() *VariableexpenseEditObjectV1Response {
	this := VariableexpenseEditObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *VariableexpenseEditObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseEditObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *VariableexpenseEditObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *VariableexpenseEditObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableexpenseEditObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *VariableexpenseEditObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *VariableexpenseEditObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o VariableexpenseEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseEditObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

type NullableVariableexpenseEditObjectV1Response struct {
	value *VariableexpenseEditObjectV1Response
	isSet bool
}

func (v NullableVariableexpenseEditObjectV1Response) Get() *VariableexpenseEditObjectV1Response {
	return v.value
}

func (v *NullableVariableexpenseEditObjectV1Response) Set(val *VariableexpenseEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseEditObjectV1Response(val *VariableexpenseEditObjectV1Response) *NullableVariableexpenseEditObjectV1Response {
	return &NullableVariableexpenseEditObjectV1Response{value: val, isSet: true}
}

func (v NullableVariableexpenseEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

