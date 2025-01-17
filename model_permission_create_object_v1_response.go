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

// checks if the PermissionCreateObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionCreateObjectV1Response{}

// PermissionCreateObjectV1Response Response for POST /1/object/permission
type PermissionCreateObjectV1Response struct {
	CommonResponse
	MPayload PermissionCreateObjectV1ResponseMPayload `json:"mPayload"`
}

type _PermissionCreateObjectV1Response PermissionCreateObjectV1Response

// NewPermissionCreateObjectV1Response instantiates a new PermissionCreateObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionCreateObjectV1Response(mPayload PermissionCreateObjectV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *PermissionCreateObjectV1Response {
	this := PermissionCreateObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewPermissionCreateObjectV1ResponseWithDefaults instantiates a new PermissionCreateObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionCreateObjectV1ResponseWithDefaults() *PermissionCreateObjectV1Response {
	this := PermissionCreateObjectV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *PermissionCreateObjectV1Response) GetMPayload() PermissionCreateObjectV1ResponseMPayload {
	if o == nil {
		var ret PermissionCreateObjectV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *PermissionCreateObjectV1Response) GetMPayloadOk() (*PermissionCreateObjectV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *PermissionCreateObjectV1Response) SetMPayload(v PermissionCreateObjectV1ResponseMPayload) {
	o.MPayload = v
}

func (o PermissionCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionCreateObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *PermissionCreateObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varPermissionCreateObjectV1Response := _PermissionCreateObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPermissionCreateObjectV1Response)

	if err != nil {
		return err
	}

	*o = PermissionCreateObjectV1Response(varPermissionCreateObjectV1Response)

	return err
}

type NullablePermissionCreateObjectV1Response struct {
	value *PermissionCreateObjectV1Response
	isSet bool
}

func (v NullablePermissionCreateObjectV1Response) Get() *PermissionCreateObjectV1Response {
	return v.value
}

func (v *NullablePermissionCreateObjectV1Response) Set(val *PermissionCreateObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionCreateObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionCreateObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionCreateObjectV1Response(val *PermissionCreateObjectV1Response) *NullablePermissionCreateObjectV1Response {
	return &NullablePermissionCreateObjectV1Response{value: val, isSet: true}
}

func (v NullablePermissionCreateObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionCreateObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


