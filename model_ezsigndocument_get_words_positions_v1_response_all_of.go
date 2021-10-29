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

// EzsigndocumentGetWordsPositionsV1ResponseAllOf struct for EzsigndocumentGetWordsPositionsV1ResponseAllOf
type EzsigndocumentGetWordsPositionsV1ResponseAllOf struct {
	// Payload for the /1/object/ezsigndocument/{pkiEzsigndocumentID}/getWordsPositions API Request
	MPayload []CustomWordPositionWordResponse `json:"mPayload"`
}

// NewEzsigndocumentGetWordsPositionsV1ResponseAllOf instantiates a new EzsigndocumentGetWordsPositionsV1ResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetWordsPositionsV1ResponseAllOf(mPayload []CustomWordPositionWordResponse) *EzsigndocumentGetWordsPositionsV1ResponseAllOf {
	this := EzsigndocumentGetWordsPositionsV1ResponseAllOf{}
	this.MPayload = mPayload
	return &this
}

// NewEzsigndocumentGetWordsPositionsV1ResponseAllOfWithDefaults instantiates a new EzsigndocumentGetWordsPositionsV1ResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetWordsPositionsV1ResponseAllOfWithDefaults() *EzsigndocumentGetWordsPositionsV1ResponseAllOf {
	this := EzsigndocumentGetWordsPositionsV1ResponseAllOf{}
	return &this
}

// GetMPayload returns the MPayload field value
func (o *EzsigndocumentGetWordsPositionsV1ResponseAllOf) GetMPayload() []CustomWordPositionWordResponse {
	if o == nil {
		var ret []CustomWordPositionWordResponse
		return ret
	}

	return o.MPayload
}

// GetMPayloadOk returns a tuple with the MPayload field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetWordsPositionsV1ResponseAllOf) GetMPayloadOk() (*[]CustomWordPositionWordResponse, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MPayload, true
}

// SetMPayload sets field value
func (o *EzsigndocumentGetWordsPositionsV1ResponseAllOf) SetMPayload(v []CustomWordPositionWordResponse) {
	o.MPayload = v
}

func (o EzsigndocumentGetWordsPositionsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mPayload"] = o.MPayload
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf struct {
	value *EzsigndocumentGetWordsPositionsV1ResponseAllOf
	isSet bool
}

func (v NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) Get() *EzsigndocumentGetWordsPositionsV1ResponseAllOf {
	return v.value
}

func (v *NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) Set(val *EzsigndocumentGetWordsPositionsV1ResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetWordsPositionsV1ResponseAllOf(val *EzsigndocumentGetWordsPositionsV1ResponseAllOf) *NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf {
	return &NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetWordsPositionsV1ResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


