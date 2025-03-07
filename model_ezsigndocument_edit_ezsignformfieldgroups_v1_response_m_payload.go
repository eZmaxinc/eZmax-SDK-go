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

// checks if the EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload{}

// EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload Payload for PUT /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignformfieldgroups
type EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload struct {
	APkiEzsignformfieldgroupID []int32 `json:"a_pkiEzsignformfieldgroupID"`
}

type _EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload

// NewEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload instantiates a new EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload(aPkiEzsignformfieldgroupID []int32) *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload {
	this := EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload{}
	this.APkiEzsignformfieldgroupID = aPkiEzsignformfieldgroupID
	return &this
}

// NewEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayloadWithDefaults() *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload {
	this := EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignformfieldgroupID returns the APkiEzsignformfieldgroupID field value
func (o *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) GetAPkiEzsignformfieldgroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignformfieldgroupID
}

// GetAPkiEzsignformfieldgroupIDOk returns a tuple with the APkiEzsignformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) GetAPkiEzsignformfieldgroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignformfieldgroupID, true
}

// SetAPkiEzsignformfieldgroupID sets field value
func (o *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) SetAPkiEzsignformfieldgroupID(v []int32) {
	o.APkiEzsignformfieldgroupID = v
}

func (o EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignformfieldgroupID"] = o.APkiEzsignformfieldgroupID
	return toSerialize, nil
}

func (o *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignformfieldgroupID",
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

	varEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload := _EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload(varEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload struct {
	value *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) Get() *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) Set(val *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload(val *EzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) *NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload {
	return &NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEditEzsignformfieldgroupsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


