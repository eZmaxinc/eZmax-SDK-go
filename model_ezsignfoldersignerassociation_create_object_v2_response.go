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

// checks if the EzsignfoldersignerassociationCreateObjectV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateObjectV2Response{}

// EzsignfoldersignerassociationCreateObjectV2Response Response for POST /2/object/ezsignfoldersignerassociation
type EzsignfoldersignerassociationCreateObjectV2Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload EzsignfoldersignerassociationCreateObjectV2ResponseMPayload `json:"mPayload"`
}

// NewEzsignfoldersignerassociationCreateObjectV2Response instantiates a new EzsignfoldersignerassociationCreateObjectV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateObjectV2Response(objDebugPayload CommonResponseObjDebugPayload, mPayload EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) *EzsignfoldersignerassociationCreateObjectV2Response {
	this := EzsignfoldersignerassociationCreateObjectV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsignfoldersignerassociationCreateObjectV2ResponseWithDefaults instantiates a new EzsignfoldersignerassociationCreateObjectV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateObjectV2ResponseWithDefaults() *EzsignfoldersignerassociationCreateObjectV2Response {
	this := EzsignfoldersignerassociationCreateObjectV2Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetMPayload() EzsignfoldersignerassociationCreateObjectV2ResponseMPayload {
	if o == nil {
		var ret EzsignfoldersignerassociationCreateObjectV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateObjectV2Response) GetMPayloadOk() (*EzsignfoldersignerassociationCreateObjectV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsignfoldersignerassociationCreateObjectV2Response) SetMPayload(v EzsignfoldersignerassociationCreateObjectV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsignfoldersignerassociationCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateObjectV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableEzsignfoldersignerassociationCreateObjectV2Response struct {
	value *EzsignfoldersignerassociationCreateObjectV2Response
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) Get() *EzsignfoldersignerassociationCreateObjectV2Response {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) Set(val *EzsignfoldersignerassociationCreateObjectV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateObjectV2Response(val *EzsignfoldersignerassociationCreateObjectV2Response) *NullableEzsignfoldersignerassociationCreateObjectV2Response {
	return &NullableEzsignfoldersignerassociationCreateObjectV2Response{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateObjectV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateObjectV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

