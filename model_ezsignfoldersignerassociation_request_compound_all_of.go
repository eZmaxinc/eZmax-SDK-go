/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.4
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfoldersignerassociationRequestCompoundAllOf struct for EzsignfoldersignerassociationRequestCompoundAllOf
type EzsignfoldersignerassociationRequestCompoundAllOf struct {
	ObjEzsignsigner *EzsignsignerRequestCompound `json:"objEzsignsigner,omitempty"`
}

// NewEzsignfoldersignerassociationRequestCompoundAllOf instantiates a new EzsignfoldersignerassociationRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldersignerassociationRequestCompoundAllOf() *EzsignfoldersignerassociationRequestCompoundAllOf {
	this := EzsignfoldersignerassociationRequestCompoundAllOf{}
	return &this
}

// NewEzsignfoldersignerassociationRequestCompoundAllOfWithDefaults instantiates a new EzsignfoldersignerassociationRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldersignerassociationRequestCompoundAllOfWithDefaults() *EzsignfoldersignerassociationRequestCompoundAllOf {
	this := EzsignfoldersignerassociationRequestCompoundAllOf{}
	return &this
}

// GetObjEzsignsigner returns the ObjEzsignsigner field value if set, zero value otherwise.
func (o *EzsignfoldersignerassociationRequestCompoundAllOf) GetObjEzsignsigner() EzsignsignerRequestCompound {
	if o == nil || o.ObjEzsignsigner == nil {
		var ret EzsignsignerRequestCompound
		return ret
	}
	return *o.ObjEzsignsigner
}

// GetObjEzsignsignerOk returns a tuple with the ObjEzsignsigner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EzsignfoldersignerassociationRequestCompoundAllOf) GetObjEzsignsignerOk() (*EzsignsignerRequestCompound, bool) {
	if o == nil || o.ObjEzsignsigner == nil {
		return nil, false
	}
	return o.ObjEzsignsigner, true
}

// HasObjEzsignsigner returns a boolean if a field has been set.
func (o *EzsignfoldersignerassociationRequestCompoundAllOf) HasObjEzsignsigner() bool {
	if o != nil && o.ObjEzsignsigner != nil {
		return true
	}

	return false
}

// SetObjEzsignsigner gets a reference to the given EzsignsignerRequestCompound and assigns it to the ObjEzsignsigner field.
func (o *EzsignfoldersignerassociationRequestCompoundAllOf) SetObjEzsignsigner(v EzsignsignerRequestCompound) {
	o.ObjEzsignsigner = &v
}

func (o EzsignfoldersignerassociationRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjEzsignsigner != nil {
		toSerialize["objEzsignsigner"] = o.ObjEzsignsigner
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldersignerassociationRequestCompoundAllOf struct {
	value *EzsignfoldersignerassociationRequestCompoundAllOf
	isSet bool
}

func (v NullableEzsignfoldersignerassociationRequestCompoundAllOf) Get() *EzsignfoldersignerassociationRequestCompoundAllOf {
	return v.value
}

func (v *NullableEzsignfoldersignerassociationRequestCompoundAllOf) Set(val *EzsignfoldersignerassociationRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldersignerassociationRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldersignerassociationRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldersignerassociationRequestCompoundAllOf(val *EzsignfoldersignerassociationRequestCompoundAllOf) *NullableEzsignfoldersignerassociationRequestCompoundAllOf {
	return &NullableEzsignfoldersignerassociationRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableEzsignfoldersignerassociationRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldersignerassociationRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


