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

// checks if the EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload{}

// EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload Payload for PUT /1/object/ezsigntemplatedocument/{pkiEzsigntemplatedocumentID}/editEzsigntemplateformfieldgroups
type EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload struct {
	APkiEzsigntemplateformfieldgroupID []int32 `json:"a_pkiEzsigntemplateformfieldgroupID"`
}

type _EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload(aPkiEzsigntemplateformfieldgroupID []int32) *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload{}
	this.APkiEzsigntemplateformfieldgroupID = aPkiEzsigntemplateformfieldgroupID
	return &this
}

// NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayloadWithDefaults instantiates a new EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayloadWithDefaults() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload {
	this := EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigntemplateformfieldgroupID returns the APkiEzsigntemplateformfieldgroupID field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) GetAPkiEzsigntemplateformfieldgroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigntemplateformfieldgroupID
}

// GetAPkiEzsigntemplateformfieldgroupIDOk returns a tuple with the APkiEzsigntemplateformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) GetAPkiEzsigntemplateformfieldgroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsigntemplateformfieldgroupID, true
}

// SetAPkiEzsigntemplateformfieldgroupID sets field value
func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) SetAPkiEzsigntemplateformfieldgroupID(v []int32) {
	o.APkiEzsigntemplateformfieldgroupID = v
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsigntemplateformfieldgroupID"] = o.APkiEzsigntemplateformfieldgroupID
	return toSerialize, nil
}

func (o *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsigntemplateformfieldgroupID",
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

	varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload := _EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload(varEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload)

	return err
}

type NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload struct {
	value *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) Get() *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) Set(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload(val *EzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload {
	return &NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatedocumentEditEzsigntemplateformfieldgroupsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


