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

// checks if the EzsigndocumentEditEzsignsignaturesV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEditEzsignsignaturesV1ResponseMPayload{}

// EzsigndocumentEditEzsignsignaturesV1ResponseMPayload Payload for PUT /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignsignatures
type EzsigndocumentEditEzsignsignaturesV1ResponseMPayload struct {
	APkiEzsignsignatureID []int32 `json:"a_pkiEzsignsignatureID"`
}

type _EzsigndocumentEditEzsignsignaturesV1ResponseMPayload EzsigndocumentEditEzsignsignaturesV1ResponseMPayload

// NewEzsigndocumentEditEzsignsignaturesV1ResponseMPayload instantiates a new EzsigndocumentEditEzsignsignaturesV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEditEzsignsignaturesV1ResponseMPayload(aPkiEzsignsignatureID []int32) *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload {
	this := EzsigndocumentEditEzsignsignaturesV1ResponseMPayload{}
	this.APkiEzsignsignatureID = aPkiEzsignsignatureID
	return &this
}

// NewEzsigndocumentEditEzsignsignaturesV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentEditEzsignsignaturesV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEditEzsignsignaturesV1ResponseMPayloadWithDefaults() *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload {
	this := EzsigndocumentEditEzsignsignaturesV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignatureID returns the APkiEzsignsignatureID field value
func (o *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) GetAPkiEzsignsignatureID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignatureID
}

// GetAPkiEzsignsignatureIDOk returns a tuple with the APkiEzsignsignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) GetAPkiEzsignsignatureIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignatureID, true
}

// SetAPkiEzsignsignatureID sets field value
func (o *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) SetAPkiEzsignsignatureID(v []int32) {
	o.APkiEzsignsignatureID = v
}

func (o EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignatureID"] = o.APkiEzsignsignatureID
	return toSerialize, nil
}

func (o *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignsignatureID",
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

	varEzsigndocumentEditEzsignsignaturesV1ResponseMPayload := _EzsigndocumentEditEzsignsignaturesV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentEditEzsignsignaturesV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentEditEzsignsignaturesV1ResponseMPayload(varEzsigndocumentEditEzsignsignaturesV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload struct {
	value *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) Get() *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) Set(val *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload(val *EzsigndocumentEditEzsignsignaturesV1ResponseMPayload) *NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload {
	return &NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEditEzsignsignaturesV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


