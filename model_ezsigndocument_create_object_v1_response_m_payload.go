/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsigndocumentCreateObjectV1ResponseMPayload Payload for the /1/object/ezsigndocument/createObject API Request
type EzsigndocumentCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsigndocumentID []int32 `json:"a_pkiEzsigndocumentID"`
}

// NewEzsigndocumentCreateObjectV1ResponseMPayload instantiates a new EzsigndocumentCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentCreateObjectV1ResponseMPayload(aPkiEzsigndocumentID []int32) *EzsigndocumentCreateObjectV1ResponseMPayload {
	this := EzsigndocumentCreateObjectV1ResponseMPayload{}
	this.APkiEzsigndocumentID = aPkiEzsigndocumentID
	return &this
}

// NewEzsigndocumentCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentCreateObjectV1ResponseMPayloadWithDefaults() *EzsigndocumentCreateObjectV1ResponseMPayload {
	this := EzsigndocumentCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsigndocumentID returns the APkiEzsigndocumentID field value
func (o *EzsigndocumentCreateObjectV1ResponseMPayload) GetAPkiEzsigndocumentID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsigndocumentID
}

// GetAPkiEzsigndocumentIDOk returns a tuple with the APkiEzsigndocumentID field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentCreateObjectV1ResponseMPayload) GetAPkiEzsigndocumentIDOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.APkiEzsigndocumentID, true
}

// SetAPkiEzsigndocumentID sets field value
func (o *EzsigndocumentCreateObjectV1ResponseMPayload) SetAPkiEzsigndocumentID(v []int32) {
	o.APkiEzsigndocumentID = v
}

func (o EzsigndocumentCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_pkiEzsigndocumentID"] = o.APkiEzsigndocumentID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentCreateObjectV1ResponseMPayload struct {
	value *EzsigndocumentCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentCreateObjectV1ResponseMPayload) Get() *EzsigndocumentCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseMPayload) Set(val *EzsigndocumentCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentCreateObjectV1ResponseMPayload(val *EzsigndocumentCreateObjectV1ResponseMPayload) *NullableEzsigndocumentCreateObjectV1ResponseMPayload {
	return &NullableEzsigndocumentCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


