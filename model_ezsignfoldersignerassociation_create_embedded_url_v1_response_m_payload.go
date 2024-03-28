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

// checks if the EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload{}

// EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload Payload for POST /1/object/ezsignfoldersignerassociation/createEmbeddedUrl
type EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload struct {
	// The embedded Url to the signing application.    Url will expire after 5 minutes.  
	SEmbeddedUrl string `json:"sEmbeddedUrl"`
}

type _EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload

// NewEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload(sEmbeddedUrl string) *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload{}
	this.SEmbeddedUrl = sEmbeddedUrl
	return &this
}

// NewEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayloadWithDefaults instantiates a new EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayloadWithDefaults() *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload {
	this := EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload{}
	return &this
}

// GetSEmbeddedUrl returns the SEmbeddedUrl field value
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) GetSEmbeddedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEmbeddedUrl
}

// GetSEmbeddedUrlOk returns a tuple with the SEmbeddedUrl field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) GetSEmbeddedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEmbeddedUrl, true
}

// SetSEmbeddedUrl sets field value
func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) SetSEmbeddedUrl(v string) {
	o.SEmbeddedUrl = v
}

func (o EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sEmbeddedUrl"] = o.SEmbeddedUrl
	return toSerialize, nil
}

func (o *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sEmbeddedUrl",
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

	varEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload := _EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload(varEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload)

	return err
}

type NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload struct {
	value *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) Get() *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) Set(val *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload(val *EzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload {
	return &NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationCreateEmbeddedUrlV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


