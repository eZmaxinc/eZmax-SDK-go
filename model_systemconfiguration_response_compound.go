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

// checks if the SystemconfigurationResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemconfigurationResponseCompound{}

// SystemconfigurationResponseCompound A Systemconfiguration Object
type SystemconfigurationResponseCompound struct {
	SystemconfigurationResponse
}

type _SystemconfigurationResponseCompound SystemconfigurationResponseCompound

// NewSystemconfigurationResponseCompound instantiates a new SystemconfigurationResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemconfigurationResponseCompound(pkiSystemconfigurationID int32, fkiSystemconfigurationtypeID int32, sSystemconfigurationtypeDescriptionX string, eSystemconfigurationNewexternaluseraction FieldESystemconfigurationNewexternaluseraction, eSystemconfigurationLanguage1 FieldESystemconfigurationLanguage1, eSystemconfigurationLanguage2 FieldESystemconfigurationLanguage2, bSystemconfigurationEzsignpersonnal bool, bSystemconfigurationSspr bool) *SystemconfigurationResponseCompound {
	this := SystemconfigurationResponseCompound{}
	this.PkiSystemconfigurationID = pkiSystemconfigurationID
	this.FkiSystemconfigurationtypeID = fkiSystemconfigurationtypeID
	this.SSystemconfigurationtypeDescriptionX = sSystemconfigurationtypeDescriptionX
	this.ESystemconfigurationNewexternaluseraction = eSystemconfigurationNewexternaluseraction
	this.ESystemconfigurationLanguage1 = eSystemconfigurationLanguage1
	this.ESystemconfigurationLanguage2 = eSystemconfigurationLanguage2
	this.BSystemconfigurationEzsignpersonnal = bSystemconfigurationEzsignpersonnal
	this.BSystemconfigurationSspr = bSystemconfigurationSspr
	return &this
}

// NewSystemconfigurationResponseCompoundWithDefaults instantiates a new SystemconfigurationResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemconfigurationResponseCompoundWithDefaults() *SystemconfigurationResponseCompound {
	this := SystemconfigurationResponseCompound{}
	return &this
}

func (o SystemconfigurationResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemconfigurationResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *SystemconfigurationResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSystemconfigurationID",
		"fkiSystemconfigurationtypeID",
		"sSystemconfigurationtypeDescriptionX",
		"eSystemconfigurationNewexternaluseraction",
		"eSystemconfigurationLanguage1",
		"eSystemconfigurationLanguage2",
		"bSystemconfigurationEzsignpersonnal",
		"bSystemconfigurationSspr",
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

	varSystemconfigurationResponseCompound := _SystemconfigurationResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSystemconfigurationResponseCompound)

	if err != nil {
		return err
	}

	*o = SystemconfigurationResponseCompound(varSystemconfigurationResponseCompound)

	return err
}

type NullableSystemconfigurationResponseCompound struct {
	value *SystemconfigurationResponseCompound
	isSet bool
}

func (v NullableSystemconfigurationResponseCompound) Get() *SystemconfigurationResponseCompound {
	return v.value
}

func (v *NullableSystemconfigurationResponseCompound) Set(val *SystemconfigurationResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemconfigurationResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemconfigurationResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemconfigurationResponseCompound(val *SystemconfigurationResponseCompound) *NullableSystemconfigurationResponseCompound {
	return &NullableSystemconfigurationResponseCompound{value: val, isSet: true}
}

func (v NullableSystemconfigurationResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemconfigurationResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


