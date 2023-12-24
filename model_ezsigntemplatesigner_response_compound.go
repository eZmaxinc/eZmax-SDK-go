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

// checks if the EzsigntemplatesignerResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzsigntemplatesignerResponseCompound{}

// EzsigntemplatesignerResponseCompound A Ezsigntemplatesigner Object
type EzsigntemplatesignerResponseCompound struct {
	// The unique ID of the Ezsigntemplatesigner
	PkiEzsigntemplatesignerID int32 `json:"pkiEzsigntemplatesignerID"`
	// The unique ID of the Ezsigntemplate
	FkiEzsigntemplateID int32 `json:"fkiEzsigntemplateID"`
	// The description of the Ezsigntemplatesigner
	SEzsigntemplatesignerDescription string `json:"sEzsigntemplatesignerDescription"`
}

type _EzsigntemplatesignerResponseCompound EzsigntemplatesignerResponseCompound

// NewEzsigntemplatesignerResponseCompound instantiates a new EzsigntemplatesignerResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsigntemplatesignerResponseCompound(pkiEzsigntemplatesignerID int32, fkiEzsigntemplateID int32, sEzsigntemplatesignerDescription string) *EzsigntemplatesignerResponseCompound {
	this := EzsigntemplatesignerResponseCompound{}
	this.PkiEzsigntemplatesignerID = pkiEzsigntemplatesignerID
	this.FkiEzsigntemplateID = fkiEzsigntemplateID
	this.SEzsigntemplatesignerDescription = sEzsigntemplatesignerDescription
	return &this
}

// NewEzsigntemplatesignerResponseCompoundWithDefaults instantiates a new EzsigntemplatesignerResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsigntemplatesignerResponseCompoundWithDefaults() *EzsigntemplatesignerResponseCompound {
	this := EzsigntemplatesignerResponseCompound{}
	return &this
}

// GetPkiEzsigntemplatesignerID returns the PkiEzsigntemplatesignerID field value
func (o *EzsigntemplatesignerResponseCompound) GetPkiEzsigntemplatesignerID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsigntemplatesignerID
}

// GetPkiEzsigntemplatesignerIDOk returns a tuple with the PkiEzsigntemplatesignerID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerResponseCompound) GetPkiEzsigntemplatesignerIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiEzsigntemplatesignerID, true
}

// SetPkiEzsigntemplatesignerID sets field value
func (o *EzsigntemplatesignerResponseCompound) SetPkiEzsigntemplatesignerID(v int32) {
	o.PkiEzsigntemplatesignerID = v
}

// GetFkiEzsigntemplateID returns the FkiEzsigntemplateID field value
func (o *EzsigntemplatesignerResponseCompound) GetFkiEzsigntemplateID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzsigntemplateID
}

// GetFkiEzsigntemplateIDOk returns a tuple with the FkiEzsigntemplateID field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerResponseCompound) GetFkiEzsigntemplateIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzsigntemplateID, true
}

// SetFkiEzsigntemplateID sets field value
func (o *EzsigntemplatesignerResponseCompound) SetFkiEzsigntemplateID(v int32) {
	o.FkiEzsigntemplateID = v
}

// GetSEzsigntemplatesignerDescription returns the SEzsigntemplatesignerDescription field value
func (o *EzsigntemplatesignerResponseCompound) GetSEzsigntemplatesignerDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsigntemplatesignerDescription
}

// GetSEzsigntemplatesignerDescriptionOk returns a tuple with the SEzsigntemplatesignerDescription field value
// and a boolean to check if the value has been set.
func (o *EzsigntemplatesignerResponseCompound) GetSEzsigntemplatesignerDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzsigntemplatesignerDescription, true
}

// SetSEzsigntemplatesignerDescription sets field value
func (o *EzsigntemplatesignerResponseCompound) SetSEzsigntemplatesignerDescription(v string) {
	o.SEzsigntemplatesignerDescription = v
}

func (o EzsigntemplatesignerResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzsigntemplatesignerResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiEzsigntemplatesignerID"] = o.PkiEzsigntemplatesignerID
	toSerialize["fkiEzsigntemplateID"] = o.FkiEzsigntemplateID
	toSerialize["sEzsigntemplatesignerDescription"] = o.SEzsigntemplatesignerDescription
	return toSerialize, nil
}

func (o *EzsigntemplatesignerResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzsigntemplatesignerID",
		"fkiEzsigntemplateID",
		"sEzsigntemplatesignerDescription",
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

	varEzsigntemplatesignerResponseCompound := _EzsigntemplatesignerResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzsigntemplatesignerResponseCompound)

	if err != nil {
		return err
	}

	*o = EzsigntemplatesignerResponseCompound(varEzsigntemplatesignerResponseCompound)

	return err
}

type NullableEzsigntemplatesignerResponseCompound struct {
	value *EzsigntemplatesignerResponseCompound
	isSet bool
}

func (v NullableEzsigntemplatesignerResponseCompound) Get() *EzsigntemplatesignerResponseCompound {
	return v.value
}

func (v *NullableEzsigntemplatesignerResponseCompound) Set(val *EzsigntemplatesignerResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsigntemplatesignerResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsigntemplatesignerResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsigntemplatesignerResponseCompound(val *EzsigntemplatesignerResponseCompound) *NullableEzsigntemplatesignerResponseCompound {
	return &NullableEzsigntemplatesignerResponseCompound{value: val, isSet: true}
}

func (v NullableEzsigntemplatesignerResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsigntemplatesignerResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


