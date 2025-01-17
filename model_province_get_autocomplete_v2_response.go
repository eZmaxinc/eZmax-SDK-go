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

// checks if the ProvinceGetAutocompleteV2Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvinceGetAutocompleteV2Response{}

// ProvinceGetAutocompleteV2Response Response for GET /2/object/province/getAutocomplete
type ProvinceGetAutocompleteV2Response struct {
	CommonResponse
	MPayload ProvinceGetAutocompleteV2ResponseMPayload `json:"mPayload"`
}

type _ProvinceGetAutocompleteV2Response ProvinceGetAutocompleteV2Response

// NewProvinceGetAutocompleteV2Response instantiates a new ProvinceGetAutocompleteV2Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvinceGetAutocompleteV2Response(mPayload ProvinceGetAutocompleteV2ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *ProvinceGetAutocompleteV2Response {
	this := ProvinceGetAutocompleteV2Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewProvinceGetAutocompleteV2ResponseWithDefaults instantiates a new ProvinceGetAutocompleteV2Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvinceGetAutocompleteV2ResponseWithDefaults() *ProvinceGetAutocompleteV2Response {
	this := ProvinceGetAutocompleteV2Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *ProvinceGetAutocompleteV2Response) GetMPayload() ProvinceGetAutocompleteV2ResponseMPayload {
	if o == nil {
		var ret ProvinceGetAutocompleteV2ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *ProvinceGetAutocompleteV2Response) GetMPayloadOk() (*ProvinceGetAutocompleteV2ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *ProvinceGetAutocompleteV2Response) SetMPayload(v ProvinceGetAutocompleteV2ResponseMPayload) {
	o.MPayload = v
}

func (o ProvinceGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvinceGetAutocompleteV2Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *ProvinceGetAutocompleteV2Response) UnmarshalJSON(data []byte) (err error) {
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

	varProvinceGetAutocompleteV2Response := _ProvinceGetAutocompleteV2Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProvinceGetAutocompleteV2Response)

	if err != nil {
		return err
	}

	*o = ProvinceGetAutocompleteV2Response(varProvinceGetAutocompleteV2Response)

	return err
}

type NullableProvinceGetAutocompleteV2Response struct {
	value *ProvinceGetAutocompleteV2Response
	isSet bool
}

func (v NullableProvinceGetAutocompleteV2Response) Get() *ProvinceGetAutocompleteV2Response {
	return v.value
}

func (v *NullableProvinceGetAutocompleteV2Response) Set(val *ProvinceGetAutocompleteV2Response) {
	v.value = val
	v.isSet = true
}

func (v NullableProvinceGetAutocompleteV2Response) IsSet() bool {
	return v.isSet
}

func (v *NullableProvinceGetAutocompleteV2Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvinceGetAutocompleteV2Response(val *ProvinceGetAutocompleteV2Response) *NullableProvinceGetAutocompleteV2Response {
	return &NullableProvinceGetAutocompleteV2Response{value: val, isSet: true}
}

func (v NullableProvinceGetAutocompleteV2Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvinceGetAutocompleteV2Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


