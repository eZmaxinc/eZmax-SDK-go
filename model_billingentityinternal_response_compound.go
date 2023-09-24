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

// checks if the BillingentityinternalResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalResponseCompound{}

// BillingentityinternalResponseCompound A Billingentityinternal Object
type BillingentityinternalResponseCompound struct {
	// The unique ID of the Billingentityinternal.
	PkiBillingentityinternalID int32 `json:"pkiBillingentityinternalID"`
	ObjBillingentityinternalDescription MultilingualBillingentityinternalDescription `json:"objBillingentityinternalDescription"`
	AObjBillingentityinternalproduct []BillingentityinternalproductResponseCompound `json:"a_objBillingentityinternalproduct"`
}

// NewBillingentityinternalResponseCompound instantiates a new BillingentityinternalResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalResponseCompound(pkiBillingentityinternalID int32, objBillingentityinternalDescription MultilingualBillingentityinternalDescription, aObjBillingentityinternalproduct []BillingentityinternalproductResponseCompound) *BillingentityinternalResponseCompound {
	this := BillingentityinternalResponseCompound{}
	this.PkiBillingentityinternalID = pkiBillingentityinternalID
	this.ObjBillingentityinternalDescription = objBillingentityinternalDescription
	this.AObjBillingentityinternalproduct = aObjBillingentityinternalproduct
	return &this
}

// NewBillingentityinternalResponseCompoundWithDefaults instantiates a new BillingentityinternalResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalResponseCompoundWithDefaults() *BillingentityinternalResponseCompound {
	this := BillingentityinternalResponseCompound{}
	return &this
}

// GetPkiBillingentityinternalID returns the PkiBillingentityinternalID field value
func (o *BillingentityinternalResponseCompound) GetPkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBillingentityinternalID
}

// GetPkiBillingentityinternalIDOk returns a tuple with the PkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalResponseCompound) GetPkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBillingentityinternalID, true
}

// SetPkiBillingentityinternalID sets field value
func (o *BillingentityinternalResponseCompound) SetPkiBillingentityinternalID(v int32) {
	o.PkiBillingentityinternalID = v
}

// GetObjBillingentityinternalDescription returns the ObjBillingentityinternalDescription field value
func (o *BillingentityinternalResponseCompound) GetObjBillingentityinternalDescription() MultilingualBillingentityinternalDescription {
	if o == nil {
		var ret MultilingualBillingentityinternalDescription
		return ret
	}

	return o.ObjBillingentityinternalDescription
}

// GetObjBillingentityinternalDescriptionOk returns a tuple with the ObjBillingentityinternalDescription field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalResponseCompound) GetObjBillingentityinternalDescriptionOk() (*MultilingualBillingentityinternalDescription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjBillingentityinternalDescription, true
}

// SetObjBillingentityinternalDescription sets field value
func (o *BillingentityinternalResponseCompound) SetObjBillingentityinternalDescription(v MultilingualBillingentityinternalDescription) {
	o.ObjBillingentityinternalDescription = v
}

// GetAObjBillingentityinternalproduct returns the AObjBillingentityinternalproduct field value
func (o *BillingentityinternalResponseCompound) GetAObjBillingentityinternalproduct() []BillingentityinternalproductResponseCompound {
	if o == nil {
		var ret []BillingentityinternalproductResponseCompound
		return ret
	}

	return o.AObjBillingentityinternalproduct
}

// GetAObjBillingentityinternalproductOk returns a tuple with the AObjBillingentityinternalproduct field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalResponseCompound) GetAObjBillingentityinternalproductOk() ([]BillingentityinternalproductResponseCompound, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjBillingentityinternalproduct, true
}

// SetAObjBillingentityinternalproduct sets field value
func (o *BillingentityinternalResponseCompound) SetAObjBillingentityinternalproduct(v []BillingentityinternalproductResponseCompound) {
	o.AObjBillingentityinternalproduct = v
}

func (o BillingentityinternalResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBillingentityinternalID"] = o.PkiBillingentityinternalID
	toSerialize["objBillingentityinternalDescription"] = o.ObjBillingentityinternalDescription
	toSerialize["a_objBillingentityinternalproduct"] = o.AObjBillingentityinternalproduct
	return toSerialize, nil
}

type NullableBillingentityinternalResponseCompound struct {
	value *BillingentityinternalResponseCompound
	isSet bool
}

func (v NullableBillingentityinternalResponseCompound) Get() *BillingentityinternalResponseCompound {
	return v.value
}

func (v *NullableBillingentityinternalResponseCompound) Set(val *BillingentityinternalResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalResponseCompound(val *BillingentityinternalResponseCompound) *NullableBillingentityinternalResponseCompound {
	return &NullableBillingentityinternalResponseCompound{value: val, isSet: true}
}

func (v NullableBillingentityinternalResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


