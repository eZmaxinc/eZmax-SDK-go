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

// checks if the EzsigndocumentGetFormDataV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigndocumentGetFormDataV1ResponseMPayload{}

// EzsigndocumentGetFormDataV1ResponseMPayload Payload for GET /1/object/ezsigndocument/{pkiEzsigndocument}/getFormData
type EzsigndocumentGetFormDataV1ResponseMPayload struct {
	ObjFormDataDocument CustomFormDataDocumentResponse `json:"objFormDataDocument"`
}

type _EzsigndocumentGetFormDataV1ResponseMPayload EzsigndocumentGetFormDataV1ResponseMPayload

// NewEzsigndocumentGetFormDataV1ResponseMPayload instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigndocumentGetFormDataV1ResponseMPayload(objFormDataDocument CustomFormDataDocumentResponse) *EzsigndocumentGetFormDataV1ResponseMPayload {
	this := EzsigndocumentGetFormDataV1ResponseMPayload{}
	this.ObjFormDataDocument = objFormDataDocument
	return &this
}

// NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults instantiates a new EzsigndocumentGetFormDataV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigndocumentGetFormDataV1ResponseMPayloadWithDefaults() *EzsigndocumentGetFormDataV1ResponseMPayload {
	this := EzsigndocumentGetFormDataV1ResponseMPayload{}
	return &this
}

// GetObjFormDataDocument returns the ObjFormDataDocument field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetObjFormDataDocument() CustomFormDataDocumentResponse {
	if o == nil {
		var ret CustomFormDataDocumentResponse
		return ret
	}

	return o.ObjFormDataDocument
}

// GetObjFormDataDocumentOk returns a tuple with the ObjFormDataDocument field value
// and a boolean to check if the value has been set.
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) GetObjFormDataDocumentOk() (*CustomFormDataDocumentResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjFormDataDocument, true
}

// SetObjFormDataDocument sets field value
func (o *EzsigndocumentGetFormDataV1ResponseMPayload) SetObjFormDataDocument(v CustomFormDataDocumentResponse) {
	o.ObjFormDataDocument = v
}

func (o EzsigndocumentGetFormDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigndocumentGetFormDataV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objFormDataDocument"] = o.ObjFormDataDocument
	return toSerialize, nil
}

func (o *EzsigndocumentGetFormDataV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objFormDataDocument",
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

	varEzsigndocumentGetFormDataV1ResponseMPayload := _EzsigndocumentGetFormDataV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigndocumentGetFormDataV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsigndocumentGetFormDataV1ResponseMPayload(varEzsigndocumentGetFormDataV1ResponseMPayload)

	return err
}

type NullableEzsigndocumentGetFormDataV1ResponseMPayload struct {
	value *EzsigndocumentGetFormDataV1ResponseMPayload
	isSet bool
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) Get() *EzsigndocumentGetFormDataV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) Set(val *EzsigndocumentGetFormDataV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigndocumentGetFormDataV1ResponseMPayload(val *EzsigndocumentGetFormDataV1ResponseMPayload) *NullableEzsigndocumentGetFormDataV1ResponseMPayload {
	return &NullableEzsigndocumentGetFormDataV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsigndocumentGetFormDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigndocumentGetFormDataV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


