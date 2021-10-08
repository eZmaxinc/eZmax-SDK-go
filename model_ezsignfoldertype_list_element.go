/*
eZmax API Definition

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.1.0
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
)

// EzsignfoldertypeListElement An Ezsignfoldertype List Element
type EzsignfoldertypeListElement struct {
	// The unique ID of the Ezsignfoldertype.
	PkiEzsignfoldertypeID int32 `json:"pkiEzsignfoldertypeID"`
	EEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel `json:"eEzsignfoldertypePrivacylevel"`
	// The name of the Ezsignfoldertype in the language of the requester
	SEzsignfoldertypeNameX string `json:"sEzsignfoldertypeNameX"`
	// Whether the Ezsignfoldertype is active or not
	BEzsignfoldertypeIsactive bool `json:"bEzsignfoldertypeIsactive"`
}

// NewEzsignfoldertypeListElement instantiates a new EzsignfoldertypeListElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzsignfoldertypeListElement(pkiEzsignfoldertypeID int32, eEzsignfoldertypePrivacylevel FieldEEzsignfoldertypePrivacylevel, sEzsignfoldertypeNameX string, bEzsignfoldertypeIsactive bool) *EzsignfoldertypeListElement {
	this := EzsignfoldertypeListElement{}
	this.PkiEzsignfoldertypeID = pkiEzsignfoldertypeID
	this.EEzsignfoldertypePrivacylevel = eEzsignfoldertypePrivacylevel
	this.SEzsignfoldertypeNameX = sEzsignfoldertypeNameX
	this.BEzsignfoldertypeIsactive = bEzsignfoldertypeIsactive
	return &this
}

// NewEzsignfoldertypeListElementWithDefaults instantiates a new EzsignfoldertypeListElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzsignfoldertypeListElementWithDefaults() *EzsignfoldertypeListElement {
	this := EzsignfoldertypeListElement{}
	return &this
}

// GetPkiEzsignfoldertypeID returns the PkiEzsignfoldertypeID field value
func (o *EzsignfoldertypeListElement) GetPkiEzsignfoldertypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiEzsignfoldertypeID
}

// GetPkiEzsignfoldertypeIDOk returns a tuple with the PkiEzsignfoldertypeID field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeListElement) GetPkiEzsignfoldertypeIDOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PkiEzsignfoldertypeID, true
}

// SetPkiEzsignfoldertypeID sets field value
func (o *EzsignfoldertypeListElement) SetPkiEzsignfoldertypeID(v int32) {
	o.PkiEzsignfoldertypeID = v
}

// GetEEzsignfoldertypePrivacylevel returns the EEzsignfoldertypePrivacylevel field value
func (o *EzsignfoldertypeListElement) GetEEzsignfoldertypePrivacylevel() FieldEEzsignfoldertypePrivacylevel {
	if o == nil {
		var ret FieldEEzsignfoldertypePrivacylevel
		return ret
	}

	return o.EEzsignfoldertypePrivacylevel
}

// GetEEzsignfoldertypePrivacylevelOk returns a tuple with the EEzsignfoldertypePrivacylevel field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeListElement) GetEEzsignfoldertypePrivacylevelOk() (*FieldEEzsignfoldertypePrivacylevel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EEzsignfoldertypePrivacylevel, true
}

// SetEEzsignfoldertypePrivacylevel sets field value
func (o *EzsignfoldertypeListElement) SetEEzsignfoldertypePrivacylevel(v FieldEEzsignfoldertypePrivacylevel) {
	o.EEzsignfoldertypePrivacylevel = v
}

// GetSEzsignfoldertypeNameX returns the SEzsignfoldertypeNameX field value
func (o *EzsignfoldertypeListElement) GetSEzsignfoldertypeNameX() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SEzsignfoldertypeNameX
}

// GetSEzsignfoldertypeNameXOk returns a tuple with the SEzsignfoldertypeNameX field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeListElement) GetSEzsignfoldertypeNameXOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SEzsignfoldertypeNameX, true
}

// SetSEzsignfoldertypeNameX sets field value
func (o *EzsignfoldertypeListElement) SetSEzsignfoldertypeNameX(v string) {
	o.SEzsignfoldertypeNameX = v
}

// GetBEzsignfoldertypeIsactive returns the BEzsignfoldertypeIsactive field value
func (o *EzsignfoldertypeListElement) GetBEzsignfoldertypeIsactive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BEzsignfoldertypeIsactive
}

// GetBEzsignfoldertypeIsactiveOk returns a tuple with the BEzsignfoldertypeIsactive field value
// and a boolean to check if the value has been set.
func (o *EzsignfoldertypeListElement) GetBEzsignfoldertypeIsactiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BEzsignfoldertypeIsactive, true
}

// SetBEzsignfoldertypeIsactive sets field value
func (o *EzsignfoldertypeListElement) SetBEzsignfoldertypeIsactive(v bool) {
	o.BEzsignfoldertypeIsactive = v
}

func (o EzsignfoldertypeListElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pkiEzsignfoldertypeID"] = o.PkiEzsignfoldertypeID
	}
	if true {
		toSerialize["eEzsignfoldertypePrivacylevel"] = o.EEzsignfoldertypePrivacylevel
	}
	if true {
		toSerialize["sEzsignfoldertypeNameX"] = o.SEzsignfoldertypeNameX
	}
	if true {
		toSerialize["bEzsignfoldertypeIsactive"] = o.BEzsignfoldertypeIsactive
	}
	return json.Marshal(toSerialize)
}

type NullableEzsignfoldertypeListElement struct {
	value *EzsignfoldertypeListElement
	isSet bool
}

func (v NullableEzsignfoldertypeListElement) Get() *EzsignfoldertypeListElement {
	return v.value
}

func (v *NullableEzsignfoldertypeListElement) Set(val *EzsignfoldertypeListElement) {
	v.value = val
	v.isSet = true
}

func (v NullableEzsignfoldertypeListElement) IsSet() bool {
	return v.isSet
}

func (v *NullableEzsignfoldertypeListElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzsignfoldertypeListElement(val *EzsignfoldertypeListElement) *NullableEzsignfoldertypeListElement {
	return &NullableEzsignfoldertypeListElement{value: val, isSet: true}
}

func (v NullableEzsignfoldertypeListElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzsignfoldertypeListElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


