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

// checks if the CustomCreditcardtransactionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomCreditcardtransactionResponse{}

// CustomCreditcardtransactionResponse A custom Creditcardtransaction Object
type CustomCreditcardtransactionResponse struct {
	ECreditcardtypeCodename *FieldECreditcardtypeCodename `json:"eCreditcardtypeCodename,omitempty"`
	// The amount of the Creditcardtransaction
	DCreditcardtransactionAmount string `json:"dCreditcardtransactionAmount"`
	// The partially decrypted credit card number used in the Creditcardtransaction
	SCreditcardtransactionPartiallydecryptednumber string `json:"sCreditcardtransactionPartiallydecryptednumber"`
	// The reference number on the creditcard service for the Creditcardtransaction
	SCreditcardtransactionReferencenumber string `json:"sCreditcardtransactionReferencenumber"`
}

// NewCustomCreditcardtransactionResponse instantiates a new CustomCreditcardtransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomCreditcardtransactionResponse(dCreditcardtransactionAmount string, sCreditcardtransactionPartiallydecryptednumber string, sCreditcardtransactionReferencenumber string) *CustomCreditcardtransactionResponse {
	this := CustomCreditcardtransactionResponse{}
	this.DCreditcardtransactionAmount = dCreditcardtransactionAmount
	this.SCreditcardtransactionPartiallydecryptednumber = sCreditcardtransactionPartiallydecryptednumber
	this.SCreditcardtransactionReferencenumber = sCreditcardtransactionReferencenumber
	return &this
}

// NewCustomCreditcardtransactionResponseWithDefaults instantiates a new CustomCreditcardtransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomCreditcardtransactionResponseWithDefaults() *CustomCreditcardtransactionResponse {
	this := CustomCreditcardtransactionResponse{}
	return &this
}

// GetECreditcardtypeCodename returns the ECreditcardtypeCodename field value if set, zero value otherwise.
func (o *CustomCreditcardtransactionResponse) GetECreditcardtypeCodename() FieldECreditcardtypeCodename {
	if o == nil || IsNil(o.ECreditcardtypeCodename) {
		var ret FieldECreditcardtypeCodename
		return ret
	}
	return *o.ECreditcardtypeCodename
}

// GetECreditcardtypeCodenameOk returns a tuple with the ECreditcardtypeCodename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCreditcardtransactionResponse) GetECreditcardtypeCodenameOk() (*FieldECreditcardtypeCodename, bool) {
	if o == nil || IsNil(o.ECreditcardtypeCodename) {
		return nil, false
	}
	return o.ECreditcardtypeCodename, true
}

// HasECreditcardtypeCodename returns a boolean if a field has been set.
func (o *CustomCreditcardtransactionResponse) HasECreditcardtypeCodename() bool {
	if o != nil && !IsNil(o.ECreditcardtypeCodename) {
		return true
	}

	return false
}

// SetECreditcardtypeCodename gets a reference to the given FieldECreditcardtypeCodename and assigns it to the ECreditcardtypeCodename field.
func (o *CustomCreditcardtransactionResponse) SetECreditcardtypeCodename(v FieldECreditcardtypeCodename) {
	o.ECreditcardtypeCodename = &v
}

// GetDCreditcardtransactionAmount returns the DCreditcardtransactionAmount field value
func (o *CustomCreditcardtransactionResponse) GetDCreditcardtransactionAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DCreditcardtransactionAmount
}

// GetDCreditcardtransactionAmountOk returns a tuple with the DCreditcardtransactionAmount field value
// and a boolean to check if the value has been set.
func (o *CustomCreditcardtransactionResponse) GetDCreditcardtransactionAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DCreditcardtransactionAmount, true
}

// SetDCreditcardtransactionAmount sets field value
func (o *CustomCreditcardtransactionResponse) SetDCreditcardtransactionAmount(v string) {
	o.DCreditcardtransactionAmount = v
}

// GetSCreditcardtransactionPartiallydecryptednumber returns the SCreditcardtransactionPartiallydecryptednumber field value
func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionPartiallydecryptednumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardtransactionPartiallydecryptednumber
}

// GetSCreditcardtransactionPartiallydecryptednumberOk returns a tuple with the SCreditcardtransactionPartiallydecryptednumber field value
// and a boolean to check if the value has been set.
func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionPartiallydecryptednumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardtransactionPartiallydecryptednumber, true
}

// SetSCreditcardtransactionPartiallydecryptednumber sets field value
func (o *CustomCreditcardtransactionResponse) SetSCreditcardtransactionPartiallydecryptednumber(v string) {
	o.SCreditcardtransactionPartiallydecryptednumber = v
}

// GetSCreditcardtransactionReferencenumber returns the SCreditcardtransactionReferencenumber field value
func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionReferencenumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SCreditcardtransactionReferencenumber
}

// GetSCreditcardtransactionReferencenumberOk returns a tuple with the SCreditcardtransactionReferencenumber field value
// and a boolean to check if the value has been set.
func (o *CustomCreditcardtransactionResponse) GetSCreditcardtransactionReferencenumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SCreditcardtransactionReferencenumber, true
}

// SetSCreditcardtransactionReferencenumber sets field value
func (o *CustomCreditcardtransactionResponse) SetSCreditcardtransactionReferencenumber(v string) {
	o.SCreditcardtransactionReferencenumber = v
}

func (o CustomCreditcardtransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomCreditcardtransactionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ECreditcardtypeCodename) {
		toSerialize["eCreditcardtypeCodename"] = o.ECreditcardtypeCodename
	}
	toSerialize["dCreditcardtransactionAmount"] = o.DCreditcardtransactionAmount
	toSerialize["sCreditcardtransactionPartiallydecryptednumber"] = o.SCreditcardtransactionPartiallydecryptednumber
	toSerialize["sCreditcardtransactionReferencenumber"] = o.SCreditcardtransactionReferencenumber
	return toSerialize, nil
}

type NullableCustomCreditcardtransactionResponse struct {
	value *CustomCreditcardtransactionResponse
	isSet bool
}

func (v NullableCustomCreditcardtransactionResponse) Get() *CustomCreditcardtransactionResponse {
	return v.value
}

func (v *NullableCustomCreditcardtransactionResponse) Set(val *CustomCreditcardtransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomCreditcardtransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomCreditcardtransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomCreditcardtransactionResponse(val *CustomCreditcardtransactionResponse) *NullableCustomCreditcardtransactionResponse {
	return &NullableCustomCreditcardtransactionResponse{value: val, isSet: true}
}

func (v NullableCustomCreditcardtransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomCreditcardtransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


