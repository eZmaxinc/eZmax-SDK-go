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
)

// checks if the FranchisebrokerGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FranchisebrokerGetAutocompleteV2ResponseMPayload{}

// FranchisebrokerGetAutocompleteV2ResponseMPayload Payload for POST /2/object/franchisebroker/getAutocomplete
type FranchisebrokerGetAutocompleteV2ResponseMPayload struct {
	// An array of Franchisebroker autocomplete element response.
	AObjFranchisebroker []FranchisebrokerAutocompleteElementResponse `json:"a_objFranchisebroker,omitempty"`
}

// NewFranchisebrokerGetAutocompleteV2ResponseMPayload instantiates a new FranchisebrokerGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisebrokerGetAutocompleteV2ResponseMPayload() *FranchisebrokerGetAutocompleteV2ResponseMPayload {
	this := FranchisebrokerGetAutocompleteV2ResponseMPayload{}
	return &this
}

// NewFranchisebrokerGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new FranchisebrokerGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisebrokerGetAutocompleteV2ResponseMPayloadWithDefaults() *FranchisebrokerGetAutocompleteV2ResponseMPayload {
	this := FranchisebrokerGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjFranchisebroker returns the AObjFranchisebroker field value if set, zero value otherwise.
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) GetAObjFranchisebroker() []FranchisebrokerAutocompleteElementResponse {
	if o == nil || IsNil(o.AObjFranchisebroker) {
		var ret []FranchisebrokerAutocompleteElementResponse
		return ret
	}
	return o.AObjFranchisebroker
}

// GetAObjFranchisebrokerOk returns a tuple with the AObjFranchisebroker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) GetAObjFranchisebrokerOk() ([]FranchisebrokerAutocompleteElementResponse, bool) {
	if o == nil || IsNil(o.AObjFranchisebroker) {
		return nil, false
	}
	return o.AObjFranchisebroker, true
}

// HasAObjFranchisebroker returns a boolean if a field has been set.
func (o *FranchisebrokerGetAutocompleteV2ResponseMPayload) HasAObjFranchisebroker() bool {
	if o != nil && !IsNil(o.AObjFranchisebroker) {
		return true
	}

	return false
}

// SetAObjFranchisebroker gets a reference to the given []FranchisebrokerAutocompleteElementResponse and assigns it to the AObjFranchisebroker field.
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
	if !IsNil(o.AObjFranchisebroker) {
		toSerialize["a_objFranchisebroker"] = o.AObjFranchisebroker
	}
	return toSerialize, nil
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


