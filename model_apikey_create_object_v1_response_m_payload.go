/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign application.  We provide SDKs for customers. They are generated using OpenAPI codegen, we encourage customers to use them as we also provide samples for them.  You can choose to build your own implementation manually or can use any compatible OpenAPI 3.0 generator like Swagger Codegen, OpenAPI codegen or any commercial generators.  If you need helping understanding how to use this API, don't waste too much time looking for it. Contact support-api@ezmax.ca, we're here to help. We are developpers so we know programmers don't like bad documentation. If you don't find what you need in the documentation, let us know, we'll improve it and put you rapidly up on track.
 *
 * API version: 1.0.29
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// ApikeyCreateObjectV1ResponseMPayload Payload for the /1/object/apikey/createObject API Request
type ApikeyCreateObjectV1ResponseMPayload struct {
	AObjApikey []ApikeyResponse `json:"a_objApikey"`
}

// NewApikeyCreateObjectV1ResponseMPayload instantiates a new ApikeyCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyCreateObjectV1ResponseMPayload(aObjApikey []ApikeyResponse, ) *ApikeyCreateObjectV1ResponseMPayload {
	this := ApikeyCreateObjectV1ResponseMPayload{}
	this.AObjApikey = aObjApikey
	return &this
}

// NewApikeyCreateObjectV1ResponseMPayloadWithDefaults instantiates a new ApikeyCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyCreateObjectV1ResponseMPayloadWithDefaults() *ApikeyCreateObjectV1ResponseMPayload {
	this := ApikeyCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAObjApikey returns the AObjApikey field value
func (o *ApikeyCreateObjectV1ResponseMPayload) GetAObjApikey() []ApikeyResponse {
	if o == nil  {
		var ret []ApikeyResponse
		return ret
	}

	return o.AObjApikey
}

// GetAObjApikeyOk returns a tuple with the AObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV1ResponseMPayload) GetAObjApikeyOk() (*[]ApikeyResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjApikey, true
}

// SetAObjApikey sets field value
func (o *ApikeyCreateObjectV1ResponseMPayload) SetAObjApikey(v []ApikeyResponse) {
	o.AObjApikey = v
}

func (o ApikeyCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_objApikey"] = o.AObjApikey
	}
	return json.Marshal(toSerialize)
}

type NullableApikeyCreateObjectV1ResponseMPayload struct {
	value *ApikeyCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableApikeyCreateObjectV1ResponseMPayload) Get() *ApikeyCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableApikeyCreateObjectV1ResponseMPayload) Set(val *ApikeyCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyCreateObjectV1ResponseMPayload(val *ApikeyCreateObjectV1ResponseMPayload) *NullableApikeyCreateObjectV1ResponseMPayload {
	return &NullableApikeyCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


