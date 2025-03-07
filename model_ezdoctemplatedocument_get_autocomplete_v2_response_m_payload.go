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

// checks if the EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload{}

// EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload Payload for POST /2/object/ezdoctemplatedocument/getAutocomplete
type EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload struct {
	// An array of Ezdoctemplatedocument autocomplete element response.
	AObjEzdoctemplatedocument []EzdoctemplatedocumentAutocompleteElementResponse `json:"a_objEzdoctemplatedocument"`
}

type _EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload

// NewEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload instantiates a new EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload(aObjEzdoctemplatedocument []EzdoctemplatedocumentAutocompleteElementResponse) *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload {
	this := EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload{}
	this.AObjEzdoctemplatedocument = aObjEzdoctemplatedocument
	return &this
}

// NewEzdoctemplatedocumentGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzdoctemplatedocumentGetAutocompleteV2ResponseMPayloadWithDefaults() *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload {
	this := EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjEzdoctemplatedocument returns the AObjEzdoctemplatedocument field value
func (o *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) GetAObjEzdoctemplatedocument() []EzdoctemplatedocumentAutocompleteElementResponse {
	if o == nil {
		var ret []EzdoctemplatedocumentAutocompleteElementResponse
		return ret
	}

	return o.AObjEzdoctemplatedocument
}

// GetAObjEzdoctemplatedocumentOk returns a tuple with the AObjEzdoctemplatedocument field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) GetAObjEzdoctemplatedocumentOk() ([]EzdoctemplatedocumentAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzdoctemplatedocument, true
}

// SetAObjEzdoctemplatedocument sets field value
func (o *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) SetAObjEzdoctemplatedocument(v []EzdoctemplatedocumentAutocompleteElementResponse) {
	o.AObjEzdoctemplatedocument = v
}

func (o EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzdoctemplatedocument"] = o.AObjEzdoctemplatedocument
	return toSerialize, nil
}

func (o *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzdoctemplatedocument",
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

	varEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload := _EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload(varEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload struct {
	value *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) Get() *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) Set(val *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload(val *EzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) *NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload {
	return &NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzdoctemplatedocumentGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


