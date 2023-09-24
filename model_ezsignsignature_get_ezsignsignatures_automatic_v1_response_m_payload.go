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
)

// checks if the EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload{}

// EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload Payload for GET /1/object/ezsigsignature/getEzsignsignaturesAutomatic
type EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	// All eEzsignsignatureType contained in the response
	AEEzsignsignatureType []FieldEEzsignsignatureType `json:"a_eEzsignsignatureType"`
	AObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse `json:"a_objEzsignfolder"`
}

// NewEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload instantiates a new EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload(aEEzsignsignatureType []FieldEEzsignsignatureType, aObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse) *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	this.AEEzsignsignatureType = aEEzsignsignatureType
	this.AObjEzsignfolder = aObjEzsignfolder
	return &this
}

// NewEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults instantiates a new EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults() *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	return &this
}

// GetAEEzsignsignatureType returns the AEEzsignsignatureType field value
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureType() []FieldEEzsignsignatureType {
	if o == nil {
		var ret []FieldEEzsignsignatureType
		return ret
	}

	return o.AEEzsignsignatureType
}

// GetAEEzsignsignatureTypeOk returns a tuple with the AEEzsignsignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureTypeOk() ([]FieldEEzsignsignatureType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AEEzsignsignatureType, true
}

// SetAEEzsignsignatureType sets field value
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAEEzsignsignatureType(v []FieldEEzsignsignatureType) {
	o.AEEzsignsignatureType = v
}

// GetAObjEzsignfolder returns the AObjEzsignfolder field value
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolder() []CustomEzsignfolderEzsignsignaturesAutomaticResponse {
	if o == nil {
		var ret []CustomEzsignfolderEzsignsignaturesAutomaticResponse
		return ret
	}

	return o.AObjEzsignfolder
}

// GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolderOk() ([]CustomEzsignfolderEzsignsignaturesAutomaticResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfolder, true
}

// SetAObjEzsignfolder sets field value
func (o *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAObjEzsignfolder(v []CustomEzsignfolderEzsignsignaturesAutomaticResponse) {
	o.AObjEzsignfolder = v
}

func (o EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_eEzsignsignatureType"] = o.AEEzsignsignatureType
	toSerialize["a_objEzsignfolder"] = o.AObjEzsignfolder
	return toSerialize, nil
}

type NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	value *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) Get() *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) Set(val *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload(val *EzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) *NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return &NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignatureGetEzsignsignaturesAutomaticV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


