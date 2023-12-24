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

// checks if the EzsigndocumentGetAttachmentsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetAttachmentsV1ResponseMPayload{}

// EzsigndocumentGetAttachmentsV1ResponseMPayload Response for GET /1/object/ezsigndocument/{pkiEzsigndocumentID}/getAttachments
type EzsigndocumentGetAttachmentsV1ResponseMPayload struct {
	AObjAttachmentdocumenttype []CustomAttachmentdocumenttypeResponse `json:"a_objAttachmentdocumenttype"`
}

type _EzsigndocumentGetAttachmentsV1ResponseMPayload EzsigndocumentGetAttachmentsV1ResponseMPayload

// NewEzsigndocumentGetAttachmentsV1ResponseMPayload instantiates a new EzsigndocumentGetAttachmentsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetAttachmentsV1ResponseMPayload(aObjAttachmentdocumenttype []CustomAttachmentdocumenttypeResponse) *EzsigndocumentGetAttachmentsV1ResponseMPayload {
	this := EzsigndocumentGetAttachmentsV1ResponseMPayload{}
	this.AObjAttachmentdocumenttype = aObjAttachmentdocumenttype
	return &this
}

// NewEzsigndocumentGetAttachmentsV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetAttachmentsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetAttachmentsV1ResponseMPayloadWithDefaults() *EzsigndocumentGetAttachmentsV1ResponseMPayload {
	this := EzsigndocumentGetAttachmentsV1ResponseMPayload{}
	return &this
}

// GetAObjAttachmentdocumenttype returns the AObjAttachmentdocumenttype field value
func (o *EzsigndocumentGetAttachmentsV1ResponseMPayload) GetAObjAttachmentdocumenttype() []CustomAttachmentdocumenttypeResponse {
	if o == nil {
		var ret []CustomAttachmentdocumenttypeResponse
		return ret
	}

	return o.AObjAttachmentdocumenttype
}

// GetAObjAttachmentdocumenttypeOk returns a tuple with the AObjAttachmentdocumenttype field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetAttachmentsV1ResponseMPayload) GetAObjAttachmentdocumenttypeOk() ([]CustomAttachmentdocumenttypeResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjAttachmentdocumenttype, true
}

// SetAObjAttachmentdocumenttype sets field value
func (o *EzsigndocumentGetAttachmentsV1ResponseMPayload) SetAObjAttachmentdocumenttype(v []CustomAttachmentdocumenttypeResponse) {
	o.AObjAttachmentdocumenttype = v
}

func (o EzsigndocumentGetAttachmentsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetAttachmentsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objAttachmentdocumenttype"] = o.AObjAttachmentdocumenttype
	return toSerialize, nil
}

func (o *EzsigndocumentGetAttachmentsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objAttachmentdocumenttype",
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

	varEzsigndocumentGetAttachmentsV1ResponseMPayload := _EzsigndocumentGetAttachmentsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetAttachmentsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetAttachmentsV1ResponseMPayload(varEzsigndocumentGetAttachmentsV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentGetAttachmentsV1ResponseMPayload struct {
	value *EzsigndocumentGetAttachmentsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) Get() *EzsigndocumentGetAttachmentsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) Set(val *EzsigndocumentGetAttachmentsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetAttachmentsV1ResponseMPayload(val *EzsigndocumentGetAttachmentsV1ResponseMPayload) *NullableEzsigndocumentGetAttachmentsV1ResponseMPayload {
	return &NullableEzsigndocumentGetAttachmentsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetAttachmentsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


