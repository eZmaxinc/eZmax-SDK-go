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

// checks if the BillingentityinternalproductRequestCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingentityinternalproductRequestCompound{}

// BillingentityinternalproductRequestCompound A Billingentityinternalproduct Object and children
type BillingentityinternalproductRequestCompound struct {
	BillingentityinternalproductRequest
}

type _BillingentityinternalproductRequestCompound BillingentityinternalproductRequestCompound

// NewBillingentityinternalproductRequestCompound instantiates a new BillingentityinternalproductRequestCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingentityinternalproductRequestCompound(fkiEzmaxproductID int32, fkiBillingentityexternalID int32) *BillingentityinternalproductRequestCompound {
	this := BillingentityinternalproductRequestCompound{}
	this.FkiEzmaxproductID = fkiEzmaxproductID
	this.FkiBillingentityexternalID = fkiBillingentityexternalID
	return &this
}

// NewBillingentityinternalproductRequestCompoundWithDefaults instantiates a new BillingentityinternalproductRequestCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingentityinternalproductRequestCompoundWithDefaults() *BillingentityinternalproductRequestCompound {
	this := BillingentityinternalproductRequestCompound{}
	return &this
}

func (o BillingentityinternalproductRequestCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingentityinternalproductRequestCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *BillingentityinternalproductRequestCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiEzmaxproductID",
		"fkiBillingentityexternalID",
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

	varBillingentityinternalproductRequestCompound := _BillingentityinternalproductRequestCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBillingentityinternalproductRequestCompound)

	if err != nil {
		return err
	}

	*o = BillingentityinternalproductRequestCompound(varBillingentityinternalproductRequestCompound)

	return err
}

type NullableBillingentityinternalproductRequestCompound struct {
	value *BillingentityinternalproductRequestCompound
	isSet bool
}

func (v NullableBillingentityinternalproductRequestCompound) Get() *BillingentityinternalproductRequestCompound {
	return v.value
}

func (v *NullableBillingentityinternalproductRequestCompound) Set(val *BillingentityinternalproductRequestCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingentityinternalproductRequestCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingentityinternalproductRequestCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingentityinternalproductRequestCompound(val *BillingentityinternalproductRequestCompound) *NullableBillingentityinternalproductRequestCompound {
	return &NullableBillingentityinternalproductRequestCompound{value: val, isSet: true}
}

func (v NullableBillingentityinternalproductRequestCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingentityinternalproductRequestCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


