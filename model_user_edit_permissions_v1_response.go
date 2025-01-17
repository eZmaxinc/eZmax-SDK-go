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

// checks if the UserEditPermissionsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEditPermissionsV1Response{}

// UserEditPermissionsV1Response Response for PUT /1/object/user/{pkiUserID}/editPermissions
type UserEditPermissionsV1Response struct {
	CommonResponse
	MPayload UserEditPermissionsV1ResponseMPayload `json:"mPayload"`
}

type _UserEditPermissionsV1Response UserEditPermissionsV1Response

// NewUserEditPermissionsV1Response instantiates a new UserEditPermissionsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEditPermissionsV1Response(mPayload UserEditPermissionsV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *UserEditPermissionsV1Response {
	this := UserEditPermissionsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewUserEditPermissionsV1ResponseWithDefaults instantiates a new UserEditPermissionsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEditPermissionsV1ResponseWithDefaults() *UserEditPermissionsV1Response {
	this := UserEditPermissionsV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *UserEditPermissionsV1Response) GetMPayload() UserEditPermissionsV1ResponseMPayload {
	if o == nil {
		var ret UserEditPermissionsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *UserEditPermissionsV1Response) GetMPayloadOk() (*UserEditPermissionsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *UserEditPermissionsV1Response) SetMPayload(v UserEditPermissionsV1ResponseMPayload) {
	o.MPayload = v
}

func (o UserEditPermissionsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEditPermissionsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *UserEditPermissionsV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
		"objDebugPayload",
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

	varUserEditPermissionsV1Response := _UserEditPermissionsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserEditPermissionsV1Response)

	if err != nil {
		return err
	}

	*o = UserEditPermissionsV1Response(varUserEditPermissionsV1Response)

	return err
}

type NullableUserEditPermissionsV1Response struct {
	value *UserEditPermissionsV1Response
	isSet bool
}

func (v NullableUserEditPermissionsV1Response) Get() *UserEditPermissionsV1Response {
	return v.value
}

func (v *NullableUserEditPermissionsV1Response) Set(val *UserEditPermissionsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEditPermissionsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEditPermissionsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEditPermissionsV1Response(val *UserEditPermissionsV1Response) *NullableUserEditPermissionsV1Response {
	return &NullableUserEditPermissionsV1Response{value: val, isSet: true}
}

func (v NullableUserEditPermissionsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEditPermissionsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


