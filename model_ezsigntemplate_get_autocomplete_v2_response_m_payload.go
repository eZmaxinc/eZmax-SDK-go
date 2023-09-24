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

// checks if the EzsigntemplateGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateGetAutocompleteV2ResponseMPayload{}

// EzsigntemplateGetAutocompleteV2ResponseMPayload Payload for POST /2/object/ezsigntemplate/getAutocomplete
type EzsigntemplateGetAutocompleteV2ResponseMPayload struct {
	// An array of Ezsigntemplate autocomplete element response.
	AObjEzsigntemplate []EzsigntemplateAutocompleteElementResponse `json:"a_objEzsigntemplate"`
}

// NewEzsigntemplateGetAutocompleteV2ResponseMPayload instantiates a new EzsigntemplateGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateGetAutocompleteV2ResponseMPayload(aObjEzsigntemplate []EzsigntemplateAutocompleteElementResponse) *EzsigntemplateGetAutocompleteV2ResponseMPayload {
	this := EzsigntemplateGetAutocompleteV2ResponseMPayload{}
	this.AObjEzsigntemplate = aObjEzsigntemplate
	return &this
}

// NewEzsigntemplateGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new EzsigntemplateGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateGetAutocompleteV2ResponseMPayloadWithDefaults() *EzsigntemplateGetAutocompleteV2ResponseMPayload {
	this := EzsigntemplateGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjEzsigntemplate returns the AObjEzsigntemplate field value
func (o *EzsigntemplateGetAutocompleteV2ResponseMPayload) GetAObjEzsigntemplate() []EzsigntemplateAutocompleteElementResponse {
	if o == nil {
		var ret []EzsigntemplateAutocompleteElementResponse
		return ret
	}

	return o.AObjEzsigntemplate
}

// GetAObjEzsigntemplateOk returns a tuple with the AObjEzsigntemplate field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateGetAutocompleteV2ResponseMPayload) GetAObjEzsigntemplateOk() ([]EzsigntemplateAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsigntemplate, true
}

// SetAObjEzsigntemplate sets field value
func (o *EzsigntemplateGetAutocompleteV2ResponseMPayload) SetAObjEzsigntemplate(v []EzsigntemplateAutocompleteElementResponse) {
	o.AObjEzsigntemplate = v
}

func (o EzsigntemplateGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsigntemplate"] = o.AObjEzsigntemplate
	return toSerialize, nil
}

type NullableEzsigntemplateGetAutocompleteV2ResponseMPayload struct {
	value *EzsigntemplateGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) Get() *EzsigntemplateGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) Set(val *EzsigntemplateGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateGetAutocompleteV2ResponseMPayload(val *EzsigntemplateGetAutocompleteV2ResponseMPayload) *NullableEzsigntemplateGetAutocompleteV2ResponseMPayload {
	return &NullableEzsigntemplateGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


