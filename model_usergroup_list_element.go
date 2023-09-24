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
)

// checks if the UsergroupListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsergroupListElement{}

// UsergroupListElement A Usergroup List Element
type UsergroupListElement struct {
	// The unique ID of the Usergroup
	PkiUsergroupID int32 `json:"pkiUsergroupID"`
	// The Name of the Usergroup in the language of the requester
	SUsergroupNameX string `json:"sUsergroupNameX"`
	// Number of users in group
	ICountUser int32 `json:"iCountUser"`
}

// NewUsergroupListElement instantiates a new UsergroupListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsergroupListElement(pkiUsergroupID int32, sUsergroupNameX string, iCountUser int32) *UsergroupListElement {
	this := UsergroupListElement{}
	this.PkiUsergroupID = pkiUsergroupID
	this.SUsergroupNameX = sUsergroupNameX
	this.ICountUser = iCountUser
	return &this
}

// NewUsergroupListElementWithDefaults instantiates a new UsergroupListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsergroupListElementWithDefaults() *UsergroupListElement {
	this := UsergroupListElement{}
	return &this
}

// GetPkiUsergroupID returns the PkiUsergroupID field value
func (o *UsergroupListElement) GetPkiUsergroupID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiUsergroupID
}

// GetPkiUsergroupIDOk returns a tuple with the PkiUsergroupID field value
// and a boolean to check if the value has been set.
func (o *UsergroupListElement) GetPkiUsergroupIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiUsergroupID, true
}

// SetPkiUsergroupID sets field value
func (o *UsergroupListElement) SetPkiUsergroupID(v int32) {
	o.PkiUsergroupID = v
}

// GetSUsergroupNameX returns the SUsergroupNameX field value
func (o *UsergroupListElement) GetSUsergroupNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SUsergroupNameX
}

// GetSUsergroupNameXOk returns a tuple with the SUsergroupNameX field value
// and a boolean to check if the value has been set.
func (o *UsergroupListElement) GetSUsergroupNameXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SUsergroupNameX, true
}

// SetSUsergroupNameX sets field value
func (o *UsergroupListElement) SetSUsergroupNameX(v string) {
	o.SUsergroupNameX = v
}

// GetICountUser returns the ICountUser field value
func (o *UsergroupListElement) GetICountUser() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICountUser
}

// GetICountUserOk returns a tuple with the ICountUser field value
// and a boolean to check if the value has been set.
func (o *UsergroupListElement) GetICountUserOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICountUser, true
}

// SetICountUser sets field value
func (o *UsergroupListElement) SetICountUser(v int32) {
	o.ICountUser = v
}

func (o UsergroupListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsergroupListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiUsergroupID"] = o.PkiUsergroupID
	toSerialize["sUsergroupNameX"] = o.SUsergroupNameX
	toSerialize["iCountUser"] = o.ICountUser
	return toSerialize, nil
}

type NullableUsergroupListElement struct {
	value *UsergroupListElement
	isSet bool
}

func (v NullableUsergroupListElement) Get() *UsergroupListElement {
	return v.value
}

func (v *NullableUsergroupListElement) Set(val *UsergroupListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableUsergroupListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableUsergroupListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsergroupListElement(val *UsergroupListElement) *NullableUsergroupListElement {
	return &NullableUsergroupListElement{value: val, isSet: true}
}

func (v NullableUsergroupListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsergroupListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

