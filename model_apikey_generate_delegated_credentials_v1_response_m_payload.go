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

// checks if the ApikeyGenerateDelegatedCredentialsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyGenerateDelegatedCredentialsV1ResponseMPayload{}

// ApikeyGenerateDelegatedCredentialsV1ResponseMPayload Payload for POST /1/object/apikey/generateDelegatedCredentials
type ApikeyGenerateDelegatedCredentialsV1ResponseMPayload struct {
	ObjApikey CustomApikey `json:"objApikey"`
}

type _ApikeyGenerateDelegatedCredentialsV1ResponseMPayload ApikeyGenerateDelegatedCredentialsV1ResponseMPayload

// NewApikeyGenerateDelegatedCredentialsV1ResponseMPayload instantiates a new ApikeyGenerateDelegatedCredentialsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyGenerateDelegatedCredentialsV1ResponseMPayload(objApikey CustomApikey) *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload {
	this := ApikeyGenerateDelegatedCredentialsV1ResponseMPayload{}
	this.ObjApikey = objApikey
	return &this
}

// NewApikeyGenerateDelegatedCredentialsV1ResponseMPayloadWithDefaults instantiates a new ApikeyGenerateDelegatedCredentialsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyGenerateDelegatedCredentialsV1ResponseMPayloadWithDefaults() *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload {
	this := ApikeyGenerateDelegatedCredentialsV1ResponseMPayload{}
	return &this
}

// GetObjApikey returns the ObjApikey field value
func (o *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) GetObjApikey() CustomApikey {
	if o == nil {
		var ret CustomApikey
		return ret
	}

	return o.ObjApikey
}

// GetObjApikeyOk returns a tuple with the ObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) GetObjApikeyOk() (*CustomApikey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjApikey, true
}

// SetObjApikey sets field value
func (o *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) SetObjApikey(v CustomApikey) {
	o.ObjApikey = v
}

func (o ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objApikey"] = o.ObjApikey
	return toSerialize, nil
}

func (o *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objApikey",
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

	varApikeyGenerateDelegatedCredentialsV1ResponseMPayload := _ApikeyGenerateDelegatedCredentialsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyGenerateDelegatedCredentialsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ApikeyGenerateDelegatedCredentialsV1ResponseMPayload(varApikeyGenerateDelegatedCredentialsV1ResponseMPayload)

	return err
}

type NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload struct {
	value *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload
	isSet bool
}

func (v NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) Get() *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload {
	return v.value
}

func (v *NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) Set(val *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload(val *ApikeyGenerateDelegatedCredentialsV1ResponseMPayload) *NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload {
	return &NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyGenerateDelegatedCredentialsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


