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

// checks if the EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}

// EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload Payload for GET /1/object/ezsignfolder/{pkiEzsignfolder}/getEzsignfoldersignerassociations
type EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload struct {
	AObjEzsignfoldersignerassociation []CustomEzsignfoldersignerassociationActionableElementResponse `json:"a_objEzsignfoldersignerassociation"`
}

type _EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload(aObjEzsignfoldersignerassociation []CustomEzsignfoldersignerassociationActionableElementResponse) *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}
	this.AObjEzsignfoldersignerassociation = aObjEzsignfoldersignerassociation
	return &this
}

// NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}
	return &this
}

// GetAObjEzsignfoldersignerassociation returns the AObjEzsignfoldersignerassociation field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) GetAObjEzsignfoldersignerassociation() []CustomEzsignfoldersignerassociationActionableElementResponse {
	if o == nil {
		var ret []CustomEzsignfoldersignerassociationActionableElementResponse
		return ret
	}

	return o.AObjEzsignfoldersignerassociation
}

// GetAObjEzsignfoldersignerassociationOk returns a tuple with the AObjEzsignfoldersignerassociation field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) GetAObjEzsignfoldersignerassociationOk() ([]CustomEzsignfoldersignerassociationActionableElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignfoldersignerassociation, true
}

// SetAObjEzsignfoldersignerassociation sets field value
func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) SetAObjEzsignfoldersignerassociation(v []CustomEzsignfoldersignerassociationActionableElementResponse) {
	o.AObjEzsignfoldersignerassociation = v
}

func (o EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignfoldersignerassociation"] = o.AObjEzsignfoldersignerassociation
	return toSerialize, nil
}

func (o *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignfoldersignerassociation",
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

	varEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload := _EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload(varEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload)

	return err
}

type NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload struct {
	value *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Get() *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Set(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload(val *EzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload {
	return &NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetEzsignfoldersignerassociationsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


