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

// checks if the EzsignbulksendResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignbulksendResponseCompound{}

// EzsignbulksendResponseCompound An Ezsignbulksend Object and children to create a complete structure
type EzsignbulksendResponseCompound struct {
	EzsignbulksendResponse
	AObjEzsignbulksenddocumentmapping []EzsignbulksenddocumentmappingResponseCompound `json:"a_objEzsignbulksenddocumentmapping"`
	AObjEzsignbulksendsignermapping []EzsignbulksendsignermappingResponse `json:"a_objEzsignbulksendsignermapping"`
}

type _EzsignbulksendResponseCompound EzsignbulksendResponseCompound

// NewEzsignbulksendResponseCompound instantiates a new EzsignbulksendResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignbulksendResponseCompound(aObjEzsignbulksenddocumentmapping []EzsignbulksenddocumentmappingResponseCompound, aObjEzsignbulksendsignermapping []EzsignbulksendsignermappingResponse, pkiEzsignbulksendID int32, fkiEzsignfoldertypeID int32, fkiLanguageID int32, sLanguageNameX string, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, sEzsignbulksendDescription string, tEzsignbulksendNote string, bEzsignbulksendNeedvalidation bool, bEzsignbulksendIsactive bool, objAudit CommonAudit) *EzsignbulksendResponseCompound {
	this := EzsignbulksendResponseCompound{}
	this.PkiEzsignbulksendID = pkiEzsignbulksendID
	this.FkiEzsignfoldertypeID = fkiEzsignfoldertypeID
	this.FkiLanguageID = fkiLanguageID
	this.SLanguageNameX = sLanguageNameX
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.SEzsignbulksendDescription = sEzsignbulksendDescription
	this.TEzsignbulksendNote = tEzsignbulksendNote
	this.BEzsignbulksendNeedvalidation = bEzsignbulksendNeedvalidation
	this.BEzsignbulksendIsactive = bEzsignbulksendIsactive
	this.ObjAudit = objAudit
	this.AObjEzsignbulksenddocumentmapping = aObjEzsignbulksenddocumentmapping
	this.AObjEzsignbulksendsignermapping = aObjEzsignbulksendsignermapping
	return &this
}

// NewEzsignbulksendResponseCompoundWithDefaults instantiates a new EzsignbulksendResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignbulksendResponseCompoundWithDefaults() *EzsignbulksendResponseCompound {
	this := EzsignbulksendResponseCompound{}
	return &this
}

// GetAObjEzsignbulksenddocumentmapping returns the AObjEzsignbulksenddocumentmapping field value
func (o *EzsignbulksendResponseCompound) GetAObjEzsignbulksenddocumentmapping() []EzsignbulksenddocumentmappingResponseCompound {
	if o == nil {
		var ret []EzsignbulksenddocumentmappingResponseCompound
		return ret
	}

	return o.AObjEzsignbulksenddocumentmapping
}

// GetAObjEzsignbulksenddocumentmappingOk returns a tuple with the AObjEzsignbulksenddocumentmapping field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendResponseCompound) GetAObjEzsignbulksenddocumentmappingOk() ([]EzsignbulksenddocumentmappingResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignbulksenddocumentmapping, true
}

// SetAObjEzsignbulksenddocumentmapping sets field value
func (o *EzsignbulksendResponseCompound) SetAObjEzsignbulksenddocumentmapping(v []EzsignbulksenddocumentmappingResponseCompound) {
	o.AObjEzsignbulksenddocumentmapping = v
}

// GetAObjEzsignbulksendsignermapping returns the AObjEzsignbulksendsignermapping field value
func (o *EzsignbulksendResponseCompound) GetAObjEzsignbulksendsignermapping() []EzsignbulksendsignermappingResponse {
	if o == nil {
		var ret []EzsignbulksendsignermappingResponse
		return ret
	}

	return o.AObjEzsignbulksendsignermapping
}

// GetAObjEzsignbulksendsignermappingOk returns a tuple with the AObjEzsignbulksendsignermapping field value
// and a boolean to check if the value has been set.
func (o *EzsignbulksendResponseCompound) GetAObjEzsignbulksendsignermappingOk() ([]EzsignbulksendsignermappingResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjEzsignbulksendsignermapping, true
}

// SetAObjEzsignbulksendsignermapping sets field value
func (o *EzsignbulksendResponseCompound) SetAObjEzsignbulksendsignermapping(v []EzsignbulksendsignermappingResponse) {
	o.AObjEzsignbulksendsignermapping = v
}

func (o EzsignbulksendResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignbulksendResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objEzsignbulksenddocumentmapping"] = o.AObjEzsignbulksenddocumentmapping
	toSerialize["a_objEzsignbulksendsignermapping"] = o.AObjEzsignbulksendsignermapping
	return toSerialize, nil
}

func (o *EzsignbulksendResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objEzsignbulksenddocumentmapping",
		"a_objEzsignbulksendsignermapping",
		"pkiEzsignbulksendID",
		"fkiEzsignfoldertypeID",
		"fkiLanguageID",
		"sLanguageNameX",
		"eEzsignfoldertypePrivacylevel",
		"sEzsignfoldertypeNameX",
		"sEzsignbulksendDescription",
		"tEzsignbulksendNote",
		"bEzsignbulksendNeedvalidation",
		"bEzsignbulksendIsactive",
		"objAudit",
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

	varEzsignbulksendResponseCompound := _EzsignbulksendResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignbulksendResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignbulksendResponseCompound(varEzsignbulksendResponseCompound)

	return err
}

type NullableEzsignbulksendResponseCompound struct {
	value *EzsignbulksendResponseCompound
	isSet bool
}

func (v NullableEzsignbulksendResponseCompound) Get() *EzsignbulksendResponseCompound {
	return v.value
}

func (v *NullableEzsignbulksendResponseCompound) Set(val *EzsignbulksendResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignbulksendResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignbulksendResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignbulksendResponseCompound(val *EzsignbulksendResponseCompound) *NullableEzsignbulksendResponseCompound {
	return &NullableEzsignbulksendResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignbulksendResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignbulksendResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


