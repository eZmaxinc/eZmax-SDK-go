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

// checks if the EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload{}

// EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload Payload for POST /1/object/ezsignfolder/{pkiEzsignfolder}/importEzsignfoldersignerassociations
type EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload struct {
	APkiEzsignfoldersignerassociationID []int32 `json:"a_pkiEzsignfoldersignerassociationID"`
}

type _EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload

// NewEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload instantiates a new EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload(aPkiEzsignfoldersignerassociationID []int32) *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload{}
	this.APkiEzsignfoldersignerassociationID = aPkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayloadWithDefaults() *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload {
	this := EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload{}
	return &this
}

// GetAPkiEzsignfoldersignerassociationID returns the APkiEzsignfoldersignerassociationID field value
func (o *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationID() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.APkiEzsignfoldersignerassociationID
}

// GetAPkiEzsignfoldersignerassociationIDOk returns a tuple with the APkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) GetAPkiEzsignfoldersignerassociationIDOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.APkiEzsignfoldersignerassociationID, true
}

// SetAPkiEzsignfoldersignerassociationID sets field value
func (o *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) SetAPkiEzsignfoldersignerassociationID(v []int32) {
	o.APkiEzsignfoldersignerassociationID = v
}

func (o EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_pkiEzsignfoldersignerassociationID"] = o.APkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_pkiEzsignfoldersignerassociationID",
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

	varEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload := _EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload)

	if err != nil {
		return err
	}

	*o = EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload(varEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload)

	return err
}

type NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload struct {
	value *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) Get() *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) Set(val *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload(val *EzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) *NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload {
	return &NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderImportEzsignfoldersignerassociationsV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


