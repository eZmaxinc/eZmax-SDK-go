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

// checks if the UsergroupexternalRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupexternalRequestCompound{}

// UsergroupexternalRequestCompound A Usergroupexternal Object and children
type UsergroupexternalRequestCompound struct {
	// The unique ID of the Usergroupexternal
	PkiUsergroupexternalID *int32 `json:"pkiUsergroupexternalID,omitempty"`
	// The name of the Usergroupexternal
	SUsergroupexternalName string `json:"sUsergroupexternalName" validate:"regexp=^.{0,64}$"`
	// The id of the Usergroupexternal
	SUsergroupexternalID string `json:"sUsergroupexternalID" validate:"regexp=^.{0,64}$"`
}

type _UsergroupexternalRequestCompound UsergroupexternalRequestCompound

// NewUsergroupexternalRequestCompound instantiates a new UsergroupexternalRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupexternalRequestCompound(sUsergroupexternalName string, sUsergroupexternalID string) *UsergroupexternalRequestCompound {
	this := UsergroupexternalRequestCompound{}
	this.SUsergroupexternalName = sUsergroupexternalName
	this.SUsergroupexternalID = sUsergroupexternalID
	return &this
}

// NewUsergroupexternalRequestCompoundWithDefaults instantiates a new UsergroupexternalRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupexternalRequestCompoundWithDefaults() *UsergroupexternalRequestCompound {
	this := UsergroupexternalRequestCompound{}
	return &this
}

// GetPkiUsergroupexternalID returns the PkiUsergroupexternalID field value if set, zero value otherwise.
func (o *UsergroupexternalRequestCompound) GetPkiUsergroupexternalID() int32 {
	if o == nil || IsNil(o.PkiUsergroupexternalID) {
		var ret int32
		return ret
	}
	return *o.PkiUsergroupexternalID
}

// GetPkiUsergroupexternalIDOk returns a tuple with the PkiUsergroupexternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsergroupexternalRequestCompound) GetPkiUsergroupexternalIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiUsergroupexternalID) {
		return nil, false
	}
	return o.PkiUsergroupexternalID, true
}

// HasPkiUsergroupexternalID returns a boolean if a field has been set.
func (o *UsergroupexternalRequestCompound) HasPkiUsergroupexternalID() bool {
	if o != nil && !IsNil(o.PkiUsergroupexternalID) {
		return true
	}

	return false
}

// SetPkiUsergroupexternalID gets a reference to the given int32 and assigns it to the PkiUsergroupexternalID field.
func (o *UsergroupexternalRequestCompound) SetPkiUsergroupexternalID(v int32) {
	o.PkiUsergroupexternalID = &v
}

// GetSUsergroupexternalName returns the SUsergroupexternalName field value
func (o *UsergroupexternalRequestCompound) GetSUsergroupexternalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalName
}

// GetSUsergroupexternalNameOk returns a tuple with the SUsergroupexternalName field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalRequestCompound) GetSUsergroupexternalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalName, true
}

// SetSUsergroupexternalName sets field value
func (o *UsergroupexternalRequestCompound) SetSUsergroupexternalName(v string) {
	o.SUsergroupexternalName = v
}

// GetSUsergroupexternalID returns the SUsergroupexternalID field value
func (o *UsergroupexternalRequestCompound) GetSUsergroupexternalID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupexternalID
}

// GetSUsergroupexternalIDOk returns a tuple with the SUsergroupexternalID field value
// and a boolean to check if the value has been set.
func (o *UsergroupexternalRequestCompound) GetSUsergroupexternalIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupexternalID, true
}

// SetSUsergroupexternalID sets field value
func (o *UsergroupexternalRequestCompound) SetSUsergroupexternalID(v string) {
	o.SUsergroupexternalID = v
}

func (o UsergroupexternalRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupexternalRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiUsergroupexternalID) {
		toSerialize["pkiUsergroupexternalID"] = o.PkiUsergroupexternalID
	}
	toSerialize["sUsergroupexternalName"] = o.SUsergroupexternalName
	toSerialize["sUsergroupexternalID"] = o.SUsergroupexternalID
	return toSerialize, nil
}

func (o *UsergroupexternalRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUsergroupexternalRequestCompound := _UsergroupexternalRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUsergroupexternalRequestCompound)

	if err != nil {
		return err
	}

	*o = UsergroupexternalRequestCompound(varUsergroupexternalRequestCompound)

	return err
}

type NullableUsergroupexternalRequestCompound struct {
	value *UsergroupexternalRequestCompound
	isSet bool
}

func (v NullableUsergroupexternalRequestCompound) Get() *UsergroupexternalRequestCompound {
	return v.value
}

func (v *NullableUsergroupexternalRequestCompound) Set(val *UsergroupexternalRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupexternalRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupexternalRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupexternalRequestCompound(val *UsergroupexternalRequestCompound) *NullableUsergroupexternalRequestCompound {
	return &NullableUsergroupexternalRequestCompound{value: val, isSet: true}
}

func (v NullableUsergroupexternalRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupexternalRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


