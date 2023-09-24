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

// checks if the TaxassignmentGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxassignmentGetAutocompleteV2ResponseMPayload{}

// TaxassignmentGetAutocompleteV2ResponseMPayload Payload for POST /2/object/taxassignment/getAutocomplete
type TaxassignmentGetAutocompleteV2ResponseMPayload struct {
	// An array of Taxassignment autocomplete element response.
	AObjTaxassignment []TaxassignmentAutocompleteElementResponse `json:"a_objTaxassignment"`
}

// NewTaxassignmentGetAutocompleteV2ResponseMPayload instantiates a new TaxassignmentGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxassignmentGetAutocompleteV2ResponseMPayload(aObjTaxassignment []TaxassignmentAutocompleteElementResponse) *TaxassignmentGetAutocompleteV2ResponseMPayload {
	this := TaxassignmentGetAutocompleteV2ResponseMPayload{}
	this.AObjTaxassignment = aObjTaxassignment
	return &this
}

// NewTaxassignmentGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new TaxassignmentGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxassignmentGetAutocompleteV2ResponseMPayloadWithDefaults() *TaxassignmentGetAutocompleteV2ResponseMPayload {
	this := TaxassignmentGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjTaxassignment returns the AObjTaxassignment field value
func (o *TaxassignmentGetAutocompleteV2ResponseMPayload) GetAObjTaxassignment() []TaxassignmentAutocompleteElementResponse {
	if o == nil {
		var ret []TaxassignmentAutocompleteElementResponse
		return ret
	}

	return o.AObjTaxassignment
}

// GetAObjTaxassignmentOk returns a tuple with the AObjTaxassignment field value
// and a boolean to check if the value has been set.
func (o *TaxassignmentGetAutocompleteV2ResponseMPayload) GetAObjTaxassignmentOk() ([]TaxassignmentAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjTaxassignment, true
}

// SetAObjTaxassignment sets field value
func (o *TaxassignmentGetAutocompleteV2ResponseMPayload) SetAObjTaxassignment(v []TaxassignmentAutocompleteElementResponse) {
	o.AObjTaxassignment = v
}

func (o TaxassignmentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxassignmentGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objTaxassignment"] = o.AObjTaxassignment
	return toSerialize, nil
}

type NullableTaxassignmentGetAutocompleteV2ResponseMPayload struct {
	value *TaxassignmentGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableTaxassignmentGetAutocompleteV2ResponseMPayload) Get() *TaxassignmentGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableTaxassignmentGetAutocompleteV2ResponseMPayload) Set(val *TaxassignmentGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxassignmentGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxassignmentGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxassignmentGetAutocompleteV2ResponseMPayload(val *TaxassignmentGetAutocompleteV2ResponseMPayload) *NullableTaxassignmentGetAutocompleteV2ResponseMPayload {
	return &NullableTaxassignmentGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableTaxassignmentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxassignmentGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

