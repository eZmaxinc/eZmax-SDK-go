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

// checks if the EzmaxinvoicingcontractResponseCompound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EzmaxinvoicingcontractResponseCompound{}

// EzmaxinvoicingcontractResponseCompound A Ezmaxinvoicingcontract Object
type EzmaxinvoicingcontractResponseCompound struct {
	EzmaxinvoicingcontractResponse
}

type _EzmaxinvoicingcontractResponseCompound EzmaxinvoicingcontractResponseCompound

// NewEzmaxinvoicingcontractResponseCompound instantiates a new EzmaxinvoicingcontractResponseCompound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEzmaxinvoicingcontractResponseCompound(pkiEzmaxinvoicingcontractID int32, eEzmaxinvoicingcontractPaymenttype FieldEEzmaxinvoicingcontractPaymenttype, iEzmaxinvoicingcontractLength int32, dtEzmaxinvoicingcontractStart string, dtEzmaxinvoicingcontractEnd string, dEzmaxinvoicingcontractLicense string, dEzmaxinvoicingcontract121qa string, bEzmaxinvoicingcontractEzsignallagents bool, objAudit CommonAudit) *EzmaxinvoicingcontractResponseCompound {
	this := EzmaxinvoicingcontractResponseCompound{}
	this.PkiEzmaxinvoicingcontractID = pkiEzmaxinvoicingcontractID
	this.EEzmaxinvoicingcontractPaymenttype = eEzmaxinvoicingcontractPaymenttype
	this.IEzmaxinvoicingcontractLength = iEzmaxinvoicingcontractLength
	this.DtEzmaxinvoicingcontractStart = dtEzmaxinvoicingcontractStart
	this.DtEzmaxinvoicingcontractEnd = dtEzmaxinvoicingcontractEnd
	this.DEzmaxinvoicingcontractLicense = dEzmaxinvoicingcontractLicense
	this.DEzmaxinvoicingcontract121qa = dEzmaxinvoicingcontract121qa
	this.BEzmaxinvoicingcontractEzsignallagents = bEzmaxinvoicingcontractEzsignallagents
	this.ObjAudit = objAudit
	return &this
}

// NewEzmaxinvoicingcontractResponseCompoundWithDefaults instantiates a new EzmaxinvoicingcontractResponseCompound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEzmaxinvoicingcontractResponseCompoundWithDefaults() *EzmaxinvoicingcontractResponseCompound {
	this := EzmaxinvoicingcontractResponseCompound{}
	return &this
}

func (o EzmaxinvoicingcontractResponseCompound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EzmaxinvoicingcontractResponseCompound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

func (o *EzmaxinvoicingcontractResponseCompound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pkiEzmaxinvoicingcontractID",
		"eEzmaxinvoicingcontractPaymenttype",
		"iEzmaxinvoicingcontractLength",
		"dtEzmaxinvoicingcontractStart",
		"dtEzmaxinvoicingcontractEnd",
		"dEzmaxinvoicingcontractLicense",
		"dEzmaxinvoicingcontract121qa",
		"bEzmaxinvoicingcontractEzsignallagents",
		"objAudit",
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

	varEzmaxinvoicingcontractResponseCompound := _EzmaxinvoicingcontractResponseCompound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEzmaxinvoicingcontractResponseCompound)

	if err != nil {
		return err
	}

	*o = EzmaxinvoicingcontractResponseCompound(varEzmaxinvoicingcontractResponseCompound)

	return err
}

type NullableEzmaxinvoicingcontractResponseCompound struct {
	value *EzmaxinvoicingcontractResponseCompound
	isSet bool
}

func (v NullableEzmaxinvoicingcontractResponseCompound) Get() *EzmaxinvoicingcontractResponseCompound {
	return v.value
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) Set(val *EzmaxinvoicingcontractResponseCompound) {
	v.value = val
	v.isSet = true
}

func (v NullableEzmaxinvoicingcontractResponseCompound) IsSet() bool {
	return v.isSet
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEzmaxinvoicingcontractResponseCompound(val *EzmaxinvoicingcontractResponseCompound) *NullableEzmaxinvoicingcontractResponseCompound {
	return &NullableEzmaxinvoicingcontractResponseCompound{value: val, isSet: true}
}

func (v NullableEzmaxinvoicingcontractResponseCompound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEzmaxinvoicingcontractResponseCompound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


