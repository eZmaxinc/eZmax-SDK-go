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

// checks if the SignatureCreateObjectV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureCreateObjectV1ResponseMPayload{}

// SignatureCreateObjectV1ResponseMPayload Payload for POST /1/object/signature
type SignatureCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiSignatureID []int32 `json:"a_pkiSignatureID"`
}

// NewSignatureCreateObjectV1ResponseMPayload instantiates a new SignatureCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureCreateObjectV1ResponseMPayload(aPkiSignatureID []int32) *SignatureCreateObjectV1ResponseMPayload {
	this := SignatureCreateObjectV1ResponseMPayload{}
	this.APkiSignatureID = aPkiSignatureID
	return &this
}

// NewSignatureCreateObjectV1ResponseMPayloadWithDefaults instantiates a new SignatureCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureCreateObjectV1ResponseMPayloadWithDefaults() *SignatureCreateObjectV1ResponseMPayload {
	this := SignatureCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiSignatureID returns the APkiSignatureID field value
func (o *SignatureCreateObjectV1ResponseMPayload) GetAPkiSignatureID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiSignatureID
}

// GetAPkiSignatureIDOk returns a tuple with the APkiSignatureID field value
// and a boolean to check if the value has been set.
func (o *SignatureCreateObjectV1ResponseMPayload) GetAPkiSignatureIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiSignatureID, true
}

// SetAPkiSignatureID sets field value
func (o *SignatureCreateObjectV1ResponseMPayload) SetAPkiSignatureID(v []int32) {
	o.APkiSignatureID = v
}

func (o SignatureCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureCreateObjectV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiSignatureID"] = o.APkiSignatureID
	return toSerialize, nil
}

type NullableSignatureCreateObjectV1ResponseMPayload struct {
	value *SignatureCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableSignatureCreateObjectV1ResponseMPayload) Get() *SignatureCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableSignatureCreateObjectV1ResponseMPayload) Set(val *SignatureCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureCreateObjectV1ResponseMPayload(val *SignatureCreateObjectV1ResponseMPayload) *NullableSignatureCreateObjectV1ResponseMPayload {
	return &NullableSignatureCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableSignatureCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

