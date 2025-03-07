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

// checks if the UserGetUsergroupsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetUsergroupsV1ResponseMPayload{}

// UserGetUsergroupsV1ResponseMPayload Response for GET /1/object/user/{pkiUserID}/getUsergroups
type UserGetUsergroupsV1ResponseMPayload struct {
	AObjUsergroup []UsergroupResponseCompound `json:"a_objUsergroup"`
}

type _UserGetUsergroupsV1ResponseMPayload UserGetUsergroupsV1ResponseMPayload

// NewUserGetUsergroupsV1ResponseMPayload instantiates a new UserGetUsergroupsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetUsergroupsV1ResponseMPayload(aObjUsergroup []UsergroupResponseCompound) *UserGetUsergroupsV1ResponseMPayload {
	this := UserGetUsergroupsV1ResponseMPayload{}
	this.AObjUsergroup = aObjUsergroup
	return &this
}

// NewUserGetUsergroupsV1ResponseMPayloadWithDefaults instantiates a new UserGetUsergroupsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetUsergroupsV1ResponseMPayloadWithDefaults() *UserGetUsergroupsV1ResponseMPayload {
	this := UserGetUsergroupsV1ResponseMPayload{}
	return &this
}

// GetAObjUsergroup returns the AObjUsergroup field value
func (o *UserGetUsergroupsV1ResponseMPayload) GetAObjUsergroup() []UsergroupResponseCompound {
	if o == nil {
		var ret []UsergroupResponseCompound
		return ret
	}

	return o.AObjUsergroup
}

// GetAObjUsergroupOk returns a tuple with the AObjUsergroup field value
// and a boolean to check if the value has been set.
func (o *UserGetUsergroupsV1ResponseMPayload) GetAObjUsergroupOk() ([]UsergroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUsergroup, true
}

// SetAObjUsergroup sets field value
func (o *UserGetUsergroupsV1ResponseMPayload) SetAObjUsergroup(v []UsergroupResponseCompound) {
	o.AObjUsergroup = v
}

func (o UserGetUsergroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetUsergroupsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUsergroup"] = o.AObjUsergroup
	return toSerialize, nil
}

func (o *UserGetUsergroupsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objUsergroup",
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

	varUserGetUsergroupsV1ResponseMPayload := _UserGetUsergroupsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGetUsergroupsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UserGetUsergroupsV1ResponseMPayload(varUserGetUsergroupsV1ResponseMPayload)

	return err
}

type NullableUserGetUsergroupsV1ResponseMPayload struct {
	value *UserGetUsergroupsV1ResponseMPayload
	isSet bool
}

func (v NullableUserGetUsergroupsV1ResponseMPayload) Get() *UserGetUsergroupsV1ResponseMPayload {
	return v.value
}

func (v *NullableUserGetUsergroupsV1ResponseMPayload) Set(val *UserGetUsergroupsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetUsergroupsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetUsergroupsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetUsergroupsV1ResponseMPayload(val *UserGetUsergroupsV1ResponseMPayload) *NullableUserGetUsergroupsV1ResponseMPayload {
	return &NullableUserGetUsergroupsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserGetUsergroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetUsergroupsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


