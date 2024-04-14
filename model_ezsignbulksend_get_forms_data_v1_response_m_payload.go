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

// checks if the EzsignbulksendGetFormsDataV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendGetFormsDataV1ResponseMPayload{}

// EzsignbulksendGetFormsDataV1ResponseMPayload Payload for GET/1/object/ezsignbulksend/{pkiEzsignbulksendID}/getFormsData
type EzsignbulksendGetFormsDataV1ResponseMPayload struct {
	AObjFormsDataFolder []CustomFormsDataFolderResponse `json:"a_objFormsDataFolder"`
}

type _EzsignbulksendGetFormsDataV1ResponseMPayload EzsignbulksendGetFormsDataV1ResponseMPayload

// NewEzsignbulksendGetFormsDataV1ResponseMPayload instantiates a new EzsignbulksendGetFormsDataV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendGetFormsDataV1ResponseMPayload(aObjFormsDataFolder []CustomFormsDataFolderResponse) *EzsignbulksendGetFormsDataV1ResponseMPayload {
	this := EzsignbulksendGetFormsDataV1ResponseMPayload{}
	this.AObjFormsDataFolder = aObjFormsDataFolder
	return &this
}

// NewEzsignbulksendGetFormsDataV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendGetFormsDataV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendGetFormsDataV1ResponseMPayloadWithDefaults() *EzsignbulksendGetFormsDataV1ResponseMPayload {
	this := EzsignbulksendGetFormsDataV1ResponseMPayload{}
	return &this
}

// GetAObjFormsDataFolder returns the AObjFormsDataFolder field value
func (o *EzsignbulksendGetFormsDataV1ResponseMPayload) GetAObjFormsDataFolder() []CustomFormsDataFolderResponse {
	if o == nil {
		var ret []CustomFormsDataFolderResponse
		return ret
	}

	return o.AObjFormsDataFolder
}

// GetAObjFormsDataFolderOk returns a tuple with the AObjFormsDataFolder field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetFormsDataV1ResponseMPayload) GetAObjFormsDataFolderOk() ([]CustomFormsDataFolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjFormsDataFolder, true
}

// SetAObjFormsDataFolder sets field value
func (o *EzsignbulksendGetFormsDataV1ResponseMPayload) SetAObjFormsDataFolder(v []CustomFormsDataFolderResponse) {
	o.AObjFormsDataFolder = v
}

func (o EzsignbulksendGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendGetFormsDataV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objFormsDataFolder"] = o.AObjFormsDataFolder
	return toSerialize, nil
}

func (o *EzsignbulksendGetFormsDataV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objFormsDataFolder",
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

	varEzsignbulksendGetFormsDataV1ResponseMPayload := _EzsignbulksendGetFormsDataV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendGetFormsDataV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignbulksendGetFormsDataV1ResponseMPayload(varEzsignbulksendGetFormsDataV1ResponseMPayload)

	return err
}

type NullableEzsignbulksendGetFormsDataV1ResponseMPayload struct {
	value *EzsignbulksendGetFormsDataV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendGetFormsDataV1ResponseMPayload) Get() *EzsignbulksendGetFormsDataV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendGetFormsDataV1ResponseMPayload) Set(val *EzsignbulksendGetFormsDataV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendGetFormsDataV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendGetFormsDataV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendGetFormsDataV1ResponseMPayload(val *EzsignbulksendGetFormsDataV1ResponseMPayload) *NullableEzsignbulksendGetFormsDataV1ResponseMPayload {
	return &NullableEzsignbulksendGetFormsDataV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendGetFormsDataV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


