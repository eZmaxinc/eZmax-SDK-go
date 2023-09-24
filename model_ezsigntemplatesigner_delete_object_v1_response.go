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

// checks if the EzsigntemplatesignerDeleteObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerDeleteObjectV1Response{}

// EzsigntemplatesignerDeleteObjectV1Response Response for DELETE /1/object/ezsigntemplatesigner/{pkiEzsigntemplatesignerID}
type EzsigntemplatesignerDeleteObjectV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsigntemplatesignerDeleteObjectV1Response instantiates a new EzsigntemplatesignerDeleteObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerDeleteObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsigntemplatesignerDeleteObjectV1Response {
	this := EzsigntemplatesignerDeleteObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsigntemplatesignerDeleteObjectV1ResponseWithDefaults instantiates a new EzsigntemplatesignerDeleteObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerDeleteObjectV1ResponseWithDefaults() *EzsigntemplatesignerDeleteObjectV1Response {
	this := EzsigntemplatesignerDeleteObjectV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplatesignerDeleteObjectV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerDeleteObjectV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplatesignerDeleteObjectV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplatesignerDeleteObjectV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerDeleteObjectV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplatesignerDeleteObjectV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplatesignerDeleteObjectV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsigntemplatesignerDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerDeleteObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

type NullableEzsigntemplatesignerDeleteObjectV1Response struct {
	value *EzsigntemplatesignerDeleteObjectV1Response
	isSet bool
}

func (v NullableEzsigntemplatesignerDeleteObjectV1Response) Get() *EzsigntemplatesignerDeleteObjectV1Response {
	return v.value
}

func (v *NullableEzsigntemplatesignerDeleteObjectV1Response) Set(val *EzsigntemplatesignerDeleteObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerDeleteObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerDeleteObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerDeleteObjectV1Response(val *EzsigntemplatesignerDeleteObjectV1Response) *NullableEzsigntemplatesignerDeleteObjectV1Response {
	return &NullableEzsigntemplatesignerDeleteObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerDeleteObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


