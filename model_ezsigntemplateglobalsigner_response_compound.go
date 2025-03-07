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

// checks if the EzsigntemplateglobalsignerResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplateglobalsignerResponseCompound{}

// EzsigntemplateglobalsignerResponseCompound A Ezsigntemplateglobalsigner Object
type EzsigntemplateglobalsignerResponseCompound struct {
	// The unique ID of the Ezsigntemplateglobalsigner
	PkiEzsigntemplateglobalsignerID int32 `json:"pkiEzsigntemplateglobalsignerID"`
	// The unique ID of the Ezsigntemplateglobal
	FkiEzsigntemplateglobalID int32 `json:"fkiEzsigntemplateglobalID"`
	// The description of the Ezsigntemplateglobalsigner
	SEzsigntemplateglobalsignerDescription string `json:"sEzsigntemplateglobalsignerDescription" validate:"regexp=^.{1,50}$"`
}

type _EzsigntemplateglobalsignerResponseCompound EzsigntemplateglobalsignerResponseCompound

// NewEzsigntemplateglobalsignerResponseCompound instantiates a new EzsigntemplateglobalsignerResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplateglobalsignerResponseCompound(pkiEzsigntemplateglobalsignerID int32, fkiEzsigntemplateglobalID int32, sEzsigntemplateglobalsignerDescription string) *EzsigntemplateglobalsignerResponseCompound {
	this := EzsigntemplateglobalsignerResponseCompound{}
	this.PkiEzsigntemplateglobalsignerID = pkiEzsigntemplateglobalsignerID
	this.FkiEzsigntemplateglobalID = fkiEzsigntemplateglobalID
	this.SEzsigntemplateglobalsignerDescription = sEzsigntemplateglobalsignerDescription
	return &this
}

// NewEzsigntemplateglobalsignerResponseCompoundWithDefaults instantiates a new EzsigntemplateglobalsignerResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplateglobalsignerResponseCompoundWithDefaults() *EzsigntemplateglobalsignerResponseCompound {
	this := EzsigntemplateglobalsignerResponseCompound{}
	return &this
}

// GetPkiEzsigntemplateglobalsignerID returns the PkiEzsigntemplateglobalsignerID field value
func (o *EzsigntemplateglobalsignerResponseCompound) GetPkiEzsigntemplateglobalsignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplateglobalsignerID
}

// GetPkiEzsigntemplateglobalsignerIDOk returns a tuple with the PkiEzsigntemplateglobalsignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalsignerResponseCompound) GetPkiEzsigntemplateglobalsignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplateglobalsignerID, true
}

// SetPkiEzsigntemplateglobalsignerID sets field value
func (o *EzsigntemplateglobalsignerResponseCompound) SetPkiEzsigntemplateglobalsignerID(v int32) {
	o.PkiEzsigntemplateglobalsignerID = v
}

// GetFkiEzsigntemplateglobalID returns the FkiEzsigntemplateglobalID field value
func (o *EzsigntemplateglobalsignerResponseCompound) GetFkiEzsigntemplateglobalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateglobalID
}

// GetFkiEzsigntemplateglobalIDOk returns a tuple with the FkiEzsigntemplateglobalID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalsignerResponseCompound) GetFkiEzsigntemplateglobalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateglobalID, true
}

// SetFkiEzsigntemplateglobalID sets field value
func (o *EzsigntemplateglobalsignerResponseCompound) SetFkiEzsigntemplateglobalID(v int32) {
	o.FkiEzsigntemplateglobalID = v
}

// GetSEzsigntemplateglobalsignerDescription returns the SEzsigntemplateglobalsignerDescription field value
func (o *EzsigntemplateglobalsignerResponseCompound) GetSEzsigntemplateglobalsignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplateglobalsignerDescription
}

// GetSEzsigntemplateglobalsignerDescriptionOk returns a tuple with the SEzsigntemplateglobalsignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplateglobalsignerResponseCompound) GetSEzsigntemplateglobalsignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplateglobalsignerDescription, true
}

// SetSEzsigntemplateglobalsignerDescription sets field value
func (o *EzsigntemplateglobalsignerResponseCompound) SetSEzsigntemplateglobalsignerDescription(v string) {
	o.SEzsigntemplateglobalsignerDescription = v
}

func (o EzsigntemplateglobalsignerResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplateglobalsignerResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplateglobalsignerID"] = o.PkiEzsigntemplateglobalsignerID
	toSerialize["fkiEzsigntemplateglobalID"] = o.FkiEzsigntemplateglobalID
	toSerialize["sEzsigntemplateglobalsignerDescription"] = o.SEzsigntemplateglobalsignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplateglobalsignerResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplateglobalsignerID",
		"fkiEzsigntemplateglobalID",
		"sEzsigntemplateglobalsignerDescription",
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

	varEzsigntemplateglobalsignerResponseCompound := _EzsigntemplateglobalsignerResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplateglobalsignerResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplateglobalsignerResponseCompound(varEzsigntemplateglobalsignerResponseCompound)

	return err
}

type NullableEzsigntemplateglobalsignerResponseCompound struct {
	value *EzsigntemplateglobalsignerResponseCompound
	isSet bool
}

func (v NullableEzsigntemplateglobalsignerResponseCompound) Get() *EzsigntemplateglobalsignerResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplateglobalsignerResponseCompound) Set(val *EzsigntemplateglobalsignerResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplateglobalsignerResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplateglobalsignerResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplateglobalsignerResponseCompound(val *EzsigntemplateglobalsignerResponseCompound) *NullableEzsigntemplateglobalsignerResponseCompound {
	return &NullableEzsigntemplateglobalsignerResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplateglobalsignerResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplateglobalsignerResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


