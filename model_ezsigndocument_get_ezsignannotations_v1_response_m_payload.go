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

// checks if the EzsigndocumentGetEzsignannotationsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetEzsignannotationsV1ResponseMPayload{}

// EzsigndocumentGetEzsignannotationsV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getEzsignannotations
type EzsigndocumentGetEzsignannotationsV1ResponseMPayload struct {
	AObjEzsignannotation []EzsignannotationResponseCompound `json:"a_objEzsignannotation"`
}

type _EzsigndocumentGetEzsignannotationsV1ResponseMPayload EzsigndocumentGetEzsignannotationsV1ResponseMPayload

// NewEzsigndocumentGetEzsignannotationsV1ResponseMPayload instantiates a new EzsigndocumentGetEzsignannotationsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetEzsignannotationsV1ResponseMPayload(aObjEzsignannotation []EzsignannotationResponseCompound) *EzsigndocumentGetEzsignannotationsV1ResponseMPayload {
	this := EzsigndocumentGetEzsignannotationsV1ResponseMPayload{}
	this.AObjEzsignannotation = aObjEzsignannotation
	return &this
}

// NewEzsigndocumentGetEzsignannotationsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetEzsignannotationsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetEzsignannotationsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetEzsignannotationsV1ResponseMPayload {
	this := EzsigndocumentGetEzsignannotationsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignannotation returns the AObjEzsignannotation field value
func (o *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) GetAObjEzsignannotation() []EzsignannotationResponseCompound {
	if o == nil {
		var ret []EzsignannotationResponseCompound
		return ret
	}

	return o.AObjEzsignannotation
}

// GetAObjEzsignannotationOk returns a tuple with the AObjEzsignannotation field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) GetAObjEzsignannotationOk() ([]EzsignannotationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignannotation, true
}

// SetAObjEzsignannotation sets field value
func (o *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) SetAObjEzsignannotation(v []EzsignannotationResponseCompound) {
	o.AObjEzsignannotation = v
}

func (o EzsigndocumentGetEzsignannotationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetEzsignannotationsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignannotation"] = o.AObjEzsignannotation
	return toSerialize, nil
}

func (o *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignannotation",
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

	varEzsigndocumentGetEzsignannotationsV1ResponseMPayload := _EzsigndocumentGetEzsignannotationsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetEzsignannotationsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetEzsignannotationsV1ResponseMPayload(varEzsigndocumentGetEzsignannotationsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload struct {
	value *EzsigndocumentGetEzsignannotationsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) Get() *EzsigndocumentGetEzsignannotationsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) Set(val *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload(val *EzsigndocumentGetEzsignannotationsV1ResponseMPayload) *NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload {
	return &NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetEzsignannotationsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


