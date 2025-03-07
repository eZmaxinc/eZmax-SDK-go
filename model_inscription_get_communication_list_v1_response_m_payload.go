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

// checks if the InscriptionGetCommunicationListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptionGetCommunicationListV1ResponseMPayload{}

// InscriptionGetCommunicationListV1ResponseMPayload Response for GET /1/object/inscription/{pkiInscriptionID}/getCommunicationList
type InscriptionGetCommunicationListV1ResponseMPayload struct {
	AObjCommunication []CustomCommunicationListElementResponse `json:"a_objCommunication"`
}

type _InscriptionGetCommunicationListV1ResponseMPayload InscriptionGetCommunicationListV1ResponseMPayload

// NewInscriptionGetCommunicationListV1ResponseMPayload instantiates a new InscriptionGetCommunicationListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptionGetCommunicationListV1ResponseMPayload(aObjCommunication []CustomCommunicationListElementResponse) *InscriptionGetCommunicationListV1ResponseMPayload {
	this := InscriptionGetCommunicationListV1ResponseMPayload{}
	this.AObjCommunication = aObjCommunication
	return &this
}

// NewInscriptionGetCommunicationListV1ResponseMPayloadWithDefaults instantiates a new InscriptionGetCommunicationListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptionGetCommunicationListV1ResponseMPayloadWithDefaults() *InscriptionGetCommunicationListV1ResponseMPayload {
	this := InscriptionGetCommunicationListV1ResponseMPayload{}
	return &this
}

// GetAObjCommunication returns the AObjCommunication field value
func (o *InscriptionGetCommunicationListV1ResponseMPayload) GetAObjCommunication() []CustomCommunicationListElementResponse {
	if o == nil {
		var ret []CustomCommunicationListElementResponse
		return ret
	}

	return o.AObjCommunication
}

// GetAObjCommunicationOk returns a tuple with the AObjCommunication field value
// and a boolean to check if the value has been set.
func (o *InscriptionGetCommunicationListV1ResponseMPayload) GetAObjCommunicationOk() ([]CustomCommunicationListElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunication, true
}

// SetAObjCommunication sets field value
func (o *InscriptionGetCommunicationListV1ResponseMPayload) SetAObjCommunication(v []CustomCommunicationListElementResponse) {
	o.AObjCommunication = v
}

func (o InscriptionGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptionGetCommunicationListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCommunication"] = o.AObjCommunication
	return toSerialize, nil
}

func (o *InscriptionGetCommunicationListV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objCommunication",
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

	varInscriptionGetCommunicationListV1ResponseMPayload := _InscriptionGetCommunicationListV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInscriptionGetCommunicationListV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = InscriptionGetCommunicationListV1ResponseMPayload(varInscriptionGetCommunicationListV1ResponseMPayload)

	return err
}

type NullableInscriptionGetCommunicationListV1ResponseMPayload struct {
	value *InscriptionGetCommunicationListV1ResponseMPayload
	isSet bool
}

func (v NullableInscriptionGetCommunicationListV1ResponseMPayload) Get() *InscriptionGetCommunicationListV1ResponseMPayload {
	return v.value
}

func (v *NullableInscriptionGetCommunicationListV1ResponseMPayload) Set(val *InscriptionGetCommunicationListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptionGetCommunicationListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptionGetCommunicationListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptionGetCommunicationListV1ResponseMPayload(val *InscriptionGetCommunicationListV1ResponseMPayload) *NullableInscriptionGetCommunicationListV1ResponseMPayload {
	return &NullableInscriptionGetCommunicationListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableInscriptionGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptionGetCommunicationListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


