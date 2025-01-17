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

// checks if the ContactRequestCompoundV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactRequestCompoundV2{}

// ContactRequestCompoundV2 A Contact Object and children to create a complete structure
type ContactRequestCompoundV2 struct {
	ContactRequestV2
	ObjContactinformations ContactinformationsRequestCompoundV2 `json:"objContactinformations"`
}

type _ContactRequestCompoundV2 ContactRequestCompoundV2

// NewContactRequestCompoundV2 instantiates a new ContactRequestCompoundV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequestCompoundV2(objContactinformations ContactinformationsRequestCompoundV2, fkiContacttitleID int32, fkiLanguageID int32, eContactType FieldEContactType, sContactFirstname string, sContactLastname string) *ContactRequestCompoundV2 {
	this := ContactRequestCompoundV2{}
	this.FkiContacttitleID = fkiContacttitleID
	this.FkiLanguageID = fkiLanguageID
	this.EContactType = eContactType
	this.SContactFirstname = sContactFirstname
	this.SContactLastname = sContactLastname
	this.ObjContactinformations = objContactinformations
	return &this
}

// NewContactRequestCompoundV2WithDefaults instantiates a new ContactRequestCompoundV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestCompoundV2WithDefaults() *ContactRequestCompoundV2 {
	this := ContactRequestCompoundV2{}
	return &this
}

// GetObjContactinformations returns the ObjContactinformations field value
func (o *ContactRequestCompoundV2) GetObjContactinformations() ContactinformationsRequestCompoundV2 {
	if o == nil {
		var ret ContactinformationsRequestCompoundV2
		return ret
	}

	return o.ObjContactinformations
}

// GetObjContactinformationsOk returns a tuple with the ObjContactinformations field value
// and a boolean to check if the value has been set.
func (o *ContactRequestCompoundV2) GetObjContactinformationsOk() (*ContactinformationsRequestCompoundV2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjContactinformations, true
}

// SetObjContactinformations sets field value
func (o *ContactRequestCompoundV2) SetObjContactinformations(v ContactinformationsRequestCompoundV2) {
	o.ObjContactinformations = v
}

func (o ContactRequestCompoundV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactRequestCompoundV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["objContactinformations"] = o.ObjContactinformations
	return toSerialize, nil
}

func (o *ContactRequestCompoundV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objContactinformations",
		"fkiContacttitleID",
		"fkiLanguageID",
		"eContactType",
		"sContactFirstname",
		"sContactLastname",
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

	varContactRequestCompoundV2 := _ContactRequestCompoundV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varContactRequestCompoundV2)

	if err != nil {
		return err
	}

	*o = ContactRequestCompoundV2(varContactRequestCompoundV2)

	return err
}

type NullableContactRequestCompoundV2 struct {
	value *ContactRequestCompoundV2
	isSet bool
}

func (v NullableContactRequestCompoundV2) Get() *ContactRequestCompoundV2 {
	return v.value
}

func (v *NullableContactRequestCompoundV2) Set(val *ContactRequestCompoundV2) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequestCompoundV2) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequestCompoundV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequestCompoundV2(val *ContactRequestCompoundV2) *NullableContactRequestCompoundV2 {
	return &NullableContactRequestCompoundV2{value: val, isSet: true}
}

func (v NullableContactRequestCompoundV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequestCompoundV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


