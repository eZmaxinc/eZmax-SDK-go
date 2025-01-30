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

// checks if the EzsignfolderGetCommunicationsendersV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderGetCommunicationsendersV1ResponseMPayload{}

// EzsignfolderGetCommunicationsendersV1ResponseMPayload Response for GET /1/object/ezsignfolder/{pkiEzsignfolderID}/getCommunicationsenders
type EzsignfolderGetCommunicationsendersV1ResponseMPayload struct {
	AObjCommunicationsenders []CustomCommunicationsenderResponse `json:"a_objCommunicationsenders"`
}

type _EzsignfolderGetCommunicationsendersV1ResponseMPayload EzsignfolderGetCommunicationsendersV1ResponseMPayload

// NewEzsignfolderGetCommunicationsendersV1ResponseMPayload instantiates a new EzsignfolderGetCommunicationsendersV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetCommunicationsendersV1ResponseMPayload(aObjCommunicationsenders []CustomCommunicationsenderResponse) *EzsignfolderGetCommunicationsendersV1ResponseMPayload {
	this := EzsignfolderGetCommunicationsendersV1ResponseMPayload{}
	this.AObjCommunicationsenders = aObjCommunicationsenders
	return &this
}

// NewEzsignfolderGetCommunicationsendersV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetCommunicationsendersV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetCommunicationsendersV1ResponseMPayloadWithDefaults() *EzsignfolderGetCommunicationsendersV1ResponseMPayload {
	this := EzsignfolderGetCommunicationsendersV1ResponseMPayload{}
	return &this
}

// GetAObjCommunicationsenders returns the AObjCommunicationsenders field value
func (o *EzsignfolderGetCommunicationsendersV1ResponseMPayload) GetAObjCommunicationsenders() []CustomCommunicationsenderResponse {
	if o == nil {
		var ret []CustomCommunicationsenderResponse
		return ret
	}

	return o.AObjCommunicationsenders
}

// GetAObjCommunicationsendersOk returns a tuple with the AObjCommunicationsenders field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetCommunicationsendersV1ResponseMPayload) GetAObjCommunicationsendersOk() ([]CustomCommunicationsenderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCommunicationsenders, true
}

// SetAObjCommunicationsenders sets field value
func (o *EzsignfolderGetCommunicationsendersV1ResponseMPayload) SetAObjCommunicationsenders(v []CustomCommunicationsenderResponse) {
	o.AObjCommunicationsenders = v
}

func (o EzsignfolderGetCommunicationsendersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderGetCommunicationsendersV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCommunicationsenders"] = o.AObjCommunicationsenders
	return toSerialize, nil
}

func (o *EzsignfolderGetCommunicationsendersV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignfolderGetCommunicationsendersV1ResponseMPayload := _EzsignfolderGetCommunicationsendersV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderGetCommunicationsendersV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfolderGetCommunicationsendersV1ResponseMPayload(varEzsignfolderGetCommunicationsendersV1ResponseMPayload)

	return err
}

type NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload struct {
	value *EzsignfolderGetCommunicationsendersV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) Get() *EzsignfolderGetCommunicationsendersV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) Set(val *EzsignfolderGetCommunicationsendersV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetCommunicationsendersV1ResponseMPayload(val *EzsignfolderGetCommunicationsendersV1ResponseMPayload) *NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload {
	return &NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetCommunicationsendersV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


