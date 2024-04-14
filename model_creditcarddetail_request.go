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

// checks if the CreditcarddetailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcarddetailRequest{}

// CreditcarddetailRequest A Creditcarddetail Object
type CreditcarddetailRequest struct {
	// The expirationmonth of the Creditcarddetail
	ICreditcarddetailExpirationmonth int32 `json:"iCreditcarddetailExpirationmonth"`
	// The expirationyear of the Creditcarddetail
	ICreditcarddetailExpirationyear int32 `json:"iCreditcarddetailExpirationyear"`
	// The civic of the Creditcarddetail
	SCreditcarddetailCivic *string `json:"sCreditcarddetailCivic,omitempty"`
	// The street of the Creditcarddetail
	SCreditcarddetailStreet *string `json:"sCreditcarddetailStreet,omitempty"`
	// The zip of the Creditcarddetail
	SCreditcarddetailZip *string `json:"sCreditcarddetailZip,omitempty"`
}

type _CreditcarddetailRequest CreditcarddetailRequest

// NewCreditcarddetailRequest instantiates a new CreditcarddetailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcarddetailRequest(iCreditcarddetailExpirationmonth int32, iCreditcarddetailExpirationyear int32) *CreditcarddetailRequest {
	this := CreditcarddetailRequest{}
	this.ICreditcarddetailExpirationmonth = iCreditcarddetailExpirationmonth
	this.ICreditcarddetailExpirationyear = iCreditcarddetailExpirationyear
	return &this
}

// NewCreditcarddetailRequestWithDefaults instantiates a new CreditcarddetailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcarddetailRequestWithDefaults() *CreditcarddetailRequest {
	this := CreditcarddetailRequest{}
	return &this
}

// GetICreditcarddetailExpirationmonth returns the ICreditcarddetailExpirationmonth field value
func (o *CreditcarddetailRequest) GetICreditcarddetailExpirationmonth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICreditcarddetailExpirationmonth
}

// GetICreditcarddetailExpirationmonthOk returns a tuple with the ICreditcarddetailExpirationmonth field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailRequest) GetICreditcarddetailExpirationmonthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICreditcarddetailExpirationmonth, true
}

// SetICreditcarddetailExpirationmonth sets field value
func (o *CreditcarddetailRequest) SetICreditcarddetailExpirationmonth(v int32) {
	o.ICreditcarddetailExpirationmonth = v
}

// GetICreditcarddetailExpirationyear returns the ICreditcarddetailExpirationyear field value
func (o *CreditcarddetailRequest) GetICreditcarddetailExpirationyear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ICreditcarddetailExpirationyear
}

// GetICreditcarddetailExpirationyearOk returns a tuple with the ICreditcarddetailExpirationyear field value
// and a boolean to check if the value has been set.
func (o *CreditcarddetailRequest) GetICreditcarddetailExpirationyearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ICreditcarddetailExpirationyear, true
}

// SetICreditcarddetailExpirationyear sets field value
func (o *CreditcarddetailRequest) SetICreditcarddetailExpirationyear(v int32) {
	o.ICreditcarddetailExpirationyear = v
}

// GetSCreditcarddetailCivic returns the SCreditcarddetailCivic field value if set, zero value otherwise.
func (o *CreditcarddetailRequest) GetSCreditcarddetailCivic() string {
	if o == nil || IsNil(o.SCreditcarddetailCivic) {
		var ret string
		return ret
	}
	return *o.SCreditcarddetailCivic
}

// GetSCreditcarddetailCivicOk returns a tuple with the SCreditcarddetailCivic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcarddetailRequest) GetSCreditcarddetailCivicOk() (*string, bool) {
	if o == nil || IsNil(o.SCreditcarddetailCivic) {
		return nil, false
	}
	return o.SCreditcarddetailCivic, true
}

// HasSCreditcarddetailCivic returns a boolean if a field has been set.
func (o *CreditcarddetailRequest) HasSCreditcarddetailCivic() bool {
	if o != nil && !IsNil(o.SCreditcarddetailCivic) {
		return true
	}

	return false
}

// SetSCreditcarddetailCivic gets a reference to the given string and assigns it to the SCreditcarddetailCivic field.
func (o *CreditcarddetailRequest) SetSCreditcarddetailCivic(v string) {
	o.SCreditcarddetailCivic = &v
}

// GetSCreditcarddetailStreet returns the SCreditcarddetailStreet field value if set, zero value otherwise.
func (o *CreditcarddetailRequest) GetSCreditcarddetailStreet() string {
	if o == nil || IsNil(o.SCreditcarddetailStreet) {
		var ret string
		return ret
	}
	return *o.SCreditcarddetailStreet
}

// GetSCreditcarddetailStreetOk returns a tuple with the SCreditcarddetailStreet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcarddetailRequest) GetSCreditcarddetailStreetOk() (*string, bool) {
	if o == nil || IsNil(o.SCreditcarddetailStreet) {
		return nil, false
	}
	return o.SCreditcarddetailStreet, true
}

// HasSCreditcarddetailStreet returns a boolean if a field has been set.
func (o *CreditcarddetailRequest) HasSCreditcarddetailStreet() bool {
	if o != nil && !IsNil(o.SCreditcarddetailStreet) {
		return true
	}

	return false
}

// SetSCreditcarddetailStreet gets a reference to the given string and assigns it to the SCreditcarddetailStreet field.
func (o *CreditcarddetailRequest) SetSCreditcarddetailStreet(v string) {
	o.SCreditcarddetailStreet = &v
}

// GetSCreditcarddetailZip returns the SCreditcarddetailZip field value if set, zero value otherwise.
func (o *CreditcarddetailRequest) GetSCreditcarddetailZip() string {
	if o == nil || IsNil(o.SCreditcarddetailZip) {
		var ret string
		return ret
	}
	return *o.SCreditcarddetailZip
}

// GetSCreditcarddetailZipOk returns a tuple with the SCreditcarddetailZip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreditcarddetailRequest) GetSCreditcarddetailZipOk() (*string, bool) {
	if o == nil || IsNil(o.SCreditcarddetailZip) {
		return nil, false
	}
	return o.SCreditcarddetailZip, true
}

// HasSCreditcarddetailZip returns a boolean if a field has been set.
func (o *CreditcarddetailRequest) HasSCreditcarddetailZip() bool {
	if o != nil && !IsNil(o.SCreditcarddetailZip) {
		return true
	}

	return false
}

// SetSCreditcarddetailZip gets a reference to the given string and assigns it to the SCreditcarddetailZip field.
func (o *CreditcarddetailRequest) SetSCreditcarddetailZip(v string) {
	o.SCreditcarddetailZip = &v
}

func (o CreditcarddetailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcarddetailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iCreditcarddetailExpirationmonth"] = o.ICreditcarddetailExpirationmonth
	toSerialize["iCreditcarddetailExpirationyear"] = o.ICreditcarddetailExpirationyear
	if !IsNil(o.SCreditcarddetailCivic) {
		toSerialize["sCreditcarddetailCivic"] = o.SCreditcarddetailCivic
	}
	if !IsNil(o.SCreditcarddetailStreet) {
		toSerialize["sCreditcarddetailStreet"] = o.SCreditcarddetailStreet
	}
	if !IsNil(o.SCreditcarddetailZip) {
		toSerialize["sCreditcarddetailZip"] = o.SCreditcarddetailZip
	}
	return toSerialize, nil
}

func (o *CreditcarddetailRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"iCreditcarddetailExpirationmonth",
		"iCreditcarddetailExpirationyear",
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

	varCreditcarddetailRequest := _CreditcarddetailRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcarddetailRequest)

	if err != nil {
		return err
	}

	*o = CreditcarddetailRequest(varCreditcarddetailRequest)

	return err
}

type NullableCreditcarddetailRequest struct {
	value *CreditcarddetailRequest
	isSet bool
}

func (v NullableCreditcarddetailRequest) Get() *CreditcarddetailRequest {
	return v.value
}

func (v *NullableCreditcarddetailRequest) Set(val *CreditcarddetailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcarddetailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcarddetailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcarddetailRequest(val *CreditcarddetailRequest) *NullableCreditcarddetailRequest {
	return &NullableCreditcarddetailRequest{value: val, isSet: true}
}

func (v NullableCreditcarddetailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcarddetailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


