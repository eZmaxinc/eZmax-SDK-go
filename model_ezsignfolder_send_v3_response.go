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

// checks if the EzsignfolderSendV3Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderSendV3Response{}

// EzsignfolderSendV3Response Response for POST /3/object/ezsignfolder/{pkiEzsignfolderID}/send
type EzsignfolderSendV3Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
}

// NewEzsignfolderSendV3Response instantiates a new EzsignfolderSendV3Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderSendV3Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignfolderSendV3Response {
	this := EzsignfolderSendV3Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignfolderSendV3ResponseWithDefaults instantiates a new EzsignfolderSendV3Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderSendV3ResponseWithDefaults() *EzsignfolderSendV3Response {
	this := EzsignfolderSendV3Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *EzsignfolderSendV3Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *EzsignfolderSendV3Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *EzsignfolderSendV3Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfolderSendV3Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *EzsignfolderSendV3Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *EzsignfolderSendV3Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

func (o EzsignfolderSendV3Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderSendV3Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	return toSerialize, nil
}

type NullableEzsignfolderSendV3Response struct {
	value *EzsignfolderSendV3Response
	isSet bool
}

func (v NullableEzsignfolderSendV3Response) Get() *EzsignfolderSendV3Response {
	return v.value
}

func (v *NullableEzsignfolderSendV3Response) Set(val *EzsignfolderSendV3Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderSendV3Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderSendV3Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderSendV3Response(val *EzsignfolderSendV3Response) *NullableEzsignfolderSendV3Response {
	return &NullableEzsignfolderSendV3Response{value: val, isSet: true}
}

func (v NullableEzsignfolderSendV3Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderSendV3Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

