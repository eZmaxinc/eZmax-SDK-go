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

// checks if the EzsignformfieldgroupsignerRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsignformfieldgroupsignerRequestCompound{}

// EzsignformfieldgroupsignerRequestCompound An Ezsignformfieldgroupsigner Object and children to create a complete structure
type EzsignformfieldgroupsignerRequestCompound struct {
	// The unique ID of the Ezsignformfieldgroupsigner
	PkiEzsignformfieldgroupsignerID *int32 `json:"pkiEzsignformfieldgroupsignerID,omitempty"`
	// The unique ID of the Ezsignfoldersignerassociation
	FkiEzsignfoldersignerassociationID int32 `json:"fkiEzsignfoldersignerassociationID"`
}

type _EzsignformfieldgroupsignerRequestCompound EzsignformfieldgroupsignerRequestCompound

// NewEzsignformfieldgroupsignerRequestCompound instantiates a new EzsignformfieldgroupsignerRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignformfieldgroupsignerRequestCompound(fkiEzsignfoldersignerassociationID int32) *EzsignformfieldgroupsignerRequestCompound {
	this := EzsignformfieldgroupsignerRequestCompound{}
	this.FkiEzsignfoldersignerassociationID = fkiEzsignfoldersignerassociationID
	return &this
}

// NewEzsignformfieldgroupsignerRequestCompoundWithDefaults instantiates a new EzsignformfieldgroupsignerRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignformfieldgroupsignerRequestCompoundWithDefaults() *EzsignformfieldgroupsignerRequestCompound {
	this := EzsignformfieldgroupsignerRequestCompound{}
	return &this
}

// GetPkiEzsignformfieldgroupsignerID returns the PkiEzsignformfieldgroupsignerID field value if set, zero value otherwise.
func (o *EzsignformfieldgroupsignerRequestCompound) GetPkiEzsignformfieldgroupsignerID() int32 {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupsignerID) {
		var ret int32
		return ret
	}
	return *o.PkiEzsignformfieldgroupsignerID
}

// GetPkiEzsignformfieldgroupsignerIDOk returns a tuple with the PkiEzsignformfieldgroupsignerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerRequestCompound) GetPkiEzsignformfieldgroupsignerIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiEzsignformfieldgroupsignerID) {
		return nil, false
	}
	return o.PkiEzsignformfieldgroupsignerID, true
}

// HasPkiEzsignformfieldgroupsignerID returns a boolean if a field has been set.
func (o *EzsignformfieldgroupsignerRequestCompound) HasPkiEzsignformfieldgroupsignerID() bool {
	if o != nil && !IsNil(o.PkiEzsignformfieldgroupsignerID) {
		return true
	}

	return false
}

// SetPkiEzsignformfieldgroupsignerID gets a reference to the given int32 and assigns it to the PkiEzsignformfieldgroupsignerID field.
func (o *EzsignformfieldgroupsignerRequestCompound) SetPkiEzsignformfieldgroupsignerID(v int32) {
	o.PkiEzsignformfieldgroupsignerID = &v
}

// GetFkiEzsignfoldersignerassociationID returns the FkiEzsignfoldersignerassociationID field value
func (o *EzsignformfieldgroupsignerRequestCompound) GetFkiEzsignfoldersignerassociationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsignfoldersignerassociationID
}

// GetFkiEzsignfoldersignerassociationIDOk returns a tuple with the FkiEzsignfoldersignerassociationID field value
// and a boolean to check if the value has been set.
func (o *EzsignformfieldgroupsignerRequestCompound) GetFkiEzsignfoldersignerassociationIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsignfoldersignerassociationID, true
}

// SetFkiEzsignfoldersignerassociationID sets field value
func (o *EzsignformfieldgroupsignerRequestCompound) SetFkiEzsignfoldersignerassociationID(v int32) {
	o.FkiEzsignfoldersignerassociationID = v
}

func (o EzsignformfieldgroupsignerRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsignformfieldgroupsignerRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiEzsignformfieldgroupsignerID) {
		toSerialize["pkiEzsignformfieldgroupsignerID"] = o.PkiEzsignformfieldgroupsignerID
	}
	toSerialize["fkiEzsignfoldersignerassociationID"] = o.FkiEzsignfoldersignerassociationID
	return toSerialize, nil
}

func (o *EzsignformfieldgroupsignerRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varEzsignformfieldgroupsignerRequestCompound := _EzsignformfieldgroupsignerRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsignformfieldgroupsignerRequestCompound)

	if err != nil {
		return err
	}

	*o = EzsignformfieldgroupsignerRequestCompound(varEzsignformfieldgroupsignerRequestCompound)

	return err
}

type NullableEzsignformfieldgroupsignerRequestCompound struct {
	value *EzsignformfieldgroupsignerRequestCompound
	isSet bool
}

func (v NullableEzsignformfieldgroupsignerRequestCompound) Get() *EzsignformfieldgroupsignerRequestCompound {
	return v.value
}

func (v *NullableEzsignformfieldgroupsignerRequestCompound) Set(val *EzsignformfieldgroupsignerRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignformfieldgroupsignerRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignformfieldgroupsignerRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignformfieldgroupsignerRequestCompound(val *EzsignformfieldgroupsignerRequestCompound) *NullableEzsignformfieldgroupsignerRequestCompound {
	return &NullableEzsignformfieldgroupsignerRequestCompound{value: val, isSet: true}
}

func (v NullableEzsignformfieldgroupsignerRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignformfieldgroupsignerRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


