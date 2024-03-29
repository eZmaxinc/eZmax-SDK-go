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

// EzsigndocumentGetDownloadUrlV1ResponseMPayload Payload for the /1/object/ezsigndocument/{pkiEzsigndocument}/getDownloadUrl API Request
type EzsigndocumentGetDownloadUrlV1ResponseMPayload struct {
	// The Url to the requested document.  Url will expire after 5 minutes.
	SDownloadUrl string `json:"sDownloadUrl"`
}

// NewEzsigndocumentGetDownloadUrlV1ResponseMPayload instantiates a new EzsigndocumentGetDownloadUrlV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetDownloadUrlV1ResponseMPayload(sDownloadUrl string) *EzsigndocumentGetDownloadUrlV1ResponseMPayload {
	this := EzsigndocumentGetDownloadUrlV1ResponseMPayload{}
	this.SDownloadUrl = sDownloadUrl
	return &this
}

// NewEzsigndocumentGetDownloadUrlV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetDownloadUrlV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetDownloadUrlV1ResponseMPayloadWithDefaults() *EzsigndocumentGetDownloadUrlV1ResponseMPayload {
	this := EzsigndocumentGetDownloadUrlV1ResponseMPayload{}
	return &this
}

// GetSDownloadUrl returns the SDownloadUrl field value
func (o *EzsigndocumentGetDownloadUrlV1ResponseMPayload) GetSDownloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SDownloadUrl
}

// GetSDownloadUrlOk returns a tuple with the SDownloadUrl field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetDownloadUrlV1ResponseMPayload) GetSDownloadUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SDownloadUrl, true
}

// SetSDownloadUrl sets field value
func (o *EzsigndocumentGetDownloadUrlV1ResponseMPayload) SetSDownloadUrl(v string) {
	o.SDownloadUrl = v
}

func (o EzsigndocumentGetDownloadUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sDownloadUrl"] = o.SDownloadUrl
	}
	return json.Marshal(toSerialize)
}

type NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload struct {
	value *EzsigndocumentGetDownloadUrlV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) Get() *EzsigndocumentGetDownloadUrlV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) Set(val *EzsigndocumentGetDownloadUrlV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetDownloadUrlV1ResponseMPayload(val *EzsigndocumentGetDownloadUrlV1ResponseMPayload) *NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload {
	return &NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetDownloadUrlV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


