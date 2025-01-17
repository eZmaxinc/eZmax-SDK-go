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

// checks if the EzsignfolderEndPrematurelyV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderEndPrematurelyV1Response{}

// EzsignfolderEndPrematurelyV1Response Response for POST /1/object/ezsignfolder/{pkiEzsignfolderID}/endPrematurely
type EzsignfolderEndPrematurelyV1Response struct {
	CommonResponse
}

type _EzsignfolderEndPrematurelyV1Response EzsignfolderEndPrematurelyV1Response

// NewEzsignfolderEndPrematurelyV1Response instantiates a new EzsignfolderEndPrematurelyV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderEndPrematurelyV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignfolderEndPrematurelyV1Response {
	this := EzsignfolderEndPrematurelyV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignfolderEndPrematurelyV1ResponseWithDefaults instantiates a new EzsignfolderEndPrematurelyV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderEndPrematurelyV1ResponseWithDefaults() *EzsignfolderEndPrematurelyV1Response {
	this := EzsignfolderEndPrematurelyV1Response{}
	return &this
}

func (o EzsignfolderEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderEndPrematurelyV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignfolderEndPrematurelyV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfolderEndPrematurelyV1Response := _EzsignfolderEndPrematurelyV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderEndPrematurelyV1Response)

	if err != nil {
		return err
	}

	*o = EzsignfolderEndPrematurelyV1Response(varEzsignfolderEndPrematurelyV1Response)

	return err
}

type NullableEzsignfolderEndPrematurelyV1Response struct {
	value *EzsignfolderEndPrematurelyV1Response
	isSet bool
}

func (v NullableEzsignfolderEndPrematurelyV1Response) Get() *EzsignfolderEndPrematurelyV1Response {
	return v.value
}

func (v *NullableEzsignfolderEndPrematurelyV1Response) Set(val *EzsignfolderEndPrematurelyV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderEndPrematurelyV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderEndPrematurelyV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderEndPrematurelyV1Response(val *EzsignfolderEndPrematurelyV1Response) *NullableEzsignfolderEndPrematurelyV1Response {
	return &NullableEzsignfolderEndPrematurelyV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderEndPrematurelyV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderEndPrematurelyV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


