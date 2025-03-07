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

// checks if the UserlogintypeGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserlogintypeGetAutocompleteV2ResponseMPayload{}

// UserlogintypeGetAutocompleteV2ResponseMPayload Payload for POST /2/object/userlogintype/getAutocomplete
type UserlogintypeGetAutocompleteV2ResponseMPayload struct {
	// An array of Userlogintype autocomplete element response.
	AObjUserlogintype []UserlogintypeAutocompleteElementResponse `json:"a_objUserlogintype"`
}

type _UserlogintypeGetAutocompleteV2ResponseMPayload UserlogintypeGetAutocompleteV2ResponseMPayload

// NewUserlogintypeGetAutocompleteV2ResponseMPayload instantiates a new UserlogintypeGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserlogintypeGetAutocompleteV2ResponseMPayload(aObjUserlogintype []UserlogintypeAutocompleteElementResponse) *UserlogintypeGetAutocompleteV2ResponseMPayload {
	this := UserlogintypeGetAutocompleteV2ResponseMPayload{}
	this.AObjUserlogintype = aObjUserlogintype
	return &this
}

// NewUserlogintypeGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new UserlogintypeGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserlogintypeGetAutocompleteV2ResponseMPayloadWithDefaults() *UserlogintypeGetAutocompleteV2ResponseMPayload {
	this := UserlogintypeGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjUserlogintype returns the AObjUserlogintype field value
func (o *UserlogintypeGetAutocompleteV2ResponseMPayload) GetAObjUserlogintype() []UserlogintypeAutocompleteElementResponse {
	if o == nil {
		var ret []UserlogintypeAutocompleteElementResponse
		return ret
	}

	return o.AObjUserlogintype
}

// GetAObjUserlogintypeOk returns a tuple with the AObjUserlogintype field value
// and a boolean to check if the value has been set.
func (o *UserlogintypeGetAutocompleteV2ResponseMPayload) GetAObjUserlogintypeOk() ([]UserlogintypeAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUserlogintype, true
}

// SetAObjUserlogintype sets field value
func (o *UserlogintypeGetAutocompleteV2ResponseMPayload) SetAObjUserlogintype(v []UserlogintypeAutocompleteElementResponse) {
	o.AObjUserlogintype = v
}

func (o UserlogintypeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserlogintypeGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objUserlogintype"] = o.AObjUserlogintype
	return toSerialize, nil
}

func (o *UserlogintypeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objUserlogintype",
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

	varUserlogintypeGetAutocompleteV2ResponseMPayload := _UserlogintypeGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserlogintypeGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = UserlogintypeGetAutocompleteV2ResponseMPayload(varUserlogintypeGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableUserlogintypeGetAutocompleteV2ResponseMPayload struct {
	value *UserlogintypeGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableUserlogintypeGetAutocompleteV2ResponseMPayload) Get() *UserlogintypeGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableUserlogintypeGetAutocompleteV2ResponseMPayload) Set(val *UserlogintypeGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUserlogintypeGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUserlogintypeGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserlogintypeGetAutocompleteV2ResponseMPayload(val *UserlogintypeGetAutocompleteV2ResponseMPayload) *NullableUserlogintypeGetAutocompleteV2ResponseMPayload {
	return &NullableUserlogintypeGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableUserlogintypeGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserlogintypeGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


