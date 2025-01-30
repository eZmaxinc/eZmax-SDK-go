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

// checks if the InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload{}

// InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload Response for GET /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationCount
type InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload struct {
	// The count of Communication.
	ICommunicationCount int32 `json:"iCommunicationCount"`
}

type _InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload

// NewInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload instantiates a new InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload(iCommunicationCount int32) *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload {
	this := InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload{}
	this.ICommunicationCount = iCommunicationCount
	return &this
}

// NewInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayloadWithDefaults instantiates a new InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayloadWithDefaults() *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload {
	this := InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload{}
	return &this
}

// GetICommunicationCount returns the ICommunicationCount field value
func (o *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) GetICommunicationCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICommunicationCount
}

// GetICommunicationCountOk returns a tuple with the ICommunicationCount field value
// and a boolean to check if the value has been set.
func (o *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) GetICommunicationCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICommunicationCount, true
}

// SetICommunicationCount sets field value
func (o *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) SetICommunicationCount(v int32) {
	o.ICommunicationCount = v
}

func (o InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iCommunicationCount"] = o.ICommunicationCount
	return toSerialize, nil
}

func (o *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload := _InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload(varInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload)

	return err
}

type NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload struct {
	value *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload
	isSet bool
}

func (v NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) Get() *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload {
	return v.value
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) Set(val *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload(val *InscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) *NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload {
	return &NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationCountV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


