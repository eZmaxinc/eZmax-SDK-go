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

// checks if the EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload{}

// EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload Payload for GET /1/object/ezsignbulksendtransmission/{pkiEzsignbulksendtransmissionID}/getEzsignsignaturesAutomatic
type EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	// All eEzsignsignatureType contained in the response
	AEEzsignsignatureType []FieldEEzsignsignatureType `json:"a_eEzsignsignatureType"`
	AObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse `json:"a_objEzsignfolder"`
}

type _EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload

// NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload instantiates a new EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload(aEEzsignsignatureType []FieldEEzsignsignatureType, aObjEzsignfolder []CustomEzsignfolderEzsignsignaturesAutomaticResponse) *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	this.AEEzsignsignatureType = aEEzsignsignatureType
	this.AObjEzsignfolder = aObjEzsignfolder
	return &this
}

// NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults instantiates a new EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayloadWithDefaults() *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload {
	this := EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload{}
	return &this
}

// GetAEEzsignsignatureType returns the AEEzsignsignatureType field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureType() []FieldEEzsignsignatureType {
	if o == nil {
		var ret []FieldEEzsignsignatureType
		return ret
	}

	return o.AEEzsignsignatureType
}

// GetAEEzsignsignatureTypeOk returns a tuple with the AEEzsignsignatureType field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAEEzsignsignatureTypeOk() ([]FieldEEzsignsignatureType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AEEzsignsignatureType, true
}

// SetAEEzsignsignatureType sets field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAEEzsignsignatureType(v []FieldEEzsignsignatureType) {
	o.AEEzsignsignatureType = v
}

// GetAObjEzsignfolder returns the AObjEzsignfolder field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolder() []CustomEzsignfolderEzsignsignaturesAutomaticResponse {
	if o == nil {
		var ret []CustomEzsignfolderEzsignsignaturesAutomaticResponse
		return ret
	}

	return o.AObjEzsignfolder
}

// GetAObjEzsignfolderOk returns a tuple with the AObjEzsignfolder field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) GetAObjEzsignfolderOk() ([]CustomEzsignfolderEzsignsignaturesAutomaticResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfolder, true
}

// SetAObjEzsignfolder sets field value
func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) SetAObjEzsignfolder(v []CustomEzsignfolderEzsignsignaturesAutomaticResponse) {
	o.AObjEzsignfolder = v
}

func (o EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_eEzsignsignatureType"] = o.AEEzsignsignatureType
	toSerialize["a_objEzsignfolder"] = o.AObjEzsignfolder
	return toSerialize, nil
}

func (o *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
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

	varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload := _EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload(varEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload)

	return err
}

type NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload struct {
	value *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) Get() *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) Set(val *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload(val *EzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload {
	return &NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendtransmissionGetEzsignsignaturesAutomaticV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


