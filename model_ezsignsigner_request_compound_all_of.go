/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.3
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignsignerRequestCompoundAllOf struct for EzsignsignerRequestCompoundAllOf
type EzsignsignerRequestCompoundAllOf struct {
	ObjContact EzsignsignerRequestCompoundContact `json:"objContact"`
}

// NewEzsignsignerRequestCompoundAllOf instantiates a new EzsignsignerRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignsignerRequestCompoundAllOf(objContact EzsignsignerRequestCompoundContact) *EzsignsignerRequestCompoundAllOf {
	this := EzsignsignerRequestCompoundAllOf{}
	this.ObjContact = objContact
	return &this
}

// NewEzsignsignerRequestCompoundAllOfWithDefaults instantiates a new EzsignsignerRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignsignerRequestCompoundAllOfWithDefaults() *EzsignsignerRequestCompoundAllOf {
	this := EzsignsignerRequestCompoundAllOf{}
	return &this
}

// GetObjContact returns the ObjContact field value
func (o *EzsignsignerRequestCompoundAllOf) GetObjContact() EzsignsignerRequestCompoundContact {
	if o == nil {
		var ret EzsignsignerRequestCompoundContact
		return ret
	}

	return o.ObjContact
}

// GetObjContactOk returns a tuple with the ObjContact field value
// and a boolean to check if the value has been set.
func (o *EzsignsignerRequestCompoundAllOf) GetObjContactOk() (*EzsignsignerRequestCompoundContact, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ObjContact, true
}

// SetObjContact sets field value
func (o *EzsignsignerRequestCompoundAllOf) SetObjContact(v EzsignsignerRequestCompoundContact) {
	o.ObjContact = v
}

func (o EzsignsignerRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["objContact"] = o.ObjContact
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignsignerRequestCompoundAllOf struct {
	value *EzsignsignerRequestCompoundAllOf
	isSet bool
}

func (v NullableEzsignsignerRequestCompoundAllOf) Get() *EzsignsignerRequestCompoundAllOf {
	return v.value
}

func (v *NullableEzsignsignerRequestCompoundAllOf) Set(val *EzsignsignerRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignsignerRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignsignerRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignsignerRequestCompoundAllOf(val *EzsignsignerRequestCompoundAllOf) *NullableEzsignsignerRequestCompoundAllOf {
	return &NullableEzsignsignerRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableEzsignsignerRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignsignerRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


