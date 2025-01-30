/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserGetEffectivePermissionsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetEffectivePermissionsV1ResponseMPayload{}

// UserGetEffectivePermissionsV1ResponseMPayload Response for GET /1/object/user/{pkiUserID}/getEffectivePermissions
type UserGetEffectivePermissionsV1ResponseMPayload struct {
	AObjModulegroup []ModulegroupResponseCompound `json:"a_objModulegroup"`
}

type _UserGetEffectivePermissionsV1ResponseMPayload UserGetEffectivePermissionsV1ResponseMPayload

// NewUserGetEffectivePermissionsV1ResponseMPayload instantiates a new UserGetEffectivePermissionsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetEffectivePermissionsV1ResponseMPayload(aObjModulegroup []ModulegroupResponseCompound) *UserGetEffectivePermissionsV1ResponseMPayload {
	this := UserGetEffectivePermissionsV1ResponseMPayload{}
	this.AObjModulegroup = aObjModulegroup
	return &this
}

// NewUserGetEffectivePermissionsV1ResponseMPayloadWithDefaults instantiates a new UserGetEffectivePermissionsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetEffectivePermissionsV1ResponseMPayloadWithDefaults() *UserGetEffectivePermissionsV1ResponseMPayload {
	this := UserGetEffectivePermissionsV1ResponseMPayload{}
	return &this
}

// GetAObjModulegroup returns the AObjModulegroup field value
func (o *UserGetEffectivePermissionsV1ResponseMPayload) GetAObjModulegroup() []ModulegroupResponseCompound {
	if o == nil {
		var ret []ModulegroupResponseCompound
		return ret
	}

	return o.AObjModulegroup
}

// GetAObjModulegroupOk returns a tuple with the AObjModulegroup field value
// and a boolean to check if the value has been set.
func (o *UserGetEffectivePermissionsV1ResponseMPayload) GetAObjModulegroupOk() ([]ModulegroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjModulegroup, true
}

// SetAObjModulegroup sets field value
func (o *UserGetEffectivePermissionsV1ResponseMPayload) SetAObjModulegroup(v []ModulegroupResponseCompound) {
	o.AObjModulegroup = v
}

func (o UserGetEffectivePermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetEffectivePermissionsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objModulegroup"] = o.AObjModulegroup
	return toSerialize, nil
}

func (o *UserGetEffectivePermissionsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objModulegroup",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserGetEffectivePermissionsV1ResponseMPayload := _UserGetEffectivePermissionsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGetEffectivePermissionsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UserGetEffectivePermissionsV1ResponseMPayload(varUserGetEffectivePermissionsV1ResponseMPayload)

	return err
}

type NullableUserGetEffectivePermissionsV1ResponseMPayload struct {
	value *UserGetEffectivePermissionsV1ResponseMPayload
	isSet bool
}

func (v NullableUserGetEffectivePermissionsV1ResponseMPayload) Get() *UserGetEffectivePermissionsV1ResponseMPayload {
	return v.value
}

func (v *NullableUserGetEffectivePermissionsV1ResponseMPayload) Set(val *UserGetEffectivePermissionsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetEffectivePermissionsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetEffectivePermissionsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetEffectivePermissionsV1ResponseMPayload(val *UserGetEffectivePermissionsV1ResponseMPayload) *NullableUserGetEffectivePermissionsV1ResponseMPayload {
	return &NullableUserGetEffectivePermissionsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserGetEffectivePermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetEffectivePermissionsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


