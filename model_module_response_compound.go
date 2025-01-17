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

// checks if the ModuleResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModuleResponseCompound{}

// ModuleResponseCompound A Module Object
type ModuleResponseCompound struct {
	ModuleResponse
	AObjModulesection []ModulesectionResponseCompound `json:"a_objModulesection,omitempty"`
}

type _ModuleResponseCompound ModuleResponseCompound

// NewModuleResponseCompound instantiates a new ModuleResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModuleResponseCompound(pkiModuleID int32, fkiModulegroupID int32, eModuleInternalname string, sModuleNameX string, bModuleRegistered bool, bModuleRegisteredapi bool) *ModuleResponseCompound {
	this := ModuleResponseCompound{}
	this.PkiModuleID = pkiModuleID
	this.FkiModulegroupID = fkiModulegroupID
	this.EModuleInternalname = eModuleInternalname
	this.SModuleNameX = sModuleNameX
	this.BModuleRegistered = bModuleRegistered
	this.BModuleRegisteredapi = bModuleRegisteredapi
	return &this
}

// NewModuleResponseCompoundWithDefaults instantiates a new ModuleResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModuleResponseCompoundWithDefaults() *ModuleResponseCompound {
	this := ModuleResponseCompound{}
	return &this
}

// GetAObjModulesection returns the AObjModulesection field value if set, zero value otherwise.
func (o *ModuleResponseCompound) GetAObjModulesection() []ModulesectionResponseCompound {
	if o == nil || IsNil(o.AObjModulesection) {
		var ret []ModulesectionResponseCompound
		return ret
	}
	return o.AObjModulesection
}

// GetAObjModulesectionOk returns a tuple with the AObjModulesection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModuleResponseCompound) GetAObjModulesectionOk() ([]ModulesectionResponseCompound, bool) {
	if o == nil || IsNil(o.AObjModulesection) {
		return nil, false
	}
	return o.AObjModulesection, true
}

// HasAObjModulesection returns a boolean if a field has been set.
func (o *ModuleResponseCompound) HasAObjModulesection() bool {
	if o != nil && !IsNil(o.AObjModulesection) {
		return true
	}

	return false
}

// SetAObjModulesection gets a reference to the given []ModulesectionResponseCompound and assigns it to the AObjModulesection field.
func (o *ModuleResponseCompound) SetAObjModulesection(v []ModulesectionResponseCompound) {
	o.AObjModulesection = v
}

func (o ModuleResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModuleResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AObjModulesection) {
		toSerialize["a_objModulesection"] = o.AObjModulesection
	}
	return toSerialize, nil
}

func (o *ModuleResponseCompound) UnmarshalJSON(data []byte) (err error) {
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

	varModuleResponseCompound := _ModuleResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModuleResponseCompound)

	if err != nil {
		return err
	}

	*o = ModuleResponseCompound(varModuleResponseCompound)

	return err
}

type NullableModuleResponseCompound struct {
	value *ModuleResponseCompound
	isSet bool
}

func (v NullableModuleResponseCompound) Get() *ModuleResponseCompound {
	return v.value
}

func (v *NullableModuleResponseCompound) Set(val *ModuleResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableModuleResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableModuleResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModuleResponseCompound(val *ModuleResponseCompound) *NullableModuleResponseCompound {
	return &NullableModuleResponseCompound{value: val, isSet: true}
}

func (v NullableModuleResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModuleResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


