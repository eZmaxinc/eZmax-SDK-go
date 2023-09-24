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

// checks if the EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response{}

// EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response Response for PUT /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/editEzsigntemplateformfieldgroups
type EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload `json:"mPayload"`
}

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseWithDefaults instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseWithDefaults() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetMPayload() EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload {
	if o == nil {
		var ret EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) GetMPayloadOk() (*EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) SetMPayload(v EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response struct {
	value *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response
	isSet bool
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) Get() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response {
	return v.value
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) Set(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response {
	return &NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

