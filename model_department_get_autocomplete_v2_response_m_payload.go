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

// checks if the DepartmentGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DepartmentGetAutocompleteV2ResponseMPayload{}

// DepartmentGetAutocompleteV2ResponseMPayload Payload for POST /2/object/department/getAutocomplete
type DepartmentGetAutocompleteV2ResponseMPayload struct {
	// An array of Department autocomplete element response.
	AObjDepartment []DepartmentAutocompleteElementResponse `json:"a_objDepartment"`
}

// NewDepartmentGetAutocompleteV2ResponseMPayload instantiates a new DepartmentGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepartmentGetAutocompleteV2ResponseMPayload(aObjDepartment []DepartmentAutocompleteElementResponse) *DepartmentGetAutocompleteV2ResponseMPayload {
	this := DepartmentGetAutocompleteV2ResponseMPayload{}
	this.AObjDepartment = aObjDepartment
	return &this
}

// NewDepartmentGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new DepartmentGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepartmentGetAutocompleteV2ResponseMPayloadWithDefaults() *DepartmentGetAutocompleteV2ResponseMPayload {
	this := DepartmentGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjDepartment returns the AObjDepartment field value
func (o *DepartmentGetAutocompleteV2ResponseMPayload) GetAObjDepartment() []DepartmentAutocompleteElementResponse {
	if o == nil {
		var ret []DepartmentAutocompleteElementResponse
		return ret
	}

	return o.AObjDepartment
}

// GetAObjDepartmentOk returns a tuple with the AObjDepartment field value
// and a boolean to check if the value has been set.
func (o *DepartmentGetAutocompleteV2ResponseMPayload) GetAObjDepartmentOk() ([]DepartmentAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjDepartment, true
}

// SetAObjDepartment sets field value
func (o *DepartmentGetAutocompleteV2ResponseMPayload) SetAObjDepartment(v []DepartmentAutocompleteElementResponse) {
	o.AObjDepartment = v
}

func (o DepartmentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DepartmentGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objDepartment"] = o.AObjDepartment
	return toSerialize, nil
}

type NullableDepartmentGetAutocompleteV2ResponseMPayload struct {
	value *DepartmentGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableDepartmentGetAutocompleteV2ResponseMPayload) Get() *DepartmentGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableDepartmentGetAutocompleteV2ResponseMPayload) Set(val *DepartmentGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableDepartmentGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableDepartmentGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepartmentGetAutocompleteV2ResponseMPayload(val *DepartmentGetAutocompleteV2ResponseMPayload) *NullableDepartmentGetAutocompleteV2ResponseMPayload {
	return &NullableDepartmentGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableDepartmentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepartmentGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

