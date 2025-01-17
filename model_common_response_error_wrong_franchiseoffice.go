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

// checks if the CommonResponseErrorWrongFranchiseoffice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonResponseErrorWrongFranchiseoffice{}

// CommonResponseErrorWrongFranchiseoffice Error Message when a Franchisebroker is not in this Franchiseoffice.
type CommonResponseErrorWrongFranchiseoffice struct {
	CommonResponseError
	// The unique ID of the Franchiseagence
	FkiFranchiseagenceID int32 `json:"fkiFranchiseagenceID"`
	// The name of the Franchiseagence
	SFranchiseagenceName string `json:"sFranchiseagenceName"`
	// The unique ID of the Franchisereoffice
	FkiFranchiseofficeID int32 `json:"fkiFranchiseofficeID"`
	// The code of the Franchiseoffice
	IFranchiseofficeCode string `json:"iFranchiseofficeCode"`
}

type _CommonResponseErrorWrongFranchiseoffice CommonResponseErrorWrongFranchiseoffice

// NewCommonResponseErrorWrongFranchiseoffice instantiates a new CommonResponseErrorWrongFranchiseoffice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonResponseErrorWrongFranchiseoffice(fkiFranchiseagenceID int32, sFranchiseagenceName string, fkiFranchiseofficeID int32, iFranchiseofficeCode string, sErrorMessage string, eErrorCode FieldEErrorCode) *CommonResponseErrorWrongFranchiseoffice {
	this := CommonResponseErrorWrongFranchiseoffice{}
	this.SErrorMessage = sErrorMessage
	this.EErrorCode = eErrorCode
	this.FkiFranchiseagenceID = fkiFranchiseagenceID
	this.SFranchiseagenceName = sFranchiseagenceName
	this.FkiFranchiseofficeID = fkiFranchiseofficeID
	this.IFranchiseofficeCode = iFranchiseofficeCode
	return &this
}

// NewCommonResponseErrorWrongFranchiseofficeWithDefaults instantiates a new CommonResponseErrorWrongFranchiseoffice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonResponseErrorWrongFranchiseofficeWithDefaults() *CommonResponseErrorWrongFranchiseoffice {
	this := CommonResponseErrorWrongFranchiseoffice{}
	return &this
}

// GetFkiFranchiseagenceID returns the FkiFranchiseagenceID field value
func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseagenceID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchiseagenceID
}

// GetFkiFranchiseagenceIDOk returns a tuple with the FkiFranchiseagenceID field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseagenceIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiFranchiseagenceID, true
}

// SetFkiFranchiseagenceID sets field value
func (o *CommonResponseErrorWrongFranchiseoffice) SetFkiFranchiseagenceID(v int32) {
	o.FkiFranchiseagenceID = v
}

// GetSFranchiseagenceName returns the SFranchiseagenceName field value
func (o *CommonResponseErrorWrongFranchiseoffice) GetSFranchiseagenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SFranchiseagenceName
}

// GetSFranchiseagenceNameOk returns a tuple with the SFranchiseagenceName field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorWrongFranchiseoffice) GetSFranchiseagenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SFranchiseagenceName, true
}

// SetSFranchiseagenceName sets field value
func (o *CommonResponseErrorWrongFranchiseoffice) SetSFranchiseagenceName(v string) {
	o.SFranchiseagenceName = v
}

// GetFkiFranchiseofficeID returns the FkiFranchiseofficeID field value
func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseofficeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FkiFranchiseofficeID
}

// GetFkiFranchiseofficeIDOk returns a tuple with the FkiFranchiseofficeID field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorWrongFranchiseoffice) GetFkiFranchiseofficeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FkiFranchiseofficeID, true
}

// SetFkiFranchiseofficeID sets field value
func (o *CommonResponseErrorWrongFranchiseoffice) SetFkiFranchiseofficeID(v int32) {
	o.FkiFranchiseofficeID = v
}

// GetIFranchiseofficeCode returns the IFranchiseofficeCode field value
func (o *CommonResponseErrorWrongFranchiseoffice) GetIFranchiseofficeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IFranchiseofficeCode
}

// GetIFranchiseofficeCodeOk returns a tuple with the IFranchiseofficeCode field value
// and a boolean to check if the value has been set.
func (o *CommonResponseErrorWrongFranchiseoffice) GetIFranchiseofficeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IFranchiseofficeCode, true
}

// SetIFranchiseofficeCode sets field value
func (o *CommonResponseErrorWrongFranchiseoffice) SetIFranchiseofficeCode(v string) {
	o.IFranchiseofficeCode = v
}

func (o CommonResponseErrorWrongFranchiseoffice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonResponseErrorWrongFranchiseoffice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fkiFranchiseagenceID"] = o.FkiFranchiseagenceID
	toSerialize["sFranchiseagenceName"] = o.SFranchiseagenceName
	toSerialize["fkiFranchiseofficeID"] = o.FkiFranchiseofficeID
	toSerialize["iFranchiseofficeCode"] = o.IFranchiseofficeCode
	return toSerialize, nil
}

func (o *CommonResponseErrorWrongFranchiseoffice) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fkiFranchiseagenceID",
		"sFranchiseagenceName",
		"fkiFranchiseofficeID",
		"iFranchiseofficeCode",
		"sErrorMessage",
		"eErrorCode",
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

	varCommonResponseErrorWrongFranchiseoffice := _CommonResponseErrorWrongFranchiseoffice{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommonResponseErrorWrongFranchiseoffice)

	if err != nil {
		return err
	}

	*o = CommonResponseErrorWrongFranchiseoffice(varCommonResponseErrorWrongFranchiseoffice)

	return err
}

type NullableCommonResponseErrorWrongFranchiseoffice struct {
	value *CommonResponseErrorWrongFranchiseoffice
	isSet bool
}

func (v NullableCommonResponseErrorWrongFranchiseoffice) Get() *CommonResponseErrorWrongFranchiseoffice {
	return v.value
}

func (v *NullableCommonResponseErrorWrongFranchiseoffice) Set(val *CommonResponseErrorWrongFranchiseoffice) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonResponseErrorWrongFranchiseoffice) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonResponseErrorWrongFranchiseoffice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonResponseErrorWrongFranchiseoffice(val *CommonResponseErrorWrongFranchiseoffice) *NullableCommonResponseErrorWrongFranchiseoffice {
	return &NullableCommonResponseErrorWrongFranchiseoffice{value: val, isSet: true}
}

func (v NullableCommonResponseErrorWrongFranchiseoffice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonResponseErrorWrongFranchiseoffice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


