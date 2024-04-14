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

// checks if the CompanyGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyGetAutocompleteV2ResponseMPayload{}

// CompanyGetAutocompleteV2ResponseMPayload Payload for POST /2/object/company/getAutocomplete
type CompanyGetAutocompleteV2ResponseMPayload struct {
	// An array of Company autocomplete element response.
	AObjCompany []CompanyAutocompleteElementResponse `json:"a_objCompany"`
}

type _CompanyGetAutocompleteV2ResponseMPayload CompanyGetAutocompleteV2ResponseMPayload

// NewCompanyGetAutocompleteV2ResponseMPayload instantiates a new CompanyGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyGetAutocompleteV2ResponseMPayload(aObjCompany []CompanyAutocompleteElementResponse) *CompanyGetAutocompleteV2ResponseMPayload {
	this := CompanyGetAutocompleteV2ResponseMPayload{}
	this.AObjCompany = aObjCompany
	return &this
}

// NewCompanyGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new CompanyGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyGetAutocompleteV2ResponseMPayloadWithDefaults() *CompanyGetAutocompleteV2ResponseMPayload {
	this := CompanyGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjCompany returns the AObjCompany field value
func (o *CompanyGetAutocompleteV2ResponseMPayload) GetAObjCompany() []CompanyAutocompleteElementResponse {
	if o == nil {
		var ret []CompanyAutocompleteElementResponse
		return ret
	}

	return o.AObjCompany
}

// GetAObjCompanyOk returns a tuple with the AObjCompany field value
// and a boolean to check if the value has been set.
func (o *CompanyGetAutocompleteV2ResponseMPayload) GetAObjCompanyOk() ([]CompanyAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCompany, true
}

// SetAObjCompany sets field value
func (o *CompanyGetAutocompleteV2ResponseMPayload) SetAObjCompany(v []CompanyAutocompleteElementResponse) {
	o.AObjCompany = v
}

func (o CompanyGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCompany"] = o.AObjCompany
	return toSerialize, nil
}

func (o *CompanyGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objCompany",
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

	varCompanyGetAutocompleteV2ResponseMPayload := _CompanyGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = CompanyGetAutocompleteV2ResponseMPayload(varCompanyGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableCompanyGetAutocompleteV2ResponseMPayload struct {
	value *CompanyGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableCompanyGetAutocompleteV2ResponseMPayload) Get() *CompanyGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableCompanyGetAutocompleteV2ResponseMPayload) Set(val *CompanyGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyGetAutocompleteV2ResponseMPayload(val *CompanyGetAutocompleteV2ResponseMPayload) *NullableCompanyGetAutocompleteV2ResponseMPayload {
	return &NullableCompanyGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableCompanyGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


