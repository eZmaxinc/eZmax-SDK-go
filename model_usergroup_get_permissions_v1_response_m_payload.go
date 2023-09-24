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

// checks if the UsergroupGetPermissionsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupGetPermissionsV1ResponseMPayload{}

// UsergroupGetPermissionsV1ResponseMPayload Response for GET /1/object/usergroup/{pkiUsergroupID}/getPermissions
type UsergroupGetPermissionsV1ResponseMPayload struct {
	AObjModulegroup []ModulegroupResponseCompound `json:"a_objModulegroup"`
}

// NewUsergroupGetPermissionsV1ResponseMPayload instantiates a new UsergroupGetPermissionsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupGetPermissionsV1ResponseMPayload(aObjModulegroup []ModulegroupResponseCompound) *UsergroupGetPermissionsV1ResponseMPayload {
	this := UsergroupGetPermissionsV1ResponseMPayload{}
	this.AObjModulegroup = aObjModulegroup
	return &this
}

// NewUsergroupGetPermissionsV1ResponseMPayloadWithDefaults instantiates a new UsergroupGetPermissionsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupGetPermissionsV1ResponseMPayloadWithDefaults() *UsergroupGetPermissionsV1ResponseMPayload {
	this := UsergroupGetPermissionsV1ResponseMPayload{}
	return &this
}

// GetAObjModulegroup returns the AObjModulegroup field value
func (o *UsergroupGetPermissionsV1ResponseMPayload) GetAObjModulegroup() []ModulegroupResponseCompound {
	if o == nil {
		var ret []ModulegroupResponseCompound
		return ret
	}

	return o.AObjModulegroup
}

// GetAObjModulegroupOk returns a tuple with the AObjModulegroup field value
// and a boolean to check if the value has been set.
func (o *UsergroupGetPermissionsV1ResponseMPayload) GetAObjModulegroupOk() ([]ModulegroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjModulegroup, true
}

// SetAObjModulegroup sets field value
func (o *UsergroupGetPermissionsV1ResponseMPayload) SetAObjModulegroup(v []ModulegroupResponseCompound) {
	o.AObjModulegroup = v
}

func (o UsergroupGetPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupGetPermissionsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objModulegroup"] = o.AObjModulegroup
	return toSerialize, nil
}

type NullableUsergroupGetPermissionsV1ResponseMPayload struct {
	value *UsergroupGetPermissionsV1ResponseMPayload
	isSet bool
}

func (v NullableUsergroupGetPermissionsV1ResponseMPayload) Get() *UsergroupGetPermissionsV1ResponseMPayload {
	return v.value
}

func (v *NullableUsergroupGetPermissionsV1ResponseMPayload) Set(val *UsergroupGetPermissionsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupGetPermissionsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupGetPermissionsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupGetPermissionsV1ResponseMPayload(val *UsergroupGetPermissionsV1ResponseMPayload) *NullableUsergroupGetPermissionsV1ResponseMPayload {
	return &NullableUsergroupGetPermissionsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUsergroupGetPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupGetPermissionsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

