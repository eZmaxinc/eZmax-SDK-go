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

// checks if the InscriptionnotauthenticatedGetCommunicationsendersV1Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptionnotauthenticatedGetCommunicationsendersV1Response{}

// InscriptionnotauthenticatedGetCommunicationsendersV1Response Response for GET /1/object/inscriptionnotauthenticated/{pkiInscriptionnotauthenticatedID}/getCommunicationrecipients
type InscriptionnotauthenticatedGetCommunicationsendersV1Response struct {
	CommonResponse
	MPayload InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload `json:"mPayload"`
}

type _InscriptionnotauthenticatedGetCommunicationsendersV1Response InscriptionnotauthenticatedGetCommunicationsendersV1Response

// NewInscriptionnotauthenticatedGetCommunicationsendersV1Response instantiates a new InscriptionnotauthenticatedGetCommunicationsendersV1Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptionnotauthenticatedGetCommunicationsendersV1Response(mPayload InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload, objDebugPayload CommonResponseObjDebugPayload) *InscriptionnotauthenticatedGetCommunicationsendersV1Response {
	this := InscriptionnotauthenticatedGetCommunicationsendersV1Response{}
	this.ObjDebugPayload = objDebugPayload
	this.MPayload = mPayload
	return &this
}

// NewInscriptionnotauthenticatedGetCommunicationsendersV1ResponseWithDefaults instantiates a new InscriptionnotauthenticatedGetCommunicationsendersV1Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptionnotauthenticatedGetCommunicationsendersV1ResponseWithDefaults() *InscriptionnotauthenticatedGetCommunicationsendersV1Response {
	this := InscriptionnotauthenticatedGetCommunicationsendersV1Response{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *InscriptionnotauthenticatedGetCommunicationsendersV1Response) GetMPayload() InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload {
	if o == nil {
		var ret InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *InscriptionnotauthenticatedGetCommunicationsendersV1Response) GetMPayloadOk() (*InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *InscriptionnotauthenticatedGetCommunicationsendersV1Response) SetMPayload(v InscriptionnotauthenticatedGetCommunicationsendersV1ResponseMPayload) {
	o.MPayload = v
}

func (o InscriptionnotauthenticatedGetCommunicationsendersV1Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptionnotauthenticatedGetCommunicationsendersV1Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mPayload"] = o.MPayload
	return toSerialize, nil
}

func (o *InscriptionnotauthenticatedGetCommunicationsendersV1Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mPayload",
		"objDebugPayload",
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

	varInscriptionnotauthenticatedGetCommunicationsendersV1Response := _InscriptionnotauthenticatedGetCommunicationsendersV1Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInscriptionnotauthenticatedGetCommunicationsendersV1Response)

	if err != nil {
		return err
	}

	*o = InscriptionnotauthenticatedGetCommunicationsendersV1Response(varInscriptionnotauthenticatedGetCommunicationsendersV1Response)

	return err
}

type NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response struct {
	value *InscriptionnotauthenticatedGetCommunicationsendersV1Response
	isSet bool
}

func (v NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) Get() *InscriptionnotauthenticatedGetCommunicationsendersV1Response {
	return v.value
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) Set(val *InscriptionnotauthenticatedGetCommunicationsendersV1Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptionnotauthenticatedGetCommunicationsendersV1Response(val *InscriptionnotauthenticatedGetCommunicationsendersV1Response) *NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response {
	return &NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response{value: val, isSet: true}
}

func (v NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptionnotauthenticatedGetCommunicationsendersV1Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


