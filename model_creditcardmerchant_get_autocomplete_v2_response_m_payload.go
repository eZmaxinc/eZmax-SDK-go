/*
eZmax API Definition (Full)

This API expose all the functionnalities for the eZmax and eZsign applications.

API version: 1.2.2
Contact: support-api@ezmax.ca
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eZmaxApi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreditcardmerchantGetAutocompleteV2ResponseMPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreditcardmerchantGetAutocompleteV2ResponseMPayload{}

// CreditcardmerchantGetAutocompleteV2ResponseMPayload Payload for POST /2/object/creditcardmerchant/getAutocomplete
type CreditcardmerchantGetAutocompleteV2ResponseMPayload struct {
	// An array of Creditcardmerchant autocomplete element response.
	AObjCreditcardmerchant []CreditcardmerchantAutocompleteElementResponse `json:"a_objCreditcardmerchant"`
}

type _CreditcardmerchantGetAutocompleteV2ResponseMPayload CreditcardmerchantGetAutocompleteV2ResponseMPayload

// NewCreditcardmerchantGetAutocompleteV2ResponseMPayload instantiates a new CreditcardmerchantGetAutocompleteV2ResponseMPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditcardmerchantGetAutocompleteV2ResponseMPayload(aObjCreditcardmerchant []CreditcardmerchantAutocompleteElementResponse) *CreditcardmerchantGetAutocompleteV2ResponseMPayload {
	this := CreditcardmerchantGetAutocompleteV2ResponseMPayload{}
	this.AObjCreditcardmerchant = aObjCreditcardmerchant
	return &this
}

// NewCreditcardmerchantGetAutocompleteV2ResponseMPayloadWithDefaults instantiates a new CreditcardmerchantGetAutocompleteV2ResponseMPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditcardmerchantGetAutocompleteV2ResponseMPayloadWithDefaults() *CreditcardmerchantGetAutocompleteV2ResponseMPayload {
	this := CreditcardmerchantGetAutocompleteV2ResponseMPayload{}
	return &this
}

// GetAObjCreditcardmerchant returns the AObjCreditcardmerchant field value
func (o *CreditcardmerchantGetAutocompleteV2ResponseMPayload) GetAObjCreditcardmerchant() []CreditcardmerchantAutocompleteElementResponse {
	if o == nil {
		var ret []CreditcardmerchantAutocompleteElementResponse
		return ret
	}

	return o.AObjCreditcardmerchant
}

// GetAObjCreditcardmerchantOk returns a tuple with the AObjCreditcardmerchant field value
// and a boolean to check if the value has been set.
func (o *CreditcardmerchantGetAutocompleteV2ResponseMPayload) GetAObjCreditcardmerchantOk() ([]CreditcardmerchantAutocompleteElementResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.AObjCreditcardmerchant, true
}

// SetAObjCreditcardmerchant sets field value
func (o *CreditcardmerchantGetAutocompleteV2ResponseMPayload) SetAObjCreditcardmerchant(v []CreditcardmerchantAutocompleteElementResponse) {
	o.AObjCreditcardmerchant = v
}

func (o CreditcardmerchantGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreditcardmerchantGetAutocompleteV2ResponseMPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a_objCreditcardmerchant"] = o.AObjCreditcardmerchant
	return toSerialize, nil
}

func (o *CreditcardmerchantGetAutocompleteV2ResponseMPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a_objCreditcardmerchant",
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

	varCreditcardmerchantGetAutocompleteV2ResponseMPayload := _CreditcardmerchantGetAutocompleteV2ResponseMPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreditcardmerchantGetAutocompleteV2ResponseMPayload)

	if err != nil {
		return err
	}

	*o = CreditcardmerchantGetAutocompleteV2ResponseMPayload(varCreditcardmerchantGetAutocompleteV2ResponseMPayload)

	return err
}

type NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload struct {
	value *CreditcardmerchantGetAutocompleteV2ResponseMPayload
	isSet bool
}

func (v NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) Get() *CreditcardmerchantGetAutocompleteV2ResponseMPayload {
	return v.value
}

func (v *NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) Set(val *CreditcardmerchantGetAutocompleteV2ResponseMPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditcardmerchantGetAutocompleteV2ResponseMPayload(val *CreditcardmerchantGetAutocompleteV2ResponseMPayload) *NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload {
	return &NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload{value: val, isSet: true}
}

func (v NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditcardmerchantGetAutocompleteV2ResponseMPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


