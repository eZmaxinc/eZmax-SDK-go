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

// checks if the UserGetPermissionsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetPermissionsV1Response{}

// UserGetPermissionsV1Response Response for GET /1/object/user/{pkiUserID}/getPermissions
type UserGetPermissionsV1Response struct {
	CommonResponse
	MPayload UserGetPermissionsV1ResponseMPayload `json:"mPayload"`
}

type _UserGetPermissionsV1Response UserGetPermissionsV1Response

// NewUserGetPermissionsV1Response instantiates a new UserGetPermissionsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetPermissionsV1Response(mPayload UserGetPermissionsV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *UserGetPermissionsV1Response {
	this := UserGetPermissionsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewUserGetPermissionsV1ResponseWithDefaults instantiates a new UserGetPermissionsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetPermissionsV1ResponseWithDefaults() *UserGetPermissionsV1Response {
	this := UserGetPermissionsV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *UserGetPermissionsV1Response) GetMPayload() UserGetPermissionsV1ResponseMPayload {
	if o == nil {
		var ret UserGetPermissionsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *UserGetPermissionsV1Response) GetMPayloadOk() (*UserGetPermissionsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *UserGetPermissionsV1Response) SetMPayload(v UserGetPermissionsV1ResponseMPayload) {
	o.MPayload = v
}

func (o UserGetPermissionsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetPermissionsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *UserGetPermissionsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varUserGetPermissionsV1Response := _UserGetPermissionsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGetPermissionsV1Response)

	if err != nil {
		return err
	}

	*o = UserGetPermissionsV1Response(varUserGetPermissionsV1Response)

	return err
}

type NullableUserGetPermissionsV1Response struct {
	value *UserGetPermissionsV1Response
	isSet bool
}

func (v NullableUserGetPermissionsV1Response) Get() *UserGetPermissionsV1Response {
	return v.value
}

func (v *NullableUserGetPermissionsV1Response) Set(val *UserGetPermissionsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetPermissionsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetPermissionsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetPermissionsV1Response(val *UserGetPermissionsV1Response) *NullableUserGetPermissionsV1Response {
	return &NullableUserGetPermissionsV1Response{value: val, isSet: true}
}

func (v NullableUserGetPermissionsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetPermissionsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


