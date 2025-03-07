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

// checks if the TimezoneGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimezoneGetAutocompleteV2ResponseMPayload{}

// TimezoneGetAutocompleteV2ResponseMPayload Payload for POST /2/object/timezone/getAutocomplete
type TimezoneGetAutocompleteV2ResponseMPayload struct {
	// An array of Timezone autocomplete element response.
	AObjTimezone []TimezoneAutocompleteElementResponse `json:"a_objTimezone"`
}

type _TimezoneGetAutocompleteV2ResponseMPayload TimezoneGetAutocompleteV2ResponseMPayload

// NewTimezoneGetAutocompleteV2ResponseMPayload instantiates a new TimezoneGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimezoneGetAutocompleteV2ResponseMPayload(aObjTimezone []TimezoneAutocompleteElementResponse) *TimezoneGetAutocompleteV2ResponseMPayload {
	this := TimezoneGetAutocompleteV2ResponseMPayload{}
	this.AObjTimezone = aObjTimezone
	return &this
}

// NewTimezoneGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new TimezoneGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimezoneGetAutocompleteV2ResponseMPayloadWithDefaults() *TimezoneGetAutocompleteV2ResponseMPayload {
	this := TimezoneGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjTimezone returns the AObjTimezone field value
func (o *TimezoneGetAutocompleteV2ResponseMPayload) GetAObjTimezone() []TimezoneAutocompleteElementResponse {
	if o == nil {
		var ret []TimezoneAutocompleteElementResponse
		return ret
	}

	return o.AObjTimezone
}

// GetAObjTimezoneOk returns a tuple with the AObjTimezone field value
// and a boolean to check if the value has been set.
func (o *TimezoneGetAutocompleteV2ResponseMPayload) GetAObjTimezoneOk() ([]TimezoneAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjTimezone, true
}

// SetAObjTimezone sets field value
func (o *TimezoneGetAutocompleteV2ResponseMPayload) SetAObjTimezone(v []TimezoneAutocompleteElementResponse) {
	o.AObjTimezone = v
}

func (o TimezoneGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimezoneGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objTimezone"] = o.AObjTimezone
	return toSerialize, nil
}

func (o *TimezoneGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objTimezone",
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

	varTimezoneGetAutocompleteV2ResponseMPayload := _TimezoneGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTimezoneGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = TimezoneGetAutocompleteV2ResponseMPayload(varTimezoneGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableTimezoneGetAutocompleteV2ResponseMPayload struct {
	value *TimezoneGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableTimezoneGetAutocompleteV2ResponseMPayload) Get() *TimezoneGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableTimezoneGetAutocompleteV2ResponseMPayload) Set(val *TimezoneGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableTimezoneGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableTimezoneGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimezoneGetAutocompleteV2ResponseMPayload(val *TimezoneGetAutocompleteV2ResponseMPayload) *NullableTimezoneGetAutocompleteV2ResponseMPayload {
	return &NullableTimezoneGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableTimezoneGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimezoneGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


