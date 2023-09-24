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

// checks if the UserGetPermissionsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetPermissionsV1ResponseMPayload{}

// UserGetPermissionsV1ResponseMPayload Response for GET /1/object/user/{pkiUserID}/getPermissions
type UserGetPermissionsV1ResponseMPayload struct {
	AObjModulegroup []ModulegroupResponseCompound `json:"a_objModulegroup"`
}

// NewUserGetPermissionsV1ResponseMPayload instantiates a new UserGetPermissionsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetPermissionsV1ResponseMPayload(aObjModulegroup []ModulegroupResponseCompound) *UserGetPermissionsV1ResponseMPayload {
	this := UserGetPermissionsV1ResponseMPayload{}
	this.AObjModulegroup = aObjModulegroup
	return &this
}

// NewUserGetPermissionsV1ResponseMPayloadWithDefaults instantiates a new UserGetPermissionsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetPermissionsV1ResponseMPayloadWithDefaults() *UserGetPermissionsV1ResponseMPayload {
	this := UserGetPermissionsV1ResponseMPayload{}
	return &this
}

// GetAObjModulegroup returns the AObjModulegroup field value
func (o *UserGetPermissionsV1ResponseMPayload) GetAObjModulegroup() []ModulegroupResponseCompound {
	if o == nil {
		var ret []ModulegroupResponseCompound
		return ret
	}

	return o.AObjModulegroup
}

// GetAObjModulegroupOk returns a tuple with the AObjModulegroup field value
// and a boolean to check if the value has been set.
func (o *UserGetPermissionsV1ResponseMPayload) GetAObjModulegroupOk() ([]ModulegroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjModulegroup, true
}

// SetAObjModulegroup sets field value
func (o *UserGetPermissionsV1ResponseMPayload) SetAObjModulegroup(v []ModulegroupResponseCompound) {
	o.AObjModulegroup = v
}

func (o UserGetPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetPermissionsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objModulegroup"] = o.AObjModulegroup
	return toSerialize, nil
}

type NullableUserGetPermissionsV1ResponseMPayload struct {
	value *UserGetPermissionsV1ResponseMPayload
	isSet bool
}

func (v NullableUserGetPermissionsV1ResponseMPayload) Get() *UserGetPermissionsV1ResponseMPayload {
	return v.value
}

func (v *NullableUserGetPermissionsV1ResponseMPayload) Set(val *UserGetPermissionsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetPermissionsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetPermissionsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetPermissionsV1ResponseMPayload(val *UserGetPermissionsV1ResponseMPayload) *NullableUserGetPermissionsV1ResponseMPayload {
	return &NullableUserGetPermissionsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserGetPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetPermissionsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


