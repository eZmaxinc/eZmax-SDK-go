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

// checks if the EzsigndocumentEndPrematurelyV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEndPrematurelyV1Response{}

// EzsigndocumentEndPrematurelyV1Response Response for POST /1/object/ezsigndocument/{pkiEzsigndocument}/endPrematurely
type EzsigndocumentEndPrematurelyV1Response struct {
	CommonResponse
}

type _EzsigndocumentEndPrematurelyV1Response EzsigndocumentEndPrematurelyV1Response

// NewEzsigndocumentEndPrematurelyV1Response instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEndPrematurelyV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults instantiates a new EzsigndocumentEndPrematurelyV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEndPrematurelyV1ResponseWithDefaults() *EzsigndocumentEndPrematurelyV1Response {
	this := EzsigndocumentEndPrematurelyV1Response{}
	return &this
}

func (o EzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEndPrematurelyV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsigndocumentEndPrematurelyV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsigndocumentEndPrematurelyV1Response := _EzsigndocumentEndPrematurelyV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentEndPrematurelyV1Response)

	if err != nil {
		return err
	}

	*o = EzsigndocumentEndPrematurelyV1Response(varEzsigndocumentEndPrematurelyV1Response)

	return err
}

type NullableEzsigndocumentEndPrematurelyV1Response struct {
	value *EzsigndocumentEndPrematurelyV1Response
	isSet bool
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) Get() *EzsigndocumentEndPrematurelyV1Response {
	return v.value
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Set(val *EzsigndocumentEndPrematurelyV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEndPrematurelyV1Response(val *EzsigndocumentEndPrematurelyV1Response) *NullableEzsigndocumentEndPrematurelyV1Response {
	return &NullableEzsigndocumentEndPrematurelyV1Response{value: val, isSet: true}
}

func (v NullableEzsigndocumentEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEndPrematurelyV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


