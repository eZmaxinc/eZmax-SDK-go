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

// checks if the UserCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCreateObjectV1ResponseMPayload{}

// UserCreateObjectV1ResponseMPayload Payload for POST /1/object/user
type UserCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiUserID []int32 `json:"a_pkiUserID"`
}

type _UserCreateObjectV1ResponseMPayload UserCreateObjectV1ResponseMPayload

// NewUserCreateObjectV1ResponseMPayload instantiates a new UserCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCreateObjectV1ResponseMPayload(aPkiUserID []int32) *UserCreateObjectV1ResponseMPayload {
	this := UserCreateObjectV1ResponseMPayload{}
	this.APkiUserID = aPkiUserID
	return &this
}

// NewUserCreateObjectV1ResponseMPayloadWithDefaults instantiates a new UserCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCreateObjectV1ResponseMPayloadWithDefaults() *UserCreateObjectV1ResponseMPayload {
	this := UserCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiUserID returns the APkiUserID field value
func (o *UserCreateObjectV1ResponseMPayload) GetAPkiUserID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiUserID
}

// GetAPkiUserIDOk returns a tuple with the APkiUserID field value
// and a boolean to check if the value has been set.
func (o *UserCreateObjectV1ResponseMPayload) GetAPkiUserIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiUserID, true
}

// SetAPkiUserID sets field value
func (o *UserCreateObjectV1ResponseMPayload) SetAPkiUserID(v []int32) {
	o.APkiUserID = v
}

func (o UserCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiUserID"] = o.APkiUserID
	return toSerialize, nil
}

func (o *UserCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiUserID",
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

	varUserCreateObjectV1ResponseMPayload := _UserCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UserCreateObjectV1ResponseMPayload(varUserCreateObjectV1ResponseMPayload)

	return err
}

type NullableUserCreateObjectV1ResponseMPayload struct {
	value *UserCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableUserCreateObjectV1ResponseMPayload) Get() *UserCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableUserCreateObjectV1ResponseMPayload) Set(val *UserCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCreateObjectV1ResponseMPayload(val *UserCreateObjectV1ResponseMPayload) *NullableUserCreateObjectV1ResponseMPayload {
	return &NullableUserCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


