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

// checks if the AuthenticationexternalCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationexternalCreateObjectV1ResponseMPayload{}

// AuthenticationexternalCreateObjectV1ResponseMPayload Payload for POST /1/object/authenticationexternal
type AuthenticationexternalCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiAuthenticationexternalID []int32 `json:"a_pkiAuthenticationexternalID"`
}

type _AuthenticationexternalCreateObjectV1ResponseMPayload AuthenticationexternalCreateObjectV1ResponseMPayload

// NewAuthenticationexternalCreateObjectV1ResponseMPayload instantiates a new AuthenticationexternalCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationexternalCreateObjectV1ResponseMPayload(aPkiAuthenticationexternalID []int32) *AuthenticationexternalCreateObjectV1ResponseMPayload {
	this := AuthenticationexternalCreateObjectV1ResponseMPayload{}
	this.APkiAuthenticationexternalID = aPkiAuthenticationexternalID
	return &this
}

// NewAuthenticationexternalCreateObjectV1ResponseMPayloadWithDefaults instantiates a new AuthenticationexternalCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationexternalCreateObjectV1ResponseMPayloadWithDefaults() *AuthenticationexternalCreateObjectV1ResponseMPayload {
	this := AuthenticationexternalCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiAuthenticationexternalID returns the APkiAuthenticationexternalID field value
func (o *AuthenticationexternalCreateObjectV1ResponseMPayload) GetAPkiAuthenticationexternalID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiAuthenticationexternalID
}

// GetAPkiAuthenticationexternalIDOk returns a tuple with the APkiAuthenticationexternalID field value
// and a boolean to check if the value has been set.
func (o *AuthenticationexternalCreateObjectV1ResponseMPayload) GetAPkiAuthenticationexternalIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiAuthenticationexternalID, true
}

// SetAPkiAuthenticationexternalID sets field value
func (o *AuthenticationexternalCreateObjectV1ResponseMPayload) SetAPkiAuthenticationexternalID(v []int32) {
	o.APkiAuthenticationexternalID = v
}

func (o AuthenticationexternalCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationexternalCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiAuthenticationexternalID"] = o.APkiAuthenticationexternalID
	return toSerialize, nil
}

func (o *AuthenticationexternalCreateObjectV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiAuthenticationexternalID",
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

	varAuthenticationexternalCreateObjectV1ResponseMPayload := _AuthenticationexternalCreateObjectV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthenticationexternalCreateObjectV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = AuthenticationexternalCreateObjectV1ResponseMPayload(varAuthenticationexternalCreateObjectV1ResponseMPayload)

	return err
}

type NullableAuthenticationexternalCreateObjectV1ResponseMPayload struct {
	value *AuthenticationexternalCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableAuthenticationexternalCreateObjectV1ResponseMPayload) Get() *AuthenticationexternalCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableAuthenticationexternalCreateObjectV1ResponseMPayload) Set(val *AuthenticationexternalCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationexternalCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationexternalCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationexternalCreateObjectV1ResponseMPayload(val *AuthenticationexternalCreateObjectV1ResponseMPayload) *NullableAuthenticationexternalCreateObjectV1ResponseMPayload {
	return &NullableAuthenticationexternalCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableAuthenticationexternalCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationexternalCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


