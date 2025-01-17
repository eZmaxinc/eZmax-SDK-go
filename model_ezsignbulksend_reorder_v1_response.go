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

// checks if the EzsignbulksendReorderV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendReorderV1Response{}

// EzsignbulksendReorderV1Response Response for POST /1/object/ezsignbulksend/{pkiEzsignbulksendID}/reorder
type EzsignbulksendReorderV1Response struct {
	CommonResponse
}

type _EzsignbulksendReorderV1Response EzsignbulksendReorderV1Response

// NewEzsignbulksendReorderV1Response instantiates a new EzsignbulksendReorderV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendReorderV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignbulksendReorderV1Response {
	this := EzsignbulksendReorderV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignbulksendReorderV1ResponseWithDefaults instantiates a new EzsignbulksendReorderV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendReorderV1ResponseWithDefaults() *EzsignbulksendReorderV1Response {
	this := EzsignbulksendReorderV1Response{}
	return &this
}

func (o EzsignbulksendReorderV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendReorderV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignbulksendReorderV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignbulksendReorderV1Response := _EzsignbulksendReorderV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendReorderV1Response)

	if err != nil {
		return err
	}

	*o = EzsignbulksendReorderV1Response(varEzsignbulksendReorderV1Response)

	return err
}

type NullableEzsignbulksendReorderV1Response struct {
	value *EzsignbulksendReorderV1Response
	isSet bool
}

func (v NullableEzsignbulksendReorderV1Response) Get() *EzsignbulksendReorderV1Response {
	return v.value
}

func (v *NullableEzsignbulksendReorderV1Response) Set(val *EzsignbulksendReorderV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendReorderV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendReorderV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendReorderV1Response(val *EzsignbulksendReorderV1Response) *NullableEzsignbulksendReorderV1Response {
	return &NullableEzsignbulksendReorderV1Response{value: val, isSet: true}
}

func (v NullableEzsignbulksendReorderV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendReorderV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


