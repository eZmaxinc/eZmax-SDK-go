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

// checks if the UsergroupexternalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupexternalResponseCompound{}

// UsergroupexternalResponseCompound A Usergroupexternal Object
type UsergroupexternalResponseCompound struct {
	// The unique ID of the Usergroupexternal
	PkiUsergroupexternalID int32 `json:"pkiUsergroupexternalID"`
	// The name of the Usergroupexternal
	SUsergroupexternalName string `json:"sUsergroupexternalName" validate:"regexp=^.{0,64}$"`
	// The id of the Usergroupexternal
	SUsergroupexternalID string `json:"sUsergroupexternalID" validate:"regexp=^.{0,64}$"`
}

type _UsergroupexternalResponseCompound UsergroupexternalResponseCompound

// NewUsergroupexternalResponseCompound instantiates a new UsergroupexternalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupexternalResponseCompound(pkiUsergroupexternalID int32, sUsergroupexternalName string, sUsergroupexternalID string) *UsergroupexternalResponseCompound {
	this := UsergroupexternalResponseCompound{}
	this.PkiUsergroupexternalID = pkiUsergroupexternalID
	this.SUsergroupexternalName = sUsergroupexternalName
	this.SUsergroupexternalID = sUsergroupexternalID
	return &this
}

// NewUsergroupexternalResponseCompoundWithDefaults instantiates a new UsergroupexternalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupexternalResponseCompoundWithDefaults() *UsergroupexternalResponseCompound {
	this := UsergroupexternalResponseCompound{}
	return &this
}

// GetPkiUsergroupexternalID returns the PkiUsergroupexternalID field value
func (o *UsergroupexternalResponseCompound) GetPkiUsergroupexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupexternalID
}

// GetPkiUsergroupexternalIDOk returns a tuple with the PkiUsergroupexternalID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalResponseCompound) GetPkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupexternalID, true
}

// SetPkiUsergroupexternalID sets field value
func (o *UsergroupexternalResponseCompound) SetPkiUsergroupexternalID(v int32) {
	o.PkiUsergroupexternalID = v
}

// GetSUsergroupexternalName returns the SUsergroupexternalName field value
func (o *UsergroupexternalResponseCompound) GetSUsergroupexternalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalName
}

// GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalResponseCompound) GetSUsergroupexternalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalName, true
}

// SetSUsergroupexternalName sets field value
func (o *UsergroupexternalResponseCompound) SetSUsergroupexternalName(v string) {
	o.SUsergroupexternalName = v
}

// GetSUsergroupexternalID returns the SUsergroupexternalID field value
func (o *UsergroupexternalResponseCompound) GetSUsergroupexternalID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalID
}

// GetSUsergroupexternalIDOk returns a tuple with the SUsergroupexternalID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalResponseCompound) GetSUsergroupexternalIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalID, true
}

// SetSUsergroupexternalID sets field value
func (o *UsergroupexternalResponseCompound) SetSUsergroupexternalID(v string) {
	o.SUsergroupexternalID = v
}

func (o UsergroupexternalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupexternalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupexternalID"] = o.PkiUsergroupexternalID
	toSerialize["sUsergroupexternalName"] = o.SUsergroupexternalName
	toSerialize["sUsergroupexternalID"] = o.SUsergroupexternalID
	return toSerialize, nil
}

func (o *UsergroupexternalResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiUsergroupexternalID",
		"sUsergroupexternalName",
		"sUsergroupexternalID",
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

	varUsergroupexternalResponseCompound := _UsergroupexternalResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupexternalResponseCompound)

	if err != nil {
		return err
	}

	*o = UsergroupexternalResponseCompound(varUsergroupexternalResponseCompound)

	return err
}

type NullableUsergroupexternalResponseCompound struct {
	value *UsergroupexternalResponseCompound
	isSet bool
}

func (v NullableUsergroupexternalResponseCompound) Get() *UsergroupexternalResponseCompound {
	return v.value
}

func (v *NullableUsergroupexternalResponseCompound) Set(val *UsergroupexternalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupexternalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupexternalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupexternalResponseCompound(val *UsergroupexternalResponseCompound) *NullableUsergroupexternalResponseCompound {
	return &NullableUsergroupexternalResponseCompound{value: val, isSet: true}
}

func (v NullableUsergroupexternalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupexternalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


