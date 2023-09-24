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

// checks if the SubnetGetObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubnetGetObjectV2ResponseMPayload{}

// SubnetGetObjectV2ResponseMPayload Payload for GET /2/object/subnet/{pkiSubnetID}
type SubnetGetObjectV2ResponseMPayload struct {
	ObjSubnet SubnetResponseCompound `json:"objSubnet"`
}

// NewSubnetGetObjectV2ResponseMPayload instantiates a new SubnetGetObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetGetObjectV2ResponseMPayload(objSubnet SubnetResponseCompound) *SubnetGetObjectV2ResponseMPayload {
	this := SubnetGetObjectV2ResponseMPayload{}
	this.ObjSubnet = objSubnet
	return &this
}

// NewSubnetGetObjectV2ResponseMPayloadWithDefaults instantiates a new SubnetGetObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetGetObjectV2ResponseMPayloadWithDefaults() *SubnetGetObjectV2ResponseMPayload {
	this := SubnetGetObjectV2ResponseMPayload{}
	return &this
}

// GetObjSubnet returns the ObjSubnet field value
func (o *SubnetGetObjectV2ResponseMPayload) GetObjSubnet() SubnetResponseCompound {
	if o == nil {
		var ret SubnetResponseCompound
		return ret
	}

	return o.ObjSubnet
}

// GetObjSubnetOk returns a tuple with the ObjSubnet field value
// and a boolean to check if the value has been set.
func (o *SubnetGetObjectV2ResponseMPayload) GetObjSubnetOk() (*SubnetResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjSubnet, true
}

// SetObjSubnet sets field value
func (o *SubnetGetObjectV2ResponseMPayload) SetObjSubnet(v SubnetResponseCompound) {
	o.ObjSubnet = v
}

func (o SubnetGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubnetGetObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objSubnet"] = o.ObjSubnet
	return toSerialize, nil
}

type NullableSubnetGetObjectV2ResponseMPayload struct {
	value *SubnetGetObjectV2ResponseMPayload
	isSet bool
}

func (v NullableSubnetGetObjectV2ResponseMPayload) Get() *SubnetGetObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableSubnetGetObjectV2ResponseMPayload) Set(val *SubnetGetObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetGetObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetGetObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetGetObjectV2ResponseMPayload(val *SubnetGetObjectV2ResponseMPayload) *NullableSubnetGetObjectV2ResponseMPayload {
	return &NullableSubnetGetObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableSubnetGetObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetGetObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


