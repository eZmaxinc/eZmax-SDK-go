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

// checks if the ModuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleResponse{}

// ModuleResponse A Module Object
type ModuleResponse struct {
	// The unique ID of the Module
	PkiModuleID int32 `json:"pkiModuleID"`
	// The unique ID of the Modulegroup
	FkiModulegroupID int32 `json:"fkiModulegroupID"`
	// The Internal name of the Module.  This is theoretically an enum field but there are so many possibles values we decided not to list them all.
	EModuleInternalname string `json:"eModuleInternalname"`
	// The Name of the Module in the language of the requester
	SModuleNameX string `json:"sModuleNameX"`
	// Whether the Module is registered or not
	BModuleRegistered bool `json:"bModuleRegistered"`
	// Whether the Module is registered or not for api use
	BModuleRegisteredapi bool `json:"bModuleRegisteredapi"`
}

type _ModuleResponse ModuleResponse

// NewModuleResponse instantiates a new ModuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleResponse(pkiModuleID int32, fkiModulegroupID int32, eModuleInternalname string, sModuleNameX string, bModuleRegistered bool, bModuleRegisteredapi bool) *ModuleResponse {
	this := ModuleResponse{}
	this.PkiModuleID = pkiModuleID
	this.FkiModulegroupID = fkiModulegroupID
	this.EModuleInternalname = eModuleInternalname
	this.SModuleNameX = sModuleNameX
	this.BModuleRegistered = bModuleRegistered
	this.BModuleRegisteredapi = bModuleRegisteredapi
	return &this
}

// NewModuleResponseWithDefaults instantiates a new ModuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleResponseWithDefaults() *ModuleResponse {
	this := ModuleResponse{}
	return &this
}

// GetPkiModuleID returns the PkiModuleID field value
func (o *ModuleResponse) GetPkiModuleID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiModuleID
}

// GetPkiModuleIDOk returns a tuple with the PkiModuleID field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetPkiModuleIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiModuleID, true
}

// SetPkiModuleID sets field value
func (o *ModuleResponse) SetPkiModuleID(v int32) {
	o.PkiModuleID = v
}

// GetFkiModulegroupID returns the FkiModulegroupID field value
func (o *ModuleResponse) GetFkiModulegroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiModulegroupID
}

// GetFkiModulegroupIDOk returns a tuple with the FkiModulegroupID field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetFkiModulegroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiModulegroupID, true
}

// SetFkiModulegroupID sets field value
func (o *ModuleResponse) SetFkiModulegroupID(v int32) {
	o.FkiModulegroupID = v
}

// GetEModuleInternalname returns the EModuleInternalname field value
func (o *ModuleResponse) GetEModuleInternalname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EModuleInternalname
}

// GetEModuleInternalnameOk returns a tuple with the EModuleInternalname field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetEModuleInternalnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EModuleInternalname, true
}

// SetEModuleInternalname sets field value
func (o *ModuleResponse) SetEModuleInternalname(v string) {
	o.EModuleInternalname = v
}

// GetSModuleNameX returns the SModuleNameX field value
func (o *ModuleResponse) GetSModuleNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SModuleNameX
}

// GetSModuleNameXOk returns a tuple with the SModuleNameX field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetSModuleNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SModuleNameX, true
}

// SetSModuleNameX sets field value
func (o *ModuleResponse) SetSModuleNameX(v string) {
	o.SModuleNameX = v
}

// GetBModuleRegistered returns the BModuleRegistered field value
func (o *ModuleResponse) GetBModuleRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BModuleRegistered
}

// GetBModuleRegisteredOk returns a tuple with the BModuleRegistered field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetBModuleRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BModuleRegistered, true
}

// SetBModuleRegistered sets field value
func (o *ModuleResponse) SetBModuleRegistered(v bool) {
	o.BModuleRegistered = v
}

// GetBModuleRegisteredapi returns the BModuleRegisteredapi field value
func (o *ModuleResponse) GetBModuleRegisteredapi() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BModuleRegisteredapi
}

// GetBModuleRegisteredapiOk returns a tuple with the BModuleRegisteredapi field value
// and a boolean to check if the value has been set.
func (o *ModuleResponse) GetBModuleRegisteredapiOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BModuleRegisteredapi, true
}

// SetBModuleRegisteredapi sets field value
func (o *ModuleResponse) SetBModuleRegisteredapi(v bool) {
	o.BModuleRegisteredapi = v
}

func (o ModuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiModuleID"] = o.PkiModuleID
	toSerialize["fkiModulegroupID"] = o.FkiModulegroupID
	toSerialize["eModuleInternalname"] = o.EModuleInternalname
	toSerialize["sModuleNameX"] = o.SModuleNameX
	toSerialize["bModuleRegistered"] = o.BModuleRegistered
	toSerialize["bModuleRegisteredapi"] = o.BModuleRegisteredapi
	return toSerialize, nil
}

func (o *ModuleResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiModuleID",
		"fkiModulegroupID",
		"eModuleInternalname",
		"sModuleNameX",
		"bModuleRegistered",
		"bModuleRegisteredapi",
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

	varModuleResponse := _ModuleResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModuleResponse)

	if err != nil {
		return err
	}

	*o = ModuleResponse(varModuleResponse)

	return err
}

type NullableModuleResponse struct {
	value *ModuleResponse
	isSet bool
}

func (v NullableModuleResponse) Get() *ModuleResponse {
	return v.value
}

func (v *NullableModuleResponse) Set(val *ModuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleResponse(val *ModuleResponse) *NullableModuleResponse {
	return &NullableModuleResponse{value: val, isSet: true}
}

func (v NullableModuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


