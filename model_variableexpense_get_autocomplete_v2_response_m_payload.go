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

// checks if the VariableexpenseGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableexpenseGetAutocompleteV2ResponseMPayload{}

// VariableexpenseGetAutocompleteV2ResponseMPayload Payload for POST /2/object/variableexpense/getAutocomplete
type VariableexpenseGetAutocompleteV2ResponseMPayload struct {
	// An array of Variableexpense autocomplete element response.
	AObjVariableexpense []VariableexpenseAutocompleteElementResponse `json:"a_objVariableexpense"`
}

type _VariableexpenseGetAutocompleteV2ResponseMPayload VariableexpenseGetAutocompleteV2ResponseMPayload

// NewVariableexpenseGetAutocompleteV2ResponseMPayload instantiates a new VariableexpenseGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableexpenseGetAutocompleteV2ResponseMPayload(aObjVariableexpense []VariableexpenseAutocompleteElementResponse) *VariableexpenseGetAutocompleteV2ResponseMPayload {
	this := VariableexpenseGetAutocompleteV2ResponseMPayload{}
	this.AObjVariableexpense = aObjVariableexpense
	return &this
}

// NewVariableexpenseGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new VariableexpenseGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableexpenseGetAutocompleteV2ResponseMPayloadWithDefaults() *VariableexpenseGetAutocompleteV2ResponseMPayload {
	this := VariableexpenseGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjVariableexpense returns the AObjVariableexpense field value
func (o *VariableexpenseGetAutocompleteV2ResponseMPayload) GetAObjVariableexpense() []VariableexpenseAutocompleteElementResponse {
	if o == nil {
		var ret []VariableexpenseAutocompleteElementResponse
		return ret
	}

	return o.AObjVariableexpense
}

// GetAObjVariableexpenseOk returns a tuple with the AObjVariableexpense field value
// and a boolean to check if the value has been set.
func (o *VariableexpenseGetAutocompleteV2ResponseMPayload) GetAObjVariableexpenseOk() ([]VariableexpenseAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjVariableexpense, true
}

// SetAObjVariableexpense sets field value
func (o *VariableexpenseGetAutocompleteV2ResponseMPayload) SetAObjVariableexpense(v []VariableexpenseAutocompleteElementResponse) {
	o.AObjVariableexpense = v
}

func (o VariableexpenseGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableexpenseGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objVariableexpense"] = o.AObjVariableexpense
	return toSerialize, nil
}

func (o *VariableexpenseGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objVariableexpense",
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

	varVariableexpenseGetAutocompleteV2ResponseMPayload := _VariableexpenseGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableexpenseGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = VariableexpenseGetAutocompleteV2ResponseMPayload(varVariableexpenseGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableVariableexpenseGetAutocompleteV2ResponseMPayload struct {
	value *VariableexpenseGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableVariableexpenseGetAutocompleteV2ResponseMPayload) Get() *VariableexpenseGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableVariableexpenseGetAutocompleteV2ResponseMPayload) Set(val *VariableexpenseGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableexpenseGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableexpenseGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableexpenseGetAutocompleteV2ResponseMPayload(val *VariableexpenseGetAutocompleteV2ResponseMPayload) *NullableVariableexpenseGetAutocompleteV2ResponseMPayload {
	return &NullableVariableexpenseGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableVariableexpenseGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableexpenseGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


