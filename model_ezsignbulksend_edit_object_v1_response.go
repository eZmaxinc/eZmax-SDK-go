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

// checks if the EzsignbulksendEditObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendEditObjectV1Response{}

// EzsignbulksendEditObjectV1Response Response for PUT /1/object/ezsignbulksend/{pkiEzsignbulksendID}
type EzsignbulksendEditObjectV1Response struct {
	CommonResponse
}

type _EzsignbulksendEditObjectV1Response EzsignbulksendEditObjectV1Response

// NewEzsignbulksendEditObjectV1Response instantiates a new EzsignbulksendEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignbulksendEditObjectV1Response {
	this := EzsignbulksendEditObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignbulksendEditObjectV1ResponseWithDefaults instantiates a new EzsignbulksendEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendEditObjectV1ResponseWithDefaults() *EzsignbulksendEditObjectV1Response {
	this := EzsignbulksendEditObjectV1Response{}
	return &this
}

func (o EzsignbulksendEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendEditObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignbulksendEditObjectV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsignbulksendEditObjectV1Response := _EzsignbulksendEditObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendEditObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignbulksendEditObjectV1Response(varEzsignbulksendEditObjectV1Response)

	return err
}

type NullableEzsignbulksendEditObjectV1Response struct {
	value *EzsignbulksendEditObjectV1Response
	isSet bool
}

func (v NullableEzsignbulksendEditObjectV1Response) Get() *EzsignbulksendEditObjectV1Response {
	return v.value
}

func (v *NullableEzsignbulksendEditObjectV1Response) Set(val *EzsignbulksendEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendEditObjectV1Response(val *EzsignbulksendEditObjectV1Response) *NullableEzsignbulksendEditObjectV1Response {
	return &NullableEzsignbulksendEditObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


