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

// checks if the GlaccountGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlaccountGetAutocompleteV2ResponseMPayload{}

// GlaccountGetAutocompleteV2ResponseMPayload Payload for POST /2/object/glaccount/getAutocomplete
type GlaccountGetAutocompleteV2ResponseMPayload struct {
	// An array of Glaccount autocomplete element response.
	AObjGlaccount []GlaccountAutocompleteElementResponse `json:"a_objGlaccount"`
}

type _GlaccountGetAutocompleteV2ResponseMPayload GlaccountGetAutocompleteV2ResponseMPayload

// NewGlaccountGetAutocompleteV2ResponseMPayload instantiates a new GlaccountGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlaccountGetAutocompleteV2ResponseMPayload(aObjGlaccount []GlaccountAutocompleteElementResponse) *GlaccountGetAutocompleteV2ResponseMPayload {
	this := GlaccountGetAutocompleteV2ResponseMPayload{}
	this.AObjGlaccount = aObjGlaccount
	return &this
}

// NewGlaccountGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new GlaccountGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlaccountGetAutocompleteV2ResponseMPayloadWithDefaults() *GlaccountGetAutocompleteV2ResponseMPayload {
	this := GlaccountGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjGlaccount returns the AObjGlaccount field value
func (o *GlaccountGetAutocompleteV2ResponseMPayload) GetAObjGlaccount() []GlaccountAutocompleteElementResponse {
	if o == nil {
		var ret []GlaccountAutocompleteElementResponse
		return ret
	}

	return o.AObjGlaccount
}

// GetAObjGlaccountOk returns a tuple with the AObjGlaccount field value
// and a boolean to check if the value has been set.
func (o *GlaccountGetAutocompleteV2ResponseMPayload) GetAObjGlaccountOk() ([]GlaccountAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjGlaccount, true
}

// SetAObjGlaccount sets field value
func (o *GlaccountGetAutocompleteV2ResponseMPayload) SetAObjGlaccount(v []GlaccountAutocompleteElementResponse) {
	o.AObjGlaccount = v
}

func (o GlaccountGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlaccountGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objGlaccount"] = o.AObjGlaccount
	return toSerialize, nil
}

func (o *GlaccountGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objGlaccount",
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

	varGlaccountGetAutocompleteV2ResponseMPayload := _GlaccountGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGlaccountGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = GlaccountGetAutocompleteV2ResponseMPayload(varGlaccountGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableGlaccountGetAutocompleteV2ResponseMPayload struct {
	value *GlaccountGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableGlaccountGetAutocompleteV2ResponseMPayload) Get() *GlaccountGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableGlaccountGetAutocompleteV2ResponseMPayload) Set(val *GlaccountGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableGlaccountGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableGlaccountGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlaccountGetAutocompleteV2ResponseMPayload(val *GlaccountGetAutocompleteV2ResponseMPayload) *NullableGlaccountGetAutocompleteV2ResponseMPayload {
	return &NullableGlaccountGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableGlaccountGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlaccountGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


