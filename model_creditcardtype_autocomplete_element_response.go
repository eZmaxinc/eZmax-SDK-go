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

// checks if the CreditcardtypeAutocompleteElementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardtypeAutocompleteElementResponse{}

// CreditcardtypeAutocompleteElementResponse Creditcardtype AutocompleteElement Response
type CreditcardtypeAutocompleteElementResponse struct {
	// The name of the Creditcardtype
	SCreditcardtypeName string `json:"sCreditcardtypeName"`
	// The unique ID of the Creditcardtype
	PkiCreditcardtypeID int32 `json:"pkiCreditcardtypeID"`
	ECreditcardtypeCodename FieldECreditcardtypeCodename `json:"eCreditcardtypeCodename"`
}

type _CreditcardtypeAutocompleteElementResponse CreditcardtypeAutocompleteElementResponse

// NewCreditcardtypeAutocompleteElementResponse instantiates a new CreditcardtypeAutocompleteElementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardtypeAutocompleteElementResponse(sCreditcardtypeName string, pkiCreditcardtypeID int32, eCreditcardtypeCodename FieldECreditcardtypeCodename) *CreditcardtypeAutocompleteElementResponse {
	this := CreditcardtypeAutocompleteElementResponse{}
	this.SCreditcardtypeName = sCreditcardtypeName
	this.PkiCreditcardtypeID = pkiCreditcardtypeID
	this.ECreditcardtypeCodename = eCreditcardtypeCodename
	return &this
}

// NewCreditcardtypeAutocompleteElementResponseWithDefaults instantiates a new CreditcardtypeAutocompleteElementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardtypeAutocompleteElementResponseWithDefaults() *CreditcardtypeAutocompleteElementResponse {
	this := CreditcardtypeAutocompleteElementResponse{}
	return &this
}

// GetSCreditcardtypeName returns the SCreditcardtypeName field value
func (o *CreditcardtypeAutocompleteElementResponse) GetSCreditcardtypeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardtypeName
}

// GetSCreditcardtypeNameOk returns a tuple with the SCreditcardtypeName field value
// and a boolean to check if the value has been set.
func (o *CreditcardtypeAutocompleteElementResponse) GetSCreditcardtypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardtypeName, true
}

// SetSCreditcardtypeName sets field value
func (o *CreditcardtypeAutocompleteElementResponse) SetSCreditcardtypeName(v string) {
	o.SCreditcardtypeName = v
}

// GetPkiCreditcardtypeID returns the PkiCreditcardtypeID field value
func (o *CreditcardtypeAutocompleteElementResponse) GetPkiCreditcardtypeID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PkiCreditcardtypeID
}

// GetPkiCreditcardtypeIDOk returns a tuple with the PkiCreditcardtypeID field value
// and a boolean to check if the value has been set.
func (o *CreditcardtypeAutocompleteElementResponse) GetPkiCreditcardtypeIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PkiCreditcardtypeID, true
}

// SetPkiCreditcardtypeID sets field value
func (o *CreditcardtypeAutocompleteElementResponse) SetPkiCreditcardtypeID(v int32) {
	o.PkiCreditcardtypeID = v
}

// GetECreditcardtypeCodename returns the ECreditcardtypeCodename field value
func (o *CreditcardtypeAutocompleteElementResponse) GetECreditcardtypeCodename() FieldECreditcardtypeCodename {
	if o == nil {
		var ret FieldECreditcardtypeCodename
		return ret
	}

	return o.ECreditcardtypeCodename
}

// GetECreditcardtypeCodenameOk returns a tuple with the ECreditcardtypeCodename field value
// and a boolean to check if the value has been set.
func (o *CreditcardtypeAutocompleteElementResponse) GetECreditcardtypeCodenameOk() (*FieldECreditcardtypeCodename, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ECreditcardtypeCodename, true
}

// SetECreditcardtypeCodename sets field value
func (o *CreditcardtypeAutocompleteElementResponse) SetECreditcardtypeCodename(v FieldECreditcardtypeCodename) {
	o.ECreditcardtypeCodename = v
}

func (o CreditcardtypeAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardtypeAutocompleteElementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sCreditcardtypeName"] = o.SCreditcardtypeName
	toSerialize["pkiCreditcardtypeID"] = o.PkiCreditcardtypeID
	toSerialize["eCreditcardtypeCodename"] = o.ECreditcardtypeCodename
	return toSerialize, nil
}

func (o *CreditcardtypeAutocompleteElementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sCreditcardtypeName",
		"pkiCreditcardtypeID",
		"eCreditcardtypeCodename",
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

	varCreditcardtypeAutocompleteElementResponse := _CreditcardtypeAutocompleteElementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardtypeAutocompleteElementResponse)

	if err != nil {
		return err
	}

	*o = CreditcardtypeAutocompleteElementResponse(varCreditcardtypeAutocompleteElementResponse)

	return err
}

type NullableCreditcardtypeAutocompleteElementResponse struct {
	value *CreditcardtypeAutocompleteElementResponse
	isSet bool
}

func (v NullableCreditcardtypeAutocompleteElementResponse) Get() *CreditcardtypeAutocompleteElementResponse {
	return v.value
}

func (v *NullableCreditcardtypeAutocompleteElementResponse) Set(val *CreditcardtypeAutocompleteElementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardtypeAutocompleteElementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardtypeAutocompleteElementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardtypeAutocompleteElementResponse(val *CreditcardtypeAutocompleteElementResponse) *NullableCreditcardtypeAutocompleteElementResponse {
	return &NullableCreditcardtypeAutocompleteElementResponse{value: val, isSet: true}
}

func (v NullableCreditcardtypeAutocompleteElementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardtypeAutocompleteElementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


