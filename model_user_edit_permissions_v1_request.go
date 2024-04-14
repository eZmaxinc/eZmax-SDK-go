/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserEditPermissionsV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEditPermissionsV1Request{}

// UserEditPermissionsV1Request Request for PUT /1/object/user/{pkiUserID}/editPermissions
type UserEditPermissionsV1Request struct {
	AObjPermission []PermissionRequestCompound `json:"a_objPermission"`
}

type _UserEditPermissionsV1Request UserEditPermissionsV1Request

// NewUserEditPermissionsV1Request instantiates a new UserEditPermissionsV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEditPermissionsV1Request(aObjPermission []PermissionRequestCompound) *UserEditPermissionsV1Request {
	this := UserEditPermissionsV1Request{}
	this.AObjPermission = aObjPermission
	return &this
}

// NewUserEditPermissionsV1RequestWithDefaults instantiates a new UserEditPermissionsV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEditPermissionsV1RequestWithDefaults() *UserEditPermissionsV1Request {
	this := UserEditPermissionsV1Request{}
	return &this
}

// GetAObjPermission returns the AObjPermission field value
func (o *UserEditPermissionsV1Request) GetAObjPermission() []PermissionRequestCompound {
	if o == nil {
		var ret []PermissionRequestCompound
		return ret
	}

	return o.AObjPermission
}

// GetAObjPermissionOk returns a tuple with the AObjPermission field value
// and a boolean to check if the value has been set.
func (o *UserEditPermissionsV1Request) GetAObjPermissionOk() ([]PermissionRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPermission, true
}

// SetAObjPermission sets field value
func (o *UserEditPermissionsV1Request) SetAObjPermission(v []PermissionRequestCompound) {
	o.AObjPermission = v
}

func (o UserEditPermissionsV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEditPermissionsV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPermission"] = o.AObjPermission
	return toSerialize, nil
}

func (o *UserEditPermissionsV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objPermission",
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

	varUserEditPermissionsV1Request := _UserEditPermissionsV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserEditPermissionsV1Request)

	if err != nil {
		return err
	}

	*o = UserEditPermissionsV1Request(varUserEditPermissionsV1Request)

	return err
}

type NullableUserEditPermissionsV1Request struct {
	value *UserEditPermissionsV1Request
	isSet bool
}

func (v NullableUserEditPermissionsV1Request) Get() *UserEditPermissionsV1Request {
	return v.value
}

func (v *NullableUserEditPermissionsV1Request) Set(val *UserEditPermissionsV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEditPermissionsV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEditPermissionsV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEditPermissionsV1Request(val *UserEditPermissionsV1Request) *NullableUserEditPermissionsV1Request {
	return &NullableUserEditPermissionsV1Request{value: val, isSet: true}
}

func (v NullableUserEditPermissionsV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEditPermissionsV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


