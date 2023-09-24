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

// checks if the ModulesectionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModulesectionResponse{}

// ModulesectionResponse A Modulesection Object
type ModulesectionResponse struct {
	// The unique ID of the Modulesection
	PkiModulesectionID int32 `json:"pkiModulesectionID"`
	// The unique ID of the Module
	FkiModuleID int32 `json:"fkiModuleID"`
	// The Internal name of the Module section.
	SModulesectionInternalname string `json:"sModulesectionInternalname"`
	// The Name of the Modulesection in the language of the requester
	SModulesectionNameX string `json:"sModulesectionNameX"`
}

// NewModulesectionResponse instantiates a new ModulesectionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModulesectionResponse(pkiModulesectionID int32, fkiModuleID int32, sModulesectionInternalname string, sModulesectionNameX string) *ModulesectionResponse {
	this := ModulesectionResponse{}
	this.PkiModulesectionID = pkiModulesectionID
	this.FkiModuleID = fkiModuleID
	this.SModulesectionInternalname = sModulesectionInternalname
	this.SModulesectionNameX = sModulesectionNameX
	return &this
}

// NewModulesectionResponseWithDefaults instantiates a new ModulesectionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModulesectionResponseWithDefaults() *ModulesectionResponse {
	this := ModulesectionResponse{}
	return &this
}

// GetPkiModulesectionID returns the PkiModulesectionID field value
func (o *ModulesectionResponse) GetPkiModulesectionID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiModulesectionID
}

// GetPkiModulesectionIDOk returns a tuple with the PkiModulesectionID field value
// and a boolean to check if the value has been set.
func (o *ModulesectionResponse) GetPkiModulesectionIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiModulesectionID, true
}

// SetPkiModulesectionID sets field value
func (o *ModulesectionResponse) SetPkiModulesectionID(v int32) {
	o.PkiModulesectionID = v
}

// GetFkiModuleID returns the FkiModuleID field value
func (o *ModulesectionResponse) GetFkiModuleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiModuleID
}

// GetFkiModuleIDOk returns a tuple with the FkiModuleID field value
// and a boolean to check if the value has been set.
func (o *ModulesectionResponse) GetFkiModuleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiModuleID, true
}

// SetFkiModuleID sets field value
func (o *ModulesectionResponse) SetFkiModuleID(v int32) {
	o.FkiModuleID = v
}

// GetSModulesectionInternalname returns the SModulesectionInternalname field value
func (o *ModulesectionResponse) GetSModulesectionInternalname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SModulesectionInternalname
}

// GetSModulesectionInternalnameOk returns a tuple with the SModulesectionInternalname field value
// and a boolean to check if the value has been set.
func (o *ModulesectionResponse) GetSModulesectionInternalnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SModulesectionInternalname, true
}

// SetSModulesectionInternalname sets field value
func (o *ModulesectionResponse) SetSModulesectionInternalname(v string) {
	o.SModulesectionInternalname = v
}

// GetSModulesectionNameX returns the SModulesectionNameX field value
func (o *ModulesectionResponse) GetSModulesectionNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SModulesectionNameX
}

// GetSModulesectionNameXOk returns a tuple with the SModulesectionNameX field value
// and a boolean to check if the value has been set.
func (o *ModulesectionResponse) GetSModulesectionNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SModulesectionNameX, true
}

// SetSModulesectionNameX sets field value
func (o *ModulesectionResponse) SetSModulesectionNameX(v string) {
	o.SModulesectionNameX = v
}

func (o ModulesectionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModulesectionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiModulesectionID"] = o.PkiModulesectionID
	toSerialize["fkiModuleID"] = o.FkiModuleID
	toSerialize["sModulesectionInternalname"] = o.SModulesectionInternalname
	toSerialize["sModulesectionNameX"] = o.SModulesectionNameX
	return toSerialize, nil
}

type NullableModulesectionResponse struct {
	value *ModulesectionResponse
	isSet bool
}

func (v NullableModulesectionResponse) Get() *ModulesectionResponse {
	return v.value
}

func (v *NullableModulesectionResponse) Set(val *ModulesectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModulesectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModulesectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModulesectionResponse(val *ModulesectionResponse) *NullableModulesectionResponse {
	return &NullableModulesectionResponse{value: val, isSet: true}
}

func (v NullableModulesectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModulesectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


