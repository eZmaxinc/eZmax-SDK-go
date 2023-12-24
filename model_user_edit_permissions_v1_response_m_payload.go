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
	"bytes"
	"fmt"
)

// checks if the UserEditPermissionsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEditPermissionsV1ResponseMPayload{}

// UserEditPermissionsV1ResponseMPayload Payload for PUT /1/object/user/{pkiUserID}/editPermissions
type UserEditPermissionsV1ResponseMPayload struct {
	APkiPermissionID []int32 `json:"a_pkiPermissionID"`
}

type _UserEditPermissionsV1ResponseMPayload UserEditPermissionsV1ResponseMPayload

// NewUserEditPermissionsV1ResponseMPayload instantiates a new UserEditPermissionsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEditPermissionsV1ResponseMPayload(aPkiPermissionID []int32) *UserEditPermissionsV1ResponseMPayload {
	this := UserEditPermissionsV1ResponseMPayload{}
	this.APkiPermissionID = aPkiPermissionID
	return &this
}

// NewUserEditPermissionsV1ResponseMPayloadWithDefaults instantiates a new UserEditPermissionsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEditPermissionsV1ResponseMPayloadWithDefaults() *UserEditPermissionsV1ResponseMPayload {
	this := UserEditPermissionsV1ResponseMPayload{}
	return &this
}

// GetAPkiPermissionID returns the APkiPermissionID field value
func (o *UserEditPermissionsV1ResponseMPayload) GetAPkiPermissionID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiPermissionID
}

// GetAPkiPermissionIDOk returns a tuple with the APkiPermissionID field value
// and a boolean to check if the value has been set.
func (o *UserEditPermissionsV1ResponseMPayload) GetAPkiPermissionIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiPermissionID, true
}

// SetAPkiPermissionID sets field value
func (o *UserEditPermissionsV1ResponseMPayload) SetAPkiPermissionID(v []int32) {
	o.APkiPermissionID = v
}

func (o UserEditPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEditPermissionsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiPermissionID"] = o.APkiPermissionID
	return toSerialize, nil
}

func (o *UserEditPermissionsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiPermissionID",
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

	varUserEditPermissionsV1ResponseMPayload := _UserEditPermissionsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserEditPermissionsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UserEditPermissionsV1ResponseMPayload(varUserEditPermissionsV1ResponseMPayload)

	return err
}

type NullableUserEditPermissionsV1ResponseMPayload struct {
	value *UserEditPermissionsV1ResponseMPayload
	isSet bool
}

func (v NullableUserEditPermissionsV1ResponseMPayload) Get() *UserEditPermissionsV1ResponseMPayload {
	return v.value
}

func (v *NullableUserEditPermissionsV1ResponseMPayload) Set(val *UserEditPermissionsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEditPermissionsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEditPermissionsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEditPermissionsV1ResponseMPayload(val *UserEditPermissionsV1ResponseMPayload) *NullableUserEditPermissionsV1ResponseMPayload {
	return &NullableUserEditPermissionsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserEditPermissionsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEditPermissionsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


