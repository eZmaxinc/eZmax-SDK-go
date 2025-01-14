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

// checks if the EzsigndocumentEditEzsignannotationsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentEditEzsignannotationsV1ResponseMPayload{}

// EzsigndocumentEditEzsignannotationsV1ResponseMPayload Payload for PUT /1/object/ezsigndocument/{pkiEzsigndocumentID}/editEzsignannotations
type EzsigndocumentEditEzsignannotationsV1ResponseMPayload struct {
	APkiEzsignannotationID []int32 `json:"a_pkiEzsignannotationID"`
}

type _EzsigndocumentEditEzsignannotationsV1ResponseMPayload EzsigndocumentEditEzsignannotationsV1ResponseMPayload

// NewEzsigndocumentEditEzsignannotationsV1ResponseMPayload instantiates a new EzsigndocumentEditEzsignannotationsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentEditEzsignannotationsV1ResponseMPayload(aPkiEzsignannotationID []int32) *EzsigndocumentEditEzsignannotationsV1ResponseMPayload {
	this := EzsigndocumentEditEzsignannotationsV1ResponseMPayload{}
	this.APkiEzsignannotationID = aPkiEzsignannotationID
	return &this
}

// NewEzsigndocumentEditEzsignannotationsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentEditEzsignannotationsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentEditEzsignannotationsV1ResponseMPayloadWithDefaults() *EzsigndocumentEditEzsignannotationsV1ResponseMPayload {
	this := EzsigndocumentEditEzsignannotationsV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignannotationID returns the APkiEzsignannotationID field value
func (o *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) GetAPkiEzsignannotationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignannotationID
}

// GetAPkiEzsignannotationIDOk returns a tuple with the APkiEzsignannotationID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) GetAPkiEzsignannotationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignannotationID, true
}

// SetAPkiEzsignannotationID sets field value
func (o *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) SetAPkiEzsignannotationID(v []int32) {
	o.APkiEzsignannotationID = v
}

func (o EzsigndocumentEditEzsignannotationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentEditEzsignannotationsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignannotationID"] = o.APkiEzsignannotationID
	return toSerialize, nil
}

func (o *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignannotationID",
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

	varEzsigndocumentEditEzsignannotationsV1ResponseMPayload := _EzsigndocumentEditEzsignannotationsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentEditEzsignannotationsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentEditEzsignannotationsV1ResponseMPayload(varEzsigndocumentEditEzsignannotationsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload struct {
	value *EzsigndocumentEditEzsignannotationsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) Get() *EzsigndocumentEditEzsignannotationsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) Set(val *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload(val *EzsigndocumentEditEzsignannotationsV1ResponseMPayload) *NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload {
	return &NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentEditEzsignannotationsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


