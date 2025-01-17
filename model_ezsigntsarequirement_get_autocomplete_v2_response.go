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

// checks if the EzsigntsarequirementGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntsarequirementGetAutocompleteV2Response{}

// EzsigntsarequirementGetAutocompleteV2Response Response for GET /2/object/ezsigntsarequirement/getAutocomplete
type EzsigntsarequirementGetAutocompleteV2Response struct {
	CommonResponse
	MPayload EzsigntsarequirementGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

type _EzsigntsarequirementGetAutocompleteV2Response EzsigntsarequirementGetAutocompleteV2Response

// NewEzsigntsarequirementGetAutocompleteV2Response instantiates a new EzsigntsarequirementGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntsarequirementGetAutocompleteV2Response(mPayload EzsigntsarequirementGetAutocompleteV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzsigntsarequirementGetAutocompleteV2Response {
	this := EzsigntsarequirementGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzsigntsarequirementGetAutocompleteV2ResponseWithDefaults instantiates a new EzsigntsarequirementGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntsarequirementGetAutocompleteV2ResponseWithDefaults() *EzsigntsarequirementGetAutocompleteV2Response {
	this := EzsigntsarequirementGetAutocompleteV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigntsarequirementGetAutocompleteV2Response) GetMPayload() EzsigntsarequirementGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret EzsigntsarequirementGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigntsarequirementGetAutocompleteV2Response) GetMPayloadOk() (*EzsigntsarequirementGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigntsarequirementGetAutocompleteV2Response) SetMPayload(v EzsigntsarequirementGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzsigntsarequirementGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntsarequirementGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzsigntsarequirementGetAutocompleteV2Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
		"objDebugPayload",
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

	varEzsigntsarequirementGetAutocompleteV2Response := _EzsigntsarequirementGetAutocompleteV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntsarequirementGetAutocompleteV2Response)

	if err != nil {
		return err
	}

	*o = EzsigntsarequirementGetAutocompleteV2Response(varEzsigntsarequirementGetAutocompleteV2Response)

	return err
}

type NullableEzsigntsarequirementGetAutocompleteV2Response struct {
	value *EzsigntsarequirementGetAutocompleteV2Response
	isSet bool
}

func (v NullableEzsigntsarequirementGetAutocompleteV2Response) Get() *EzsigntsarequirementGetAutocompleteV2Response {
	return v.value
}

func (v *NullableEzsigntsarequirementGetAutocompleteV2Response) Set(val *EzsigntsarequirementGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntsarequirementGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntsarequirementGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntsarequirementGetAutocompleteV2Response(val *EzsigntsarequirementGetAutocompleteV2Response) *NullableEzsigntsarequirementGetAutocompleteV2Response {
	return &NullableEzsigntsarequirementGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullableEzsigntsarequirementGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntsarequirementGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


