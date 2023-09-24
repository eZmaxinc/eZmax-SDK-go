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
)

// checks if the ApikeyCreateObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyCreateObjectV2ResponseMPayload{}

// ApikeyCreateObjectV2ResponseMPayload Payload for POST /2/object/apikey
type ApikeyCreateObjectV2ResponseMPayload struct {
	AObjApikey []ApikeyResponseCompound `json:"a_objApikey"`
}

// NewApikeyCreateObjectV2ResponseMPayload instantiates a new ApikeyCreateObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyCreateObjectV2ResponseMPayload(aObjApikey []ApikeyResponseCompound) *ApikeyCreateObjectV2ResponseMPayload {
	this := ApikeyCreateObjectV2ResponseMPayload{}
	this.AObjApikey = aObjApikey
	return &this
}

// NewApikeyCreateObjectV2ResponseMPayloadWithDefaults instantiates a new ApikeyCreateObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyCreateObjectV2ResponseMPayloadWithDefaults() *ApikeyCreateObjectV2ResponseMPayload {
	this := ApikeyCreateObjectV2ResponseMPayload{}
	return &this
}

// GetAObjApikey returns the AObjApikey field value
func (o *ApikeyCreateObjectV2ResponseMPayload) GetAObjApikey() []ApikeyResponseCompound {
	if o == nil {
		var ret []ApikeyResponseCompound
		return ret
	}

	return o.AObjApikey
}

// GetAObjApikeyOk returns a tuple with the AObjApikey field value
// and a boolean to check if the value has been set.
func (o *ApikeyCreateObjectV2ResponseMPayload) GetAObjApikeyOk() ([]ApikeyResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjApikey, true
}

// SetAObjApikey sets field value
func (o *ApikeyCreateObjectV2ResponseMPayload) SetAObjApikey(v []ApikeyResponseCompound) {
	o.AObjApikey = v
}

func (o ApikeyCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyCreateObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objApikey"] = o.AObjApikey
	return toSerialize, nil
}

type NullableApikeyCreateObjectV2ResponseMPayload struct {
	value *ApikeyCreateObjectV2ResponseMPayload
	isSet bool
}

func (v NullableApikeyCreateObjectV2ResponseMPayload) Get() *ApikeyCreateObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableApikeyCreateObjectV2ResponseMPayload) Set(val *ApikeyCreateObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyCreateObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyCreateObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyCreateObjectV2ResponseMPayload(val *ApikeyCreateObjectV2ResponseMPayload) *NullableApikeyCreateObjectV2ResponseMPayload {
	return &NullableApikeyCreateObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyCreateObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


