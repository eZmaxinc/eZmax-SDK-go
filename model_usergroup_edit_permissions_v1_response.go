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

// checks if the UsergroupEditPermissionsV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupEditPermissionsV1Response{}

// UsergroupEditPermissionsV1Response Response for PUT /1/object/usergroup/{pkiUsergroupID}/editPermissions
type UsergroupEditPermissionsV1Response struct {
	CommonResponse
	MPayload UsergroupEditPermissionsV1ResponseMPayload `json:"mPayload"`
}

type _UsergroupEditPermissionsV1Response UsergroupEditPermissionsV1Response

// NewUsergroupEditPermissionsV1Response instantiates a new UsergroupEditPermissionsV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupEditPermissionsV1Response(mPayload UsergroupEditPermissionsV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *UsergroupEditPermissionsV1Response {
	this := UsergroupEditPermissionsV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewUsergroupEditPermissionsV1ResponseWithDefaults instantiates a new UsergroupEditPermissionsV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupEditPermissionsV1ResponseWithDefaults() *UsergroupEditPermissionsV1Response {
	this := UsergroupEditPermissionsV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *UsergroupEditPermissionsV1Response) GetMPayload() UsergroupEditPermissionsV1ResponseMPayload {
	if o == nil {
		var ret UsergroupEditPermissionsV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *UsergroupEditPermissionsV1Response) GetMPayloadOk() (*UsergroupEditPermissionsV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *UsergroupEditPermissionsV1Response) SetMPayload(v UsergroupEditPermissionsV1ResponseMPayload) {
	o.MPayload = v
}

func (o UsergroupEditPermissionsV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupEditPermissionsV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *UsergroupEditPermissionsV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varUsergroupEditPermissionsV1Response := _UsergroupEditPermissionsV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupEditPermissionsV1Response)

	if err != nil {
		return err
	}

	*o = UsergroupEditPermissionsV1Response(varUsergroupEditPermissionsV1Response)

	return err
}

type NullableUsergroupEditPermissionsV1Response struct {
	value *UsergroupEditPermissionsV1Response
	isSet bool
}

func (v NullableUsergroupEditPermissionsV1Response) Get() *UsergroupEditPermissionsV1Response {
	return v.value
}

func (v *NullableUsergroupEditPermissionsV1Response) Set(val *UsergroupEditPermissionsV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupEditPermissionsV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupEditPermissionsV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupEditPermissionsV1Response(val *UsergroupEditPermissionsV1Response) *NullableUsergroupEditPermissionsV1Response {
	return &NullableUsergroupEditPermissionsV1Response{value: val, isSet: true}
}

func (v NullableUsergroupEditPermissionsV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupEditPermissionsV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


