/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// FranchisereferalincomeRequestCompoundAllOf struct for FranchisereferalincomeRequestCompoundAllOf
type FranchisereferalincomeRequestCompoundAllOf struct {
	ObjAddress *AddressRequest `json:"objAddress,omitempty"`
	AObjContact []ContactRequestCompound `json:"a_objContact"`
}

// NewFranchisereferalincomeRequestCompoundAllOf instantiates a new FranchisereferalincomeRequestCompoundAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFranchisereferalincomeRequestCompoundAllOf(aObjContact []ContactRequestCompound) *FranchisereferalincomeRequestCompoundAllOf {
	this := FranchisereferalincomeRequestCompoundAllOf{}
	this.AObjContact = aObjContact
	return &this
}

// NewFranchisereferalincomeRequestCompoundAllOfWithDefaults instantiates a new FranchisereferalincomeRequestCompoundAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFranchisereferalincomeRequestCompoundAllOfWithDefaults() *FranchisereferalincomeRequestCompoundAllOf {
	this := FranchisereferalincomeRequestCompoundAllOf{}
	return &this
}

// GetObjAddress returns the ObjAddress field value if set, zero value otherwise.
func (o *FranchisereferalincomeRequestCompoundAllOf) GetObjAddress() AddressRequest {
	if o == nil || o.ObjAddress == nil {
		var ret AddressRequest
		return ret
	}
	return *o.ObjAddress
}

// GetObjAddressOk returns a tuple with the ObjAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompoundAllOf) GetObjAddressOk() (*AddressRequest, bool) {
	if o == nil || o.ObjAddress == nil {
		return nil, false
	}
	return o.ObjAddress, true
}

// HasObjAddress returns a boolean if a field has been set.
func (o *FranchisereferalincomeRequestCompoundAllOf) HasObjAddress() bool {
	if o != nil && o.ObjAddress != nil {
		return true
	}

	return false
}

// SetObjAddress gets a reference to the given AddressRequest and assigns it to the ObjAddress field.
func (o *FranchisereferalincomeRequestCompoundAllOf) SetObjAddress(v AddressRequest) {
	o.ObjAddress = &v
}

// GetAObjContact returns the AObjContact field value
func (o *FranchisereferalincomeRequestCompoundAllOf) GetAObjContact() []ContactRequestCompound {
	if o == nil {
		var ret []ContactRequestCompound
		return ret
	}

	return o.AObjContact
}

// GetAObjContactOk returns a tuple with the AObjContact field value
// and a boolean to check if the value has been set.
func (o *FranchisereferalincomeRequestCompoundAllOf) GetAObjContactOk() (*[]ContactRequestCompound, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AObjContact, true
}

// SetAObjContact sets field value
func (o *FranchisereferalincomeRequestCompoundAllOf) SetAObjContact(v []ContactRequestCompound) {
	o.AObjContact = v
}

func (o FranchisereferalincomeRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjAddress != nil {
		toSerialize["objAddress"] = o.ObjAddress
	}
	if true {
		toSerialize["a_objContact"] = o.AObjContact
	}
	return json.Marshal(toSerialize)
}

type NullableFranchisereferalincomeRequestCompoundAllOf struct {
	value *FranchisereferalincomeRequestCompoundAllOf
	isSet bool
}

func (v NullableFranchisereferalincomeRequestCompoundAllOf) Get() *FranchisereferalincomeRequestCompoundAllOf {
	return v.value
}

func (v *NullableFranchisereferalincomeRequestCompoundAllOf) Set(val *FranchisereferalincomeRequestCompoundAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFranchisereferalincomeRequestCompoundAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFranchisereferalincomeRequestCompoundAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFranchisereferalincomeRequestCompoundAllOf(val *FranchisereferalincomeRequestCompoundAllOf) *NullableFranchisereferalincomeRequestCompoundAllOf {
	return &NullableFranchisereferalincomeRequestCompoundAllOf{value: val, isSet: true}
}

func (v NullableFranchisereferalincomeRequestCompoundAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFranchisereferalincomeRequestCompoundAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


