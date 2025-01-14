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

// checks if the SupplyGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupplyGetAutocompleteV2ResponseMPayload{}

// SupplyGetAutocompleteV2ResponseMPayload Payload for POST /2/object/supply/getAutocomplete
type SupplyGetAutocompleteV2ResponseMPayload struct {
	// An array of Supply autocomplete element response.
	AObjSupply []SupplyAutocompleteElementResponse `json:"a_objSupply"`
}

type _SupplyGetAutocompleteV2ResponseMPayload SupplyGetAutocompleteV2ResponseMPayload

// NewSupplyGetAutocompleteV2ResponseMPayload instantiates a new SupplyGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupplyGetAutocompleteV2ResponseMPayload(aObjSupply []SupplyAutocompleteElementResponse) *SupplyGetAutocompleteV2ResponseMPayload {
	this := SupplyGetAutocompleteV2ResponseMPayload{}
	this.AObjSupply = aObjSupply
	return &this
}

// NewSupplyGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new SupplyGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupplyGetAutocompleteV2ResponseMPayloadWithDefaults() *SupplyGetAutocompleteV2ResponseMPayload {
	this := SupplyGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjSupply returns the AObjSupply field value
func (o *SupplyGetAutocompleteV2ResponseMPayload) GetAObjSupply() []SupplyAutocompleteElementResponse {
	if o == nil {
		var ret []SupplyAutocompleteElementResponse
		return ret
	}

	return o.AObjSupply
}

// GetAObjSupplyOk returns a tuple with the AObjSupply field value
// and a boolean to check if the value has been set.
func (o *SupplyGetAutocompleteV2ResponseMPayload) GetAObjSupplyOk() ([]SupplyAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjSupply, true
}

// SetAObjSupply sets field value
func (o *SupplyGetAutocompleteV2ResponseMPayload) SetAObjSupply(v []SupplyAutocompleteElementResponse) {
	o.AObjSupply = v
}

func (o SupplyGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupplyGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objSupply"] = o.AObjSupply
	return toSerialize, nil
}

func (o *SupplyGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objSupply",
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

	varSupplyGetAutocompleteV2ResponseMPayload := _SupplyGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSupplyGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = SupplyGetAutocompleteV2ResponseMPayload(varSupplyGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableSupplyGetAutocompleteV2ResponseMPayload struct {
	value *SupplyGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableSupplyGetAutocompleteV2ResponseMPayload) Get() *SupplyGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableSupplyGetAutocompleteV2ResponseMPayload) Set(val *SupplyGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSupplyGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSupplyGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupplyGetAutocompleteV2ResponseMPayload(val *SupplyGetAutocompleteV2ResponseMPayload) *NullableSupplyGetAutocompleteV2ResponseMPayload {
	return &NullableSupplyGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableSupplyGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupplyGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


