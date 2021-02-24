/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.31
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxinc/eZmax-SDK-go

import (
	"encoding/json"
)

// EzsignfolderCreateObjectV1ResponseMPayload Payload for the /1/object/ezsignfolder/createObject API Request
type EzsignfolderCreateObjectV1ResponseMPayload struct {
	// An array of unique IDs representing the object that were requested to be created.  They are returned in the same order as the array containing the objects to be created that was sent in the request.
	APkiEzsignfolderID []int32 `json:"a_pkiEzsignfolderID"`
}

// NewEzsignfolderCreateObjectV1ResponseMPayload instantiates a new EzsignfolderCreateObjectV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderCreateObjectV1ResponseMPayload(aPkiEzsignfolderID []int32) *EzsignfolderCreateObjectV1ResponseMPayload {
	this := EzsignfolderCreateObjectV1ResponseMPayload{}
	this.APkiEzsignfolderID = aPkiEzsignfolderID
	return &this
}

// NewEzsignfolderCreateObjectV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderCreateObjectV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderCreateObjectV1ResponseMPayloadWithDefaults() *EzsignfolderCreateObjectV1ResponseMPayload {
	this := EzsignfolderCreateObjectV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignfolderID returns the APkiEzsignfolderID field value
func (o *EzsignfolderCreateObjectV1ResponseMPayload) GetAPkiEzsignfolderID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfolderID
}

// GetAPkiEzsignfolderIDOk returns a tuple with the APkiEzsignfolderID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderCreateObjectV1ResponseMPayload) GetAPkiEzsignfolderIDOk() (*[]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.APkiEzsignfolderID, true
}

// SetAPkiEzsignfolderID sets field value
func (o *EzsignfolderCreateObjectV1ResponseMPayload) SetAPkiEzsignfolderID(v []int32) {
	o.APkiEzsignfolderID = v
}

func (o EzsignfolderCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a_pkiEzsignfolderID"] = o.APkiEzsignfolderID
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfolderCreateObjectV1ResponseMPayload struct {
	value *EzsignfolderCreateObjectV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderCreateObjectV1ResponseMPayload) Get() *EzsignfolderCreateObjectV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderCreateObjectV1ResponseMPayload) Set(val *EzsignfolderCreateObjectV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderCreateObjectV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderCreateObjectV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderCreateObjectV1ResponseMPayload(val *EzsignfolderCreateObjectV1ResponseMPayload) *NullableEzsignfolderCreateObjectV1ResponseMPayload {
	return &NullableEzsignfolderCreateObjectV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderCreateObjectV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderCreateObjectV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


