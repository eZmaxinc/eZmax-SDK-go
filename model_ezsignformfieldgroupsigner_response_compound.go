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

// checks if the EzsignformfieldgroupsignerResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupsignerResponseCompound{}

// EzsignformfieldgroupsignerResponseCompound An Ezsignformfieldgroupsigner Object and children to create a complete structure
type EzsignformfieldgroupsignerResponseCompound struct {
	// The unique ID of the Ezsignformfieldgroupsigner
	PkiEzsignformfieldgroupsignerID int32 `json:"pkiEzsignformfieldgroupsignerID"`
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
}

type _EzsignformfieldgroupsignerResponseCompound EzsignformfieldgroupsignerResponseCompound

// NewEzsignformfieldgroupsignerResponseCompound instantiates a new EzsignformfieldgroupsignerResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupsignerResponseCompound(pkiEzsignformfieldgroupsignerID int32, fkiEzsignfoldersignerassociationID int32) *EzsignformfieldgroupsignerResponseCompound {
	this := EzsignformfieldgroupsignerResponseCompound{}
	this.PkiEzsignformfieldgroupsignerID = pkiEzsignformfieldgroupsignerID
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignformfieldgroupsignerResponseCompoundWithDefaults instantiates a new EzsignformfieldgroupsignerResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupsignerResponseCompoundWithDefaults() *EzsignformfieldgroupsignerResponseCompound {
	this := EzsignformfieldgroupsignerResponseCompound{}
	return &this
}

// GetPkiEzsignformfieldgroupsignerID returns the PkiEzsignformfieldgroupsignerID field value
func (o *EzsignformfieldgroupsignerResponseCompound) GetPkiEzsignformfieldgroupsignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignformfieldgroupsignerID
}

// GetPkiEzsignformfieldgroupsignerIDOk returns a tuple with the PkiEzsignformfieldgroupsignerID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerResponseCompound) GetPkiEzsignformfieldgroupsignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsignformfieldgroupsignerID, true
}

// SetPkiEzsignformfieldgroupsignerID sets field value
func (o *EzsignformfieldgroupsignerResponseCompound) SetPkiEzsignformfieldgroupsignerID(v int32) {
	o.PkiEzsignformfieldgroupsignerID = v
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *EzsignformfieldgroupsignerResponseCompound) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerResponseCompound) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignformfieldgroupsignerResponseCompound) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

func (o EzsignformfieldgroupsignerResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupsignerResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsignformfieldgroupsignerID"] = o.PkiEzsignformfieldgroupsignerID
	toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignformfieldgroupsignerResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsignformfieldgroupsignerID",
		"fkiEzsignfoldersignerassociationID",
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

	varEzsignformfieldgroupsignerResponseCompound := _EzsignformfieldgroupsignerResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldgroupsignerResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsignformfieldgroupsignerResponseCompound(varEzsignformfieldgroupsignerResponseCompound)

	return err
}

type NullableEzsignformfieldgroupsignerResponseCompound struct {
	value *EzsignformfieldgroupsignerResponseCompound
	isSet bool
}

func (v NullableEzsignformfieldgroupsignerResponseCompound) Get() *EzsignformfieldgroupsignerResponseCompound {
	return v.value
}

func (v *NullableEzsignformfieldgroupsignerResponseCompound) Set(val *EzsignformfieldgroupsignerResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupsignerResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupsignerResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupsignerResponseCompound(val *EzsignformfieldgroupsignerResponseCompound) *NullableEzsignformfieldgroupsignerResponseCompound {
	return &NullableEzsignformfieldgroupsignerResponseCompound{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupsignerResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupsignerResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


