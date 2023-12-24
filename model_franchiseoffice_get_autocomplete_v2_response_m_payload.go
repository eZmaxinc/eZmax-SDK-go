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

// checks if the FranchiseofficeGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchiseofficeGetAutocompleteV2ResponseMPayload{}

// FranchiseofficeGetAutocompleteV2ResponseMPayload Payload for POST /2/object/franchiseoffice/getAutocomplete
type FranchiseofficeGetAutocompleteV2ResponseMPayload struct {
	// An array of Franchiseoffice autocomplete element response.
	AObjFranchiseoffice []FranchiseofficeAutocompleteElementResponse `json:"a_objFranchiseoffice"`
}

type _FranchiseofficeGetAutocompleteV2ResponseMPayload FranchiseofficeGetAutocompleteV2ResponseMPayload

// NewFranchiseofficeGetAutocompleteV2ResponseMPayload instantiates a new FranchiseofficeGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchiseofficeGetAutocompleteV2ResponseMPayload(aObjFranchiseoffice []FranchiseofficeAutocompleteElementResponse) *FranchiseofficeGetAutocompleteV2ResponseMPayload {
	this := FranchiseofficeGetAutocompleteV2ResponseMPayload{}
	this.AObjFranchiseoffice = aObjFranchiseoffice
	return &this
}

// NewFranchiseofficeGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new FranchiseofficeGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchiseofficeGetAutocompleteV2ResponseMPayloadWithDefaults() *FranchiseofficeGetAutocompleteV2ResponseMPayload {
	this := FranchiseofficeGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjFranchiseoffice returns the AObjFranchiseoffice field value
func (o *FranchiseofficeGetAutocompleteV2ResponseMPayload) GetAObjFranchiseoffice() []FranchiseofficeAutocompleteElementResponse {
	if o == nil {
		var ret []FranchiseofficeAutocompleteElementResponse
		return ret
	}

	return o.AObjFranchiseoffice
}

// GetAObjFranchiseofficeOk returns a tuple with the AObjFranchiseoffice field value
// and a boolean to check if the value has been set.
func (o *FranchiseofficeGetAutocompleteV2ResponseMPayload) GetAObjFranchiseofficeOk() ([]FranchiseofficeAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFranchiseoffice, true
}

// SetAObjFranchiseoffice sets field value
func (o *FranchiseofficeGetAutocompleteV2ResponseMPayload) SetAObjFranchiseoffice(v []FranchiseofficeAutocompleteElementResponse) {
	o.AObjFranchiseoffice = v
}

func (o FranchiseofficeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FranchiseofficeGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objFranchiseoffice"] = o.AObjFranchiseoffice
	return toSerialize, nil
}

func (o *FranchiseofficeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objFranchiseoffice",
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

	varFranchiseofficeGetAutocompleteV2ResponseMPayload := _FranchiseofficeGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFranchiseofficeGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = FranchiseofficeGetAutocompleteV2ResponseMPayload(varFranchiseofficeGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableFranchiseofficeGetAutocompleteV2ResponseMPayload struct {
	value *FranchiseofficeGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) Get() *FranchiseofficeGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) Set(val *FranchiseofficeGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchiseofficeGetAutocompleteV2ResponseMPayload(val *FranchiseofficeGetAutocompleteV2ResponseMPayload) *NullableFranchiseofficeGetAutocompleteV2ResponseMPayload {
	return &NullableFranchiseofficeGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchiseofficeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


