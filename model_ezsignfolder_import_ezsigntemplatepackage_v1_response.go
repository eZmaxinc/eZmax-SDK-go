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

// checks if the EzsignfolderImportEzsigntemplatepackageV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderImportEzsigntemplatepackageV1Response{}

// EzsignfolderImportEzsigntemplatepackageV1Response Response for POST/1/object/ezsignfolder/{pkiEzsignfolderID}/importEzsigntemplatepackage
type EzsignfolderImportEzsigntemplatepackageV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload `json:"mPayload"`
}

// NewEzsignfolderImportEzsigntemplatepackageV1Response instantiates a new EzsignfolderImportEzsigntemplatepackageV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderImportEzsigntemplatepackageV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) *EzsignfolderImportEzsigntemplatepackageV1Response {
	this := EzsignfolderImportEzsigntemplatepackageV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignfolderImportEzsigntemplatepackageV1ResponseWithDefaults instantiates a new EzsignfolderImportEzsigntemplatepackageV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderImportEzsigntemplatepackageV1ResponseWithDefaults() *EzsignfolderImportEzsigntemplatepackageV1Response {
	this := EzsignfolderImportEzsigntemplatepackageV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetMPayload() EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload {
	if o == nil {
		var ret EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) GetMPayloadOk() (*EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfolderImportEzsigntemplatepackageV1Response) SetMPayload(v EzsignfolderImportEzsigntemplatepackageV1ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfolderImportEzsigntemplatepackageV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderImportEzsigntemplatepackageV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignfolderImportEzsigntemplatepackageV1Response struct {
	value *EzsignfolderImportEzsigntemplatepackageV1Response
	isSet bool
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1Response) Get() *EzsignfolderImportEzsigntemplatepackageV1Response {
	return v.value
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1Response) Set(val *EzsignfolderImportEzsigntemplatepackageV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderImportEzsigntemplatepackageV1Response(val *EzsignfolderImportEzsigntemplatepackageV1Response) *NullableEzsignfolderImportEzsigntemplatepackageV1Response {
	return &NullableEzsignfolderImportEzsigntemplatepackageV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderImportEzsigntemplatepackageV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderImportEzsigntemplatepackageV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

