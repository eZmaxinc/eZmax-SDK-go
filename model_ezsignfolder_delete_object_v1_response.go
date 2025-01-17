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

// checks if the EzsignfolderDeleteObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderDeleteObjectV1Response{}

// EzsignfolderDeleteObjectV1Response Response for DELETE /1/object/ezsignfolder/{pkiEzsignfolderID}
type EzsignfolderDeleteObjectV1Response struct {
	CommonResponse
}

type _EzsignfolderDeleteObjectV1Response EzsignfolderDeleteObjectV1Response

// NewEzsignfolderDeleteObjectV1Response instantiates a new EzsignfolderDeleteObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderDeleteObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignfolderDeleteObjectV1Response {
	this := EzsignfolderDeleteObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignfolderDeleteObjectV1ResponseWithDefaults instantiates a new EzsignfolderDeleteObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderDeleteObjectV1ResponseWithDefaults() *EzsignfolderDeleteObjectV1Response {
	this := EzsignfolderDeleteObjectV1Response{}
	return &this
}

func (o EzsignfolderDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderDeleteObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignfolderDeleteObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfolderDeleteObjectV1Response := _EzsignfolderDeleteObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderDeleteObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignfolderDeleteObjectV1Response(varEzsignfolderDeleteObjectV1Response)

	return err
}

type NullableEzsignfolderDeleteObjectV1Response struct {
	value *EzsignfolderDeleteObjectV1Response
	isSet bool
}

func (v NullableEzsignfolderDeleteObjectV1Response) Get() *EzsignfolderDeleteObjectV1Response {
	return v.value
}

func (v *NullableEzsignfolderDeleteObjectV1Response) Set(val *EzsignfolderDeleteObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderDeleteObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderDeleteObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderDeleteObjectV1Response(val *EzsignfolderDeleteObjectV1Response) *NullableEzsignfolderDeleteObjectV1Response {
	return &NullableEzsignfolderDeleteObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignfolderDeleteObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderDeleteObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


