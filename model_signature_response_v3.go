/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.1
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SignatureResponseV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureResponseV3{}

// SignatureResponseV3 A Signature Object
type SignatureResponseV3 struct {
	// The unique ID of the Signature
	PkiSignatureID int32 `json:"pkiSignatureID"`
	// The unique ID of the Font
	FkiFontID int32 `json:"fkiFontID"`
	ESignaturePreference FieldESignaturePreference `json:"eSignaturePreference"`
	// Whether the signature has a SVG or not
	BSignatureSvg bool `json:"bSignatureSvg"`
	// Whether the initials has a SVG or not
	BSignatureSvginitials bool `json:"bSignatureSvginitials"`
}

type _SignatureResponseV3 SignatureResponseV3

// NewSignatureResponseV3 instantiates a new SignatureResponseV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureResponseV3(pkiSignatureID int32, fkiFontID int32, eSignaturePreference FieldESignaturePreference, bSignatureSvg bool, bSignatureSvginitials bool) *SignatureResponseV3 {
	this := SignatureResponseV3{}
	this.PkiSignatureID = pkiSignatureID
	this.FkiFontID = fkiFontID
	this.ESignaturePreference = eSignaturePreference
	this.BSignatureSvg = bSignatureSvg
	this.BSignatureSvginitials = bSignatureSvginitials
	return &this
}

// NewSignatureResponseV3WithDefaults instantiates a new SignatureResponseV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureResponseV3WithDefaults() *SignatureResponseV3 {
	this := SignatureResponseV3{}
	return &this
}

// GetPkiSignatureID returns the PkiSignatureID field value
func (o *SignatureResponseV3) GetPkiSignatureID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSignatureID
}

// GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseV3) GetPkiSignatureIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSignatureID, true
}

// SetPkiSignatureID sets field value
func (o *SignatureResponseV3) SetPkiSignatureID(v int32) {
	o.PkiSignatureID = v
}

// GetFkiFontID returns the FkiFontID field value
func (o *SignatureResponseV3) GetFkiFontID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFontID
}

// GetFkiFontIDOk returns a tuple with the FkiFontID field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseV3) GetFkiFontIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiFontID, true
}

// SetFkiFontID sets field value
func (o *SignatureResponseV3) SetFkiFontID(v int32) {
	o.FkiFontID = v
}

// GetESignaturePreference returns the ESignaturePreference field value
func (o *SignatureResponseV3) GetESignaturePreference() FieldESignaturePreference {
	if o == nil {
		var ret FieldESignaturePreference
		return ret
	}

	return o.ESignaturePreference
}

// GetESignaturePreferenceOk returns a tuple with the ESignaturePreference field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseV3) GetESignaturePreferenceOk() (*FieldESignaturePreference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ESignaturePreference, true
}

// SetESignaturePreference sets field value
func (o *SignatureResponseV3) SetESignaturePreference(v FieldESignaturePreference) {
	o.ESignaturePreference = v
}

// GetBSignatureSvg returns the BSignatureSvg field value
func (o *SignatureResponseV3) GetBSignatureSvg() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSignatureSvg
}

// GetBSignatureSvgOk returns a tuple with the BSignatureSvg field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseV3) GetBSignatureSvgOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSignatureSvg, true
}

// SetBSignatureSvg sets field value
func (o *SignatureResponseV3) SetBSignatureSvg(v bool) {
	o.BSignatureSvg = v
}

// GetBSignatureSvginitials returns the BSignatureSvginitials field value
func (o *SignatureResponseV3) GetBSignatureSvginitials() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.BSignatureSvginitials
}

// GetBSignatureSvginitialsOk returns a tuple with the BSignatureSvginitials field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseV3) GetBSignatureSvginitialsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BSignatureSvginitials, true
}

// SetBSignatureSvginitials sets field value
func (o *SignatureResponseV3) SetBSignatureSvginitials(v bool) {
	o.BSignatureSvginitials = v
}

func (o SignatureResponseV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureResponseV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSignatureID"] = o.PkiSignatureID
	toSerialize["fkiFontID"] = o.FkiFontID
	toSerialize["eSignaturePreference"] = o.ESignaturePreference
	toSerialize["bSignatureSvg"] = o.BSignatureSvg
	toSerialize["bSignatureSvginitials"] = o.BSignatureSvginitials
	return toSerialize, nil
}

func (o *SignatureResponseV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSignatureID",
		"fkiFontID",
		"eSignaturePreference",
		"bSignatureSvg",
		"bSignatureSvginitials",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSignatureResponseV3 := _SignatureResponseV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureResponseV3)

	if err != nil {
		return err
	}

	*o = SignatureResponseV3(varSignatureResponseV3)

	return err
}

type NullableSignatureResponseV3 struct {
	value *SignatureResponseV3
	isSet bool
}

func (v NullableSignatureResponseV3) Get() *SignatureResponseV3 {
	return v.value
}

func (v *NullableSignatureResponseV3) Set(val *SignatureResponseV3) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureResponseV3) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureResponseV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureResponseV3(val *SignatureResponseV3) *NullableSignatureResponseV3 {
	return &NullableSignatureResponseV3{value: val, isSet: true}
}

func (v NullableSignatureResponseV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureResponseV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


