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

// checks if the UserCreateObjectV1Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCreateObjectV1Request{}

// UserCreateObjectV1Request Request for POST /1/object/user
type UserCreateObjectV1Request struct {
	AObjUser []UserRequestCompound `json:"a_objUser"`
}

type _UserCreateObjectV1Request UserCreateObjectV1Request

// NewUserCreateObjectV1Request instantiates a new UserCreateObjectV1Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateObjectV1Request(aObjUser []UserRequestCompound) *UserCreateObjectV1Request {
	this := UserCreateObjectV1Request{}
	this.AObjUser = aObjUser
	return &this
}

// NewUserCreateObjectV1RequestWithDefaults instantiates a new UserCreateObjectV1Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateObjectV1RequestWithDefaults() *UserCreateObjectV1Request {
	this := UserCreateObjectV1Request{}
	return &this
}

// GetAObjUser returns the AObjUser field value
func (o *UserCreateObjectV1Request) GetAObjUser() []UserRequestCompound {
	if o == nil {
		var ret []UserRequestCompound
		return ret
	}

	return o.AObjUser
}

// GetAObjUserOk returns a tuple with the AObjUser field value
// and a boolean to check if the value has been set.
func (o *UserCreateObjectV1Request) GetAObjUserOk() ([]UserRequestCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUser, true
}

// SetAObjUser sets field value
func (o *UserCreateObjectV1Request) SetAObjUser(v []UserRequestCompound) {
	o.AObjUser = v
}

func (o UserCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreateObjectV1Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUser"] = o.AObjUser
	return toSerialize, nil
}

func (o *UserCreateObjectV1Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objUser",
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

	varUserCreateObjectV1Request := _UserCreateObjectV1Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCreateObjectV1Request)

	if err != nil {
		return err
	}

	*o = UserCreateObjectV1Request(varUserCreateObjectV1Request)

	return err
}

type NullableUserCreateObjectV1Request struct {
	value *UserCreateObjectV1Request
	isSet bool
}

func (v NullableUserCreateObjectV1Request) Get() *UserCreateObjectV1Request {
	return v.value
}

func (v *NullableUserCreateObjectV1Request) Set(val *UserCreateObjectV1Request) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateObjectV1Request) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateObjectV1Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateObjectV1Request(val *UserCreateObjectV1Request) *NullableUserCreateObjectV1Request {
	return &NullableUserCreateObjectV1Request{value: val, isSet: true}
}

func (v NullableUserCreateObjectV1Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateObjectV1Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


