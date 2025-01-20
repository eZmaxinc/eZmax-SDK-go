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

// checks if the ApikeyRegenerateV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyRegenerateV1ResponseMPayload{}

// ApikeyRegenerateV1ResponseMPayload Response for GET /1/object/apikey/{pkiApikeyID}/regenerate
type ApikeyRegenerateV1ResponseMPayload struct {
	// An Apikey Object and children to create a complete structure
	ObjApikey ApikeyResponse `json:"objApikey"`
}

type _ApikeyRegenerateV1ResponseMPayload ApikeyRegenerateV1ResponseMPayload

// NewApikeyRegenerateV1ResponseMPayload instantiates a new ApikeyRegenerateV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyRegenerateV1ResponseMPayload(objApikey ApikeyResponse) *ApikeyRegenerateV1ResponseMPayload {
	this := ApikeyRegenerateV1ResponseMPayload{}
	this.ObjApikey = objApikey
	return &this
}

// NewApikeyRegenerateV1ResponseMPayloadWithDefaults instantiates a new ApikeyRegenerateV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyRegenerateV1ResponseMPayloadWithDefaults() *ApikeyRegenerateV1ResponseMPayload {
	this := ApikeyRegenerateV1ResponseMPayload{}
	return &this
}

// GetObjApikey returns the ObjApikey field value
func (o *ApikeyRegenerateV1ResponseMPayload) GetObjApikey() ApikeyResponse {
	if o == nil {
		var ret ApikeyResponse
		return ret
	}

	return o.ObjApikey
}

// GetObjApikeyOk returns a tuple with the ObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyRegenerateV1ResponseMPayload) GetObjApikeyOk() (*ApikeyResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjApikey, true
}

// SetObjApikey sets field value
func (o *ApikeyRegenerateV1ResponseMPayload) SetObjApikey(v ApikeyResponse) {
	o.ObjApikey = v
}

func (o ApikeyRegenerateV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyRegenerateV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objApikey"] = o.ObjApikey
	return toSerialize, nil
}

func (o *ApikeyRegenerateV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varApikeyRegenerateV1ResponseMPayload := _ApikeyRegenerateV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyRegenerateV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ApikeyRegenerateV1ResponseMPayload(varApikeyRegenerateV1ResponseMPayload)

	return err
}

type NullableApikeyRegenerateV1ResponseMPayload struct {
	value *ApikeyRegenerateV1ResponseMPayload
	isSet bool
}

func (v NullableApikeyRegenerateV1ResponseMPayload) Get() *ApikeyRegenerateV1ResponseMPayload {
	return v.value
}

func (v *NullableApikeyRegenerateV1ResponseMPayload) Set(val *ApikeyRegenerateV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyRegenerateV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyRegenerateV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyRegenerateV1ResponseMPayload(val *ApikeyRegenerateV1ResponseMPayload) *NullableApikeyRegenerateV1ResponseMPayload {
	return &NullableApikeyRegenerateV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyRegenerateV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyRegenerateV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


