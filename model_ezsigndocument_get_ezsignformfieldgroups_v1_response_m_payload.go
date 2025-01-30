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

// checks if the EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload{}

// EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getEzsignformfieldgroups
type EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload struct {
	AObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound `json:"a_objEzsignformfieldgroup"`
}

type _EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload

// NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload instantiates a new EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload(aObjEzsignformfieldgroup []EzsignformfieldgroupResponseCompound) *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload {
	this := EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload{}
	this.AObjEzsignformfieldgroup = aObjEzsignformfieldgroup
	return &this
}

// NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload {
	this := EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignformfieldgroup returns the AObjEzsignformfieldgroup field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) GetAObjEzsignformfieldgroup() []EzsignformfieldgroupResponseCompound {
	if o == nil {
		var ret []EzsignformfieldgroupResponseCompound
		return ret
	}

	return o.AObjEzsignformfieldgroup
}

// GetAObjEzsignformfieldgroupOk returns a tuple with the AObjEzsignformfieldgroup field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) GetAObjEzsignformfieldgroupOk() ([]EzsignformfieldgroupResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignformfieldgroup, true
}

// SetAObjEzsignformfieldgroup sets field value
func (o *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) SetAObjEzsignformfieldgroup(v []EzsignformfieldgroupResponseCompound) {
	o.AObjEzsignformfieldgroup = v
}

func (o EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignformfieldgroup"] = o.AObjEzsignformfieldgroup
	return toSerialize, nil
}

func (o *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignformfieldgroup",
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

	varEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload := _EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload(varEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload struct {
	value *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) Get() *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) Set(val *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload(val *EzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) *NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload {
	return &NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignformfieldgroupsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


