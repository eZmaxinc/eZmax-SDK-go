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

// checks if the ApikeyGetSubnetsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApikeyGetSubnetsV1ResponseMPayload{}

// ApikeyGetSubnetsV1ResponseMPayload Response for GET /1/object/apikey/{pkiApikeyID}/getSubnets
type ApikeyGetSubnetsV1ResponseMPayload struct {
	AObjSubnet []SubnetResponseCompound `json:"a_objSubnet"`
}

type _ApikeyGetSubnetsV1ResponseMPayload ApikeyGetSubnetsV1ResponseMPayload

// NewApikeyGetSubnetsV1ResponseMPayload instantiates a new ApikeyGetSubnetsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApikeyGetSubnetsV1ResponseMPayload(aObjSubnet []SubnetResponseCompound) *ApikeyGetSubnetsV1ResponseMPayload {
	this := ApikeyGetSubnetsV1ResponseMPayload{}
	this.AObjSubnet = aObjSubnet
	return &this
}

// NewApikeyGetSubnetsV1ResponseMPayloadWithDefaults instantiates a new ApikeyGetSubnetsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApikeyGetSubnetsV1ResponseMPayloadWithDefaults() *ApikeyGetSubnetsV1ResponseMPayload {
	this := ApikeyGetSubnetsV1ResponseMPayload{}
	return &this
}

// GetAObjSubnet returns the AObjSubnet field value
func (o *ApikeyGetSubnetsV1ResponseMPayload) GetAObjSubnet() []SubnetResponseCompound {
	if o == nil {
		var ret []SubnetResponseCompound
		return ret
	}

	return o.AObjSubnet
}

// GetAObjSubnetOk returns a tuple with the AObjSubnet field value
// and a boolean to check if the value has been set.
func (o *ApikeyGetSubnetsV1ResponseMPayload) GetAObjSubnetOk() ([]SubnetResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjSubnet, true
}

// SetAObjSubnet sets field value
func (o *ApikeyGetSubnetsV1ResponseMPayload) SetAObjSubnet(v []SubnetResponseCompound) {
	o.AObjSubnet = v
}

func (o ApikeyGetSubnetsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApikeyGetSubnetsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objSubnet"] = o.AObjSubnet
	return toSerialize, nil
}

func (o *ApikeyGetSubnetsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objSubnet",
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

	varApikeyGetSubnetsV1ResponseMPayload := _ApikeyGetSubnetsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApikeyGetSubnetsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ApikeyGetSubnetsV1ResponseMPayload(varApikeyGetSubnetsV1ResponseMPayload)

	return err
}

type NullableApikeyGetSubnetsV1ResponseMPayload struct {
	value *ApikeyGetSubnetsV1ResponseMPayload
	isSet bool
}

func (v NullableApikeyGetSubnetsV1ResponseMPayload) Get() *ApikeyGetSubnetsV1ResponseMPayload {
	return v.value
}

func (v *NullableApikeyGetSubnetsV1ResponseMPayload) Set(val *ApikeyGetSubnetsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyGetSubnetsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyGetSubnetsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyGetSubnetsV1ResponseMPayload(val *ApikeyGetSubnetsV1ResponseMPayload) *NullableApikeyGetSubnetsV1ResponseMPayload {
	return &NullableApikeyGetSubnetsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableApikeyGetSubnetsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyGetSubnetsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


