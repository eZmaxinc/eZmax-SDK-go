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

// checks if the EzsignsignatureCreateObjectV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureCreateObjectV2ResponseMPayload{}

// EzsignsignatureCreateObjectV2ResponseMPayload Payload for POST /2/object/ezsignsignature
type EzsignsignatureCreateObjectV2ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignsignatureID []int32 `json:"a_pkiEzsignsignatureID"`
}

// NewEzsignsignatureCreateObjectV2ResponseMPayload instantiates a new EzsignsignatureCreateObjectV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureCreateObjectV2ResponseMPayload(aPkiEzsignsignatureID []int32) *EzsignsignatureCreateObjectV2ResponseMPayload {
	this := EzsignsignatureCreateObjectV2ResponseMPayload{}
	this.APkiEzsignsignatureID = aPkiEzsignsignatureID
	return &this
}

// NewEzsignsignatureCreateObjectV2ResponseMPayloadWithDefaults instantiates a new EzsignsignatureCreateObjectV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureCreateObjectV2ResponseMPayloadWithDefaults() *EzsignsignatureCreateObjectV2ResponseMPayload {
	this := EzsignsignatureCreateObjectV2ResponseMPayload{}
	return &this
}

// GetAPkiEzsignsignatureID returns the APkiEzsignsignatureID field value
func (o *EzsignsignatureCreateObjectV2ResponseMPayload) GetAPkiEzsignsignatureID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignsignatureID
}

// GetAPkiEzsignsignatureIDOk returns a tuple with the APkiEzsignsignatureID field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureCreateObjectV2ResponseMPayload) GetAPkiEzsignsignatureIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignsignatureID, true
}

// SetAPkiEzsignsignatureID sets field value
func (o *EzsignsignatureCreateObjectV2ResponseMPayload) SetAPkiEzsignsignatureID(v []int32) {
	o.APkiEzsignsignatureID = v
}

func (o EzsignsignatureCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureCreateObjectV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignsignatureID"] = o.APkiEzsignsignatureID
	return toSerialize, nil
}

type NullableEzsignsignatureCreateObjectV2ResponseMPayload struct {
	value *EzsignsignatureCreateObjectV2ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignatureCreateObjectV2ResponseMPayload) Get() *EzsignsignatureCreateObjectV2ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignatureCreateObjectV2ResponseMPayload) Set(val *EzsignsignatureCreateObjectV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureCreateObjectV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureCreateObjectV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureCreateObjectV2ResponseMPayload(val *EzsignsignatureCreateObjectV2ResponseMPayload) *NullableEzsignsignatureCreateObjectV2ResponseMPayload {
	return &NullableEzsignsignatureCreateObjectV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignatureCreateObjectV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureCreateObjectV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


