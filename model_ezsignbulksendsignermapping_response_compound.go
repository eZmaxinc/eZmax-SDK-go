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

// checks if the EzsignbulksendsignermappingResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendsignermappingResponseCompound{}

// EzsignbulksendsignermappingResponseCompound A Ezsignbulksendsignermapping Object
type EzsignbulksendsignermappingResponseCompound struct {
	EzsignbulksendsignermappingResponse
}

type _EzsignbulksendsignermappingResponseCompound EzsignbulksendsignermappingResponseCompound

// NewEzsignbulksendsignermappingResponseCompound instantiates a new EzsignbulksendsignermappingResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendsignermappingResponseCompound(pkiEzsignbulksendsignermappingID int32, fkiEzsignbulksendID int32, sEzsignbulksendsignermappingDescription string) *EzsignbulksendsignermappingResponseCompound {
	this := EzsignbulksendsignermappingResponseCompound{}
	this.PkiEzsignbulksendsignermappingID = pkiEzsignbulksendsignermappingID
	this.FkiEzsignbulksendID = fkiEzsignbulksendID
	this.SEzsignbulksendsignermappingDescription = sEzsignbulksendsignermappingDescription
	return &this
}

// NewEzsignbulksendsignermappingResponseCompoundWithDefaults instantiates a new EzsignbulksendsignermappingResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendsignermappingResponseCompoundWithDefaults() *EzsignbulksendsignermappingResponseCompound {
	this := EzsignbulksendsignermappingResponseCompound{}
	return &this
}

func (o EzsignbulksendsignermappingResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendsignermappingResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignbulksendsignermappingResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignbulksendsignermappingID",
		"fkiEzsignbulksendID",
		"sEzsignbulksendsignermappingDescription",
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

	varEzsignbulksendsignermappingResponseCompound := _EzsignbulksendsignermappingResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendsignermappingResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignbulksendsignermappingResponseCompound(varEzsignbulksendsignermappingResponseCompound)

	return err
}

type NullableEzsignbulksendsignermappingResponseCompound struct {
	value *EzsignbulksendsignermappingResponseCompound
	isSet bool
}

func (v NullableEzsignbulksendsignermappingResponseCompound) Get() *EzsignbulksendsignermappingResponseCompound {
	return v.value
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) Set(val *EzsignbulksendsignermappingResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendsignermappingResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendsignermappingResponseCompound(val *EzsignbulksendsignermappingResponseCompound) *NullableEzsignbulksendsignermappingResponseCompound {
	return &NullableEzsignbulksendsignermappingResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignbulksendsignermappingResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendsignermappingResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


