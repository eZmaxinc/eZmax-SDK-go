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

// checks if the UserGetSubnetsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGetSubnetsV1Response{}

// UserGetSubnetsV1Response Response for GET /1/object/user/{pkiUserID}/getSubnets
type UserGetSubnetsV1Response struct {
	CommonResponse
	MPayload UserGetSubnetsV1ResponseMPayload `json:"mPayload"`
}

type _UserGetSubnetsV1Response UserGetSubnetsV1Response

// NewUserGetSubnetsV1Response instantiates a new UserGetSubnetsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGetSubnetsV1Response(mPayload UserGetSubnetsV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *UserGetSubnetsV1Response {
	this := UserGetSubnetsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewUserGetSubnetsV1ResponseWithDefaults instantiates a new UserGetSubnetsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGetSubnetsV1ResponseWithDefaults() *UserGetSubnetsV1Response {
	this := UserGetSubnetsV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *UserGetSubnetsV1Response) GetMPayload() UserGetSubnetsV1ResponseMPayload {
	if o == nil {
		var ret UserGetSubnetsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *UserGetSubnetsV1Response) GetMPayloadOk() (*UserGetSubnetsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *UserGetSubnetsV1Response) SetMPayload(v UserGetSubnetsV1ResponseMPayload) {
	o.MPayload = v
}

func (o UserGetSubnetsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGetSubnetsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *UserGetSubnetsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varUserGetSubnetsV1Response := _UserGetSubnetsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserGetSubnetsV1Response)

	if err != nil {
		return err
	}

	*o = UserGetSubnetsV1Response(varUserGetSubnetsV1Response)

	return err
}

type NullableUserGetSubnetsV1Response struct {
	value *UserGetSubnetsV1Response
	isSet bool
}

func (v NullableUserGetSubnetsV1Response) Get() *UserGetSubnetsV1Response {
	return v.value
}

func (v *NullableUserGetSubnetsV1Response) Set(val *UserGetSubnetsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGetSubnetsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGetSubnetsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGetSubnetsV1Response(val *UserGetSubnetsV1Response) *NullableUserGetSubnetsV1Response {
	return &NullableUserGetSubnetsV1Response{value: val, isSet: true}
}

func (v NullableUserGetSubnetsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGetSubnetsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


