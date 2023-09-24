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

// checks if the SignatureRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureRequestCompound{}

// SignatureRequestCompound A Signature Object and children
type SignatureRequestCompound struct {
	// The unique ID of the Signature
	PkiSignatureID *int32 `json:"pkiSignatureID,omitempty"`
	// The svg of the Signature
	TSignatureSvg string `json:"tSignatureSvg"`
}

// NewSignatureRequestCompound instantiates a new SignatureRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestCompound(tSignatureSvg string) *SignatureRequestCompound {
	this := SignatureRequestCompound{}
	this.TSignatureSvg = tSignatureSvg
	return &this
}

// NewSignatureRequestCompoundWithDefaults instantiates a new SignatureRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestCompoundWithDefaults() *SignatureRequestCompound {
	this := SignatureRequestCompound{}
	return &this
}

// GetPkiSignatureID returns the PkiSignatureID field value if set, zero value otherwise.
func (o *SignatureRequestCompound) GetPkiSignatureID() int32 {
	if o == nil || IsNil(o.PkiSignatureID) {
		var ret int32
		return ret
	}
	return *o.PkiSignatureID
}

// GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestCompound) GetPkiSignatureIDOk() (*int32, bool) {
	if o == nil || IsNil(o.PkiSignatureID) {
		return nil, false
	}
	return o.PkiSignatureID, true
}

// HasPkiSignatureID returns a boolean if a field has been set.
func (o *SignatureRequestCompound) HasPkiSignatureID() bool {
	if o != nil && !IsNil(o.PkiSignatureID) {
		return true
	}

	return false
}

// SetPkiSignatureID gets a reference to the given int32 and assigns it to the PkiSignatureID field.
func (o *SignatureRequestCompound) SetPkiSignatureID(v int32) {
	o.PkiSignatureID = &v
}

// GetTSignatureSvg returns the TSignatureSvg field value
func (o *SignatureRequestCompound) GetTSignatureSvg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TSignatureSvg
}

// GetTSignatureSvgOk returns a tuple with the TSignatureSvg field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestCompound) GetTSignatureSvgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TSignatureSvg, true
}

// SetTSignatureSvg sets field value
func (o *SignatureRequestCompound) SetTSignatureSvg(v string) {
	o.TSignatureSvg = v
}

func (o SignatureRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PkiSignatureID) {
		toSerialize["pkiSignatureID"] = o.PkiSignatureID
	}
	toSerialize["tSignatureSvg"] = o.TSignatureSvg
	return toSerialize, nil
}

type NullableSignatureRequestCompound struct {
	value *SignatureRequestCompound
	isSet bool
}

func (v NullableSignatureRequestCompound) Get() *SignatureRequestCompound {
	return v.value
}

func (v *NullableSignatureRequestCompound) Set(val *SignatureRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestCompound(val *SignatureRequestCompound) *NullableSignatureRequestCompound {
	return &NullableSignatureRequestCompound{value: val, isSet: true}
}

func (v NullableSignatureRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


