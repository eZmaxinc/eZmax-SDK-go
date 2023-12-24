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
	"bytes"
	"fmt"
)

// checks if the LanguageGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LanguageGetAutocompleteV2ResponseMPayload{}

// LanguageGetAutocompleteV2ResponseMPayload Payload for POST /2/object/language/getAutocomplete
type LanguageGetAutocompleteV2ResponseMPayload struct {
	// An array of Language autocomplete element response.
	AObjLanguage []LanguageAutocompleteElementResponse `json:"a_objLanguage"`
}

type _LanguageGetAutocompleteV2ResponseMPayload LanguageGetAutocompleteV2ResponseMPayload

// NewLanguageGetAutocompleteV2ResponseMPayload instantiates a new LanguageGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLanguageGetAutocompleteV2ResponseMPayload(aObjLanguage []LanguageAutocompleteElementResponse) *LanguageGetAutocompleteV2ResponseMPayload {
	this := LanguageGetAutocompleteV2ResponseMPayload{}
	this.AObjLanguage = aObjLanguage
	return &this
}

// NewLanguageGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new LanguageGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLanguageGetAutocompleteV2ResponseMPayloadWithDefaults() *LanguageGetAutocompleteV2ResponseMPayload {
	this := LanguageGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjLanguage returns the AObjLanguage field value
func (o *LanguageGetAutocompleteV2ResponseMPayload) GetAObjLanguage() []LanguageAutocompleteElementResponse {
	if o == nil {
		var ret []LanguageAutocompleteElementResponse
		return ret
	}

	return o.AObjLanguage
}

// GetAObjLanguageOk returns a tuple with the AObjLanguage field value
// and a boolean to check if the value has been set.
func (o *LanguageGetAutocompleteV2ResponseMPayload) GetAObjLanguageOk() ([]LanguageAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjLanguage, true
}

// SetAObjLanguage sets field value
func (o *LanguageGetAutocompleteV2ResponseMPayload) SetAObjLanguage(v []LanguageAutocompleteElementResponse) {
	o.AObjLanguage = v
}

func (o LanguageGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LanguageGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objLanguage"] = o.AObjLanguage
	return toSerialize, nil
}

func (o *LanguageGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objLanguage",
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

	varLanguageGetAutocompleteV2ResponseMPayload := _LanguageGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLanguageGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = LanguageGetAutocompleteV2ResponseMPayload(varLanguageGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableLanguageGetAutocompleteV2ResponseMPayload struct {
	value *LanguageGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableLanguageGetAutocompleteV2ResponseMPayload) Get() *LanguageGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableLanguageGetAutocompleteV2ResponseMPayload) Set(val *LanguageGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableLanguageGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableLanguageGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLanguageGetAutocompleteV2ResponseMPayload(val *LanguageGetAutocompleteV2ResponseMPayload) *NullableLanguageGetAutocompleteV2ResponseMPayload {
	return &NullableLanguageGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableLanguageGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLanguageGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


