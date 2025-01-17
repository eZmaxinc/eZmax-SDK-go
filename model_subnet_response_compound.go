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

// checks if the SubnetResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubnetResponseCompound{}

// SubnetResponseCompound A Subnet Object
type SubnetResponseCompound struct {
	SubnetResponse
}

type _SubnetResponseCompound SubnetResponseCompound

// NewSubnetResponseCompound instantiates a new SubnetResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubnetResponseCompound(pkiSubnetID int32, objSubnetDescription MultilingualSubnetDescription, iSubnetNetwork int64, iSubnetMask int64) *SubnetResponseCompound {
	this := SubnetResponseCompound{}
	this.PkiSubnetID = pkiSubnetID
	this.ObjSubnetDescription = objSubnetDescription
	this.ISubnetNetwork = iSubnetNetwork
	this.ISubnetMask = iSubnetMask
	return &this
}

// NewSubnetResponseCompoundWithDefaults instantiates a new SubnetResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubnetResponseCompoundWithDefaults() *SubnetResponseCompound {
	this := SubnetResponseCompound{}
	return &this
}

func (o SubnetResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubnetResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *SubnetResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSubnetID",
		"objSubnetDescription",
		"iSubnetNetwork",
		"iSubnetMask",
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

	varSubnetResponseCompound := _SubnetResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubnetResponseCompound)

	if err != nil {
		return err
	}

	*o = SubnetResponseCompound(varSubnetResponseCompound)

	return err
}

type NullableSubnetResponseCompound struct {
	value *SubnetResponseCompound
	isSet bool
}

func (v NullableSubnetResponseCompound) Get() *SubnetResponseCompound {
	return v.value
}

func (v *NullableSubnetResponseCompound) Set(val *SubnetResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSubnetResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSubnetResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubnetResponseCompound(val *SubnetResponseCompound) *NullableSubnetResponseCompound {
	return &NullableSubnetResponseCompound{value: val, isSet: true}
}

func (v NullableSubnetResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubnetResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


