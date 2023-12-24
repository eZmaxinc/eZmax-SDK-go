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

// checks if the EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload{}

// EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload Payload for GET /1/object/ezsignbulksend/{pkiEzsignbulksendID}/getEzsignsignaturesAutomatic
type EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	// All eEzsignsignatureType contained in the response
	AEEzsignsignatureType []FieldEEzsignsignatureType `json:"a_eEzsignsignatureType"`
	AObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse `json:"a_objEzsignfolder"`
}

type _EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload

// NewEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload instantiates a new EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload(aEEzsignsignatureType []FieldEEzsignsignatureType, aObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse) *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	this.AEEzsignsignatureType = aEEzsignsignatureType
	this.AObjEzsignfolder = aObjEzsignfolder
	return &this
}

// NewEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults() *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	return &this
}

// GetAEEzsignsignatureType returns the AEEzsignsignatureType field value
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureType() []FieldEEzsignsignatureType {
	if o == nil {
		var ret []FieldEEzsignsignatureType
		return ret
	}

	return o.AEEzsignsignatureType
}

// GetAEEzsignsignatureTypeOk returns a tuple with the AEEzsignsignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureTypeOk() ([]FieldEEzsignsignatureType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AEEzsignsignatureType, true
}

// SetAEEzsignsignatureType sets field value
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAEEzsignsignatureType(v []FieldEEzsignsignatureType) {
	o.AEEzsignsignatureType = v
}

// GetAObjEzsignfolder returns the AObjEzsignfolder field value
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolder() []CustomEzsignfolderEzsignsignaturesAutomaticResponse {
	if o == nil {
		var ret []CustomEzsignfolderEzsignsignaturesAutomaticResponse
		return ret
	}

	return o.AObjEzsignfolder
}

// GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolderOk() ([]CustomEzsignfolderEzsignsignaturesAutomaticResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfolder, true
}

// SetAObjEzsignfolder sets field value
func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAObjEzsignfolder(v []CustomEzsignfolderEzsignsignaturesAutomaticResponse) {
	o.AObjEzsignfolder = v
}

func (o EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_eEzsignsignatureType"] = o.AEEzsignsignatureType
	toSerialize["a_objEzsignfolder"] = o.AObjEzsignfolder
	return toSerialize, nil
}

func (o *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_eEzsignsignatureType",
		"a_objEzsignfolder",
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

	varEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload := _EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload(varEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload)

	return err
}

type NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	value *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) Get() *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) Set(val *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload(val *EzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) *NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return &NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendGetEzsignsignaturesAutomaticV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


