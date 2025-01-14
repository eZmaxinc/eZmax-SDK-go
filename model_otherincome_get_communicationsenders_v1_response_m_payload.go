/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OtherincomeGetCommunicationsendersV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OtherincomeGetCommunicationsendersV1ResponseMPayload{}

// OtherincomeGetCommunicationsendersV1ResponseMPayload Response for GET /1/object/otherincome/{pkiOtherincomeID}/getCommunicationsenders
type OtherincomeGetCommunicationsendersV1ResponseMPayload struct {
	AObjCommunicationsenders []CustomCommunicationsenderResponse `json:"a_objCommunicationsenders"`
}

type _OtherincomeGetCommunicationsendersV1ResponseMPayload OtherincomeGetCommunicationsendersV1ResponseMPayload

// NewOtherincomeGetCommunicationsendersV1ResponseMPayload instantiates a new OtherincomeGetCommunicationsendersV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOtherincomeGetCommunicationsendersV1ResponseMPayload(aObjCommunicationsenders []CustomCommunicationsenderResponse) *OtherincomeGetCommunicationsendersV1ResponseMPayload {
	this := OtherincomeGetCommunicationsendersV1ResponseMPayload{}
	this.AObjCommunicationsenders = aObjCommunicationsenders
	return &this
}

// NewOtherincomeGetCommunicationsendersV1ResponseMPayloadWithDefaults instantiates a new OtherincomeGetCommunicationsendersV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOtherincomeGetCommunicationsendersV1ResponseMPayloadWithDefaults() *OtherincomeGetCommunicationsendersV1ResponseMPayload {
	this := OtherincomeGetCommunicationsendersV1ResponseMPayload{}
	return &this
}

// GetAObjCommunicationsenders returns the AObjCommunicationsenders field value
func (o *OtherincomeGetCommunicationsendersV1ResponseMPayload) GetAObjCommunicationsenders() []CustomCommunicationsenderResponse {
	if o == nil {
		var ret []CustomCommunicationsenderResponse
		return ret
	}

	return o.AObjCommunicationsenders
}

// GetAObjCommunicationsendersOk returns a tuple with the AObjCommunicationsenders field value
// and a boolean to check if the value has been set.
func (o *OtherincomeGetCommunicationsendersV1ResponseMPayload) GetAObjCommunicationsendersOk() ([]CustomCommunicationsenderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunicationsenders, true
}

// SetAObjCommunicationsenders sets field value
func (o *OtherincomeGetCommunicationsendersV1ResponseMPayload) SetAObjCommunicationsenders(v []CustomCommunicationsenderResponse) {
	o.AObjCommunicationsenders = v
}

func (o OtherincomeGetCommunicationsendersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OtherincomeGetCommunicationsendersV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCommunicationsenders"] = o.AObjCommunicationsenders
	return toSerialize, nil
}

func (o *OtherincomeGetCommunicationsendersV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objCommunicationsenders",
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

	varOtherincomeGetCommunicationsendersV1ResponseMPayload := _OtherincomeGetCommunicationsendersV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOtherincomeGetCommunicationsendersV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = OtherincomeGetCommunicationsendersV1ResponseMPayload(varOtherincomeGetCommunicationsendersV1ResponseMPayload)

	return err
}

type NullableOtherincomeGetCommunicationsendersV1ResponseMPayload struct {
	value *OtherincomeGetCommunicationsendersV1ResponseMPayload
	isSet bool
}

func (v NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) Get() *OtherincomeGetCommunicationsendersV1ResponseMPayload {
	return v.value
}

func (v *NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) Set(val *OtherincomeGetCommunicationsendersV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOtherincomeGetCommunicationsendersV1ResponseMPayload(val *OtherincomeGetCommunicationsendersV1ResponseMPayload) *NullableOtherincomeGetCommunicationsendersV1ResponseMPayload {
	return &NullableOtherincomeGetCommunicationsendersV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOtherincomeGetCommunicationsendersV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


