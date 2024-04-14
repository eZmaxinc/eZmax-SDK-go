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

// checks if the PhonetypeGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhonetypeGetAutocompleteV2ResponseMPayload{}

// PhonetypeGetAutocompleteV2ResponseMPayload Payload for POST /2/object/phonetype/getAutocomplete
type PhonetypeGetAutocompleteV2ResponseMPayload struct {
	// An array of Phonetype autocomplete element response.
	AObjPhonetype []PhonetypeAutocompleteElementResponse `json:"a_objPhonetype"`
}

type _PhonetypeGetAutocompleteV2ResponseMPayload PhonetypeGetAutocompleteV2ResponseMPayload

// NewPhonetypeGetAutocompleteV2ResponseMPayload instantiates a new PhonetypeGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhonetypeGetAutocompleteV2ResponseMPayload(aObjPhonetype []PhonetypeAutocompleteElementResponse) *PhonetypeGetAutocompleteV2ResponseMPayload {
	this := PhonetypeGetAutocompleteV2ResponseMPayload{}
	this.AObjPhonetype = aObjPhonetype
	return &this
}

// NewPhonetypeGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new PhonetypeGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhonetypeGetAutocompleteV2ResponseMPayloadWithDefaults() *PhonetypeGetAutocompleteV2ResponseMPayload {
	this := PhonetypeGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjPhonetype returns the AObjPhonetype field value
func (o *PhonetypeGetAutocompleteV2ResponseMPayload) GetAObjPhonetype() []PhonetypeAutocompleteElementResponse {
	if o == nil {
		var ret []PhonetypeAutocompleteElementResponse
		return ret
	}

	return o.AObjPhonetype
}

// GetAObjPhonetypeOk returns a tuple with the AObjPhonetype field value
// and a boolean to check if the value has been set.
func (o *PhonetypeGetAutocompleteV2ResponseMPayload) GetAObjPhonetypeOk() ([]PhonetypeAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjPhonetype, true
}

// SetAObjPhonetype sets field value
func (o *PhonetypeGetAutocompleteV2ResponseMPayload) SetAObjPhonetype(v []PhonetypeAutocompleteElementResponse) {
	o.AObjPhonetype = v
}

func (o PhonetypeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhonetypeGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objPhonetype"] = o.AObjPhonetype
	return toSerialize, nil
}

func (o *PhonetypeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objPhonetype",
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

	varPhonetypeGetAutocompleteV2ResponseMPayload := _PhonetypeGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhonetypeGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = PhonetypeGetAutocompleteV2ResponseMPayload(varPhonetypeGetAutocompleteV2ResponseMPayload)

	return err
}

type NullablePhonetypeGetAutocompleteV2ResponseMPayload struct {
	value *PhonetypeGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullablePhonetypeGetAutocompleteV2ResponseMPayload) Get() *PhonetypeGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullablePhonetypeGetAutocompleteV2ResponseMPayload) Set(val *PhonetypeGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullablePhonetypeGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePhonetypeGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhonetypeGetAutocompleteV2ResponseMPayload(val *PhonetypeGetAutocompleteV2ResponseMPayload) *NullablePhonetypeGetAutocompleteV2ResponseMPayload {
	return &NullablePhonetypeGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullablePhonetypeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhonetypeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


