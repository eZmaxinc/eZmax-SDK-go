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

// checks if the UsergroupmembershipResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupmembershipResponseCompound{}

// UsergroupmembershipResponseCompound A Usergroupmembership Object
type UsergroupmembershipResponseCompound struct {
	UsergroupmembershipResponse
}

type _UsergroupmembershipResponseCompound UsergroupmembershipResponseCompound

// NewUsergroupmembershipResponseCompound instantiates a new UsergroupmembershipResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupmembershipResponseCompound(pkiUsergroupmembershipID int32, fkiUsergroupID int32, sUsergroupNameX string) *UsergroupmembershipResponseCompound {
	this := UsergroupmembershipResponseCompound{}
	this.PkiUsergroupmembershipID = pkiUsergroupmembershipID
	this.FkiUsergroupID = fkiUsergroupID
	this.SUsergroupNameX = sUsergroupNameX
	return &this
}

// NewUsergroupmembershipResponseCompoundWithDefaults instantiates a new UsergroupmembershipResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupmembershipResponseCompoundWithDefaults() *UsergroupmembershipResponseCompound {
	this := UsergroupmembershipResponseCompound{}
	return &this
}

func (o UsergroupmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupmembershipResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *UsergroupmembershipResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupmembershipID",
		"fkiUsergroupID",
		"sUsergroupNameX",
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

	varUsergroupmembershipResponseCompound := _UsergroupmembershipResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupmembershipResponseCompound)

	if err != nil {
		return err
	}

	*o = UsergroupmembershipResponseCompound(varUsergroupmembershipResponseCompound)

	return err
}

type NullableUsergroupmembershipResponseCompound struct {
	value *UsergroupmembershipResponseCompound
	isSet bool
}

func (v NullableUsergroupmembershipResponseCompound) Get() *UsergroupmembershipResponseCompound {
	return v.value
}

func (v *NullableUsergroupmembershipResponseCompound) Set(val *UsergroupmembershipResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupmembershipResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupmembershipResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupmembershipResponseCompound(val *UsergroupmembershipResponseCompound) *NullableUsergroupmembershipResponseCompound {
	return &NullableUsergroupmembershipResponseCompound{value: val, isSet: true}
}

func (v NullableUsergroupmembershipResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupmembershipResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


