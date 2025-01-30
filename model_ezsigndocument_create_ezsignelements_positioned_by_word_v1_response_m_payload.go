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

// checks if the EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload{}

// EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload Payload for POST /1/object/ezsigndocument/{pkiEzsigndocumentID}/createEzsignelementsPositionedByWord
type EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload struct {
	APkiEzsignsignatureID []int32 `json:"a_pkiEzsignsignatureID"`
	APkiEzsignformfieldgroupID []int32 `json:"a_pkiEzsignformfieldgroupID"`
}

type _EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload(aPkiEzsignsignatureID []int32, aPkiEzsignformfieldgroupID []int32) *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload{}
	this.APkiEzsignsignatureID = aPkiEzsignsignatureID
	this.APkiEzsignformfieldgroupID = aPkiEzsignformfieldgroupID
	return &this
}

// NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayloadWithDefaults() *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload {
	this := EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignatureID returns the APkiEzsignsignatureID field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) GetAPkiEzsignsignatureID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignatureID
}

// GetAPkiEzsignsignatureIDOk returns a tuple with the APkiEzsignsignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) GetAPkiEzsignsignatureIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignatureID, true
}

// SetAPkiEzsignsignatureID sets field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) SetAPkiEzsignsignatureID(v []int32) {
	o.APkiEzsignsignatureID = v
}

// GetAPkiEzsignformfieldgroupID returns the APkiEzsignformfieldgroupID field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) GetAPkiEzsignformfieldgroupID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignformfieldgroupID
}

// GetAPkiEzsignformfieldgroupIDOk returns a tuple with the APkiEzsignformfieldgroupID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) GetAPkiEzsignformfieldgroupIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignformfieldgroupID, true
}

// SetAPkiEzsignformfieldgroupID sets field value
func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) SetAPkiEzsignformfieldgroupID(v []int32) {
	o.APkiEzsignformfieldgroupID = v
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignatureID"] = o.APkiEzsignsignatureID
	toSerialize["a_pkiEzsignformfieldgroupID"] = o.APkiEzsignformfieldgroupID
	return toSerialize, nil
}

func (o *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignsignatureID",
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

	varEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload := _EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload(varEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload struct {
	value *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) Get() *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) Set(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload(val *EzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload {
	return &NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateEzsignelementsPositionedByWordV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


