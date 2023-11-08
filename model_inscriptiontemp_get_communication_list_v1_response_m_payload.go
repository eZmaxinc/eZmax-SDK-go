/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// checks if the InscriptiontempGetCommunicationListV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptiontempGetCommunicationListV1ResponseMPayload{}

// InscriptiontempGetCommunicationListV1ResponseMPayload Response for GET /1/object/inscriptiontemp/{pkiInscriptiontempID}/getCommunicationList
type InscriptiontempGetCommunicationListV1ResponseMPayload struct {
	AObjCommunication []CustomCommunicationListElementResponse `json:"a_objCommunication"`
}

// NewInscriptiontempGetCommunicationListV1ResponseMPayload instantiates a new InscriptiontempGetCommunicationListV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptiontempGetCommunicationListV1ResponseMPayload(aObjCommunication []CustomCommunicationListElementResponse) *InscriptiontempGetCommunicationListV1ResponseMPayload {
	this := InscriptiontempGetCommunicationListV1ResponseMPayload{}
	this.AObjCommunication = aObjCommunication
	return &this
}

// NewInscriptiontempGetCommunicationListV1ResponseMPayloadWithDefaults instantiates a new InscriptiontempGetCommunicationListV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptiontempGetCommunicationListV1ResponseMPayloadWithDefaults() *InscriptiontempGetCommunicationListV1ResponseMPayload {
	this := InscriptiontempGetCommunicationListV1ResponseMPayload{}
	return &this
}

// GetAObjCommunication returns the AObjCommunication field value
func (o *InscriptiontempGetCommunicationListV1ResponseMPayload) GetAObjCommunication() []CustomCommunicationListElementResponse {
	if o == nil {
		var ret []CustomCommunicationListElementResponse
		return ret
	}

	return o.AObjCommunication
}

// GetAObjCommunicationOk returns a tuple with the AObjCommunication field value
// and a boolean to check if the value has been set.
func (o *InscriptiontempGetCommunicationListV1ResponseMPayload) GetAObjCommunicationOk() ([]CustomCommunicationListElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunication, true
}

// SetAObjCommunication sets field value
func (o *InscriptiontempGetCommunicationListV1ResponseMPayload) SetAObjCommunication(v []CustomCommunicationListElementResponse) {
	o.AObjCommunication = v
}

func (o InscriptiontempGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptiontempGetCommunicationListV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCommunication"] = o.AObjCommunication
	return toSerialize, nil
}

type NullableInscriptiontempGetCommunicationListV1ResponseMPayload struct {
	value *InscriptiontempGetCommunicationListV1ResponseMPayload
	isSet bool
}

func (v NullableInscriptiontempGetCommunicationListV1ResponseMPayload) Get() *InscriptiontempGetCommunicationListV1ResponseMPayload {
	return v.value
}

func (v *NullableInscriptiontempGetCommunicationListV1ResponseMPayload) Set(val *InscriptiontempGetCommunicationListV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptiontempGetCommunicationListV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptiontempGetCommunicationListV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptiontempGetCommunicationListV1ResponseMPayload(val *InscriptiontempGetCommunicationListV1ResponseMPayload) *NullableInscriptiontempGetCommunicationListV1ResponseMPayload {
	return &NullableInscriptiontempGetCommunicationListV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableInscriptiontempGetCommunicationListV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptiontempGetCommunicationListV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


