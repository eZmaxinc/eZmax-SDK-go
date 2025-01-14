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

// checks if the InscriptiontempGetCommunicationCountV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptiontempGetCommunicationCountV1ResponseMPayload{}

// InscriptiontempGetCommunicationCountV1ResponseMPayload Response for GET /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationCount
type InscriptiontempGetCommunicationCountV1ResponseMPayload struct {
	// The count of Communication.
	ICommunicationCount int32 `json:"iCommunicationCount"`
}

type _InscriptiontempGetCommunicationCountV1ResponseMPayload InscriptiontempGetCommunicationCountV1ResponseMPayload

// NewInscriptiontempGetCommunicationCountV1ResponseMPayload instantiates a new InscriptiontempGetCommunicationCountV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptiontempGetCommunicationCountV1ResponseMPayload(iCommunicationCount int32) *InscriptiontempGetCommunicationCountV1ResponseMPayload {
	this := InscriptiontempGetCommunicationCountV1ResponseMPayload{}
	this.ICommunicationCount = iCommunicationCount
	return &this
}

// NewInscriptiontempGetCommunicationCountV1ResponseMPayloadWithDefaults instantiates a new InscriptiontempGetCommunicationCountV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptiontempGetCommunicationCountV1ResponseMPayloadWithDefaults() *InscriptiontempGetCommunicationCountV1ResponseMPayload {
	this := InscriptiontempGetCommunicationCountV1ResponseMPayload{}
	return &this
}

// GetICommunicationCount returns the ICommunicationCount field value
func (o *InscriptiontempGetCommunicationCountV1ResponseMPayload) GetICommunicationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICommunicationCount
}

// GetICommunicationCountOk returns a tuple with the ICommunicationCount field value
// and a boolean to check if the value has been set.
func (o *InscriptiontempGetCommunicationCountV1ResponseMPayload) GetICommunicationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICommunicationCount, true
}

// SetICommunicationCount sets field value
func (o *InscriptiontempGetCommunicationCountV1ResponseMPayload) SetICommunicationCount(v int32) {
	o.ICommunicationCount = v
}

func (o InscriptiontempGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptiontempGetCommunicationCountV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iCommunicationCount"] = o.ICommunicationCount
	return toSerialize, nil
}

func (o *InscriptiontempGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iCommunicationCount",
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

	varInscriptiontempGetCommunicationCountV1ResponseMPayload := _InscriptiontempGetCommunicationCountV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInscriptiontempGetCommunicationCountV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = InscriptiontempGetCommunicationCountV1ResponseMPayload(varInscriptiontempGetCommunicationCountV1ResponseMPayload)

	return err
}

type NullableInscriptiontempGetCommunicationCountV1ResponseMPayload struct {
	value *InscriptiontempGetCommunicationCountV1ResponseMPayload
	isSet bool
}

func (v NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) Get() *InscriptiontempGetCommunicationCountV1ResponseMPayload {
	return v.value
}

func (v *NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) Set(val *InscriptiontempGetCommunicationCountV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptiontempGetCommunicationCountV1ResponseMPayload(val *InscriptiontempGetCommunicationCountV1ResponseMPayload) *NullableInscriptiontempGetCommunicationCountV1ResponseMPayload {
	return &NullableInscriptiontempGetCommunicationCountV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptiontempGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


