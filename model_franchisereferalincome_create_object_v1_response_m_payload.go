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

// FranchisereferalincomeCreateObjectV1ResponseMPayload Payload for the /1/object/franchisereferalincome/createObject API Request
type FranchisereferalincomeCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiFranchisereferalincomeID []int32 `json:"a_pkiFranchisereferalincomeID"`
}

// NewFranchisereferalincomeCreateObjectV1ResponseMPayload instantiates a new FranchisereferalincomeCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeCreateObjectV1ResponseMPayload(aPkiFranchisereferalincomeID []int32) *FranchisereferalincomeCreateObjectV1ResponseMPayload {
	this := FranchisereferalincomeCreateObjectV1ResponseMPayload{}
	this.APkiFranchisereferalincomeID = aPkiFranchisereferalincomeID
	return &this
}

// NewFranchisereferalincomeCreateObjectV1ResponseMPayloadWithDefaults instantiates a new FranchisereferalincomeCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeCreateObjectV1ResponseMPayloadWithDefaults() *FranchisereferalincomeCreateObjectV1ResponseMPayload {
	this := FranchisereferalincomeCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiFranchisereferalincomeID returns the APkiFranchisereferalincomeID field value
func (o *FranchisereferalincomeCreateObjectV1ResponseMPayload) GetAPkiFranchisereferalincomeID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiFranchisereferalincomeID
}

// GetAPkiFranchisereferalincomeIDOk returns a tuple with the APkiFranchisereferalincomeID field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeCreateObjectV1ResponseMPayload) GetAPkiFranchisereferalincomeIDOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.APkiFranchisereferalincomeID, true
}

// SetAPkiFranchisereferalincomeID sets field value
func (o *FranchisereferalincomeCreateObjectV1ResponseMPayload) SetAPkiFranchisereferalincomeID(v []int32) {
	o.APkiFranchisereferalincomeID = v
}

func (o FranchisereferalincomeCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_pkiFranchisereferalincomeID"] = o.APkiFranchisereferalincomeID
	}
	return json.Marshal(toSerialize)
}

type NullableFranchisereferalincomeCreateObjectV1ResponseMPayload struct {
	value *FranchisereferalincomeCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) Get() *FranchisereferalincomeCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) Set(val *FranchisereferalincomeCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeCreateObjectV1ResponseMPayload(val *FranchisereferalincomeCreateObjectV1ResponseMPayload) *NullableFranchisereferalincomeCreateObjectV1ResponseMPayload {
	return &NullableFranchisereferalincomeCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


