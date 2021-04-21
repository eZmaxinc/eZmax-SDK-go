/*
 * eZmax API Definition
 *
 * This API expose all the functionnalities for the eZmax and eZsign applications.
 *
 * API version: 1.0.41
 * Contact: support-api@ezmax.ca
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// ContactRequestCompoundAllOf struct for ContactRequestCompoundAllOf
type ContactRequestCompoundAllOf struct {
	ObjContactinformations ContactinformationsRequestCompound `json:"objContactinformations"`
}

// NewContactRequestCompoundAllOf instantiates a new ContactRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequestCompoundAllOf(objContactinformations ContactinformationsRequestCompound) *ContactRequestCompoundAllOf {
	this := ContactRequestCompoundAllOf{}
	this.ObjContactinformations = objContactinformations
	return &this
}

// NewContactRequestCompoundAllOfWithDefaults instantiates a new ContactRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestCompoundAllOfWithDefaults() *ContactRequestCompoundAllOf {
	this := ContactRequestCompoundAllOf{}
	return &this
}

// GetObjContactinformations returns the ObjContactinformations field value
func (o *ContactRequestCompoundAllOf) GetObjContactinformations() ContactinformationsRequestCompound {
	if o == nil {
		var ret ContactinformationsRequestCompound
		return ret
	}

	return o.ObjContactinformations
}

// GetObjContactinformationsOk returns a tuple with the ObjContactinformations field value
// and a boolean to check if the value has been set.
func (o *ContactRequestCompoundAllOf) GetObjContactinformationsOk() (*ContactinformationsRequestCompound, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjContactinformations, true
}

// SetObjContactinformations sets field value
func (o *ContactRequestCompoundAllOf) SetObjContactinformations(v ContactinformationsRequestCompound) {
	o.ObjContactinformations = v
}

func (o ContactRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objContactinformations"] = o.ObjContactinformations
	}
	return json.Marshal(toSerialize)
}

type NullableContactRequestCompoundAllOf struct {
	value *ContactRequestCompoundAllOf
	isSet bool
}

func (v NullableContactRequestCompoundAllOf) Get() *ContactRequestCompoundAllOf {
	return v.value
}

func (v *NullableContactRequestCompoundAllOf) Set(val *ContactRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequestCompoundAllOf(val *ContactRequestCompoundAllOf) *NullableContactRequestCompoundAllOf {
	return &NullableContactRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableContactRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


