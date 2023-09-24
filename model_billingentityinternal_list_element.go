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

// checks if the BillingentityinternalListElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalListElement{}

// BillingentityinternalListElement A Billingentityinternal List Element
type BillingentityinternalListElement struct {
	// The unique ID of the Billingentityinternal.
	PkiBillingentityinternalID int32 `json:"pkiBillingentityinternalID"`
	// The description of the Billingentityinternal in the language of the requester
	SBillingentityinternalDescriptionX string `json:"sBillingentityinternalDescriptionX"`
}

// NewBillingentityinternalListElement instantiates a new BillingentityinternalListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalListElement(pkiBillingentityinternalID int32, sBillingentityinternalDescriptionX string) *BillingentityinternalListElement {
	this := BillingentityinternalListElement{}
	this.PkiBillingentityinternalID = pkiBillingentityinternalID
	this.SBillingentityinternalDescriptionX = sBillingentityinternalDescriptionX
	return &this
}

// NewBillingentityinternalListElementWithDefaults instantiates a new BillingentityinternalListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalListElementWithDefaults() *BillingentityinternalListElement {
	this := BillingentityinternalListElement{}
	return &this
}

// GetPkiBillingentityinternalID returns the PkiBillingentityinternalID field value
func (o *BillingentityinternalListElement) GetPkiBillingentityinternalID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiBillingentityinternalID
}

// GetPkiBillingentityinternalIDOk returns a tuple with the PkiBillingentityinternalID field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalListElement) GetPkiBillingentityinternalIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiBillingentityinternalID, true
}

// SetPkiBillingentityinternalID sets field value
func (o *BillingentityinternalListElement) SetPkiBillingentityinternalID(v int32) {
	o.PkiBillingentityinternalID = v
}

// GetSBillingentityinternalDescriptionX returns the SBillingentityinternalDescriptionX field value
func (o *BillingentityinternalListElement) GetSBillingentityinternalDescriptionX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SBillingentityinternalDescriptionX
}

// GetSBillingentityinternalDescriptionXOk returns a tuple with the SBillingentityinternalDescriptionX field value
// and a boolean to check if the value has been set.
func (o *BillingentityinternalListElement) GetSBillingentityinternalDescriptionXOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SBillingentityinternalDescriptionX, true
}

// SetSBillingentityinternalDescriptionX sets field value
func (o *BillingentityinternalListElement) SetSBillingentityinternalDescriptionX(v string) {
	o.SBillingentityinternalDescriptionX = v
}

func (o BillingentityinternalListElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalListElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiBillingentityinternalID"] = o.PkiBillingentityinternalID
	toSerialize["sBillingentityinternalDescriptionX"] = o.SBillingentityinternalDescriptionX
	return toSerialize, nil
}

type NullableBillingentityinternalListElement struct {
	value *BillingentityinternalListElement
	isSet bool
}

func (v NullableBillingentityinternalListElement) Get() *BillingentityinternalListElement {
	return v.value
}

func (v *NullableBillingentityinternalListElement) Set(val *BillingentityinternalListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalListElement(val *BillingentityinternalListElement) *NullableBillingentityinternalListElement {
	return &NullableBillingentityinternalListElement{value: val, isSet: true}
}

func (v NullableBillingentityinternalListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

