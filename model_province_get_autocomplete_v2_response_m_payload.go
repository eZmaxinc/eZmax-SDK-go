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

// checks if the ProvinceGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvinceGetAutocompleteV2ResponseMPayload{}

// ProvinceGetAutocompleteV2ResponseMPayload Payload for POST /2/object/province/getAutocomplete
type ProvinceGetAutocompleteV2ResponseMPayload struct {
	// An array of Province autocomplete element response.
	AObjProvince []ProvinceAutocompleteElementResponse `json:"a_objProvince"`
}

type _ProvinceGetAutocompleteV2ResponseMPayload ProvinceGetAutocompleteV2ResponseMPayload

// NewProvinceGetAutocompleteV2ResponseMPayload instantiates a new ProvinceGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvinceGetAutocompleteV2ResponseMPayload(aObjProvince []ProvinceAutocompleteElementResponse) *ProvinceGetAutocompleteV2ResponseMPayload {
	this := ProvinceGetAutocompleteV2ResponseMPayload{}
	this.AObjProvince = aObjProvince
	return &this
}

// NewProvinceGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new ProvinceGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvinceGetAutocompleteV2ResponseMPayloadWithDefaults() *ProvinceGetAutocompleteV2ResponseMPayload {
	this := ProvinceGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjProvince returns the AObjProvince field value
func (o *ProvinceGetAutocompleteV2ResponseMPayload) GetAObjProvince() []ProvinceAutocompleteElementResponse {
	if o == nil {
		var ret []ProvinceAutocompleteElementResponse
		return ret
	}

	return o.AObjProvince
}

// GetAObjProvinceOk returns a tuple with the AObjProvince field value
// and a boolean to check if the value has been set.
func (o *ProvinceGetAutocompleteV2ResponseMPayload) GetAObjProvinceOk() ([]ProvinceAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjProvince, true
}

// SetAObjProvince sets field value
func (o *ProvinceGetAutocompleteV2ResponseMPayload) SetAObjProvince(v []ProvinceAutocompleteElementResponse) {
	o.AObjProvince = v
}

func (o ProvinceGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvinceGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objProvince"] = o.AObjProvince
	return toSerialize, nil
}

func (o *ProvinceGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objProvince",
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

	varProvinceGetAutocompleteV2ResponseMPayload := _ProvinceGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvinceGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ProvinceGetAutocompleteV2ResponseMPayload(varProvinceGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableProvinceGetAutocompleteV2ResponseMPayload struct {
	value *ProvinceGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableProvinceGetAutocompleteV2ResponseMPayload) Get() *ProvinceGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableProvinceGetAutocompleteV2ResponseMPayload) Set(val *ProvinceGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableProvinceGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableProvinceGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvinceGetAutocompleteV2ResponseMPayload(val *ProvinceGetAutocompleteV2ResponseMPayload) *NullableProvinceGetAutocompleteV2ResponseMPayload {
	return &NullableProvinceGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableProvinceGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvinceGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


