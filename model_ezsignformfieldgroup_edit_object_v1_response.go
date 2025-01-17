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

// checks if the EzsignformfieldgroupEditObjectV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupEditObjectV1Response{}

// EzsignformfieldgroupEditObjectV1Response Response for PUT /1/object/ezsignformfieldgroup/{pkiEzsignfoldersignerassociationID}
type EzsignformfieldgroupEditObjectV1Response struct {
	CommonResponse
}

type _EzsignformfieldgroupEditObjectV1Response EzsignformfieldgroupEditObjectV1Response

// NewEzsignformfieldgroupEditObjectV1Response instantiates a new EzsignformfieldgroupEditObjectV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupEditObjectV1Response(objDebugPayload CommonResponseObjDebugPayload) *EzsignformfieldgroupEditObjectV1Response {
	this := EzsignformfieldgroupEditObjectV1Response{}
	this.ObjDebugPayload = objDebugPayload
	return &this
}

// NewEzsignformfieldgroupEditObjectV1ResponseWithDefaults instantiates a new EzsignformfieldgroupEditObjectV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupEditObjectV1ResponseWithDefaults() *EzsignformfieldgroupEditObjectV1Response {
	this := EzsignformfieldgroupEditObjectV1Response{}
	return &this
}

func (o EzsignformfieldgroupEditObjectV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupEditObjectV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzsignformfieldgroupEditObjectV1Response) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignformfieldgroupEditObjectV1Response := _EzsignformfieldgroupEditObjectV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldgroupEditObjectV1Response)

	if err != nil {
		return err
	}

	*o = EzsignformfieldgroupEditObjectV1Response(varEzsignformfieldgroupEditObjectV1Response)

	return err
}

type NullableEzsignformfieldgroupEditObjectV1Response struct {
	value *EzsignformfieldgroupEditObjectV1Response
	isSet bool
}

func (v NullableEzsignformfieldgroupEditObjectV1Response) Get() *EzsignformfieldgroupEditObjectV1Response {
	return v.value
}

func (v *NullableEzsignformfieldgroupEditObjectV1Response) Set(val *EzsignformfieldgroupEditObjectV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupEditObjectV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupEditObjectV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupEditObjectV1Response(val *EzsignformfieldgroupEditObjectV1Response) *NullableEzsignformfieldgroupEditObjectV1Response {
	return &NullableEzsignformfieldgroupEditObjectV1Response{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupEditObjectV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupEditObjectV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


