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

// checks if the FranchisebrokerGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchisebrokerGetAutocompleteV2ResponseMPayload{}

// FranchisebrokerGetAutocompleteV2ResponseMPayload Payload for POST /2/object/franchisebroker/getAutocomplete
type FranchisebrokerGetAutocompleteV2ResponseMPayload struct {
	// An array of Franchisebroker autocomplete element response.
	AObjFranchisebroker []FranchisebrokerAutocompleteElementResponse `json:"a_objFranchisebroker"`
}

type _FranchisebrokerGetAutocompleteV2ResponseMPayload FranchisebrokerGetAutocompleteV2ResponseMPayload

// NewFranchisebrokerGetAutocompleteV2ResponseMPayload instantiates a new FranchisebrokerGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisebrokerGetAutocompleteV2ResponseMPayload(aObjFranchisebroker []FranchisebrokerAutocompleteElementResponse) *FranchisebrokerGetAutocompleteV2ResponseMPayload {
	this := FranchisebrokerGetAutocompleteV2ResponseMPayload{}
	this.AObjFranchisebroker = aObjFranchisebroker
	return &this
}

// NewFranchisebrokerGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new FranchisebrokerGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisebrokerGetAutocompleteV2ResponseMPayloadWithDefaults() *FranchisebrokerGetAutocompleteV2ResponseMPayload {
	this := FranchisebrokerGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjFranchisebroker returns the AObjFranchisebroker field value
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) GetAObjFranchisebroker() []FranchisebrokerAutocompleteElementResponse {
	if o == nil {
		var ret []FranchisebrokerAutocompleteElementResponse
		return ret
	}

	return o.AObjFranchisebroker
}

// GetAObjFranchisebrokerOk returns a tuple with the AObjFranchisebroker field value
// and a boolean to check if the value has been set.
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) GetAObjFranchisebrokerOk() ([]FranchisebrokerAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFranchisebroker, true
}

// SetAObjFranchisebroker sets field value
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) SetAObjFranchisebroker(v []FranchisebrokerAutocompleteElementResponse) {
	o.AObjFranchisebroker = v
}

func (o FranchisebrokerGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchisebrokerGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objFranchisebroker"] = o.AObjFranchisebroker
	return toSerialize, nil
}

func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objFranchisebroker",
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

	varFranchisebrokerGetAutocompleteV2ResponseMPayload := _FranchisebrokerGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFranchisebrokerGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = FranchisebrokerGetAutocompleteV2ResponseMPayload(varFranchisebrokerGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableFranchisebrokerGetAutocompleteV2ResponseMPayload struct {
	value *FranchisebrokerGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) Get() *FranchisebrokerGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) Set(val *FranchisebrokerGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisebrokerGetAutocompleteV2ResponseMPayload(val *FranchisebrokerGetAutocompleteV2ResponseMPayload) *NullableFranchisebrokerGetAutocompleteV2ResponseMPayload {
	return &NullableFranchisebrokerGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisebrokerGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


