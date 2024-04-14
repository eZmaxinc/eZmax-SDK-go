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

// checks if the EzsignSuggestSignersV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignSuggestSignersV1ResponseMPayload{}

// EzsignSuggestSignersV1ResponseMPayload Payload for GET /1/module/ezsign/suggestSigners
type EzsignSuggestSignersV1ResponseMPayload struct {
	AObjEzsignfoldersignerassociation []EzsignfoldersignerassociationResponseCompound `json:"a_objEzsignfoldersignerassociation"`
	AObjUserTeam []CustomUserResponse `json:"a_objUserTeam"`
	AObjUser []CustomUserResponse `json:"a_objUser"`
}

type _EzsignSuggestSignersV1ResponseMPayload EzsignSuggestSignersV1ResponseMPayload

// NewEzsignSuggestSignersV1ResponseMPayload instantiates a new EzsignSuggestSignersV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignSuggestSignersV1ResponseMPayload(aObjEzsignfoldersignerassociation []EzsignfoldersignerassociationResponseCompound, aObjUserTeam []CustomUserResponse, aObjUser []CustomUserResponse) *EzsignSuggestSignersV1ResponseMPayload {
	this := EzsignSuggestSignersV1ResponseMPayload{}
	this.AObjEzsignfoldersignerassociation = aObjEzsignfoldersignerassociation
	this.AObjUserTeam = aObjUserTeam
	this.AObjUser = aObjUser
	return &this
}

// NewEzsignSuggestSignersV1ResponseMPayloadWithDefaults instantiates a new EzsignSuggestSignersV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignSuggestSignersV1ResponseMPayloadWithDefaults() *EzsignSuggestSignersV1ResponseMPayload {
	this := EzsignSuggestSignersV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignfoldersignerassociation returns the AObjEzsignfoldersignerassociation field value
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjEzsignfoldersignerassociation() []EzsignfoldersignerassociationResponseCompound {
	if o == nil {
		var ret []EzsignfoldersignerassociationResponseCompound
		return ret
	}

	return o.AObjEzsignfoldersignerassociation
}

// GetAObjEzsignfoldersignerassociationOk returns a tuple with the AObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjEzsignfoldersignerassociationOk() ([]EzsignfoldersignerassociationResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfoldersignerassociation, true
}

// SetAObjEzsignfoldersignerassociation sets field value
func (o *EzsignSuggestSignersV1ResponseMPayload) SetAObjEzsignfoldersignerassociation(v []EzsignfoldersignerassociationResponseCompound) {
	o.AObjEzsignfoldersignerassociation = v
}

// GetAObjUserTeam returns the AObjUserTeam field value
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjUserTeam() []CustomUserResponse {
	if o == nil {
		var ret []CustomUserResponse
		return ret
	}

	return o.AObjUserTeam
}

// GetAObjUserTeamOk returns a tuple with the AObjUserTeam field value
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjUserTeamOk() ([]CustomUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUserTeam, true
}

// SetAObjUserTeam sets field value
func (o *EzsignSuggestSignersV1ResponseMPayload) SetAObjUserTeam(v []CustomUserResponse) {
	o.AObjUserTeam = v
}

// GetAObjUser returns the AObjUser field value
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjUser() []CustomUserResponse {
	if o == nil {
		var ret []CustomUserResponse
		return ret
	}

	return o.AObjUser
}

// GetAObjUserOk returns a tuple with the AObjUser field value
// and a boolean to check if the value has been set.
func (o *EzsignSuggestSignersV1ResponseMPayload) GetAObjUserOk() ([]CustomUserResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjUser, true
}

// SetAObjUser sets field value
func (o *EzsignSuggestSignersV1ResponseMPayload) SetAObjUser(v []CustomUserResponse) {
	o.AObjUser = v
}

func (o EzsignSuggestSignersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignSuggestSignersV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignfoldersignerassociation"] = o.AObjEzsignfoldersignerassociation
	toSerialize["a_objUserTeam"] = o.AObjUserTeam
	toSerialize["a_objUser"] = o.AObjUser
	return toSerialize, nil
}

func (o *EzsignSuggestSignersV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignfoldersignerassociation",
		"a_objUserTeam",
		"a_objUser",
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

	varEzsignSuggestSignersV1ResponseMPayload := _EzsignSuggestSignersV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignSuggestSignersV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignSuggestSignersV1ResponseMPayload(varEzsignSuggestSignersV1ResponseMPayload)

	return err
}

type NullableEzsignSuggestSignersV1ResponseMPayload struct {
	value *EzsignSuggestSignersV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignSuggestSignersV1ResponseMPayload) Get() *EzsignSuggestSignersV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignSuggestSignersV1ResponseMPayload) Set(val *EzsignSuggestSignersV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignSuggestSignersV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignSuggestSignersV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignSuggestSignersV1ResponseMPayload(val *EzsignSuggestSignersV1ResponseMPayload) *NullableEzsignSuggestSignersV1ResponseMPayload {
	return &NullableEzsignSuggestSignersV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignSuggestSignersV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignSuggestSignersV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


