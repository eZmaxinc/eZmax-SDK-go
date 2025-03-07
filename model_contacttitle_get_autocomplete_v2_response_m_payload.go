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

// checks if the ContacttitleGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContacttitleGetAutocompleteV2ResponseMPayload{}

// ContacttitleGetAutocompleteV2ResponseMPayload Payload for POST /2/object/contacttitle/getAutocomplete
type ContacttitleGetAutocompleteV2ResponseMPayload struct {
	// An array of Contacttitle autocomplete element response.
	AObjContacttitle []ContacttitleAutocompleteElementResponse `json:"a_objContacttitle"`
}

type _ContacttitleGetAutocompleteV2ResponseMPayload ContacttitleGetAutocompleteV2ResponseMPayload

// NewContacttitleGetAutocompleteV2ResponseMPayload instantiates a new ContacttitleGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContacttitleGetAutocompleteV2ResponseMPayload(aObjContacttitle []ContacttitleAutocompleteElementResponse) *ContacttitleGetAutocompleteV2ResponseMPayload {
	this := ContacttitleGetAutocompleteV2ResponseMPayload{}
	this.AObjContacttitle = aObjContacttitle
	return &this
}

// NewContacttitleGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new ContacttitleGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContacttitleGetAutocompleteV2ResponseMPayloadWithDefaults() *ContacttitleGetAutocompleteV2ResponseMPayload {
	this := ContacttitleGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjContacttitle returns the AObjContacttitle field value
func (o *ContacttitleGetAutocompleteV2ResponseMPayload) GetAObjContacttitle() []ContacttitleAutocompleteElementResponse {
	if o == nil {
		var ret []ContacttitleAutocompleteElementResponse
		return ret
	}

	return o.AObjContacttitle
}

// GetAObjContacttitleOk returns a tuple with the AObjContacttitle field value
// and a boolean to check if the value has been set.
func (o *ContacttitleGetAutocompleteV2ResponseMPayload) GetAObjContacttitleOk() ([]ContacttitleAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjContacttitle, true
}

// SetAObjContacttitle sets field value
func (o *ContacttitleGetAutocompleteV2ResponseMPayload) SetAObjContacttitle(v []ContacttitleAutocompleteElementResponse) {
	o.AObjContacttitle = v
}

func (o ContacttitleGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContacttitleGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objContacttitle"] = o.AObjContacttitle
	return toSerialize, nil
}

func (o *ContacttitleGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objContacttitle",
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

	varContacttitleGetAutocompleteV2ResponseMPayload := _ContacttitleGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContacttitleGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = ContacttitleGetAutocompleteV2ResponseMPayload(varContacttitleGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableContacttitleGetAutocompleteV2ResponseMPayload struct {
	value *ContacttitleGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableContacttitleGetAutocompleteV2ResponseMPayload) Get() *ContacttitleGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableContacttitleGetAutocompleteV2ResponseMPayload) Set(val *ContacttitleGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableContacttitleGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableContacttitleGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContacttitleGetAutocompleteV2ResponseMPayload(val *ContacttitleGetAutocompleteV2ResponseMPayload) *NullableContacttitleGetAutocompleteV2ResponseMPayload {
	return &NullableContacttitleGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableContacttitleGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContacttitleGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


