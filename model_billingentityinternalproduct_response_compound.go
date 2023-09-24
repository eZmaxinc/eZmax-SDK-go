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

// checks if the BillingentityinternalproductResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalproductResponseCompound{}

// BillingentityinternalproductResponseCompound A Billingentityinternalproduct Object
type BillingentityinternalproductResponseCompound struct {
	// The unique ID of the Billingentityinternalproduct
	PkiBillingentityinternalproductID int32 `json:"pkiBillingentityinternalproductID"`
	// The unique ID of the Billingentityinternal.
	FkiBillingentityinternalID int32 `json:"fkiBillingentityinternalID"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX string `json:"sBillingentityinternalDescriptionX"`
	// The unique ID of the Ezmaxproduct
	FkiEzmaxproductID int32 `json:"fkiEzmaxproductID"`
	// The description of the Ezmaxproduct in the language of the requester
	SEzmaxproductDescriptionX string `json:"sEzmaxproductDescriptionX"`
	// The unique ID of the Billingentityexternal
	FkiBillingentityexternalID int32 `json:"fkiBillingentityexternalID"`
	// The description of the Billingentityexternal
	SBillingentityexternalDescription string `json:"sBillingentityexternalDescription"`
}

// NewBillingentityinternalproductResponseCompound instantiates a new BillingentityinternalproductResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalproductResponseCompound(pkiBillingentityinternalproductID int32, fkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string, fkiEzmaxproductID int32, sEzmaxproductDescriptionX string, fkiBillingentityexternalID int32, sBillingentityexternalDescription string) *BillingentityinternalproductResponseCompound {
	this := BillingentityinternalproductResponseCompound{}
	this.PkiBillingentityinternalproductID = pkiBillingentityinternalproductID
	this.FkiBillingentityinternalID = fkiBillingentityinternalID
	this.SBillingentityinternalDescriptionX = sBillingentityinternalDescriptionX
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.SEzmaxproductDescriptionX = sEzmaxproductDescriptionX
	this.FkiBillingentityexternalID = fkiBillingentityexternalID
	this.SBillingentityexternalDescription = sBillingentityexternalDescription
	return &this
}

// NewBillingentityinternalproductResponseCompoundWithDefaults instantiates a new BillingentityinternalproductResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalproductResponseCompoundWithDefaults() *BillingentityinternalproductResponseCompound {
	this := BillingentityinternalproductResponseCompound{}
	return &this
}

// GetPkiBillingentityinternalproductID returns the PkiBillingentityinternalproductID field value
func (o *BillingentityinternalproductResponseCompound) GetPkiBillingentityinternalproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBillingentityinternalproductID
}

// GetPkiBillingentityinternalproductIDOk returns a tuple with the PkiBillingentityinternalproductID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetPkiBillingentityinternalproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBillingentityinternalproductID, true
}

// SetPkiBillingentityinternalproductID sets field value
func (o *BillingentityinternalproductResponseCompound) SetPkiBillingentityinternalproductID(v int32) {
	o.PkiBillingentityinternalproductID = v
}

// GetFkiBillingentityinternalID returns the FkiBillingentityinternalID field value
func (o *BillingentityinternalproductResponseCompound) GetFkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityinternalID
}

// GetFkiBillingentityinternalIDOk returns a tuple with the FkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetFkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityinternalID, true
}

// SetFkiBillingentityinternalID sets field value
func (o *BillingentityinternalproductResponseCompound) SetFkiBillingentityinternalID(v int32) {
	o.FkiBillingentityinternalID = v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value
func (o *BillingentityinternalproductResponseCompound) GetSBillingentityinternalDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityinternalDescriptionX, true
}

// SetSBillingentityinternalDescriptionX sets field value
func (o *BillingentityinternalproductResponseCompound) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = v
}

// GetFkiEzmaxproductID returns the FkiEzmaxproductID field value
func (o *BillingentityinternalproductResponseCompound) GetFkiEzmaxproductID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiEzmaxproductID
}

// GetFkiEzmaxproductIDOk returns a tuple with the FkiEzmaxproductID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetFkiEzmaxproductIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiEzmaxproductID, true
}

// SetFkiEzmaxproductID sets field value
func (o *BillingentityinternalproductResponseCompound) SetFkiEzmaxproductID(v int32) {
	o.FkiEzmaxproductID = v
}

// GetSEzmaxproductDescriptionX returns the SEzmaxproductDescriptionX field value
func (o *BillingentityinternalproductResponseCompound) GetSEzmaxproductDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzmaxproductDescriptionX
}

// GetSEzmaxproductDescriptionXOk returns a tuple with the SEzmaxproductDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetSEzmaxproductDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SEzmaxproductDescriptionX, true
}

// SetSEzmaxproductDescriptionX sets field value
func (o *BillingentityinternalproductResponseCompound) SetSEzmaxproductDescriptionX(v string) {
	o.SEzmaxproductDescriptionX = v
}

// GetFkiBillingentityexternalID returns the FkiBillingentityexternalID field value
func (o *BillingentityinternalproductResponseCompound) GetFkiBillingentityexternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiBillingentityexternalID
}

// GetFkiBillingentityexternalIDOk returns a tuple with the FkiBillingentityexternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetFkiBillingentityexternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiBillingentityexternalID, true
}

// SetFkiBillingentityexternalID sets field value
func (o *BillingentityinternalproductResponseCompound) SetFkiBillingentityexternalID(v int32) {
	o.FkiBillingentityexternalID = v
}

// GetSBillingentityexternalDescription returns the SBillingentityexternalDescription field value
func (o *BillingentityinternalproductResponseCompound) GetSBillingentityexternalDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityexternalDescription
}

// GetSBillingentityexternalDescriptionOk returns a tuple with the SBillingentityexternalDescription field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalproductResponseCompound) GetSBillingentityexternalDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityexternalDescription, true
}

// SetSBillingentityexternalDescription sets field value
func (o *BillingentityinternalproductResponseCompound) SetSBillingentityexternalDescription(v string) {
	o.SBillingentityexternalDescription = v
}

func (o BillingentityinternalproductResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalproductResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBillingentityinternalproductID"] = o.PkiBillingentityinternalproductID
	toSerialize["fkiBillingentityinternalID"] = o.FkiBillingentityinternalID
	toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	toSerialize["fkiEzmaxproductID"] = o.FkiEzmaxproductID
	toSerialize["sEzmaxproductDescriptionX"] = o.SEzmaxproductDescriptionX
	toSerialize["fkiBillingentityexternalID"] = o.FkiBillingentityexternalID
	toSerialize["sBillingentityexternalDescription"] = o.SBillingentityexternalDescription
	return toSerialize, nil
}

type NullableBillingentityinternalproductResponseCompound struct {
	value *BillingentityinternalproductResponseCompound
	isSet bool
}

func (v NullableBillingentityinternalproductResponseCompound) Get() *BillingentityinternalproductResponseCompound {
	return v.value
}

func (v *NullableBillingentityinternalproductResponseCompound) Set(val *BillingentityinternalproductResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalproductResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalproductResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalproductResponseCompound(val *BillingentityinternalproductResponseCompound) *NullableBillingentityinternalproductResponseCompound {
	return &NullableBillingentityinternalproductResponseCompound{value: val, isSet: true}
}

func (v NullableBillingentityinternalproductResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalproductResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


