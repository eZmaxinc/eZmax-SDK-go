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
	"bytes"
	"fmt"
)

// checks if the EzsigndocumentGetDownloadUrlV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetDownloadUrlV1ResponseMPayload{}

// EzsigndocumentGetDownloadUrlV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getDownloadUrl
type EzsigndocumentGetDownloadUrlV1ResponseMPayload struct {
	// The Url to the requested document.  Url will expire after 5 minutes.
	SDownloadUrl string `json:"sDownloadUrl"`
}

type _EzsigndocumentGetDownloadUrlV1ResponseMPayload EzsigndocumentGetDownloadUrlV1ResponseMPayload

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
	if o == nil {
		return nil, false
	}
	return &o.SDownloadUrl, true
}

// SetSDownloadUrl sets field value
func (o *EzsigndocumentGetDownloadUrlV1ResponseMPayload) SetSDownloadUrl(v string) {
	o.SDownloadUrl = v
}

func (o EzsigndocumentGetDownloadUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetDownloadUrlV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sDownloadUrl"] = o.SDownloadUrl
	return toSerialize, nil
}

func (o *EzsigndocumentGetDownloadUrlV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sDownloadUrl",
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

	varEzsigndocumentGetDownloadUrlV1ResponseMPayload := _EzsigndocumentGetDownloadUrlV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetDownloadUrlV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetDownloadUrlV1ResponseMPayload(varEzsigndocumentGetDownloadUrlV1ResponseMPayload)

	return err
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


