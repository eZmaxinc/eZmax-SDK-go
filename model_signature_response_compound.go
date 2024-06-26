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
	"bytes"
	"fmt"
)

// checks if the SignatureResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureResponseCompound{}

// SignatureResponseCompound A Signature Object
type SignatureResponseCompound struct {
	// The unique ID of the Signature
	PkiSignatureID int32 `json:"pkiSignatureID"`
	// The URL of the SVG file for the Signature
	SSignatureUrl string `json:"sSignatureUrl"`
}

type _SignatureResponseCompound SignatureResponseCompound

// NewSignatureResponseCompound instantiates a new SignatureResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureResponseCompound(pkiSignatureID int32, sSignatureUrl string) *SignatureResponseCompound {
	this := SignatureResponseCompound{}
	this.PkiSignatureID = pkiSignatureID
	this.SSignatureUrl = sSignatureUrl
	return &this
}

// NewSignatureResponseCompoundWithDefaults instantiates a new SignatureResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureResponseCompoundWithDefaults() *SignatureResponseCompound {
	this := SignatureResponseCompound{}
	return &this
}

// GetPkiSignatureID returns the PkiSignatureID field value
func (o *SignatureResponseCompound) GetPkiSignatureID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiSignatureID
}

// GetPkiSignatureIDOk returns a tuple with the PkiSignatureID field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseCompound) GetPkiSignatureIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiSignatureID, true
}

// SetPkiSignatureID sets field value
func (o *SignatureResponseCompound) SetPkiSignatureID(v int32) {
	o.PkiSignatureID = v
}

// GetSSignatureUrl returns the SSignatureUrl field value
func (o *SignatureResponseCompound) GetSSignatureUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SSignatureUrl
}

// GetSSignatureUrlOk returns a tuple with the SSignatureUrl field value
// and a boolean to check if the value has been set.
func (o *SignatureResponseCompound) GetSSignatureUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SSignatureUrl, true
}

// SetSSignatureUrl sets field value
func (o *SignatureResponseCompound) SetSSignatureUrl(v string) {
	o.SSignatureUrl = v
}

func (o SignatureResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pkiSignatureID"] = o.PkiSignatureID
	toSerialize["sSignatureUrl"] = o.SSignatureUrl
	return toSerialize, nil
}

func (o *SignatureResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiSignatureID",
		"sSignatureUrl",
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

	varSignatureResponseCompound := _SignatureResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureResponseCompound)

	if err != nil {
		return err
	}

	*o = SignatureResponseCompound(varSignatureResponseCompound)

	return err
}

type NullableSignatureResponseCompound struct {
	value *SignatureResponseCompound
	isSet bool
}

func (v NullableSignatureResponseCompound) Get() *SignatureResponseCompound {
	return v.value
}

func (v *NullableSignatureResponseCompound) Set(val *SignatureResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureResponseCompound(val *SignatureResponseCompound) *NullableSignatureResponseCompound {
	return &NullableSignatureResponseCompound{value: val, isSet: true}
}

func (v NullableSignatureResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


