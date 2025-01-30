/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EzsignfolderGetAttachmentCountV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderGetAttachmentCountV1ResponseMPayload{}

// EzsignfolderGetAttachmentCountV1ResponseMPayload Response for GET /1/object/ezsignfolder/{pkiEzsignfolderID}/getAttachmentCount
type EzsignfolderGetAttachmentCountV1ResponseMPayload struct {
	// The count of Attachment.
	IAttachmentCount int32 `json:"iAttachmentCount"`
}

type _EzsignfolderGetAttachmentCountV1ResponseMPayload EzsignfolderGetAttachmentCountV1ResponseMPayload

// NewEzsignfolderGetAttachmentCountV1ResponseMPayload instantiates a new EzsignfolderGetAttachmentCountV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetAttachmentCountV1ResponseMPayload(iAttachmentCount int32) *EzsignfolderGetAttachmentCountV1ResponseMPayload {
	this := EzsignfolderGetAttachmentCountV1ResponseMPayload{}
	this.IAttachmentCount = iAttachmentCount
	return &this
}

// NewEzsignfolderGetAttachmentCountV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetAttachmentCountV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetAttachmentCountV1ResponseMPayloadWithDefaults() *EzsignfolderGetAttachmentCountV1ResponseMPayload {
	this := EzsignfolderGetAttachmentCountV1ResponseMPayload{}
	return &this
}

// GetIAttachmentCount returns the IAttachmentCount field value
func (o *EzsignfolderGetAttachmentCountV1ResponseMPayload) GetIAttachmentCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IAttachmentCount
}

// GetIAttachmentCountOk returns a tuple with the IAttachmentCount field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetAttachmentCountV1ResponseMPayload) GetIAttachmentCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IAttachmentCount, true
}

// SetIAttachmentCount sets field value
func (o *EzsignfolderGetAttachmentCountV1ResponseMPayload) SetIAttachmentCount(v int32) {
	o.IAttachmentCount = v
}

func (o EzsignfolderGetAttachmentCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderGetAttachmentCountV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iAttachmentCount"] = o.IAttachmentCount
	return toSerialize, nil
}

func (o *EzsignfolderGetAttachmentCountV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iAttachmentCount",
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

	varEzsignfolderGetAttachmentCountV1ResponseMPayload := _EzsignfolderGetAttachmentCountV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderGetAttachmentCountV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfolderGetAttachmentCountV1ResponseMPayload(varEzsignfolderGetAttachmentCountV1ResponseMPayload)

	return err
}

type NullableEzsignfolderGetAttachmentCountV1ResponseMPayload struct {
	value *EzsignfolderGetAttachmentCountV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) Get() *EzsignfolderGetAttachmentCountV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) Set(val *EzsignfolderGetAttachmentCountV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetAttachmentCountV1ResponseMPayload(val *EzsignfolderGetAttachmentCountV1ResponseMPayload) *NullableEzsignfolderGetAttachmentCountV1ResponseMPayload {
	return &NullableEzsignfolderGetAttachmentCountV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetAttachmentCountV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


