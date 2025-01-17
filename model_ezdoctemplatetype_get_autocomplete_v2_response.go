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

// checks if the EzdoctemplatetypeGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzdoctemplatetypeGetAutocompleteV2Response{}

// EzdoctemplatetypeGetAutocompleteV2Response Response for GET /2/object/ezdoctemplatetype/getAutocomplete
type EzdoctemplatetypeGetAutocompleteV2Response struct {
	CommonResponse
	MPayload EzdoctemplatetypeGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

type _EzdoctemplatetypeGetAutocompleteV2Response EzdoctemplatetypeGetAutocompleteV2Response

// NewEzdoctemplatetypeGetAutocompleteV2Response instantiates a new EzdoctemplatetypeGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzdoctemplatetypeGetAutocompleteV2Response(mPayload EzdoctemplatetypeGetAutocompleteV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *EzdoctemplatetypeGetAutocompleteV2Response {
	this := EzdoctemplatetypeGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewEzdoctemplatetypeGetAutocompleteV2ResponseWithDefaults instantiates a new EzdoctemplatetypeGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzdoctemplatetypeGetAutocompleteV2ResponseWithDefaults() *EzdoctemplatetypeGetAutocompleteV2Response {
	this := EzdoctemplatetypeGetAutocompleteV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzdoctemplatetypeGetAutocompleteV2Response) GetMPayload() EzdoctemplatetypeGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret EzdoctemplatetypeGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzdoctemplatetypeGetAutocompleteV2Response) GetMPayloadOk() (*EzdoctemplatetypeGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzdoctemplatetypeGetAutocompleteV2Response) SetMPayload(v EzdoctemplatetypeGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o EzdoctemplatetypeGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzdoctemplatetypeGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *EzdoctemplatetypeGetAutocompleteV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzdoctemplatetypeGetAutocompleteV2Response := _EzdoctemplatetypeGetAutocompleteV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzdoctemplatetypeGetAutocompleteV2Response)

	if err != nil {
		return err
	}

	*o = EzdoctemplatetypeGetAutocompleteV2Response(varEzdoctemplatetypeGetAutocompleteV2Response)

	return err
}

type NullableEzdoctemplatetypeGetAutocompleteV2Response struct {
	value *EzdoctemplatetypeGetAutocompleteV2Response
	isSet bool
}

func (v NullableEzdoctemplatetypeGetAutocompleteV2Response) Get() *EzdoctemplatetypeGetAutocompleteV2Response {
	return v.value
}

func (v *NullableEzdoctemplatetypeGetAutocompleteV2Response) Set(val *EzdoctemplatetypeGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzdoctemplatetypeGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzdoctemplatetypeGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzdoctemplatetypeGetAutocompleteV2Response(val *EzdoctemplatetypeGetAutocompleteV2Response) *NullableEzdoctemplatetypeGetAutocompleteV2Response {
	return &NullableEzdoctemplatetypeGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullableEzdoctemplatetypeGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzdoctemplatetypeGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


