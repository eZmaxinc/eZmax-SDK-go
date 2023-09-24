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

// checks if the UserGetEffectivePermissionsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetEffectivePermissionsV1Response{}

// UserGetEffectivePermissionsV1Response Response for GET /1/object/user/{pkiUserID}/getEffectivePermissions
type UserGetEffectivePermissionsV1Response struct {
	ObjDebugPayload CommonResponseObjDebugPayload `json:"objDebugPayload"`
	ObjDebug *CommonResponseObjDebug `json:"objDebug,omitempty"`
	MPayload UserGetEffectivePermissionsV1ResponseMPayload `json:"mPayload"`
}

// NewUserGetEffectivePermissionsV1Response instantiates a new UserGetEffectivePermissionsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetEffectivePermissionsV1Response(objDebugPayload CommonResponseObjDebugPayload, mPayload UserGetEffectivePermissionsV1ResponseMPayload) *UserGetEffectivePermissionsV1Response {
	this := UserGetEffectivePermissionsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewUserGetEffectivePermissionsV1ResponseWithDefaults instantiates a new UserGetEffectivePermissionsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetEffectivePermissionsV1ResponseWithDefaults() *UserGetEffectivePermissionsV1Response {
	this := UserGetEffectivePermissionsV1Response{}
	return &this
}

// GetObjDebugPayload returns the ObjDebugPayload field value
func (o *UserGetEffectivePermissionsV1Response) GetObjDebugPayload() CommonResponseObjDebugPayload {
	if o == nil {
		var ret CommonResponseObjDebugPayload
		return ret
	}

	return o.ObjDebugPayload
}

// GetObjDebugPayloadOk returns a tuple with the ObjDebugPayload field value
// and a boolean to check if the value has been set.
func (o *UserGetEffectivePermissionsV1Response) GetObjDebugPayloadOk() (*CommonResponseObjDebugPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjDebugPayload, true
}

// SetObjDebugPayload sets field value
func (o *UserGetEffectivePermissionsV1Response) SetObjDebugPayload(v CommonResponseObjDebugPayload) {
	o.ObjDebugPayload = v
}

// GetObjDebug returns the ObjDebug field value if set, zero value otherwise.
func (o *UserGetEffectivePermissionsV1Response) GetObjDebug() CommonResponseObjDebug {
	if o == nil || IsNil(o.ObjDebug) {
		var ret CommonResponseObjDebug
		return ret
	}
	return *o.ObjDebug
}

// GetObjDebugOk returns a tuple with the ObjDebug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGetEffectivePermissionsV1Response) GetObjDebugOk() (*CommonResponseObjDebug, bool) {
	if o == nil || IsNil(o.ObjDebug) {
		return nil, false
	}
	return o.ObjDebug, true
}

// HasObjDebug returns a boolean if a field has been set.
func (o *UserGetEffectivePermissionsV1Response) HasObjDebug() bool {
	if o != nil && !IsNil(o.ObjDebug) {
		return true
	}

	return false
}

// SetObjDebug gets a reference to the given CommonResponseObjDebug and assigns it to the ObjDebug field.
func (o *UserGetEffectivePermissionsV1Response) SetObjDebug(v CommonResponseObjDebug) {
	o.ObjDebug = &v
}

// GetMPayload returns the MPayload field value
func (o *UserGetEffectivePermissionsV1Response) GetMPayload() UserGetEffectivePermissionsV1ResponseMPayload {
	if o == nil {
		var ret UserGetEffectivePermissionsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *UserGetEffectivePermissionsV1Response) GetMPayloadOk() (*UserGetEffectivePermissionsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *UserGetEffectivePermissionsV1Response) SetMPayload(v UserGetEffectivePermissionsV1ResponseMPayload) {
	o.MPayload = v
}

func (o UserGetEffectivePermissionsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetEffectivePermissionsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objDebugPayload"] = o.ObjDebugPayload
	if !IsNil(o.ObjDebug) {
		toSerialize["objDebug"] = o.ObjDebug
	}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

type NullableUserGetEffectivePermissionsV1Response struct {
	value *UserGetEffectivePermissionsV1Response
	isSet bool
}

func (v NullableUserGetEffectivePermissionsV1Response) Get() *UserGetEffectivePermissionsV1Response {
	return v.value
}

func (v *NullableUserGetEffectivePermissionsV1Response) Set(val *UserGetEffectivePermissionsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetEffectivePermissionsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetEffectivePermissionsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetEffectivePermissionsV1Response(val *UserGetEffectivePermissionsV1Response) *NullableUserGetEffectivePermissionsV1Response {
	return &NullableUserGetEffectivePermissionsV1Response{value: val, isSet: true}
}

func (v NullableUserGetEffectivePermissionsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetEffectivePermissionsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


