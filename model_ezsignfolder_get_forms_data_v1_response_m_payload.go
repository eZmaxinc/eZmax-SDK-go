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

// checks if the EzsignfolderGetFormsDataV1ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignfolderGetFormsDataV1ResponseMPayload{}

// EzsignfolderGetFormsDataV1ResponseMPayload Payload for GET /1/object/ezsignfolder/{pkiEzsigndocument}/getFormsData
type EzsignfolderGetFormsDataV1ResponseMPayload struct {
	ObjFormsDataFolder CustomFormsDataFolderResponse `json:"objFormsDataFolder"`
}

// NewEzsignfolderGetFormsDataV1ResponseMPayload instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfolderGetFormsDataV1ResponseMPayload(objFormsDataFolder CustomFormsDataFolderResponse) *EzsignfolderGetFormsDataV1ResponseMPayload {
	this := EzsignfolderGetFormsDataV1ResponseMPayload{}
	this.ObjFormsDataFolder = objFormsDataFolder
	return &this
}

// NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults instantiates a new EzsignfolderGetFormsDataV1ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfolderGetFormsDataV1ResponseMPayloadWithDefaults() *EzsignfolderGetFormsDataV1ResponseMPayload {
	this := EzsignfolderGetFormsDataV1ResponseMPayload{}
	return &this
}

// GetObjFormsDataFolder returns the ObjFormsDataFolder field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetObjFormsDataFolder() CustomFormsDataFolderResponse {
	if o == nil {
		var ret CustomFormsDataFolderResponse
		return ret
	}

	return o.ObjFormsDataFolder
}

// GetObjFormsDataFolderOk returns a tuple with the ObjFormsDataFolder field value
// and a boolean to check if the value has been set.
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) GetObjFormsDataFolderOk() (*CustomFormsDataFolderResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjFormsDataFolder, true
}

// SetObjFormsDataFolder sets field value
func (o *EzsignfolderGetFormsDataV1ResponseMPayload) SetObjFormsDataFolder(v CustomFormsDataFolderResponse) {
	o.ObjFormsDataFolder = v
}

func (o EzsignfolderGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignfolderGetFormsDataV1ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objFormsDataFolder"] = o.ObjFormsDataFolder
	return toSerialize, nil
}

type NullableEzsignfolderGetFormsDataV1ResponseMPayload struct {
	value *EzsignfolderGetFormsDataV1ResponseMPayload
	isSet bool
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) Get() *EzsignfolderGetFormsDataV1ResponseMPayload {
	return v.value
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) Set(val *EzsignfolderGetFormsDataV1ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfolderGetFormsDataV1ResponseMPayload(val *EzsignfolderGetFormsDataV1ResponseMPayload) *NullableEzsignfolderGetFormsDataV1ResponseMPayload {
	return &NullableEzsignfolderGetFormsDataV1ResponseMPayload{value: val, isSet: true}
}

func (v NullableEzsignfolderGetFormsDataV1ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfolderGetFormsDataV1ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


